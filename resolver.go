package smartmei_backend

import (
	"context"
	"time"

	"bitbucket.org/challen_ge/smartmei_backend/internal/cotacao"
	"bitbucket.org/challen_ge/smartmei_backend/internal/crawler"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

type resultadoConsultado struct {
	tarifa *crawler.Tarifa
	err    error
}

func (r *queryResolver) TarifaAtual(ctx context.Context) (*Tarifa, error) {
	var err error
	agora := time.Now()

	chResultado := make(chan *resultadoConsultado)

	go func(chR chan *resultadoConsultado) {
		tarifa, err := crawler.ConsultarTarifa()
		chR <- &resultadoConsultado{tarifa, err}
	}(chResultado)

	var dto cotacao.DTO
	dto, err = cotacao.ConsultarCotacao()
	if err != nil {
		return nil, err
	}

	conversor := cotacao.NewConversor(dto)

	resultado := <-chResultado
	if resultado.err != nil {
		return nil, resultado.err
	}

	tarifaConsultada := resultado.tarifa

	tarifa := &Tarifa{
		Data:       agora.Format("2006-01-02 15:04:05"),
		Descricao:  tarifaConsultada.Descricao,
		ValorReal:  *tarifaConsultada.Valor,
		ValorDolar: conversor.ParaDolar(*tarifaConsultada.Valor),
		ValorEuro:  conversor.ParaEuro(*tarifaConsultada.Valor),
	}
	return tarifa, err
}
