package main

import (
	"fmt"
)

func main() {
	// placeholder variable
	settings := AppSettings{}

	// filling the variable with the settings file and env vars
	if err := ReadFromFileAndEnv(&settings); err != nil {
		panic(err)
	}

	// do something with the settings
	fmt.Printf("%+v", settings)
}
