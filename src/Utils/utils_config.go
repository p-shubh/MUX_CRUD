package Utils

import (
	"flag"
	"github.com/gorilla/sessions"
)

var configFile = flag.String("f", "config/DataCollector.json", "path for the configfile")
var store = sessions.NewCookieStore([]byte("some-encrypted-key"))
