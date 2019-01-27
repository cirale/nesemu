package nes

type CPUBus struct{
    RAM *RAM
    PRGROM []byte
    PRGROMSize uint
    ppu *PPU
    keypad *KeyPad
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
        
    }else if addr == 0x4016 {
        // APU I/O, PAD
        return bus.keypad.Read()
        
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
        
    }else if addr == 0x4016 {
        // APU I/O, PAD
        bus.keypad.Write(data)
        
    }else if 0x4020 <= addr && addr <= 0x5fff {
        // extended rom
        
    }else if 0x6000 <= addr && addr <= 0x7fff {
        // extended ram
        
    }
}


func NewBus(ram *RAM, rom *GameROM, ppu *PPU, keypad *KeyPad) *CPUBus {
    var bus CPUBus
    bus.RAM = ram
    bus.PRGROM = rom.program
    bus.PRGROMSize = rom.PRGROMSize
    bus.ppu = ppu
    bus.keypad = keypad
    return &bus
}

type PPUBus struct {
    VRAM *RAM
    CHRROM []byte
    CHRROMSize uint
    PalletesTable *RAM
}

func NewPPUBus(ram *RAM, rom *GameROM) *PPUBus {
    var bus PPUBus
    bus.VRAM = ram
    bus.CHRROM = rom.character
    bus.CHRROMSize = rom.CHRROMSize
    bus.PalletesTable = NewRAM(0x20)

    return &bus
}

func (bus *PPUBus) ReadByte(addr uint16) byte {
    if 0x00 <= addr && addr <= 0x0fff {
        return bus.CHRROM[addr]
    }else if 0x1000 <= addr && addr <= 0x1fff {
        return bus.CHRROM[addr - 0x1000]
    }else if 0x2000 <= addr && addr <= 0x2fff {
        return bus.VRAM.Read(addr - 0x2000)
    }else if 0x3000 <= addr && addr <= 0x3eff {
        return bus.VRAM.Read(addr - 0x3000)
    }else if 0x3f00 <= addr && addr <= 0x3fff {
        address := (addr - 0x3f00) % 0x20
        if address >= 0x0010 && address % 0x4 == 0 {
            address -= 0x10
        }
        return bus.PalletesTable.Read(address)
    }else{
        return 0
    }
}

func (bus *PPUBus) WriteByte(addr uint16, data byte){
    if 0x00 <= addr && addr <= 0x0fff {
        bus.CHRROM[addr] = data
    }else if 0x1000 <= addr && addr <= 0x1fff {
        bus.CHRROM[addr - 0x1000] = data
    }else if 0x2000 <= addr && addr <= 0x2fff {
        bus.VRAM.Write(addr - 0x2000, data)
    }else if 0x3000 <= addr && addr <= 0x3eff {
        bus.VRAM.Write(addr - 0x3000, data)
    }else if 0x3f00 <= addr && addr <= 0x3fff {
        address := (addr - 0x3f00) % 0x20
        if address >= 0x0010 && address % 0x4 == 0 {
            address -= 0x10
        }
        bus.PalletesTable.Write(address, data)
    }
}
