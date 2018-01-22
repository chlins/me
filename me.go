package main

import (
	"os"

	"github.com/chlins/me/core"
)

func main() {
	ender := make(chan os.Signal, 1)
	app := core.NewApp("6001")
	app.Start()
	<-ender
}
