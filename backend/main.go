package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Akagi201/light"
	"github.com/gohttp/logger"
	"github.com/gohttp/serve"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Host string `long:"host" default:"0.0.0.0" description:"Host to bind to"`
	Port uint16 `long:"port" default:"3201" description:"Port to bind to"`
}

func main() {
	_, err := flags.Parse(&opts)
	if err != nil {
		if !strings.Contains(err.Error(), "Usage") {
			log.Printf("error: %v\n", err.Error())
			os.Exit(1)
		} else {
			log.Printf("%v\n", err.Error())
			os.Exit(0)
		}
	}

	app := light.New()

	app.Use(logger.New())
	app.Use(serve.New("public"))

	log.Printf("HTTP listening at %s:%d", opts.Host, opts.Port)
	app.Listen(fmt.Sprintf("%s:%d", opts.Host, opts.Port))
}
