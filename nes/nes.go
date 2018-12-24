package nes

type NES struct {
    RAM *RAM
    ROM *GameROM
    CPU *CPU
}

func Start(rom []byte) *NES {
    var nes NES
    nes.ROM = ParseROM(rom)
    nes.CPU = NewCPU()
    nes.RAM = NewRAM()

    return &nes
}
