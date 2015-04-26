package rope

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func StdEnvelope(body string) string {

	t := `<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
   <soapenv:Header/>
   <soapenv:Body>
      %s
   </soapenv:Body>
</soapenv:Envelope>`

	return fmt.Sprintf(t, body)
}

type Service interface {
	Endpoint() string
	RequestTemplate() string

	RequestBody() string
}

type BasicAuth struct {
	Username string
	Password string
}

type SoapClient struct {
	UseTLS bool
	Auth   *BasicAuth
}

func (sc *SoapClient) SendServiceRequest(service Service) (string, error) {
	return sc.SendRequest(service.RequestBody(), service.Endpoint())
}

func (sc *SoapClient) SendRequest(body, url string) (string, error) {
	bodyBytes := []byte(body)
	bytebuffer := bytes.NewBuffer(bodyBytes)

	c := &http.Client{}
	if sc.UseTLS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{},
		}
		c.Transport = tr
	}

	req, err := http.NewRequest("POST", url, bytebuffer)

	if err != nil {
		return "", err
	}

	if sc.Auth != nil {
		req.SetBasicAuth(sc.Auth.Username, sc.Auth.Password)
	}

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}

}
