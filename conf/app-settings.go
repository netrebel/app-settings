package conf

//https://travix.io/making-your-go-app-configurable-bb5e5f4a9df9

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
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
	file, err := os.Open("conf/appsettings.json")

	if err != nil {
		return err
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		return errors.Wrap(err, "Failed to read appsettings")
	}

	err = json.Unmarshal(data, settings)

	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal appsettings")
	}

	err = envconfig.Init(settings)

	if err != nil {
		err = errors.Wrap(err, "Failed to update with env vars")
	}
	return err
}
