package ginsitemap

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func New(path string) gin.HandlerFunc {
	path = filepath.FromSlash(path)
	if len(path) > 0 && !os.IsPathSeparator(path[0]) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		path = filepath.Join(wd, path)
	}

	info, err := os.Stat(path)
	if err != nil || info == nil || info.IsDir() {
		panic("Invalid sitemap path: " + path)
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	reader := bytes.NewReader(file)

	return func(c *gin.Context) {
		if c.Request.RequestURI != "/sitemap.xml" {
			return
		}
		if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
			status := http.StatusOK
			if c.Request.Method != "OPTIONS" {
				status = http.StatusMethodNotAllowed
			}
			c.Header("Allow", "GET,HEAD,OPTIONS")
			c.AbortWithStatus(status)
			return
		}
		c.Header("Content-Type", "text/xml")
		http.ServeContent(c.Writer, c.Request, "sitemap.xml", info.ModTime(), reader)
		return
	}
}
