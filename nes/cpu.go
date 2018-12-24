package nes

// ----------------------------------------
// Register
// ----------------------------------------

type Register struct{
    A byte
    X byte
    Y byte
    S byte
    P struct {
        N bool
        V bool
        R bool
        B bool
        D bool
        I bool
        Z bool
        C bool
    }
    PC uint16
}

func (r Register)reset(bus *Bus){
    r.S = 0xFD
    r.PC = bus.ReadWord(0xfffc)
    r.P.N = false
    r.P.V = false
    r.P.R = true
    r.P.B = false
    r.P.D = false
    r.P.I = false
    r.P.Z = true
    r.P.C = false
}

func NewRegister() *Register{
    var register Register

    return &register
}

// ----------------------------------------
// CPU
// ----------------------------------------

type CPU struct {
    Register *Register
    Bus *Bus
}

func NewCPU(ram *RAM, rom *GameROM) *CPU{
    var cpu CPU
    cpu.Register = NewRegister()
    cpu.Bus = NewBus(ram, rom)

    cpu.reset()
    return &cpu
}

func (cpu CPU) reset(){
    cpu.Register.reset(cpu.Bus)
}

func (cpu CPU) push(data byte){
    cpu.Bus.WriteByte(0x0100 | (uint16(cpu.Register.S & 0xff)), data)
    cpu.Register.S--
}

func (cpu CPU) pop() byte {
    cpu.Register.S++
    return cpu.Bus.ReadByte(0x0100 | (uint16(cpu.Register.S & 0xff)))
}

// ----------------------------------------
// InstructionDecoder
// ----------------------------------------

type AddressingMode int
const (
    Accumulator AddressingMode = iota
    Immediate
    Absolute
    ZeroPage
    IndexedZeroPageX
    IndexedZeroPageY
    IndexedAbsoluteX
    IndexedAbsoluteY
    Implied
    Relative
    IndeirectX
    IndeirectY
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
    CLC; SEC; CLI; SEI; CKD; SED; CLV
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
    if opcode == 0x00 {
        inst.Inst = BRK
        inst.Mode = Implied
    }else if opcode == 0x01 {
        inst.Inst = ORA
        inst.Mode = IndeirectX
    }else if opcode == 0x05 {
        inst.Inst = ORA
        inst.Mode = ZeroPage
    }else if opcode == 0x06 {
        inst.Inst = ASL
        inst.Mode = ZeroPage
    }else if opcode == 0x08 {
        inst.Inst = PHP
        inst.Mode = Implied
    }
    
    return inst 
}
