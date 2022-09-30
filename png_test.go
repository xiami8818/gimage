package gimage

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

const (
	_brokenImg  = "resources/images/broken1.png"
	_brokenImg2 = "resources/images/broken2.png"
	_brokenImg3 = "resources/images/broken3.jpeg"
	_brokenImg4 = "resources/images/broken4.png"
)

func openImg(path string, t *testing.T) []byte {
	file, err := os.Open(path)
	if err != nil {
		t.Fatal("open file error: ", err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal("read file error: ", err)
	}
	return content
}

func TestFormatPng(t *testing.T) {
	img := openImg(_brokenImg4, t)
	imageByte, err := FixPng(img)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("fixed.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.Write(imageByte)
	if err != nil {
		t.Error(err)
	}
	err = file.Close()
	if err != nil {
		t.Error(err)
	}
	t.Log("success")
}

func TestImage(t *testing.T) {
	data := bytes.Buffer{}
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewNRGBA(rect)
	err := png.Encode(&data, img)
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.Create("test.png")
	if err != nil {
		t.Fatal(err)
	}
	_, err = file.Write(data.Bytes())
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestOpenImg(t *testing.T) {
	file, err := os.Open("../../../resources/images/broken3.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	_, err = png.Decode(file)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}

func TestZlib(t *testing.T) {
	img := openImg(_brokenImg, t)
	rd := bytes.NewReader(img)
	reader, err := zlib.NewReader(rd)
	if err != nil {
		t.Fatal(err)
	}
	read, err := reader.Read(img)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(read)
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	println(len(content))
}
