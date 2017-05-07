package common

type config struct {
	Server string
}

var Config config

func init() {
	Config.Server = "127.0.0.1:8000" // XXX: Read from config file
}
