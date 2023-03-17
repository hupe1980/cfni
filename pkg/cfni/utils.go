package cfni

import (
	"bytes"
	"encoding/hex"
	"text/template"
)

func executeTemplate(name string, data any) (*bytes.Buffer, error) {
	tpl, err := template.ParseFS(templates, name)
	if err != nil {
		return nil, err
	}

	result := new(bytes.Buffer)
	if err = tpl.Execute(result, data); err != nil {
		return nil, err
	}

	return result, nil
}

func xor(input, key string) (output string) {
	for i := 0; i < len(input); i++ {
		output += string(input[i] ^ key[i%len(key)])
	}

	return output
}

func hexify(input string) string {
	return hex.EncodeToString([]byte(input))
}
