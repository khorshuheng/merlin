package custom

import "context"

type Jaeger struct {

}

func (j Jaeger) Initialize(configJson string) error {
	panic("implement me")
}

func (j Jaeger) Transform(ctx context.Context, request []byte) ([]byte, error) {
	panic("implement me")
}

func (j Jaeger) CleanUp() {
	panic("implement me")
}
