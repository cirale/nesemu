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

func (r *Register)reset(bus *Bus){
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

func (cpu *CPU) reset(){
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

func (cpu CPU) fetch() byte {
    res := cpu.Bus.ReadByte(cpu.Register.PC)
    cpu.Register.PC++
    return res
}

func (cpu CPU) FetchAddress(mode AddressingMode) uint16 {
    var address uint16
    
    switch mode{
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
        
    case AbsoluteY:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = ((upper << 8) | lower) + uint16(cpu.Register.Y)
        
    case Relative:
        address = uint16(cpu.fetch()) + cpu.Register.PC
        
    case IndirectX:
        base := uint16((cpu.fetch() + cpu.Register.X) & 0xff)
        address = cpu.Bus.ReadWord(base)

    case IndirectY:
        base := uint16(cpu.fetch() & 0xff)
        address = cpu.Bus.ReadWord(base) + uint16(cpu.Register.Y)

    case AbsoluteIndirect:
        lower := uint16(cpu.fetch())
        upper := uint16(cpu.fetch())
        address = (upper << 8) | lower
    }
    
    return address
}

func (cpu *CPU)FetchOperand(mode AddressingMode) byte {
    var operand byte
    if mode == Immediate{
        operand = cpu.fetch()
    }else{
        operand = cpu.Bus.ReadByte(cpu.FetchAddress(mode))
    }
    return operand
    
}

func (cpu *CPU) ExecInstruction(inst InstructionSet){
    switch inst.Inst {
    case ADC:
        operand := cpu.FetchOperand(inst.Mode)
        operated := cpu.Register.A + operand + map[bool]byte{true:1,false:0}[cpu.Register.P.C]
        overflow := ((cpu.Register.A >> 7) & (operand >> 7) & ^(operated >> 7)) + (^(cpu.Register.A >> 7) & ^(operand >> 7) & (operated >> 7))   
        cpu.Register.P.N = (operated & 0x80 != 0)
        cpu.Register.P.Z = (operated == 0)
        cpu.Register.P.V = (overflow == 1)
        cpu.Register.P.C = (operated < cpu.Register.A || operated < operand) 
        cpu.Register.A = operated & 0xff

    case SBC:
        operand := cpu.FetchOperand(inst.Mode)
        operated := cpu.Register.A - operand - map[bool]byte{true:0,false:1}[cpu.Register.P.C]
        overflow := ((cpu.Register.A >> 7) & ^(operand >> 7) & ^(operated >> 7)) + (^(cpu.Register.A >> 7) & (operand >> 7) & (operated >> 7))
        cpu.Register.P.N = (operated & 0x80 != 0)
        cpu.Register.P.Z = (operated == 0)
        cpu.Register.P.V = (overflow == 1)
        cpu.Register.P.C = (operated <= 0x80) 
        cpu.Register.A = operated & 0xff

    case AND:
        operand := cpu.FetchOperand(inst.Mode)
        cpu.Register.A &= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case ORA:
        operand := cpu.FetchOperand(inst.Mode)
        cpu.Register.A |= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case EOR:
        operand := cpu.FetchOperand(inst.Mode)
        cpu.Register.A ^= operand
        cpu.Register.P.N = (cpu.Register.A & 0x80 != 0)
        cpu.Register.P.Z = (cpu.Register.A == 0)

    case ASL:
        var operand byte
        var address uint16
        if inst.Mode == Accumulator {
            operand = cpu.Register.A
        }else{
            address = cpu.FetchAddress(inst.Mode)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand << 1) & 0xfe
        cpu.Register.P.C = (operand & 0x80) != 0
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
            address = cpu.FetchAddress(inst.Mode)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand >> 1) & 0x7f 
        cpu.Register.P.C = (operand & 0x01) != 0
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
            address = cpu.FetchAddress(inst.Mode)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand << 1) & 0xfe | map[bool]byte{true:1,false:0}[cpu.Register.P.C]  
        cpu.Register.P.C = (operand & 0x01) != 0
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
            address = cpu.FetchAddress(inst.Mode)
            operand = cpu.Bus.ReadByte(address)
        }
        operated := (operand >> 1) & 0x7f | (map[bool]byte{true:1,false:0}[cpu.Register.P.C] << 7)
        cpu.Register.P.C = (operand & 0x01) != 0
        if inst.Mode == Accumulator {
            cpu.Register.A = operated
        }else{
            cpu.Bus.WriteByte(address, operated)
        }
    }

    
}
