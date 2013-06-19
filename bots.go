package bots
/*
Main config for bots
*/
import "log"

import "github.com/mgmtech/gobots/registry"

// Import the bots
//import parrot "github.com/mgmtech/gobots/parrot"
//import burt "github.com/mgmtech/gobots/burt"
//import webvu "github.com/mgmtech/gobots/webvu"


var Roster = registry.BotRegistry{
//    RegEntry(burt.Registry),
//    RegEntry(webvu.Registry)
}

func Start() {
    for k, v := range Roster {
            log.Print(k,v)
    }
}
