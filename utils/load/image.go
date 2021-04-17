package load

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type Image struct {
	imagPath string
	name     string
}

func NewImage(imagePath, name string) *Image {
	return &Image{
		imagPath: imagePath,
		name:     name,
	}
}

func (i *Image) Download() error {
	resp, err := http.Get(i.imagPath)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	out, err := os.Create(i.name)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, bytes.NewReader(body))
	return err
}
