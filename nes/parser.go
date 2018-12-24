package nes

import "log"

type GameROM struct {
    program []byte
    character []byte
}

// ParseROM : nes rom parser
func ParseROM(nes []byte) *GameROM {

    const nesHeaderSize = 0x0010
    //unit of program rom
    const programROMPageSize = 0x4000
    //unit of character rom
    const characterROMPageSize = 0x2000

    // fmt.Printf("[% x]", nesFile)
    var programROMPages uint = uint(nes[4])
    log.Printf("program rom size:%d", programROMPages)
    var characterROMPages uint = uint(nes[5])
    log.Printf("character rom size:%d", characterROMPages)

    var characterROMStart uint = nesHeaderSize + programROMPages * programROMPageSize
    log.Printf("character rom start at %d", characterROMStart)
    var characterROMEnd uint = characterROMStart + characterROMPages * characterROMPageSize
    log.Printf("character rom end at %d", characterROMEnd)

    return &GameROM{
        program : nes[nesHeaderSize:characterROMStart - 1],
        character : nes[characterROMStart:characterROMEnd],
    }
}
