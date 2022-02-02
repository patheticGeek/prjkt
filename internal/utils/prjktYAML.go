package utils

import (
	"errors"
	types "internal/types"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func ReadRepoPrjktYAML(destination string) ([]byte, error) {
	prjktYAMLPath := filepath.Join(destination, "/prjkt.yaml")

	// Check if prjkt.yaml exists
	if _, err := os.Stat(prjktYAMLPath); errors.Is(err, os.ErrNotExist) {
		return []byte{}, nil
	}

	// Read the prjkt.yaml file
	fileData, err := ioutil.ReadFile(prjktYAMLPath)
	if err != nil {
		return []byte{}, err
	}

	return fileData, nil
}

func GetDefaultActionUrl(name string) string {
	// Templates are stored here: https://github.com/patheticGeek/prjkt-templates/defaults
	return "https://raw.githubusercontent.com/patheticGeek/prjkt-templates/main/defaults/" + name + ".yaml"
}

func GetDefaultAction(name string) ([]byte, error) {
	// Get the data
	url := GetDefaultActionUrl(name)
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Read request body into fileData
	fileData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return []byte(string(fileData)), nil
}

func ParsePrjktYAML(fileData []byte) (types.PrjktYAML, error) {
	data := types.PrjktYAML{}

	err := yaml.Unmarshal(fileData, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
