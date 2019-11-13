package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strconv"
	"time"
)

//Tarifa representa preço de Transferência
type Tarifa struct {
	Descricao string
	Valor     *float64
}

var rgTarifa = regexp.MustCompile(`(\d{1,})</div>`)

func carregarSite(url string, client *http.Client) (body string, cookies []*http.Cookie, err error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Host", "applicant-test.us-east-1.elasticbeanstalk.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:68.0) Gecko/20100101 Firefox/68.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")

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

func recuperarTarifa(body string) (descricao, valor string) {
	t := rgTarifa.FindStringSubmatch(body)
	if len(t) > 1 {
		valor = t[1]
	}
	return
}

func consultarTarifa(url string) (t *Tarifa) {

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout: time.Second * 30,
		Jar:     cookieJar,
	}

	body, _, err := carregarSite(url, client)
	if err != nil {
		log.Fatal(err)
	}

	descricaoTarifa, valorTarifa := recuperarTarifa(body)

	valorDecimal, err := strconv.ParseFloat(valorTarifa, 64)
	if err != nil {
		valorDecimal = 0
	}

	t = &Tarifa{
		Descricao: descricaoTarifa,
		Valor:     &valorDecimal,
	}
	return
}
