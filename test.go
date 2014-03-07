package main

import (
	"./fetcher"
	"net/url"
)

var execChans chan bool

func blood() {
    f := fetcher.NewFetcher("activity.wexinfruit.com")

    f.Get("/") // create session
    data := url.Values {
        "task_from": {"self"},
        "target": {"http://golang.org"},
        "ac": {"http"},
    }
    _, body, err := f.PostForm("/140303_wx", data)
    if err != nil { 
        return 
    }
    println(string(body))
    <-execChans
}


func main(){
    execChans = make(chan bool, 1000)
    for{
        execChans <- true
        go blood()
    }
}
