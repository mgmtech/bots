package bots
/*
Main config for bots
*/
import "log"

import registry "github.com/mgmtech/gobot/bots/registry"

// Import the bots
import parrot "github.com/mgmtech/gobot/bots/parrot"
//import burt "github.com/mgmtech/gobot/bots/burt"
//import webvu "github.com/mgmtech/gobot/bots/webvu"


// main botmap
var Registry = registry.BotRegistry{
	"parrot": parrot.Registry,
 //   "burt": burt.Registry,
   // "webvu": webvu.Registry,
}

func Start() {
    for k, v := range Registry {
    
            log.Print(k,v)
    }
}
