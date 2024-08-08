package main

type Component interface {
	Accept(visitor ComponentVisitor)
}

type ComponentVisitor interface {
	VisitorDB(c *DB)
	VisitorServer(c *Server)
}
