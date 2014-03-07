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
        "region": {"yy"},
        "apartment_num": {"23"},
    }
    _, _, err := f.PostForm("/140303_wx", data)
    if err != nil { 
        return 
    }
    //println(string(body))
    <-execChans
}


func main(){
    execChans = make(chan bool, 1000)
    for{
        execChans <- true
        go blood()
    }
}
