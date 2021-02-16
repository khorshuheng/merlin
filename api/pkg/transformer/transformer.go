package transformer

import "context"

type Transformer interface {
	Initialize(configJson string) error
	Transform(ctx context.Context, request []byte) ([]byte, error)
	CleanUp()
}