package nes

type RAM struct{
    ram []byte
}

func (ram RAM) Read(addr uint16) byte {
    return ram.ram[addr]
}

func (ram RAM) Write(addr uint16, data byte){
    ram.ram[addr] = data
}

func NewRAM() *RAM {
    var ram RAM
    ram.ram = make([]byte,0x10000)
    ram.ram[0xfffc] = 0x00
    ram.ram[0xfffd] = 0x80

    return &ram
}
