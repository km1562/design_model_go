package main

import (
	"fmt"
	"sync"
)

var loc = sync.Mutex{}

type single struct{

}

var SingleInstance *single

fund Gets