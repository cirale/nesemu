package nes

import "log"

type RAM struct{
    ram []byte
    size uint16
}

func (ram RAM) Read(addr uint16) byte {
    if addr >= ram.size {
        log.Printf("error: address 0x%x is out of range, ram size is 0x%x", addr, ram.size)
        return 0
    }
    return ram.ram[addr]
}

func (ram RAM) Write(addr uint16, data byte){
    if addr >= ram.size {
        log.Printf("error: address 0x%x is out of range, ram size is 0x%x", addr, ram.size)
        return
    }
    ram.ram[addr] = data
}

func NewRAM(size uint16) *RAM {
    var ram RAM
    ram.size = size
    ram.ram = make([]byte,size)

    return &ram
}
