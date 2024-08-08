/*
	This pattern hides the complexity of the subsystem
	by allowing the customer to interact with the system
	through a simple interface.
*/

package main

import (
	"fmt"
)

// Database subsystem
type Database struct{}

func (db *Database) Connect() {
	fmt.Println("Connecting to the database...")
}

func (db *Database) Disconnect() {
	fmt.Println("Disconnecting from the database...")
}

// Logger subsystem
type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println("Logging message:", message)
}

// RequestHandler subsystem
type RequestHandler struct{}

func (rh *RequestHandler) HandleRequest(request string) {
	fmt.Println("Handling request:", request)
}

// Facade
type SystemFacade struct {
	db         *Database
	log        *Logger
	reqHandler *RequestHandler
}

func NewSystemFacade() *SystemFacade {
	return &SystemFacade{
		db:         &Database{},
		log:        &Logger{},
		reqHandler: &RequestHandler{},
	}
}

func (sf *SystemFacade) Start() {
	sf.db.Connect()
	sf.log.Log("System started")
}

func (sf *SystemFacade) ProcessRequest(request string) {
	sf.reqHandler.HandleRequest(request)
	sf.log.Log("Processed request: " + request)
}

func (sf *SystemFacade) Stop() {
	sf.db.Disconnect()
	sf.log.Log("System stopped")
}
