package main

import (
	"fmt"

	"github.com/netrebel/app-settings/conf"
)

func main() {
	// placeholder variable
	settings := conf.AppSettings{}

	// filling the variable with the settings file and env vars
	if err := conf.ReadFromFileAndEnv(&settings); err != nil {
		panic(err)
	}

	// do something with the settings
	fmt.Printf("%+v", settings)
}
