package nes

type RAM struct{
    ram []byte
}
func (ram RAM) ReadWord(addr uint16) uint16 {
    return 0
}

func NewRAM() *RAM {
    var ram RAM
    ram.ram = make([]byte,0x10000)

    return &ram
}
