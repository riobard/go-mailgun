package mailgun

import (
	"encoding/json"
	"net/url"
)

type Email struct {
	From      string
	To        []string
	Cc        []string
	Bcc       []string
	Subject   string
	Html      string
	Text      string
	Headers   url.Values
	Options   url.Values
	Variables url.Values
}

func (mg Mailgun) Send(domain string, email *Email) (msgId string, err error) {
	v := url.Values{}
	v.Set("from", email.From)
	for _, to := range email.To {
		v.Add("to", to)
	}
	for _, cc := range email.Cc {
		v.Add("cc", cc)
	}
	for _, bcc := range email.Bcc {
		v.Add("bcc", bcc)
	}
	v.Set("subject", email.Subject)
	v.Set("text", email.Text)
	v.Set("html", email.Html)

	for k, ls := range email.Headers {
		for _, e := range ls {
			v.Add("h:"+k, e)
		}
	}
	for k, ls := range email.Options {
		for _, e := range ls {
			v.Add("o:"+k, e)
		}
	}
	for k, ls := range email.Variables {
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
