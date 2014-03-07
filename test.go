package main

import (
	"./fetcher"
	"net/url"
        "regexp"
)

var execChans chan bool
var regExp = regexp.MustCompile(`form.tk.value = "[a-z]{19}"`)

func blood() {
    f := fetcher.NewFetcher("activity.wexinfruit.com")

    _, body, err := f.Get("/140303_wx") // create session
    param := regExp.FindAllString(string(body), -1)
    
    data := url.Values {
	"tk": {param[0][17:36]},
        "r": {"yy"},
        "n": {"5"},
	"submit_vote": {"投票"},
    }
    _, _, err = f.PostForm("/140303_wx", data)
    if err != nil { 
        return 
    }
    //println(err)
    <-execChans
}


func main(){
    execChans = make(chan bool, 1000)
    for{
        execChans <- true
        go blood()
    }
}
