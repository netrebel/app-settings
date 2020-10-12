package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vrischmann/envconfig"
)

// AppSettings - JSON model
type AppSettings struct {
	Log struct {
		MinFilter string `envconfig:"optional"`
	}
	Cors struct {
		Origins []string `envconfig:"optional"`
	}
	SomeService struct {
		URL string `json:"Url" envconfig:"optional"`
	}
	SomePublisher struct {
		Env     string `envconfig:"optional"`
		Project string `envconfig:"optional"`
		Topic   string `envconfig:"optional"`
	}
	Google struct {
		Application struct {
			Credentials string
		}
	}
	SomeXML struct {
		Storage struct {
			BucketID string `json:"BucketId" envconfig:"optional"`
		}
		AuthBasicToken string `envconfig:"optional"`
	} `json:"SomeXml"`
	DefaultProvider string `envconfig:"optional"`
}

// ReadFromFileAndEnv reads the settings from a local file and applies any existing environment variables to it.
func ReadFromFileAndEnv(settings interface{}) error {
	file, err := os.Open("appsettings.json")

	if err != nil {
		return err
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return fmt.Errorf("Failed to read appsettings.json: %v", err)
	}

	err = json.Unmarshal(data, settings)

	if err != nil {
		return fmt.Errorf("Failed to unmarshal appsettings: %v", err)
	}

	err = envconfig.Init(settings)

	if err != nil {
		err = fmt.Errorf("Failed to update with env vars: %v", err)
	}
	return err
}
