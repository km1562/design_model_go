package main

import (
	"fmt"
	"sync"
)

var loc = sync.Mutex{}

type single struct{

}

var SingleInstance *single

func getSingleInstance() *single{
	if SingleInstance == nil{
		loc.Lock()
		defer loc.Unlock()
		if SingleInstance == nil{
			SingleInstance = &single{}
		}else{
			fmt.Println("SingleInstance is already created!")
		}
	}else{
		fmt.Println("SingleInstance is already created!")
	}
	return SingleInstance
}

