package nes

import "log"

type PPURegister struct {
    PPUCTRL struct{
        VBlankNMI bool
        PPUMasterSlave bool
        SpriteSize bool
        BackgroundTableAddress bool
        SpriteTableAddress bool
        PPUAddressIncrement bool
        NameTableAddress byte
    }
    PPUMASK struct {
        BackgroundColor byte
        SpriteEnable bool
        BackgroundEnable bool
        SpriteMask bool
        BackgroundMask bool
        DisplayType bool
    }
    PPUSTATUS struct {
        VBlack bool
        SpriteHit bool
        ScanLineSprite bool
    }
    OAMADDR byte
    OAMDATA byte
    PPUSCROLLX byte
    PPUSCROLLY byte
    PPUADDR uint16
    PPUDATA byte
    
    //register write order status
    ScrollXWrited bool
    PPUAddressUpperBitWrited bool


}

func NewPPURegister() *PPURegister {
    var register PPURegister
    register.ByteToPPUCTRL(0)
    register.ByteToPPUMASK(0)
    register.ByteToPPUSTATUS(0)
    register.OAMADDR = 0
    register.OAMDATA = 0
    register.PPUSCROLLX = 0
    register.PPUSCROLLY = 0
    register.PPUADDR = 0
    register.PPUDATA = 0
    register.ScrollXWrited = false
    register.PPUAddressUpperBitWrited = false
    return &register
}

func (register *PPURegister) Reset() {
    register.ByteToPPUCTRL(0)
    register.ByteToPPUMASK(0)
    register.OAMADDR = 0
    register.OAMDATA = 0
    register.PPUSCROLLX = 0
    register.PPUSCROLLY = 0
    register.PPUDATA = 0
    register.ScrollXWrited = false
    register.PPUAddressUpperBitWrited = false
}

//probably will not use
func (register *PPURegister)PPUCTRLToByte() byte {
    boolToByte := map[bool]byte{true:1,false:0}
    var status byte
    status |= boolToByte[register.PPUCTRL.VBlankNMI] << 7
    status |= boolToByte[register.PPUCTRL.PPUMasterSlave] << 6
    status |= boolToByte[register.PPUCTRL.SpriteSize] << 5
    status |= boolToByte[register.PPUCTRL.BackgroundTableAddress] << 4
    status |= boolToByte[register.PPUCTRL.SpriteTableAddress] << 3
    status |= boolToByte[register.PPUCTRL.PPUAddressIncrement] << 2
    status |= register.PPUCTRL.NameTableAddress & 0x03
    return status
}

func (register *PPURegister) ByteToPPUCTRL(status byte){
    byteToBool := map[byte]bool{1:true,0:false}
    register.PPUCTRL.VBlankNMI = byteToBool[(status >> 7) & 0x01]
    register.PPUCTRL.PPUMasterSlave = byteToBool[(status >> 6) & 0x01]
    register.PPUCTRL.SpriteSize = byteToBool[(status >> 5) & 0x01]
    register.PPUCTRL.BackgroundTableAddress = byteToBool[(status >> 4) & 0x01]
    register.PPUCTRL.SpriteTableAddress = byteToBool[(status >> 3) & 0x01]
    register.PPUCTRL.PPUAddressIncrement = byteToBool[(status >> 2) & 0x01]
    register.PPUCTRL.NameTableAddress = status & 0x03 
}

//probably will not use
func (register *PPURegister)PPUMASKToByte() byte {
    boolToByte := map[bool]byte{true:1,false:0}
    var status byte
    status |= register.PPUMASK.BackgroundColor << 5
    status |= boolToByte[register.PPUMASK.SpriteEnable] << 4
    status |= boolToByte[register.PPUMASK.BackgroundEnable] << 3
    status |= boolToByte[register.PPUMASK.BackgroundMask] << 2
    status |= boolToByte[register.PPUMASK.DisplayType]
    return status
}

func (register *PPURegister) ByteToPPUMASK(status byte){
    byteToBool := map[byte]bool{1:true,0:false}
    register.PPUMASK.BackgroundColor = (status & 0xE0) >> 5 
    register.PPUMASK.SpriteEnable = byteToBool[(status >> 4) & 0x01]
    register.PPUMASK.BackgroundEnable = byteToBool[(status >> 3) & 0x01]
    register.PPUMASK.SpriteMask = byteToBool[(status >> 2) & 0x01]
    register.PPUMASK.BackgroundMask = byteToBool[(status >> 1) & 0x01]
    register.PPUMASK.DisplayType = byteToBool[status & 0x01]
}

