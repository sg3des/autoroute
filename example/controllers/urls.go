package controllers

import (
	"net/http"
)

type Urls struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Urls) Index() {
	c.W.Write([]byte("Urls Index"))
}

func (c *Urls) Echo() {
	c.W.Write([]byte("Urls Echo"))
}
