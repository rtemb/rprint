package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/rtemb/rprint/version"
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

// Root url
func Root(c *router.Control) {
	fmt.Fprintf(c.Writer, "Processing URL %s...", c.Request.URL.Path)
}

// ServiceInfo defines HTTP API response giving service information
type ServiceInfo struct {
	Host    string       `json:"host"`
	Runtime *RuntimeInfo `json:"runtime"`
	Version string       `json:"version"`
	Repo    string       `json:"repo"`
	Commit  string       `json:"commit"`
}

// RuntimeInfo defines runtime part of service information
type RuntimeInfo struct {
	Compiler   string `json:"compilier"`
	CPU        int    `json:"cpu"`
	Memory     string `json:"memory"`
	Goroutines int    `json:"goroutines"`
}

// Info provides JSON API response giving service information
func Info(c *router.Control) {
	host, _ := os.Hostname()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)

	rt := &RuntimeInfo{
		CPU:        runtime.NumCPU(),
		Memory:     fmt.Sprintf("%.2fMB", float64(m.Alloc)/(1<<(10*2))),
		Goroutines: runtime.NumGoroutine(),
	}

	info := ServiceInfo{
		Host:    host,
		Runtime: rt,
		Version: version.RELEASE,
		Repo:    version.REPO,
		Commit:  version.COMMIT,
	}

	c.Code(http.StatusOK).Body(info)
}

// Healthz provides JSON API response giving live service or down
func Healthz(c *router.Control) {
	c.Code(http.StatusOK).Body(http.StatusText(http.StatusOK))
}