func (register *PPURegister)PPUSTATUSToByte() byte {
    boolToByte := map[bool]byte{true:1,false:0}
    var status byte
    status |= boolToByte[register.PPUSTATUS.VBlack] << 7
    status |= boolToByte[register.PPUSTATUS.SpriteHit] << 6
    status |= boolToByte[register.PPUSTATUS.ScanLineSprite] << 5
    return status
}

//probably will not use
func (register *PPURegister) ByteToPPUSTATUS(status byte){
    byteToBool := map[byte]bool{1:true,0:false}
    register.PPUSTATUS.VBlack = byteToBool[(status >> 7) & 0x01]
    register.PPUSTATUS.SpriteHit = byteToBool[(status >> 6) & 0x01]
    register.PPUSTATUS.ScanLineSprite = byteToBool[(status >> 5) & 0x01]
}

type PPU struct {
    Register *PPURegister
    Bus *PPUBus
    SpriteRam *RAM

    Clock int
    
    //buffer for ppudata access
    PPUDataBuffer byte
}

func NewPPU(vram *RAM, rom *GameROM) *PPU {
    var ppu PPU
    ppu.Bus = NewPPUBus(vram, rom)
    ppu.SpriteRam = NewRAM(0x100)
    ppu.Register = NewPPURegister()
    return &ppu
}

func (ppu *PPU) ReadPPURegister(addr uint16) byte{
    reg := (addr - 0x2000) % 8
    if reg == 0x0002 {
        ppu.Register.ScrollXWrited = false
        return ppu.Register.PPUSTATUSToByte()
    }else if reg == 0x0004 {
        return ppu.Register.OAMDATA
    }else if reg == 0x0007 {
        address := ppu.Register.PPUADDR
        if !ppu.Register.PPUCTRL.PPUAddressIncrement {
            ppu.Register.PPUADDR += 0x01
        }else{
            ppu.Register.PPUADDR += 0x20
        }
        if address <= 0x3eff {
            res := ppu.PPUDataBuffer
            ppu.PPUDataBuffer = ppu.Bus.ReadByte(address)
            return res
        }else{ 
           return ppu.Bus.ReadByte(address)
        }  
    }else{
        log.Printf("error: read access to write only register:%x",addr)
        return 0
    }
}

func (ppu *PPU) WritePPURegister(addr uint16, data byte){
    reg := (addr - 0x2000) % 8
    if reg == 0x0000 {
        ppu.Register.ByteToPPUCTRL(data)
    }else if reg == 0x0001 {
        ppu.Register.ByteToPPUMASK(data)
    }else if reg == 0x0003 {
        ppu.Register.OAMADDR = data
    }else if reg == 0x0004 {
        ppu.SpriteRam.Write(uint16(ppu.Register.OAMADDR),data)
        ppu.Register.OAMADDR++
    }else if reg ==0x0005 {
        if !ppu.Register.ScrollXWrited {
            ppu.Register.PPUSCROLLX = data
            ppu.Register.ScrollXWrited = true
        }else{
            ppu.Register.PPUSCROLLY = data
            ppu.Register.ScrollXWrited = false
        }
    }else if reg == 0x0006 {
        if !ppu.Register.PPUAddressUpperBitWrited {
            ppu.Register.PPUADDR = uint16(data) << 8
            ppu.Register.PPUAddressUpperBitWrited = true
        }else{
            ppu.Register.PPUADDR |= uint16(data)
            ppu.Register.PPUAddressUpperBitWrited = false
        }
    }else if reg == 0x0007 {
        ppu.Bus.WriteByte(ppu.Register.PPUADDR, data)
        if !ppu.Register.PPUCTRL.PPUAddressIncrement {
            ppu.Register.PPUADDR += 0x01
        }else{
            ppu.Register.PPUADDR += 0x20
        }
    }else{
        log.Printf("error: write access to read only register:%x",addr)
    }
}

type Tile struct {
    sprite []byte
    pallete byte
}

func NewTile(posX byte, posY byte, offset uint16, bus PPUBus) Tile {
    var tile Tile
    tile.sprite = make([]byte,16)

    SpriteAddress := uint16(posY * 0x20 + posX) + offset
    BlockID := int((posX % 4) / 2) + int((posY % 4) / 2) * 2
    AttrAddress := SpriteAddress + 0x03c0
    
    for i := uint16(0); i < 16; i++ {
        tile.sprite = append(tile.sprite, bus.ReadByte(SpriteAddress + i))
    }
    tile.pallete = (bus.ReadByte(AttrAddress) >> byte(BlockID * 2)) & 0x3 
    
    return tile
}
