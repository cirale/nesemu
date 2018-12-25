package nes

// ----------------------------------------
// InstructionDecoder
// ----------------------------------------

type AddressingMode int
const (
    Accumulator AddressingMode = iota
    Immediate
    Absolute
    ZeroPage
    ZeroPageX
    ZeroPageY
    AbsoluteX
    AbsoluteY
    Implied
    Relative
    IndirectX
    IndirectY
    AbsoluteIndirect
)

type Instruction int
const (
    // Calculation
    ADC Instruction = iota; SBC
    // Logical operation
    AND; ORA; EOR
    // Shift
    ASL; LSR; ROL; ROR
    // Branch
    BCC; BCS; BEQ; BNE; BVC; BVS; BPL; BMI
    // Test
    BIT
    // Jump
    JMP; JSR; RTS
    // Interrupt
    BRK; RTI
    // Compare
    CMP; CPX; CPY
    // Increment,Decrement
    INC; DEC; INX; DEX; INY; DEY
    // Flag
    CLC; SEC; CLI; SEI; CLD; SED; CLV
    // Load
    LDA; LDX; LDY
    // Store
    STA; STX; STY
    // Register transfer
    TAX; TXA; TAY; TYA; TSX; TXS;
    // Stack
    PHA; PLA; PHP; PLP
    // No Operation
    NOP
)


type InstructionSet struct{
    Inst Instruction
    Mode AddressingMode
}

