package bots
/*
Main config for bots
*/
import "log"

import "github.com/mgmtech/gobots/registry"

// Import the bots
import parrot "github.com/mgmtech/gobots/parrot2"
//import burt "github.com/mgmtech/gobots/burt"
//import webvu "github.com/mgmtech/gobots/webvu"


var Registry = registry.BotRegistry{
    "parrot": parrot.Registry,
//    RegEntry(burt.Registry),
//    RegEntry(webvu.Registry)
}

func Start() {
    for k, v := range Registry {
            log.Print(k,v)
            go Registry[k].SrvStart()
    }
}
