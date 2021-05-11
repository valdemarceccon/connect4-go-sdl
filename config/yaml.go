package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const filename = "config.yaml"

// Connect4Config stores configuration for connect for
type Connect4Config struct {
	Dimensions Dimension
	Rows       int32
	Columns    int32
	Gap        int32
	Padding    int32
	Border     int32
}

// Dimension store height and width
type Dimension struct {
	Width  int32
	Height int32
}

// Load loads the configuration from file
func Load() (*Connect4Config, error) {

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return &Connect4Config{
			Dimensions: Dimension{
				Width:  800,
				Height: 600,
			},
			Gap:     10,
			Padding: 10,
			Border:  10,
			Rows:    6,
			Columns: 7,
		}, nil
	}

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	t := Connect4Config{}
	err = yaml.Unmarshal([]byte(data), &t)

	if err != nil {
		return nil, err
	}

	return &t, nil
}
