package cotacao

import (
	"encoding/json"
	"net/http"
)

const urlBaseUser = "https://api.exchangeratesapi.io/latest?base=BRL&symbols=USD,EUR"

// Conversor aplica conversões de valores em Real para outras moedas
// Usa cotação do dia para aplicar conversões
type Conversor struct {
	cotacao DTO
}

// NewConversor cria instância de Conversor
func NewConversor(dto DTO) (c Conversor) {
	c = Conversor{
		cotacao: dto,
	}
	return
}

//ParaEuro converte valor em real para Eudo
func (c Conversor) ParaEuro(valor float64) (dolar float64) {
	dolar = valor * c.cotacao.Rates.Euro
	return
}

//ParaDolar converte valor em real para Dolar
func (c Conversor) ParaDolar(valor float64) (dolar float64) {
	dolar = valor * c.cotacao.Rates.Dolar
	return
}

// DTO mapeia resposta da API de câmbio
type DTO struct {
	Rates Rates `json:"rates"`
}

// Rates mapeia resposta da API de câmbio
type Rates struct {
	Euro  float64 `json:"EUR"`
	Dolar float64 `json:"USD"`
}

func consultarCotacao() (dto DTO, err error) {

	err = executarRequisicao(urlBaseUser, &dto)

	return
}

func executarRequisicao(url string, objetoRetorno interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&objetoRetorno)
	return err
}
