package bots
/*
Main config for bots
*/
import "log"

import "github.com/mgmtech/bots/registry"

// Import the bots
import "github.com/mgmtech/bots/parrot"
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
