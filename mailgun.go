package mailgun

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	API_VERSION  = 2
	API_ENDPOINT = "api.mailgun.net"
)

type Mailgun struct {
	httpClient *http.Client
	key        string
}

func Open(key string) *Mailgun {
	return &Mailgun{httpClient: &http.Client{}, key: key}
}

// make an api request
func (mg *Mailgun) api(method string, path string, fields url.Values) (body []byte, err error) {
	var req *http.Request
	url := fmt.Sprintf("https://%s/v%d%s", API_ENDPOINT, API_VERSION, path)

	if method == "POST" && fields != nil {
		req, err = http.NewRequest(method, url, strings.NewReader(fields.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	} else {
		if fields != nil {
			url += "?" + fields.Encode()
		}
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return
	}
	req.SetBasicAuth("api", mg.key)
	rsp, err := mg.httpClient.Do(req)
	if err != nil {
		return
	}
	defer rsp.Body.Close()
	body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return
	}
	if rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		err = fmt.Errorf("mailgun error: %d %s", rsp.StatusCode, body)
	}
	return
}
