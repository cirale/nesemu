package nes

import (
    "bufio"
    "os"
    "log"
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
        cycle := nes.CPU.run()
        log.Printf("Instruction Finish. Cycle:%d", cycle)
        nes.PPU.run(cycle*3)
        nes.CPU.Register.ShowRegister()
        nes.PPU.Register.ShowPPURegister()
        if debug {
            bufio.NewScanner(os.Stdin).Scan()
        }
    }
    
}

func (nes *NES) Reset(){
    nes.CPU.reset()
}
