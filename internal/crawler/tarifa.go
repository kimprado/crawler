package crawler

import (
	"flag"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var url string

func init() {
	// urlSite := flag.String("url", "https://www.smartmei.com.br", "Caminho para site")
	// flag.Parse()
	// url = *urlSite
}

//Tarifa representa preço de Transferência
type Tarifa struct {
	Descricao string
	Valor     *float64
}

var rgTarifa = regexp.MustCompile(`tarifas-2-2-2">\s{0,}(R\$\s{0,}\d{1,},\d{1,})`)
var rgValorTarifa = regexp.MustCompile(`(R\$\s{0,})(\d{1,})(,)(\d{1,})`)

func carregarSite(url string, client *http.Client) (body string, cookies []*http.Cookie, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	var bb []byte
	bb, err = ioutil.ReadAll(resp.Body)

	body = string(bb)
	cookies = resp.Cookies()
	return
}

func recuperarTarifa(body string) (descricao string) {
	t := rgTarifa.FindStringSubmatch(body)
	if len(t) > 1 {
		descricao = t[1]
	}
	return
}

func ConsultarTarifa() (t *Tarifa, err error) {

	if url == "" {
		urlSite := flag.String("url", "http://www.smartmei.com.br", "Caminho para site")
		flag.Parse()
		url = *urlSite
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	body, _, err := carregarSite(url, client)
	if err != nil {
		return
	}

	descricaoTarifa := recuperarTarifa(body)

	valorTarifa := strings.TrimSpace(descricaoTarifa)
	valorTarifa = rgValorTarifa.ReplaceAllString(valorTarifa, "$2.$4")

	valorDecimal, err := strconv.ParseFloat(valorTarifa, 64)
	if err != nil {
		return
	}

	t = &Tarifa{
		Descricao: descricaoTarifa,
		Valor:     &valorDecimal,
	}
	return
}
