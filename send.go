package mailgun

import (
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
)

type Mail interface {
	From() string
	To() []string
	Cc() []string
	Bcc() []string
	Subject() string
	Html() string
	Text() string
	Headers() url.Values
	Options() url.Values
	Variables() url.Values
}

var EMAIL_DOMAIN_RE *regexp.Regexp

func init() {
	EMAIL_DOMAIN_RE = regexp.MustCompile(`[^<>]+<?.+@([^<>]+)>?`)
}

func (mg Mailgun) Send(m Mail) (msgId string, err error) {
	match := EMAIL_DOMAIN_RE.FindStringSubmatch(m.From())
	if len(match) != 2 {
		err = fmt.Errorf("invalid From address: %s", m.From())
		return
	}
	domain := match[1]
	v := url.Values{}
	v.Set("from", m.From())
	for _, to := range m.To() {
		v.Add("to", to)
	}
	for _, cc := range m.Cc() {
		v.Add("cc", cc)
	}
	for _, bcc := range m.Bcc() {
		v.Add("bcc", bcc)
	}
	v.Set("subject", m.Subject())
	v.Set("html", m.Html())
	v.Set("text", m.Text())

	for k, ls := range m.Headers() {
		for _, e := range ls {
			v.Add("h:"+k, e)
		}
	}
	for k, ls := range m.Options() {
		for _, e := range ls {
			v.Add("o:"+k, e)
		}
	}
	for k, ls := range m.Variables() {
		for _, e := range ls {
			v.Add("v:"+k, e)
		}
	}

	rsp, err := mg.api("POST", "/"+domain+"/messages", v)
	if err != nil {
		return
	}
	var res struct {
		Message string `json:"message"`
		Id      string `json:"id"`
	}
	err = json.Unmarshal(rsp, &res)
	msgId = res.Id
	return
}
