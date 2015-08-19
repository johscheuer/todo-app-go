package main

import (
	"flag"
	"log"
	"net/http"
)

var slaveService *string
var masterService *string
var protocol *string
var framework *string
var domain *string

func main() {
	slaveService = flag.String("slaveService", "none", "Define the service name of the redis slave")
	masterService = flag.String("masterService", "none", "Define the service name of the redis slave")
	protocol = flag.String("protocol", "tcp", "Define the protocl the service is using")
	framework = flag.String("framework", "marathon", "Define the framework which runs the service")
	domain = flag.String("domain", "mesos", "Define the domain of Mesos-DNS")
	//	port := flag.String("port", "8080", "Port to listen on")
	flag.Parse()

	Pong()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
