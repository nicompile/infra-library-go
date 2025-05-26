package config

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Config struct {
	Name      string `yaml:"name"`
	Type      string `yaml:"type"`
	GoVersion string `yaml:"goVersion"`
	AppRoot   string `yaml:"appRoot"`
	Target    string `yaml:"target"`
}

func GetAll(dir string) ([]Config, error) {
	yamls, err := getYamls(dir)
	if err != nil {
		return nil, err
	}
	configs := make([]Config, len(yamls))

	for i, path := range yamls {
		c, err := getConfig(path)
		if err != nil {
			return nil, err
		}
		configs[i] = c
	}

	return configs, nil
}

func getYamls(dir string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() && filepath.Ext(path) == ".yaml" {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func getConfig(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	return parse(file)
}

func parse(r io.Reader) (Config, error) {
	c := Config{}
	err := yaml.NewDecoder(r).Decode(&c)
	return c, err
}
