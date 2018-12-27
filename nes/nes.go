package nes

import (
    "bufio"
    "os"
)

type NES struct {
    RAM *RAM
    ROM *GameROM
    CPU *CPU
}

func NewNES(rom []byte) *NES{
    var nes NES
    nes.ROM = ParseROM(rom)
    nes.RAM = NewRAM()
    nes.CPU = NewCPU(nes.RAM, nes.ROM)
    return &nes
    
}

func (nes *NES) Start(debug bool){
    for {
        nes.CPU.run()
        nes.CPU.Register.ShowRegister()
        if debug {
            bufio.NewScanner(os.Stdin).Scan()
        }
    }
    
}

func (nes *NES) Reset(){
    nes.CPU.reset()
}
