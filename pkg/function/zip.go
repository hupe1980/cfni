package function

import (
	"archive/zip"
	"bytes"
	"io"
)

func (c *Client) zipFile(handler []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buf)

	zf, err := zipWriter.Create("lambda_function.py")
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(zf, bytes.NewReader(handler)); err != nil {
		return nil, err
	}

	zipWriter.Close()

	return buf.Bytes(), nil
}
