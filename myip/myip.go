package main

/*
MyIpBot
---------

MyIpBot simply listens for http requests and remits the headers and raw data.

The only argument is the channel_name which defaults to the dev channel

*/

import (

	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/mgmtech/gobots/registry"
	zmq "github.com/pebbe/zmq3"
)

type RegEntry registry.RegEntry

var Registry = RegEntry{
	Name:     "myip",
	Port:     557, // RPC Port 0mq
	Fend:     "",
	Bend:     "ipc://myip.ipc",
	Commands: nil,
	Settings: map[string]string{
		"MYIP_PORT":   "8666",
	},
}

const (
	githubTemplates = `
        {{ define "git-url" }}http://www.github.com{{ end }}
        {{ define "git-repo" }}{{ .Repository.Url }}{{ end }}
        {{ define "git-compare" }}{{ .Repository.Url }}/compare/{{ .CompBranch }}...{{ .After }}{{ end }}`
)

var tmplGit = template.Must(template.New("git").Parse(githubTemplates))

func info(msg string) { log.Printf("INFO (Parrot)-> %v", msg) }

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

	// Handle the post-receive from github.com..
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
    
            resp_str := fmt.Sprintf("%v", r)

			log.Printf("%v", resp_str)
			client.Send(resp_str, 0)
		})

	log.Fatal(http.ListenAndServe(":"+Registry.Settings["MYIP_PORT"], nil))

}


func main () {

    SrvStart()
}
