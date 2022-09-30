package server

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

var xmlContentType = []string{"application/xml; charset=utf-8"}
var plainContentType = []string{"text/plain; charset=utf-8"}

func writeContextType(c *gin.Context, value []string) {
	if val := c.Request.Header.Get("Content-Type"); len(val) == 0 {
		c.Request.Header.Add("Content-Type", value[0])
	}
}

// Render render from bytes
func (srv *Server) Render(bytes []byte) {
	srv.GContext.Writer.WriteString(string(bytes))
}

// String render from string
func (srv *Server) String(str string) {
	writeContextType(srv.GContext, plainContentType)
	srv.Render([]byte(str))
}

// XML render to xml
func (srv *Server) XML(obj interface{}) {
	writeContextType(srv.GContext, xmlContentType)
	bytes, err := xml.Marshal(obj)
	if err != nil {
		panic(err)
	}
	srv.Render(bytes)
}

// Query returns the keyed url query value if it exists
func (srv *Server) Query(key string) string {
	value, _ := srv.GetQuery(key)
	return value
}

// GetQuery is like Query(), it returns the keyed url query value
func (srv *Server) GetQuery(key string) (string, bool) {
	req := srv.GContext
	if values, ok := req.Request.URL.Query()[key]; ok && len(values) > 0 {
		return values[0], true
	}
	return "", false
}
