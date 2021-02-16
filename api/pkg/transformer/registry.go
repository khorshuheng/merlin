package transformer

import (
	"errors"
	"fmt"
	"github.com/gojek/merlin/pkg/transformer/custom"
)

func Create(id string) (Transformer, error) {
	switch id {
	case "jaeger":
		return custom.Jaeger{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Unknown transformer %s", id))
	}
}

