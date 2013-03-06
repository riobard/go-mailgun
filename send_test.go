package mailgun

import (
	"net/url"
	"testing"
)

type mail struct {
	from      string
	to        []string
	cc        []string
	bcc       []string
	subject   string
	html      string
	text      string
	headers   url.Values
	options   url.Values
	variables url.Values
}

func (m *mail) From() string          { return m.from }
func (m *mail) To() []string          { return m.to }
func (m *mail) Cc() []string          { return m.cc }
func (m *mail) Bcc() []string         { return m.bcc }
func (m *mail) Subject() string       { return m.subject }
func (m *mail) Html() string          { return m.html }
func (m *mail) Text() string          { return m.text }
func (m *mail) Headers() url.Values   { return m.headers }
func (m *mail) Options() url.Values   { return m.options }
func (m *mail) Variables() url.Values { return m.variables }

func TestSend(t *testing.T) {
	m := &mail{
		from:    *from,
		to:      []string{*to},
		subject: "test mail",
		text:    "this is a test mail",
	}
	m.headers = url.Values{}

	id, err := mg.Send(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MessageID = %s", id)
}
