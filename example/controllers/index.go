package controllers

import (
	"net/http"
)

type Main struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Main) Index() {
	c.W.Write([]byte("Main Index"))
}

func (c *Main) Echo() {
	c.W.Write([]byte("Main Echo"))
}
