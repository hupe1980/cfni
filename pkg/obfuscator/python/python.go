package python

import (
	"github.com/hupe1980/cfni/pkg/obfuscator"
)

type Python struct{}

func New() obfuscator.Obfuscator {
	return &Python{}
}

func (o *Python) Obfuscate(code string) (string, error) {
	return code, nil
}
