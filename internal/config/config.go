package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func ValidateConfig[T any](config *T) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(config)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		logrus.Error(errors)
		return err
	}
	return nil
}

func ReadConfigFromYAML[T any](path string) (*T, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file from %s, due to: %w", path, err)
	}
	var conf T
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from %s, due to: %w", path, err)
	}

	return &conf, nil
}