func (cpu CPU)Decode(opcode byte) InstructionSet{
    var inst InstructionSet
    if opcode == 0x6d {
        inst.Inst = ADC
        inst.Mode = Absolute
    }else if opcode == 0x7d {
        inst.Inst = ADC
        inst.Mode = AbsoluteX
    }else if opcode == 0x79 {
        inst.Inst = ADC
        inst.Mode = AbsoluteY
    }else if opcode == 0x69 {
        inst.Inst = ADC
        inst.Mode = Immediate
    }else if opcode == 0x61 {
        inst.Inst = ADC
        inst.Mode = IndirectX
    }else if opcode == 0x71 {
        inst.Inst = ADC
        inst.Mode = IndirectY
    }else if opcode == 0x65 {
        inst.Inst = ADC
        inst.Mode = ZeroPage
    }else if opcode == 0x75 {
        inst.Inst = ADC
        inst.Mode = ZeroPageX
    }else if opcode == 0x2d {
        inst.Inst = AND
        inst.Mode = Absolute
    }else if opcode == 0x3d {
        inst.Inst = AND
        inst.Mode = AbsoluteX
    }else if opcode == 0x39 {
        inst.Inst = AND
        inst.Mode = AbsoluteY
    }else if opcode == 0x29 {
        inst.Inst = AND
        inst.Mode = Immediate
    }else if opcode == 0x21 {
        inst.Inst = AND
        inst.Mode = IndirectX
    }else if opcode == 0x31 {
        inst.Inst = AND
        inst.Mode = IndirectY
    }else if opcode == 0x25 {
        inst.Inst = AND
        inst.Mode = ZeroPage
    }else if opcode == 0x35 {
        inst.Inst = AND
        inst.Mode = ZeroPageX
    }else if opcode == 0x0e {
        inst.Inst = ASL
        inst.Mode = Absolute
    }else if opcode == 0x1e {
        inst.Inst = ASL
        inst.Mode = AbsoluteX
    }else if opcode == 0x0a {
        inst.Inst = ASL
        inst.Mode = Accumulator
    }else if opcode == 0x06 {
        inst.Inst = ASL
        inst.Mode = ZeroPage
    }else if opcode == 0x16 {
        inst.Inst = ASL
        inst.Mode = ZeroPageX
    }else if opcode == 0x90 {
        inst.Inst = BCC
        inst.Mode = Relative
    }else if opcode == 0xb0 {
        inst.Inst = BCS
        inst.Mode = Relative
    }else if opcode == 0xf0 {
        inst.Inst = BEQ
        inst.Mode = Relative
    }else if opcode == 0x2c {
        inst.Inst = BIT
        inst.Mode = Absolute
    }else if opcode == 0x24 {
        inst.Inst = BIT
        inst.Mode = ZeroPage
    }else if opcode == 0x30 {
        inst.Inst = BMI
        inst.Mode = Relative
    }else if opcode == 0xd0 {
        inst.Inst = BNE
        inst.Mode = Relative
    }else if opcode == 0x10 {
        inst.Inst = BPL
        inst.Mode = Relative
    }else if opcode == 0x00 {
        inst.Inst = BRK
        inst.Mode = Implied
    }else if opcode == 0x50 {
        inst.Inst = BVC
        inst.Mode = Relative
    }else if opcode == 0x70 {
        inst.Inst = BVS
        inst.Mode = Relative
    }else if opcode == 0x18 {
        inst.Inst = CLC
        inst.Mode = Implied
    }else if opcode == 0xd8 {
        inst.Inst = CLD
        inst.Mode = Implied
    }else if opcode == 0x58 {
        inst.Inst = CLI
        inst.Mode = Implied
    }else if opcode == 0xb8 {
        inst.Inst = CLV
        inst.Mode = Implied
    }else if opcode == 0xcd {
        inst.Inst = CMP
        inst.Mode = Absolute
    }else if opcode == 0xdd {
        inst.Inst = CMP
        inst.Mode = AbsoluteX
    }else if opcode == 0xd9 {
        inst.Inst = CMP
        inst.Mode = AbsoluteY
    }else if opcode == 0xc9 {
        inst.Inst = CMP
        inst.Mode = Immediate
    }else if opcode == 0xc1 {
        inst.Inst = CMP
        inst.Mode = IndirectX
    }else if opcode == 0xd1 {
        inst.Inst = CMP
        inst.Mode = IndirectY
    }else if opcode == 0xc5 {
        inst.Inst = CMP
        inst.Mode = ZeroPage
    }else if opcode == 0xd5 {
        inst.Inst = CMP
        inst.Mode = ZeroPageX
    }else if opcode == 0xec {
        inst.Inst = CPX
        inst.Mode = Absolute
    }else if opcode == 0xe0 {
        inst.Inst = CPX
        inst.Mode = Immediate
    }else if opcode == 0xe4 {
        inst.Inst = CPX
        inst.Mode = ZeroPage
    }else if opcode == 0xcc {
        inst.Inst = CPY
        inst.Mode = Absolute
    }else if opcode == 0xc0 {
        inst.Inst = CPY
        inst.Mode = Immediate
    }else if opcode == 0xc4 {
        inst.Inst = CPY
        inst.Mode = ZeroPage
    }else if opcode == 0xce {
        inst.Inst = DEC
        inst.Mode = Absolute
    }else if opcode == 0xde {
        inst.Inst = DEC
        inst.Mode = AbsoluteX
    }else if opcode == 0xc6 {
        inst.Inst = DEC
        inst.Mode = ZeroPage
    }else if opcode == 0xd6 {
        inst.Inst = DEC
        inst.Mode = ZeroPageX
    }else if opcode == 0xca {
        inst.Inst = DEX
        inst.Mode = Implied
    }else if opcode == 0x88 {
        inst.Inst = DEY
        inst.Mode = Implied
    }else if opcode == 0x4d {
        inst.Inst = EOR
        inst.Mode = Absolute
    }else if opcode == 0x5d {
        inst.Inst = EOR
        inst.Mode = AbsoluteX
    }else if opcode == 0x59 {
        inst.Inst = EOR
        inst.Mode = AbsoluteY
    }else if opcode == 0x49 {
        inst.Inst = EOR
        inst.Mode = Immediate
    }else if opcode == 0x41 {
        inst.Inst = EOR
        inst.Mode = IndirectX
    }else if opcode == 0x51 {
        inst.Inst = EOR
        inst.Mode = IndirectY
    }else if opcode == 0x45 {
        inst.Inst = EOR
        inst.Mode = ZeroPage
    }else if opcode == 0x55 {
        inst.Inst = EOR
        inst.Mode = ZeroPageX
    }else if opcode == 0xee {
        inst.Inst = INC
        inst.Mode = Absolute
    }else if opcode == 0xfe {
        inst.Inst = INC
        inst.Mode = AbsoluteX
    }else if opcode == 0xe6 {
        inst.Inst = INC
        inst.Mode = ZeroPage
    }else if opcode == 0xf6 {
        inst.Inst = INC
        inst.Mode = ZeroPageX
    }else if opcode == 0xe8 {
        inst.Inst = INX
        inst.Mode = Implied
    }else if opcode == 0xc8 {
        inst.Inst = INY
        inst.Mode = Implied
    }else if opcode == 0x4c {
        inst.Inst = JMP
        inst.Mode = Absolute
    }else if opcode == 0x6c {
        inst.Inst = JMP
        inst.Mode = AbsoluteIndirect
    }else if opcode == 0x20 {
        inst.Inst = JSR
        inst.Mode = Absolute
    }else if opcode == 0xad {
        inst.Inst = LDA
        inst.Mode = Absolute
    }else if opcode == 0xbd {
        inst.Inst = LDA
        inst.Mode = AbsoluteX
    }else if opcode == 0xb9 {
        inst.Inst = LDA
        inst.Mode = AbsoluteY
    }else if opcode == 0xa9 {
        inst.Inst = LDA
        inst.Mode = Immediate
    }else if opcode == 0xa1 {
        inst.Inst = LDA
        inst.Mode = IndirectX
    }else if opcode == 0xb1 {
        inst.Inst = LDA
        inst.Mode = IndirectY
    }else if opcode == 0xa5 {
        inst.Inst = LDA
        inst.Mode = ZeroPage
    }else if opcode == 0xb5 {
        inst.Inst = LDA
        inst.Mode = ZeroPageX
    }else if opcode == 0xae {
        inst.Inst = LDX
        inst.Mode = Absolute
    }else if opcode == 0xbe {
        inst.Inst = LDX
        inst.Mode = AbsoluteY
    }else if opcode == 0xa2 {
        inst.Inst = LDX
        inst.Mode = Immediate
    }else if opcode == 0xa6 {
        inst.Inst = LDX
        inst.Mode = ZeroPage
    }else if opcode == 0xb6 {
        inst.Inst = LDX
        inst.Mode = ZeroPageY
    }else if opcode == 0xac {
        inst.Inst = LDY
        inst.Mode = Absolute
    }else if opcode == 0xbc {
        inst.Inst = LDY
        inst.Mode = AbsoluteX
    }else if opcode == 0xa0 {
        inst.Inst = LDY
        inst.Mode = Immediate
    }else if opcode == 0xa4 {
        inst.Inst = LDY
        inst.Mode = ZeroPage
    }else if opcode == 0xb4 {
        inst.Inst = LDY
        inst.Mode = ZeroPageX
    }else if opcode == 0x4e {
        inst.Inst = LSR
        inst.Mode = Absolute
    }else if opcode == 0x5e {
        inst.Inst = LSR
        inst.Mode = AbsoluteX
    }else if opcode == 0x4a {
        inst.Inst = LSR
        inst.Mode = Accumulator
    }else if opcode == 0x46 {
        inst.Inst = LSR
        inst.Mode = ZeroPage
    }else if opcode == 0x56 {
        inst.Inst = LSR
        inst.Mode = ZeroPageX
    }else if opcode == 0xea {
        inst.Inst = NOP
        inst.Mode = Implied
    }else if opcode == 0x0d {
        inst.Inst = ORA
        inst.Mode = Absolute
    }else if opcode == 0x1d {
        inst.Inst = ORA
        inst.Mode = AbsoluteX
    }else if opcode == 0x19 {
        inst.Inst = ORA
        inst.Mode = AbsoluteY
    }else if opcode == 0x09 {
        inst.Inst = ORA
        inst.Mode = Immediate
    }else if opcode == 0x01 {
        inst.Inst = ORA
        inst.Mode = IndirectX
    }else if opcode == 0x11 {
        inst.Inst = ORA
        inst.Mode = IndirectY
    }else if opcode == 0x05 {
        inst.Inst = ORA
        inst.Mode = ZeroPage
    }else if opcode == 0x15 {
        inst.Inst = ORA
        inst.Mode = ZeroPageX
    }else if opcode == 0x48 {
        inst.Inst = PHA
        inst.Mode = Implied
    }else if opcode == 0x08 {
        inst.Inst = PHP
        inst.Mode = Implied
    }else if opcode == 0x68 {
        inst.Inst = PLA
        inst.Mode = Implied
    }else if opcode == 0x28 {
        inst.Inst = PLP
        inst.Mode = Implied
    }else if opcode == 0x2e {
        inst.Inst = ROL
        inst.Mode = Absolute
    }else if opcode == 0x3e {
        inst.Inst = ROL
        inst.Mode = AbsoluteX
    }else if opcode == 0x2a {
        inst.Inst = ROL
        inst.Mode = Accumulator
    }else if opcode == 0x26 {
        inst.Inst = ROL
        inst.Mode = ZeroPage
    }else if opcode == 0x36 {
        inst.Inst = ROL
        inst.Mode = ZeroPageX
    }else if opcode == 0x6e {
        inst.Inst = ROR
        inst.Mode = Absolute
    }else if opcode == 0x7e {
        inst.Inst = ROR
        inst.Mode = AbsoluteX
    }else if opcode == 0x6a {
        inst.Inst = ROR
        inst.Mode = Accumulator
    }else if opcode == 0x66 {
        inst.Inst = ROR
        inst.Mode = ZeroPage
    }else if opcode == 0x76 {
        inst.Inst = ROR
        inst.Mode = ZeroPageX
    }else if opcode == 0x40 {
        inst.Inst = RTI
        inst.Mode = Implied
    }else if opcode == 0x60 {
        inst.Inst = RTS
        inst.Mode = Implied
    }else if opcode == 0xed {
        inst.Inst = SBC
        inst.Mode = Absolute
    }else if opcode == 0xfd {
        inst.Inst = SBC
        inst.Mode = AbsoluteX
    }else if opcode == 0xf9 {
        inst.Inst = SBC
        inst.Mode = AbsoluteY
    }else if opcode == 0xe9 {
        inst.Inst = SBC
        inst.Mode = Immediate
    }else if opcode == 0xe1 {
        inst.Inst = SBC
        inst.Mode = IndirectX
    }else if opcode == 0xf1 {
        inst.Inst = SBC
        inst.Mode = IndirectY
    }else if opcode == 0xe5 {
        inst.Inst = SBC
        inst.Mode = ZeroPage
    }else if opcode == 0xf5 {
        inst.Inst = SBC
        inst.Mode = ZeroPageX
    }else if opcode == 0x38 {
        inst.Inst = SEC
        inst.Mode = Implied
    }else if opcode == 0xf8 {
        inst.Inst = SED
        inst.Mode = Implied
    }else if opcode == 0x78 {
        inst.Inst = SEI
        inst.Mode = Implied
    }else if opcode == 0x8d {
        inst.Inst = STA
        inst.Mode = Absolute
    }else if opcode == 0x9d {
        inst.Inst = STA
        inst.Mode = AbsoluteX
    }else if opcode == 0x99 {
        inst.Inst = STA
        inst.Mode = AbsoluteY
    }else if opcode == 0x81 {
        inst.Inst = STA
        inst.Mode = IndirectX
    }else if opcode == 0x91 {
        inst.Inst = STA
        inst.Mode = IndirectY
    }else if opcode == 0x85 {
        inst.Inst = STA
        inst.Mode = ZeroPage
    }else if opcode == 0x95 {
        inst.Inst = STA
        inst.Mode = ZeroPageX
    }else if opcode == 0x8e {
        inst.Inst = STX
        inst.Mode = Absolute
    }else if opcode == 0x86 {
        inst.Inst = STX
        inst.Mode = ZeroPage
    }else if opcode == 0x96 {
        inst.Inst = STX
        inst.Mode = ZeroPageY
    }else if opcode == 0x8c {
        inst.Inst = STY
        inst.Mode = Absolute
    }else if opcode == 0x84 {
        inst.Inst = STY
        inst.Mode = ZeroPage
    }else if opcode == 0x94 {
        inst.Inst = STY
        inst.Mode = ZeroPageX
    }else if opcode == 0xaa {
        inst.Inst = TAX
        inst.Mode = Implied
    }else if opcode == 0xa8 {
        inst.Inst = TAY
        inst.Mode = Implied
    }else if opcode == 0xba {
        inst.Inst = TSX
        inst.Mode = Implied
    }else if opcode == 0x8a {
        inst.Inst = TXA
        inst.Mode = Implied
    }else if opcode == 0x9a {
        inst.Inst = TXS
        inst.Mode = Implied
    }else if opcode == 0x98 {
        inst.Inst = TYA
        inst.Mode = Implied
    }
    return inst 
}
