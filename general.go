package main

import (
	"github.com/takama/router"
)

// logger provides a log of requests
func logger(c *router.Control) {
	remoteAddr := c.Request.Header.Get("X-Forwarded-For")
	if remoteAddr == "" {
		remoteAddr = c.Request.RemoteAddr
	}
	log.Infof("%s %s %s", remoteAddr, c.Request.Method, c.Request.URL.Path)
}
