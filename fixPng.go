package gimage

import "bytes"

func FixPng(data []byte) ([]byte, error) {
	reader := bytes.NewReader(data)
	decode, err := Decode(reader)
	if err != nil && err != pixelError {
		return nil, err
	}
	out := bytes.Buffer{}
	err = Encode(&out, decode)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
