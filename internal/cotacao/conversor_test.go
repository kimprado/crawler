package cotacao

import (
	"testing"
)

func TestConverterMoedas(t *testing.T) {
	testesConversao := []struct {
		valor         float64
		cotacaoDolar  float64
		cotacaoEuro   float64
		dolarEsperado float64
		euroEsperado  float64
	}{
		{7, 0.2401456353, 0.2180169181, 1.6810194471, 1.5261184267},
		{7, 0.3083175565, 0.2524296352, 2.1582228955, 1.7670074464},
	}

	for _, tc := range testesConversao {

		c := NewConversor(DTO{
			Rates: Rates{
				Dolar: tc.cotacaoDolar,
				Euro:  tc.cotacaoEuro,
			},
		})

		if c.ParaDolar(tc.valor) != tc.dolarEsperado {
			t.Errorf("ParaDolar(%v) esperado %v, mas recebido %v\n", tc.valor, tc.dolarEsperado, c.ParaDolar(tc.valor))
		}
		if c.ParaEuro(tc.valor) != tc.euroEsperado {
			t.Errorf("ParaEuro(%v) esperado %v, mas recebido %v\n", tc.valor, tc.euroEsperado, c.ParaEuro(tc.valor))
		}
	}
}
