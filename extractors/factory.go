package extractors

import (
	"errors"
)

func CreateExtractor(extractorName string) (Extractor, error) {
	if extractorName == "ContainerHosts" {
		return &containerHosts{}, nil
	}

	if extractorName == "Traefik" {
		return &traefik{}, nil
	}

	return nil, errors.New("unsupported extractor")
}
