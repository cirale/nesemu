package nes

import (
    "log"
    "image"
    "image/png"
    "bytes"
    "fmt"
    "encoding/base64"
    "gopkg.in/olahol/melody.v1"
    
)

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

func (regiser PPURegister)ShowPPURegister(){
    log.Printf("debug: PPUCTRL:0x%02x, PPUMASK:0x%02x, PPUSTATUS:0x%02x, OAMADDR:0x%02x, ", regiser.PPUCTRLToByte(),regiser.PPUMASKToByte(),regiser.PPUSTATUSToByte(),regiser.OAMADDR)
    log.Printf("debug: PPUSCROLLX:0x%02x, PPUSCROLL:0x%02x, PPUADDR:0x%04x", regiser.PPUSCROLLX,regiser.PPUSCROLLY,regiser.PPUADDR)
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
    Line int

    //Image *image.RGBA
    Image []*Tile
    colors [][]byte
    ws *melody.Melody
    
    //buffer for ppudata access
    PPUDataBuffer byte
}

func NewPPU(vram *RAM, rom *GameROM, ws *melody.Melody) *PPU {
    var ppu PPU
    ppu.Bus = NewPPUBus(vram, rom)
    ppu.SpriteRam = NewRAM(0x100)
    ppu.Register = NewPPURegister()
    ppu.Image = make([]*Tile,32*30)
    ppu.ws = ws
    ppu.Reset()
    
    return &ppu
}

func (ppu *PPU) Reset(){
    ppu.Register.Reset()
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

func (ppu *PPU) TilesToPix() []byte {
    res := make([]byte,256*240*4)
    for _, tile := range ppu.Image {
        tile.TileToImage(res,ppu.Bus)
    }
    return res
}

func (ppu *PPU) run(cpuCycle int){
    ppu.Clock += cpuCycle
    if ppu.Clock >= 341 {
        ppu.Clock -= 341
        ppu.Line++
        log.Printf("Line:%d", ppu.Line)
    }
    log.Printf("debug: PPUClock:%d", ppu.Clock)
    if ppu.Line > 0 && ppu.Line <= 240 && ppu.Line % 8 == 0 {
        for i := 0; i < 32; i++ {
            ppu.Image[(ppu.Line / 8 - 1) * 32 + i] = NewTile(i, (ppu.Line / 8) - 1, 0x2000, ppu.Bus) 
        }
    }
    
    if ppu.Line == 241 {
        ppu.Register.PPUSTATUS.VBlack = true
        //log.Printf("error: Image:%v", ppu.Image)
        i := image.NewRGBA(image.Rect(0, 0, 256, 240))
        i.Pix = ppu.TilesToPix()
        
        var buffer bytes.Buffer
        if err := png.Encode(&buffer, i); err != nil {
            panic(err)
        }
        imgEnc := fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(buffer.Bytes()))
        var msg Message
        msg.Msgtype = "image"
        msg.Data = imgEnc
        if s, err := msg.toBytes(); err != nil{
            log.Fatal(err)
        }else{
            ppu.ws.Broadcast(s)
        }
    }
    if ppu.Line >= 260 {
        ppu.Line -= 260
        ppu.Register.PPUSTATUS.VBlack = false
    }

}

type Tile struct {
    sprite []byte
    pallete byte
    posX int
    posY int
    offset uint16
}

func NewTile(posX int, posY int, offset uint16, bus *PPUBus) *Tile {
    var tile Tile
    tile.posX = posX
    tile.posY = posY
    tile.offset = offset

    SpriteNo := bus.ReadByte(uint16(posY * 0x20 + posX) + offset)
    SpriteAddress := uint16(SpriteNo) * 16
    BlockID := int((posX % 4) / 2) + int((posY % 4) / 2) * 2
    AttrAddress := uint16(posX / 4) + (uint16(posY / 4) * 8) + 0x03c0 + offset
    
    for i := uint16(0); i < 16; i++ {
        tile.sprite = append(tile.sprite, bus.ReadByte(SpriteAddress + i))
    }
    tile.pallete = (bus.ReadByte(AttrAddress) >> byte(BlockID * 2)) & 0x3
    
    return &tile
}

