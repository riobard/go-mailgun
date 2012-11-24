package mailgun

import (
	"flag"
)

var key = flag.String("key", "", "Mailgun key")
var domain = flag.String("domain", "", "Test domain")
var from = flag.String("from", "", "Test mail sender address")
var to = flag.String("to", "", "Test mail recipient address")
var mg *Mailgun

func init() {
	flag.Parse()
	mg = Open(*key)
}
