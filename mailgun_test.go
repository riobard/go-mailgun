package mailgun

import (
	"flag"
)

var (
	key    = flag.String("key", "", "Mailgun key")
	domain = flag.String("domain", "", "Test domain")
	from   = flag.String("from", "", "Test mail sender address")
	to     = flag.String("to", "", "Test mail recipient address")
	mg     *Mailgun
)

func init() {
	flag.Parse()
	mg = Open(*key)
}
