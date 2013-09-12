package bots
/*
Main config for bots
*/
import "log"

import "github.com/mgmtech/gobots/registry"



var Registry = registry.BotRegistry{
    "parrot": parrot.Registry,
}

func Start() {
    for k, v := range Registry {
            log.Print(k,v)
            go Registry[k].SrvStart()
    }
}
