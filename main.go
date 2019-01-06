package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "github.com/comail/colog"    
    "./nes"
)

func main(){
    if len(os.Args) < 2 {
        fmt.Printf("usage: %s [romfile_path] ([debug(0or1)])", os.Args[0])
    }else{
        debug := false
        if len(os.Args) == 3 {
            if os.Args[2] == "1" {
                debug = true
            }
        }
        bytes, err := ioutil.ReadFile(os.Args[1])
        if err != nil {
            log.Fatal(err)
        }

        colog.SetDefaultLevel(colog.LDebug)
        if debug {
            colog.SetMinLevel(colog.LDebug)
        }else {
            colog.SetMinLevel(colog.LError)
        }
        colog.SetFormatter(&colog.StdFormatter{
            Colors: true,
            Flag:   log.Ldate | log.Ltime | log.Lshortfile,
        })
        colog.Register()
        
        n := nes.NewNES(bytes)
        n.Start(debug)
        
        
    }
}
