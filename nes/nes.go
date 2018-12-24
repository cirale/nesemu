package nes

type NES struct {
    RAM *RAM
    ROM *GameROM
    CPU *CPU
}

func Start(rom []byte) *NES {
    var nes NES
    nes.ROM = ParseROM(rom)
    nes.RAM = NewRAM()
    nes.CPU = NewCPU(nes.RAM, nes.ROM)

    return &nes
}
