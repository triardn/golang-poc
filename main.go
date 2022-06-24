package main

import (
	"github.com/kitabisa/golang-poc/cmd"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	zlog.Logger = zlog.With().Caller().Logger()

    cmd.Execute()
}
