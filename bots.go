package bots
/*
Main config for bots
*/
import "log"

import registry "github.com/mgmtech/gobots/registry"

import parrot "github.com/mgmtech/gobots/parrot"

type RegEntry registry.RegEntry

var Registry = registry.BotRegistry{
    "parrot": parrot.RegistryEntry,
}

func Start() {
    for k, v := range Registry {
            log.Print(k,v)
            go Registry[k].SrvStart()
    }
}
