package main

/*
ParrotBot
---------

Parrot is a simple GoBot which uses ZeroMq to communicate using a Publisher socket
, clients connect to the backend url as Subscribers. A repo configured to push
to this servers registry.Settings["GITPUSHPORT"] will have a message remitted and
a compare link against registry.Settings["GITDIFFBRANCH"]. Eventually the settings
may be expanded to contain a list of repos, and branches to diff against.
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	registry "github.com/mgmtech/gobots/registry"
	zmq "github.com/pebbe/zmq3"
)


var RegistryEntry = registry.RegEntry{
	Name:     "parrot",
	Port:     556,
	Fend:     "",
	Bend:     "ipc://parrotbackend.ipc",
	Commands: nil,
	Settings: map[string]string{
		"GITPUSHPORT":   "8085",
		"GITDIFFBRANCH": "develop",
	},
}

/* Structs to map to the git post-receiver web hook payload */
type GitAuthor struct {
	Name  string
	Email string
}

type GitRepo struct {
	Name  string
	Url   string
	Owner GitAuthor
}

type GitCommit struct {
	Message   string
	Timestamp string
	Url       string
	Author    GitAuthor
}

type GitWebHookPayload struct {
	Before     string
	After      string
	Commits    []GitCommit
	Repository GitRepo
	CompBranch string
}

/* String template for url functions */
const (
	githubTemplates = `
        {{ define "git-url" }}http://www.github.com{{ end }}
        {{ define "git-repo" }}{{ .Repository.Url }}{{ end }}
        {{ define "git-compare" }}{{ .Repository.Url }}/compare/{{ .CompBranch }}...{{ .After }}{{ end }}`
)

var tmplGit = template.Must(template.New("git").Parse(githubTemplates))

func info(msg string) { log.Printf("INFO (Parrot)-> %v", msg) }

func (payload GitWebHookPayload) String() string {

	buff := bytes.NewBufferString("")

	if err := tmplGit.ExecuteTemplate(buff, "git-compare", payload); err != nil {
		log.Print("Error executing template")
		log.Print(err)
		return fmt.Sprintf("Error %v", err) // ERROR ! XXX: implement error handling fool
	}

	return fmt.Sprintf("%v", buff)
}

func CliStart() *zmq.Socket {
	client, err := zmq.NewSocket(zmq.SUB)
	if err != nil {
		log.Fatal("Problem connection to front-end")
	}

	client.Connect(RegistryEntry.Bend)
	client.SetSubscribe("")

	return client
}

func SrvStart() {
	// Link up publisher socket, could use Multicast here..
	client, err := zmq.NewSocket(zmq.PUB)
	if err != nil {
		log.Fatal("FAiled to connect push front-end %v", RegistryEntry.Bend)
	}
	defer client.Close()
	client.Bind(RegistryEntry.Bend)

	// Handle the post-receive from github.com..
	http.HandleFunc("/post-receive",
		func(w http.ResponseWriter, r *http.Request) {
			payload := r.FormValue("payload")
			info("Received github.com payload from github")

			var m GitWebHookPayload

			err := json.Unmarshal([]byte(payload), &m)

			if err != nil {
				log.Println("Error unpacking json:", err)
			}

            log.Print(len(m.Commits))
			m.CompBranch = RegistryEntry.Settings["GITDIFFBRANCH"]
			var resp_str = fmt.Sprintf("(parrot) %v, %v '%v'-> %v",
				 m.Repository.Name, m.Commits[0].Author.Name, m.Commits[0].Message, m)

			log.Printf("%v", resp_str)
			client.Send(resp_str, 0)
		})

	log.Fatal(http.ListenAndServe(":"+RegistryEntry.Settings["GITPUSHPORT"], nil))

}


func main () {

    SrvStart()
}
