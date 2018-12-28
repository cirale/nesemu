package nes

type CPUBus struct{
    RAM *RAM
    PRGROM []byte
    PRGROMSize uint
    ppu *PPU
}

func (bus *CPUBus) ReadByte(addr uint16) byte{
    if 0x0000 <= addr && addr <= 0x07ff {
        // WRAM
        return bus.RAM.Read(addr)
        
    }else if 0x0800 <= addr && addr <= 0x1fff {
        // WRAM Mirror
        return bus.RAM.Read(addr - 0x0800)
        
    }else if 0x2000 <= addr && addr <= 0x3fff {
        // PPU Register
        return bus.ppu.ReadPPURegister(addr)
        
    }else if 0x4000 <= addr && addr <= 0x401f {
        // APU I/O, PAD
        return 0
        
    }else if 0x4020 <= addr && addr <= 0x5fff {
        // extended rom
        return 0
        
    }else if 0x6000 <= addr && addr <= 0x7fff {
        // extended ram
        return 0
        
    }else if 0x8000 <= addr && addr <= 0xBfff {
        // program rom
        return bus.PRGROM[addr - 0x8000]
        
    }else if 0xc000 <= addr && addr <= 0xffff {
        // program rom
        if bus.PRGROMSize < 2 {
            return bus.PRGROM[addr - 0xc000]    
        }
        return bus.PRGROM[addr - 0x8000]
    }
    return 0
}
 
func (bus *CPUBus) ReadWord(addr uint16) uint16 {
    var Lower uint16 = uint16(bus.ReadByte(addr))
    var Upper uint16 = uint16(bus.ReadByte(addr+1))
    return (Upper << 8) | Lower 
}

func (bus *CPUBus) WriteByte(addr uint16, data byte){
    if 0x0000 <= addr && addr <= 0x07ff {
        // WRAM
        bus.RAM.Write(addr,data)
        
    }else if 0x0800 <= addr && addr <= 0x1fff {
        // WRAM Mirror
        bus.RAM.Write(addr - 0x0800,data)
        
    }else if 0x2000 <= addr && addr <= 0x3fff {
        // PPU Register
        bus.ppu.WritePPURegister(addr, data)
        
    }else if 0x4000 <= addr && addr <= 0x401f {
        // APU I/O, PAD
        
    }else if 0x4020 <= addr && addr <= 0x5fff {
        // extended rom
        
    }else if 0x6000 <= addr && addr <= 0x7fff {
        // extended ram
        
    }
}


func NewBus(ram *RAM, rom *GameROM, ppu *PPU) *CPUBus {
    var bus CPUBus
    bus.RAM = ram
    bus.PRGROM = rom.program
    bus.PRGROMSize = rom.PRGROMSize
    bus.ppu = ppu
    return &bus
}

type PPUBus struct {
    VRAM *RAM
    CHRROM []byte
    CHRROMSize uint
}

func NewPPUBus(ram *RAM, rom *GameROM) *PPUBus {
    var bus PPUBus
    bus.VRAM = ram
    bus.CHRROM = rom.character
    bus.CHRROMSize = rom.CHRROMSize
    
    return &bus
}

func (bus *PPUBus) ReadByte(addr uint16) byte {
    return bus.VRAM.Read(addr)
}

func (bus *PPUBus) WriteByte(addr uint16, data byte){
    
}
