package js

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hupe1980/cfni/pkg/obfuscator"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

type JS struct{}

func New() obfuscator.Obfuscator {
	return &JS{}
}

func (o *JS) Obfuscate(code string) (string, error) {
	r := bytes.NewBufferString(code)
	w := &bytes.Buffer{}

	m := js.Minifier{}
	if err := m.Minify(minify.New(), w, r, nil); err != nil {
		return "", err
	}

	c := w.String()
	if !strings.HasSuffix(c, ";") {
		c = fmt.Sprintf("%s;", c)
	}

	return c, nil
}
