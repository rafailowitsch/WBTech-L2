package main

type DB struct {
	adress   string
	dbname   string
	userinfo string
}

func (c *DB) Accept(visitor ComponentVisitor) {
	visitor.VisitorDB(c)
}

type Server struct {
	adress  string
	timeout string
}

func (c *Server) Accept(visitor ComponentVisitor) {
	visitor.VisitorServer(c)
}
