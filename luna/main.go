package main

import (
	"io/ioutil"

	eg "github.com/h8liu/luna/luna/hello"
)

func main() {
	const PageSize = 4096

	code, data := eg.Img()
	n := len(code)
	if n%4 != 0 {
		panic("code buffer not aligned")
	}
	if n > PageSize {
		panic("code buffer too large")
	}
	if len(data) > PageSize {
		panic("data buffer too large")
	}

	img := make([]byte, PageSize*2)
	copy(img[:PageSize], code)
	copy(img[PageSize:], data)

	e := ioutil.WriteFile("luna.img", img, 0666)
	if e != nil {
		panic(e)
	}
}
