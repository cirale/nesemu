package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "log"
    "github.com/comail/colog"    
    "./nes"
    "net/http"
    "github.com/gin-gonic/gin"
    "gopkg.in/olahol/melody.v1"
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

        ws := melody.New()
        n := nes.NewNES(bytes, ws)
        ServerInit(ws, n)

        //n.Start(debug)
        
    }
}

func ServerInit(m *melody.Melody, n *nes.NES){
    router := gin.Default()

    rg := router.Group("nesemu")
    rg.GET("/",func(ctx *gin.Context){
        http.ServeFile(ctx.Writer, ctx.Request, "./view/index.html")
        if !n.Running {
            go n.Start(false)
        }
    })
    
    rg.GET("/ws",func(ctx *gin.Context){
        m.HandleRequest(ctx.Writer, ctx.Request)
    })

    m.HandleConnect(func(s *melody.Session) {
        log.Printf("debug: websocket connection open. [session: %#v]\n", s)
    })

    m.HandleDisconnect(func(s *melody.Session) {
        log.Printf("debug: websocket connection close. [session: %#v]\n", s)
    })

    m.HandleMessage(func(s *melody.Session, msg []byte){
        n.KeyPad.Write(msg)
    })

    router.Run(":8989")
}
