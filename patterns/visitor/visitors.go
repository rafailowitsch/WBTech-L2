package main

import "strings"

type InfoVisitors struct {
	info string
}

func (i *InfoVisitors) VisitorDB(c *DB) {
	var builder strings.Builder

	builder.WriteString(c.adress)
	builder.WriteString(c.dbname)
	builder.WriteString(c.userinfo)

	result := builder.String()

	i.info = result
}

func (i *InfoVisitors) VisitorServer(c *Server) {
	var builder strings.Builder

	builder.WriteString(c.adress)
	builder.WriteString(c.timeout)

	result := builder.String()

	i.info = result
}
