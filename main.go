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
        
        n := nes.Start(bytes) 
        n.CPU.Bus.WriteByte(0x0000,0x7f)
        //fmt.Printf("%x",n.CPU.Bus.ReadWord(0x0000))
        var i nes.InstructionSet
        i.Inst = nes.ADC
        i.Mode = nes.Immediate
        n.CPU.Register.A = 1
        n.CPU.ExecInstruction(i)
        fmt.Printf("%x",n.CPU.Register.A)
        fmt.Printf("N:%v", n.CPU.Register.P.N)
        fmt.Printf("Z:%v", n.CPU.Register.P.Z)
        fmt.Printf("C:%v", n.CPU.Register.P.C)
        fmt.Printf("V:%v", n.CPU.Register.P.V)
    }
}

