package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "./nes"
)

func main(){
    if len(os.Args) != 2 {
        fmt.Printf("usage: %s [romfile_path]", os.Args[0])
    }else{
        bytes, err := ioutil.ReadFile(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }
        n := nes.NewNES(bytes)
        n.Start(true)
    }
}

