package nes

import (
    "log"
)

// ----------------------------------------
// Register
// ----------------------------------------

type CPURegister struct{
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

func (r *CPURegister)reset(bus *CPUBus){
    r.A = 0
    r.X = 0
    r.Y = 0
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

func NewRegister() *CPURegister{
    var register CPURegister
    return &register
}

func (register CPURegister)StatusToByte() byte {
    boolToByte := map[bool]byte{true:1,false:0}
    var status byte
    status |= boolToByte[register.P.N] << 7
    status |= boolToByte[register.P.V] << 6
    status |= boolToByte[register.P.R] << 5
    status |= boolToByte[register.P.B] << 4
    status |= boolToByte[register.P.D] << 3
    status |= boolToByte[register.P.I] << 2
    status |= boolToByte[register.P.Z] << 1
    status |= boolToByte[register.P.C]

    return status
}

func (register *CPURegister) ByteToStatus(status byte){
    byteToBool := map[byte]bool{1:true,0:false}
    register.P.N = byteToBool[(status >> 7) & 0x01]
    register.P.V = byteToBool[(status >> 6) & 0x01]
    register.P.R = byteToBool[(status >> 5) & 0x01]
    register.P.B = byteToBool[(status >> 4) & 0x01]
    register.P.D = byteToBool[(status >> 3) & 0x01]
    register.P.I = byteToBool[(status >> 2) & 0x01]
    register.P.Z = byteToBool[(status >> 1) & 0x01]
    register.P.C = byteToBool[status & 0x01]
}

func (r CPURegister) ShowRegister() {
    
    log.Printf("debug: A:0x%02x, X:0x%02x, Y:0x%02x, S:0x%02x, PC:0x%04x\n", r.A, r.X, r.Y, r.S, r.PC)
    log.Printf("debug: N:%v, V:%v, R:%v, B:%v, D:%v, I:%v, Z:%v, C:%v\n",r.P.N, r.P.V, r.P.R, r.P.B, r.P.D, r.P.I, r.P.Z, r.P.C)
}

// ----------------------------------------
// CPU
// ----------------------------------------

type CPU struct {
    Register *CPURegister
    Bus *CPUBus
}

func NewCPU(ram *RAM, rom *GameROM, ppu *PPU, keypad *KeyPad) *CPU{
    var cpu CPU
    cpu.Register = NewRegister()
    cpu.Bus = NewBus(ram, rom, ppu, keypad)

    cpu.reset()
    return &cpu
}

func (cpu *CPU) reset(){
    cpu.Register.reset(cpu.Bus)
}

func (cpu *CPU) interruptNMI(){
    
}

func (cpu *CPU) interruptIRQ(){
    
}

func (cpu *CPU) push(data byte){
    cpu.Bus.WriteByte(0x0100 | (uint16(cpu.Register.S & 0xff)), data)
    cpu.Register.S--
}

func (cpu *CPU) pop() byte {
    cpu.Register.S++
    return cpu.Bus.ReadByte(0x0100 | (uint16(cpu.Register.S & 0xff)))
}

func (cpu *CPU) fetch() byte {
    res := cpu.Bus.ReadByte(cpu.Register.PC)
    cpu.Register.PC++
    return res
}

func (cpu *CPU) run() int {
    opcode := cpu.fetch()
    log.Printf("debug: Fetch opcode:0x%02x, at 0x%04x", opcode,cpu.Register.PC-1)
    inst := cpu.Decode(opcode)
    return cpu.ExecInstruction(&inst)
}

func (cpu CPU) FetchAddress(inst *InstructionSet) uint16 {
    var address uint16
    
    switch inst.Mode {
    case Accumulator:
        // dummy
        address = 0x00

    case Implied:
        // dummy
        address = 0x00
        
    case Immediate:
        // dummy
        address = 0x00
        
    case ZeroPage:
        address = uint16(cpu.fetch() & 0xff)
        
    case ZeroPageX:
        address = uint16((cpu.fetch() + cpu.Register.X) & 0xff)

    case ZeroPageY:
        address = uint16((cpu.fetch() + cpu.Register.Y) & 0xff)
        
    case Absolute:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = (upper << 8) | lower

    case AbsoluteX:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = ((upper << 8) | lower) + uint16(cpu.Register.X)
        if address & 0xff00 != cpu.Register.PC & 0xff00 {
            inst.Cycle++
        }
        
    case AbsoluteY:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = ((upper << 8) | lower) + uint16(cpu.Register.Y)
        if address & 0xff00 != cpu.Register.PC & 0xff00 {
            inst.Cycle++
        }
        
    case Relative:
        address = uint16(int8(cpu.fetch())) + cpu.Register.PC
        if address & 0xff00 != cpu.Register.PC & 0xff00 {
            inst.Cycle++
        }
        
    case IndirectX:
        base := uint16((cpu.fetch() + cpu.Register.X) & 0xff)
        address = cpu.Bus.ReadWord(base)

    case IndirectY:
        base := uint16(cpu.fetch() & 0xff)
        address = cpu.Bus.ReadWord(base) + uint16(cpu.Register.Y)
        if address & 0xff00 != cpu.Register.PC & 0xff00 {
            inst.Cycle++
        }

    case AbsoluteIndirect:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = (upper << 8) | lower
    }
    log.Printf("debug: Address Fetch:0x%04x", address)
    return address
}

func (cpu *CPU)FetchOperand(inst *InstructionSet) byte {
    var operand byte
    if inst.Mode == Immediate{
        operand = cpu.fetch()
    }else{
        operand = cpu.Bus.ReadByte(cpu.FetchAddress(inst))
    }
    log.Printf("debug: Operand Fetch:0x%02x", operand)
    return operand
    
}

func (cpu *CPU) ExecInstruction(inst *InstructionSet) int {
    log.Printf("debug: Exec Instruction:%v, AddressingMode:%v", inst.Inst, inst.Mode)
    
    switch inst.Inst {
    case ADC:
        operand := cpu.FetchOperand(inst)
        operated := cpu.Register.A + operand + map[bool]byte{true:1,false:0}[cpu.Register.P.C]
        overflow := ((cpu.Register.A >> 7) & (operand >> 7) & ^(operated >> 7)) + (^(cpu.Register.A >> 7) & ^(operand >> 7) & (operated >> 7))   
        cpu.Register.P.N = (operated & 0x80 != 0)
        cpu.Register.P.Z = (operated == 0)
        cpu.Register.P.V = (overflow == 1)
        cpu.Register.P.C = (operated < cpu.Register.A || operated < operand) 
        cpu.Register.A = operated & 0xff

    case SBC:
        operand := cpu.FetchOperand(inst)
        operated := cpu.Register.A - operand - map[bool]byte{true:0,false:1}[cpu.Register.P.C]
        overflow := ((cpu.Register.A >> 7) & ^(operand >> 7) & ^(operated >> 7)) + (^(cpu.Register.A >> 7) & (operand >> 7) & (operated >> 7))
        cpu.Register.P.N = (operated & 0x80 != 0)
        cpu.Register.P.Z = (operated == 0)
        cpu.Register.P.V = (overflow == 1)
        cpu.Register.P.C = (operated <= 0x80) 
        cpu.Register.A = operated & 0xff

    case AND:
        operand := cpu.FetchOperand(inst)
        cpu.Register.A &= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case ORA:
        operand := cpu.FetchOperand(inst)
        cpu.Register.A |= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case EOR:
        operand := cpu.FetchOperand(inst)
        cpu.Register.A ^= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case ASL:
        var operand byte
        var address uint16
        if inst.Mode == Accumulator {
            operand = cpu.Register.A
        }else{
            address = cpu.FetchAddress(inst)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand << 1) & 0xfe
        cpu.Register.P.C = ((operand & 0x80) != 0)
        if inst.Mode == Accumulator {
            cpu.Register.A = operated
        }else{
            cpu.Bus.WriteByte(address, operated)
        }
        
    case LSR:
        var operand byte
        var address uint16
        if inst.Mode == Accumulator {
            operand = cpu.Register.A
        }else{
            address = cpu.FetchAddress(inst)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand >> 1) & 0x7f 
        cpu.Register.P.C = ((operand & 0x01) != 0)
        if inst.Mode == Accumulator {
            cpu.Register.A = operated
        }else{
            cpu.Bus.WriteByte(address, operated)
        }
        
    case ROL:
        var operand byte
        var address uint16
        if inst.Mode == Accumulator {
            operand = cpu.Register.A
        }else{
            address = cpu.FetchAddress(inst)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand << 1) & 0xfe | map[bool]byte{true:1,false:0}[cpu.Register.P.C]  
        cpu.Register.P.C = ((operand & 0x80) != 0)
        if inst.Mode == Accumulator {
            cpu.Register.A = operated
        }else{
            cpu.Bus.WriteByte(address, operated)
        }

    case ROR:
        var operand byte
        var address uint16
        if inst.Mode == Accumulator {
            operand = cpu.Register.A
        }else{
            address = cpu.FetchAddress(inst)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand >> 1) & 0x7f | (map[bool]byte{true:1,false:0}[cpu.Register.P.C] << 7)
        cpu.Register.P.C = ((operand & 0x01) != 0)
        if inst.Mode == Accumulator {
            cpu.Register.A = operated
        }else{
            cpu.Bus.WriteByte(address, operated)
        }

    case BCC:
        address := cpu.FetchAddress(inst)
        if !cpu.Register.P.C {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BCS:
        address := cpu.FetchAddress(inst)
        if cpu.Register.P.C {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BEQ:
        address := cpu.FetchAddress(inst)
        if cpu.Register.P.Z {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BNE:
        address := cpu.FetchAddress(inst)
        if !cpu.Register.P.Z {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BVC:
        address := cpu.FetchAddress(inst)
        if !cpu.Register.P.V {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BVS:
        address := cpu.FetchAddress(inst)
        if cpu.Register.P.V {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BPL:
        address := cpu.FetchAddress(inst)
        if !cpu.Register.P.N {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BMI:
        address := cpu.FetchAddress(inst)
        if cpu.Register.P.N {
            cpu.Register.PC = address
            inst.Cycle++
        }
    case BIT:
        operand := cpu.FetchOperand(inst)
        result := operand & cpu.Register.A
        if result == 0 {
            cpu.Register.P.Z = true
        }
        if operand & 0x80 != 0 {
            cpu.Register.P.N = true
        }else{
            cpu.Register.P.N = false
        }
        if operand & 0x40 != 0 {
            cpu.Register.P.V = true
        }else{
            cpu.Register.P.V = false
        }

    case JMP:
        address := cpu.FetchAddress(inst)
        cpu.Register.PC = address

    case JSR:
        cpu.push(byte((cpu.Register.PC >> 8) & 0xff))
        cpu.push(byte(cpu.Register.PC & 0xff))
        cpu.Register.PC = cpu.FetchAddress(inst)
        
    case RTS:
        lower := uint16(cpu.pop())
        upper := uint16(cpu.pop())
        cpu.Register.PC = (upper << 8) | lower
        cpu.Register.PC++

    case BRK:
        cpu.Register.PC++
        if !cpu.Register.P.I {
            cpu.Register.P.B = true
            cpu.push(byte((cpu.Register.PC >> 8) & 0xff))
            cpu.push(byte(cpu.Register.PC & 0xff))
            cpu.push(cpu.Register.StatusToByte())
            cpu.Register.P.I = true
            cpu.Register.PC = cpu.Bus.ReadWord(0xfffe)
        }
        
    case RTI:
        cpu.Register.ByteToStatus(cpu.pop())
        lower := uint16(cpu.pop())
        upper := uint16(cpu.pop())
        cpu.Register.PC = (upper << 8) | lower

    case CMP:
        operand := cpu.FetchOperand(inst)
        operated := cpu.Register.A - operand
        cpu.Register.P.C = ((operated & 0x80) == 0)
        cpu.Register.P.Z = operated == 0
        cpu.Register.P.N = ((operated & 0x80) != 0)
        
    case CPX:
        operand := cpu.FetchOperand(inst)
        operated := cpu.Register.X - operand
        cpu.Register.P.C = ((operated & 0x80) == 0)
        cpu.Register.P.Z = operated == 0
        cpu.Register.P.N = ((operated & 0x80) != 0)
        
    case CPY:
        operand := cpu.FetchOperand(inst)
        operated := cpu.Register.Y - operand
        cpu.Register.P.C = ((operated & 0x80) == 0)
        cpu.Register.P.Z = operated == 0
        cpu.Register.P.N = ((operated & 0x80) != 0)

    case INC:
        address := cpu.FetchAddress(inst)
        operand := cpu.Bus.ReadByte(address)
        operand++
        cpu.Register.P.Z = operand == 0
        cpu.Register.P.N = ((operand & 0x80) != 0)        
        cpu.Bus.WriteByte(address, operand)
        
    case DEC:
        address := cpu.FetchAddress(inst)
        operand := cpu.Bus.ReadByte(address)
        operand--
        cpu.Register.P.Z = operand == 0
        cpu.Register.P.N = ((operand & 0x80) != 0)        
        cpu.Bus.WriteByte(address, operand)

    case INX:
        cpu.Register.X++
        cpu.Register.P.Z = cpu.Register.X == 0
        cpu.Register.P.N = ((cpu.Register.X & 0x80) != 0)        

    case DEX:
        cpu.Register.X--
        cpu.Register.P.Z = cpu.Register.X == 0
        cpu.Register.P.N = ((cpu.Register.X & 0x80) != 0)        
        
    case INY:
        cpu.Register.Y++
        cpu.Register.P.Z = cpu.Register.Y == 0
        cpu.Register.P.N = ((cpu.Register.Y & 0x80) != 0)        

    case DEY:
        cpu.Register.Y--
        cpu.Register.P.Z = cpu.Register.Y == 0
        cpu.Register.P.N = ((cpu.Register.Y & 0x80) != 0)

    case CLC:
        cpu.Register.P.C = false

    case SEC:
        cpu.Register.P.C = true
        
    case CLI:
        cpu.Register.P.I = false

    case SEI:
        cpu.Register.P.I = true
        
    case CLD:
        cpu.Register.P.D = false

    case SED:
        cpu.Register.P.D = true

    case CLV:
        cpu.Register.P.V = false

    case LDA:
        operand := cpu.FetchOperand(inst)
        cpu.Register.A = operand
        cpu.Register.P.Z = operand == 0
        cpu.Register.P.N = ((operand & 0x80) != 0)

    case LDX:
        operand := cpu.FetchOperand(inst)
        cpu.Register.X = operand
        cpu.Register.P.Z = operand == 0
        cpu.Register.P.N = ((operand & 0x80) != 0)

    case LDY:
        operand := cpu.FetchOperand(inst)
        cpu.Register.Y = operand
        cpu.Register.P.Z = operand == 0
        cpu.Register.P.N = ((operand & 0x80) != 0)

    case STA:
        address := cpu.FetchAddress(inst)
        cpu.Bus.WriteByte(address, cpu.Register.A)
        
    case STX:
        address := cpu.FetchAddress(inst)
        cpu.Bus.WriteByte(address, cpu.Register.X)

    case STY:
        address := cpu.FetchAddress(inst)
        cpu.Bus.WriteByte(address, cpu.Register.Y)

    case TAX:
        cpu.Register.X = cpu.Register.A
        cpu.Register.P.Z = cpu.Register.X == 0
        cpu.Register.P.N = ((cpu.Register.X & 0x80) != 0)

    case TXA:
        cpu.Register.A = cpu.Register.X
        cpu.Register.P.Z = cpu.Register.A == 0
        cpu.Register.P.N = ((cpu.Register.A & 0x80) != 0)

    case TAY:
        cpu.Register.Y = cpu.Register.A
        cpu.Register.P.Z = cpu.Register.Y == 0
        cpu.Register.P.N = ((cpu.Register.Y & 0x80) != 0)

    case TYA:
        cpu.Register.A = cpu.Register.Y
        cpu.Register.P.Z = cpu.Register.A == 0
        cpu.Register.P.N = ((cpu.Register.A & 0x80) != 0)

    case TSX:
        cpu.Register.X = cpu.Register.S
        cpu.Register.P.Z = cpu.Register.X == 0
        cpu.Register.P.N = ((cpu.Register.X & 0x80) != 0)

    case TXS:
        cpu.Register.S = cpu.Register.X
        cpu.Register.P.Z = cpu.Register.S == 0
        cpu.Register.P.N = ((cpu.Register.S & 0x80) != 0)

    case PHA:
        cpu.push(cpu.Register.A)

    case PLA:
        cpu.Register.A = cpu.pop()
        cpu.Register.P.Z = cpu.Register.A == 0
        cpu.Register.P.N = ((cpu.Register.A & 0x80) != 0)

    case PHP:
        cpu.push(cpu.Register.StatusToByte())

    case PLP:
        cpu.Register.ByteToStatus(cpu.pop())

    }
    
    return inst.Cycle
    
}
