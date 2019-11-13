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

func (r *queryResolver) TarifaAtual(ctx context.Context) (*Tarifa, error) {
	var err error
	agora := time.Now()

	var dto cotacao.DTO

	tarifaConsultada := crawler.ConsultarTarifa()
	dto, err = cotacao.ConsultarCotacao()

	conversor := cotacao.NewConversor(dto)

	tarifa := &Tarifa{
		Data:       agora.Format("2006-01-02 15:04:05"),
		Descricao:  tarifaConsultada.Descricao,
		ValorReal:  *tarifaConsultada.Valor,
		ValorDolar: conversor.ParaDolar(*tarifaConsultada.Valor),
		ValorEuro:  conversor.ParaEuro(*tarifaConsultada.Valor),
	}
	return tarifa, err
}
