package nes

import "log"

type GameROM struct {
    program []byte
    character []byte
    PRGROMSize uint
    CHRROMSize uint
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
    log.Printf("debug: program rom size:%x", programROMPages)
    var characterROMPages uint = uint(nes[5])
    log.Printf("debug: character rom size:%x", characterROMPages)

    var characterROMStart uint = nesHeaderSize + programROMPages * programROMPageSize
    log.Printf("debug: character rom start at 0x%x", characterROMStart)
    var characterROMEnd uint = characterROMStart + characterROMPages * characterROMPageSize
    log.Printf("debug: character rom end at 0x%x", characterROMEnd)

    return &GameROM{
        program : nes[nesHeaderSize:characterROMStart - 1],
        character : nes[characterROMStart:characterROMEnd],
        PRGROMSize : programROMPages,
        CHRROMSize : characterROMPages,
    }
}
