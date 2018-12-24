package nes

type Bus struct{
    RAM *RAM
    PRGROM []byte
}

func (bus Bus) ReadByte(addr uint16) byte{
    if 0x0000 <= addr && addr <= 0x07ff {
        // WRAM
        return bus.RAM.Read(addr)
        
    }else if 0x0800 <= addr && addr <= 0x1fff {
        // WRAM Mirror
        return bus.RAM.Read(addr - 0x8000)
        
    }else if 0x2000 <= addr && addr <= 0x2007 {
        // PPU Register
        return 0
        
    }else if 0x2008 <= addr && addr <= 0x3fff {
        // PPU Register Mirror
        return 0
        
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
        if len(bus.PRGROM) > 0x4000 {
            return bus.PRGROM[addr - 0xc000]    
        }
        return bus.PRGROM[addr - 0x8000]
    }
    return 0
}
 
func (bus Bus) ReadWord(addr uint16) uint16 {
    var Lower uint16 = uint16(bus.RAM.Read(addr))
    var Upper uint16 = uint16(bus.RAM.Read(addr+1))
    return (Upper << 8) | Lower 
}

func (bus Bus) WriteByte(addr uint16, data byte){
    bus.RAM.Write(addr, data)
}


func NewBus(ram *RAM, rom *GameROM) *Bus {
    var bus Bus
    bus.RAM = ram
    bus.PRGROM = rom.program
    
    return &bus
}
