package utils

import (
	types "internal/types"

	"gopkg.in/yaml.v3"
)

func ParsePrjktYAML(fileData []byte) (types.PrjktYAML, error) {
	data := types.PrjktYAML{}

	err := yaml.Unmarshal(fileData, &data)

	if err != nil {
		return data, err
	}

	return data, nil
}