func (tile Tile) TileToImage(Image []byte, bus *PPUBus) {
    var ColorPallete []byte
    //make color pallete
    for i := uint16(0); i < 4; i++ {
        ColorPallete = append(ColorPallete, bus.ReadByte(uint16(0x3f00) + uint16(tile.pallete)*4 + i))
    }

    for i := 0; i < 8; i++ {
        LowerBits := tile.sprite[i]
        UpperBits := tile.sprite[i + 8]
        for j := 0; j < 8; j++ {
            color := (LowerBits >> byte(7-j)) & 0x1 | (((UpperBits >> byte(7-j)) & 0x1) << 1)
            RGB := ColorRGB[ColorPallete[color]]
            X := (tile.posX * 8) + j
            Y := (tile.posY * 8) + i
            pos := (Y * 256 + X) * 4
            Image[pos] = RGB[0]; Image[pos+1] = RGB[1]; Image[pos+2] = RGB[2];Image[pos+3] = 255
        }
    }

}

var ColorRGB [][]byte =  [][]byte{
    {0x80, 0x80, 0x80}, {0x00, 0x3D, 0xA6}, {0x00, 0x12, 0xB0}, {0x44, 0x00, 0x96},
    {0xA1, 0x00, 0x5E}, {0xC7, 0x00, 0x28}, {0xBA, 0x06, 0x00}, {0x8C, 0x17, 0x00},
    {0x5C, 0x2F, 0x00}, {0x10, 0x45, 0x00}, {0x05, 0x4A, 0x00}, {0x00, 0x47, 0x2E},
    {0x00, 0x41, 0x66}, {0x00, 0x00, 0x00}, {0x05, 0x05, 0x05}, {0x05, 0x05, 0x05},
    {0xC7, 0xC7, 0xC7}, {0x00, 0x77, 0xFF}, {0x21, 0x55, 0xFF}, {0x82, 0x37, 0xFA},
    {0xEB, 0x2F, 0xB5}, {0xFF, 0x29, 0x50}, {0xFF, 0x22, 0x00}, {0xD6, 0x32, 0x00},
    {0xC4, 0x62, 0x00}, {0x35, 0x80, 0x00}, {0x05, 0x8F, 0x00}, {0x00, 0x8A, 0x55},
    {0x00, 0x99, 0xCC}, {0x21, 0x21, 0x21}, {0x09, 0x09, 0x09}, {0x09, 0x09, 0x09},
    {0xFF, 0xFF, 0xFF}, {0x0F, 0xD7, 0xFF}, {0x69, 0xA2, 0xFF}, {0xD4, 0x80, 0xFF},
    {0xFF, 0x45, 0xF3}, {0xFF, 0x61, 0x8B}, {0xFF, 0x88, 0x33}, {0xFF, 0x9C, 0x12},
    {0xFA, 0xBC, 0x20}, {0x9F, 0xE3, 0x0E}, {0x2B, 0xF0, 0x35}, {0x0C, 0xF0, 0xA4},
    {0x05, 0xFB, 0xFF}, {0x5E, 0x5E, 0x5E}, {0x0D, 0x0D, 0x0D}, {0x0D, 0x0D, 0x0D},
    {0xFF, 0xFF, 0xFF}, {0xA6, 0xFC, 0xFF}, {0xB3, 0xEC, 0xFF}, {0xDA, 0xAB, 0xEB},
    {0xFF, 0xA8, 0xF9}, {0xFF, 0xAB, 0xB3}, {0xFF, 0xD2, 0xB0}, {0xFF, 0xEF, 0xA6},
    {0xFF, 0xF7, 0x9C}, {0xD7, 0xE8, 0x95}, {0xA6, 0xED, 0xAF}, {0xA2, 0xF2, 0xDA},
    {0x99, 0xFF, 0xFC}, {0xDD, 0xDD, 0xDD}, {0x11, 0x11, 0x11}, {0x11, 0x11, 0x11},
}
