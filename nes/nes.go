package nes

import (
    "bufio"
    "os"
)

type NES struct {
    RAM *RAM
    VRAM *RAM
    ROM *GameROM
    CPU *CPU
    PPU *PPU
}

func NewNES(rom []byte) *NES{
    var nes NES
    nes.ROM = ParseROM(rom)
    nes.RAM = NewRAM(0x2000)
    nes.VRAM = NewRAM(0x2000)
    nes.PPU = NewPPU(nes.VRAM, nes.ROM)
    nes.CPU = NewCPU(nes.RAM, nes.ROM,nes.PPU)
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
