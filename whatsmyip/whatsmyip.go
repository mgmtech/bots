package main

/*
WhatsMyIpBot
---------

WhatsMyIpBot is a simple GoBot which uses ZeroMq to communicate using a Publisher socket
, clients connect to the backend url as Subscribers. 

When you make a request to a valid whatsmyip url it will remit it via 0mq to GoBot.

*/

import (
	"log"
	"net/http"

	"github.com/mgmtech/gobots/registry"
	zmq "github.com/pebbe/zmq3"
)

type RegEntry registry.RegEntry

var Registry = RegEntry{
	Name:     "whatsmyip",
	Port:     557,
	Fend:     "",
	Bend:     "ipc://ipbackend.ipc",
	Commands: nil,
	Settings: map[string]string{
	},
}

func CliStart() *zmq.Socket {
	client, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatal("Problem connection to front-end")
	}

	client.Connect(Registry.Bend)
	client.SetSubscribe("")

	return client
}

func SrvStart() {
	// Link up publisher socket, could use Multicast here..
	client, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		log.Fatal("FAiled to connect push front-end %v", Registry.Bend)
	}
	defer client.Close()
	client.Bind(Registry.Bend)

	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
        // send the meta-data needed
        resp_str := "NTSH"
			client.Send(resp_str, 0)
		})

	log.Fatal(http.ListenAndServe(":"+Registry.Settings["GITPUSHPORT"], nil))

}


func main () {

    SrvStart()
}
