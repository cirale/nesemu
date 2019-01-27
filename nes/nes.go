package nes

import (
    "bufio"
    "os"
    "log"
    "gopkg.in/olahol/melody.v1"
    "encoding/json"
    "time"
)

type NES struct {
    RAM *RAM
    VRAM *RAM
    ROM *GameROM
    CPU *CPU
    PPU *PPU
    KeyPad *KeyPad

    Running bool
}

type Message struct {
    Msgtype string `json:"msgtype"`
    Data string `json:"data"`
}

func (msg Message) toBytes() ([]byte, error){
    return json.Marshal(msg)
}

func NewNES(rom []byte, ws *melody.Melody) *NES{
    var nes NES
    nes.ROM = ParseROM(rom)
    nes.RAM = NewRAM(0x2000)
    nes.VRAM = NewRAM(0x2000)
    nes.KeyPad = NewKeyPad()
    nes.PPU = NewPPU(nes.VRAM, nes.ROM, ws)
    nes.CPU = NewCPU(nes.RAM, nes.ROM,nes.PPU, nes.KeyPad)
    
    nes.Running = false
    return &nes
}

func (nes *NES) Start(debug bool){
    nes.Running = true
    for {
        cycle := nes.CPU.run()
        log.Printf("Instruction Finish. Cycle:%d", cycle)
        nes.PPU.run(cycle*3)
        time.Sleep(time.Duration(cycle / 1790000) * time.Second)
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
