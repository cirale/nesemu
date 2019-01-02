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

func (a AddressingMode) String() string {
    switch a {
    case Accumulator:
        return "Accumulator"
    case Immediate:
        return "Immediate"
    case Absolute:
        return "Absolute"
    case ZeroPage:
        return "ZeroPage"
    case ZeroPageX:
        return "ZeroPageX"
    case ZeroPageY:
        return "ZeroPageY"
    case AbsoluteX:
        return "AbsoluteX"
    case AbsoluteY:
        return "AbsoluteY"
    case Implied:
        return "Implied"
    case Relative:
        return "Relative"
    case IndirectX:
        return "IndirectX"
    case IndirectY:
        return "IndirectY"
    case AbsoluteIndirect:
        return "AbsoluteIndirect"
    default:
        return ""
    }
}

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

func (i Instruction) String() string {
    switch i {
    case ADC:
        return "ADC"
    case SBC:
        return "SBC"
    case AND:
        return "AND"
    case ORA:
        return "ORA"
    case EOR:
        return "EOR"
    case ASL:
        return "ASL"
    case LSR:
        return "LSR"
    case ROL:
        return "ROL"
    case ROR:
        return "ROR"
    case BCC:
        return "BCC"
    case BCS:
        return "BCS"
    case BEQ:
        return "BEQ"
    case BNE:
        return "BNE"
    case BVC:
        return "BVC"
    case BVS:
        return "BVS"
    case BPL:
        return "BPL"
    case BMI:
        return "BMI"
    case BIT:
        return "BIT"
    case JMP:
        return "JMP"
    case JSR:
        return "JSR"
    case RTS:
        return "RTS"
    case BRK:
        return "BRK"
    case RTI:
        return "RTI"
    case CMP:
        return "CMP"
    case CPX:
        return "CPX"
    case CPY:
        return "CPY"
    case INC:
        return "INC"
    case DEC:
        return "DEC"
    case INX:
        return "INX"
    case DEX:
        return "DEX"
    case INY:
        return "INY"
    case DEY:
        return "DEY"
    case CLC:
        return "CLC"
    case SEC:
        return "SEC"
    case CLI:
        return "CLI"
    case SEI:
        return "SEI"
    case CLD:
        return "CLD"
    case SED:
        return "SED"
    case CLV:
        return "CLV"
    case LDA:
        return "LDA"
    case LDX:
        return "LDX"
    case LDY:
        return "LDY"
    case STA:
        return "STA"
    case STX:
        return "STX"
    case STY:
        return "STY"
    case TAX:
        return "TAX"
    case TXA:
        return "TXA"
    case TAY:
        return "TAY"
    case TYA:
        return "TYA"
    case TSX:
        return "TSX"
    case TXS:
        return "TXS"
    case PHA:
        return "PHA"
    case PLA:
        return "PLA"
    case PHP:
        return "PHP"
    case PLP:
        return "PLP"
    case NOP:
        return "NOP"
    default:
        return ""
    }
}


type InstructionSet struct{
    Inst Instruction
    Mode AddressingMode
    Cycle int
}

func (cpu CPU)Decode(opcode byte) InstructionSet{
    var inst InstructionSet
    switch opcode {
    case 0x69:
        inst.Inst = ADC
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0x65:
        inst.Inst = ADC
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x75:
        inst.Inst = ADC
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x6d:
        inst.Inst = ADC
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x7d:
        inst.Inst = ADC
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x79:
        inst.Inst = ADC
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0x61:
        inst.Inst = ADC
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0x71:
        inst.Inst = ADC
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0xe9:
        inst.Inst = SBC
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xe5:
        inst.Inst = SBC
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xf5:
        inst.Inst = SBC
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0xed:
        inst.Inst = SBC
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xfd:
        inst.Inst = SBC
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0xf9:
        inst.Inst = SBC
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0xe1:
        inst.Inst = SBC
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0xf1:
        inst.Inst = SBC
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0x29:
        inst.Inst = AND
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0x25:
        inst.Inst = AND
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x35:
        inst.Inst = AND
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x2d:
        inst.Inst = AND
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x3d:
        inst.Inst = AND
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x39:
        inst.Inst = AND
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0x21:
        inst.Inst = AND
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0x31:
        inst.Inst = AND
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0x09:
        inst.Inst = ORA
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0x05:
        inst.Inst = ORA
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x15:
        inst.Inst = ORA
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x0d:
        inst.Inst = ORA
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x1d:
        inst.Inst = ORA
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x19:
        inst.Inst = ORA
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0x01:
        inst.Inst = ORA
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0x11:
        inst.Inst = ORA
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0x49:
        inst.Inst = EOR
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0x45:
        inst.Inst = EOR
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x55:
        inst.Inst = EOR
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x4d:
        inst.Inst = EOR
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x5d:
        inst.Inst = EOR
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x59:
        inst.Inst = EOR
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0x41:
        inst.Inst = EOR
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0x51:
        inst.Inst = EOR
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0x0a:
        inst.Inst = ASL
        inst.Mode = Accumulator
        inst.Cycle = 2
    case 0x06:
        inst.Inst = ASL
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0x16:
        inst.Inst = ASL
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0x0e:
        inst.Inst = ASL
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0x1e:
        inst.Inst = ASL
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0x4a:
        inst.Inst = LSR
        inst.Mode = Accumulator
        inst.Cycle = 2
    case 0x46:
        inst.Inst = LSR
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0x56:
        inst.Inst = LSR
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0x4e:
        inst.Inst = LSR
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0x5e:
        inst.Inst = LSR
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0x2a:
        inst.Inst = ROL
        inst.Mode = Accumulator
        inst.Cycle = 2
    case 0x26:
        inst.Inst = ROL
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0x36:
        inst.Inst = ROL
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0x2e:
        inst.Inst = ROL
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0x3e:
        inst.Inst = ROL
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0x6a:
        inst.Inst = ROR
        inst.Mode = Accumulator
        inst.Cycle = 2
    case 0x66:
        inst.Inst = ROR
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0x76:
        inst.Inst = ROR
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0x6e:
        inst.Inst = ROR
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0x7e:
        inst.Inst = ROR
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0x90:
        inst.Inst = BCC
        inst.Mode = Relative
        inst.Cycle = 2
    case 0xb0:
        inst.Inst = BCS
        inst.Mode = Relative
        inst.Cycle = 2
    case 0xf0:
        inst.Inst = BEQ
        inst.Mode = Relative
        inst.Cycle = 2
    case 0xd0:
        inst.Inst = BNE
        inst.Mode = Relative
        inst.Cycle = 2
    case 0x50:
        inst.Inst = BVC
        inst.Mode = Relative
        inst.Cycle = 2
    case 0x70:
        inst.Inst = BVS
        inst.Mode = Relative
        inst.Cycle = 2
    case 0x10:
        inst.Inst = BPL
        inst.Mode = Relative
        inst.Cycle = 2
    case 0x30:
        inst.Inst = BMI
        inst.Mode = Relative
        inst.Cycle = 2
    case 0x24:
        inst.Inst = BIT
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x2c:
        inst.Inst = BIT
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x4c:
        inst.Inst = JMP
        inst.Mode = Absolute
        inst.Cycle = 3
    case 0x6c:
        inst.Inst = JMP
        inst.Mode = AbsoluteIndirect
        inst.Cycle = 5
    case 0x20:
        inst.Inst = JSR
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0x60:
        inst.Inst = RTS
        inst.Mode = Implied
        inst.Cycle = 6
    case 0x00:
        inst.Inst = BRK
        inst.Mode = Implied
        inst.Cycle = 7
    case 0x40:
        inst.Inst = RTI
        inst.Mode = Implied
        inst.Cycle = 6
    case 0xc9:
        inst.Inst = CMP
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xc5:
        inst.Inst = CMP
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xd5:
        inst.Inst = CMP
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0xcd:
        inst.Inst = CMP
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xdd:
        inst.Inst = CMP
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0xd9:
        inst.Inst = CMP
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0xc1:
        inst.Inst = CMP
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0xd1:
        inst.Inst = CMP
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0xe0:
        inst.Inst = CPX
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xe4:
        inst.Inst = CPX
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xec:
        inst.Inst = CPX
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xc0:
        inst.Inst = CPY
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xc4:
        inst.Inst = CPY
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xcc:
        inst.Inst = CPY
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xe6:
        inst.Inst = INC
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0xf6:
        inst.Inst = INC
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0xee:
        inst.Inst = INC
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0xfe:
        inst.Inst = INC
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0xc6:
        inst.Inst = DEC
        inst.Mode = ZeroPage
        inst.Cycle = 5
    case 0xd6:
        inst.Inst = DEC
        inst.Mode = ZeroPageX
        inst.Cycle = 6
    case 0xce:
        inst.Inst = DEC
        inst.Mode = Absolute
        inst.Cycle = 6
    case 0xde:
        inst.Inst = DEC
        inst.Mode = AbsoluteX
        inst.Cycle = 6
    case 0xe8:
        inst.Inst = INX
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xca:
        inst.Inst = DEX
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xc8:
        inst.Inst = INY
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x88:
        inst.Inst = DEY
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x18:
        inst.Inst = CLC
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x38:
        inst.Inst = SEC
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x58:
        inst.Inst = CLI
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x78:
        inst.Inst = SEI
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xd8:
        inst.Inst = CLD
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xf8:
        inst.Inst = SED
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xb8:
        inst.Inst = CLV
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xa9:
        inst.Inst = LDA
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xa5:
        inst.Inst = LDA
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xb5:
        inst.Inst = LDA
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0xad:
        inst.Inst = LDA
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xbd:
        inst.Inst = LDA
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0xb9:
        inst.Inst = LDA
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0xa1:
        inst.Inst = LDA
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0xb1:
        inst.Inst = LDA
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0xa2:
        inst.Inst = LDX
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xa6:
        inst.Inst = LDX
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xb6:
        inst.Inst = LDX
        inst.Mode = ZeroPageY
        inst.Cycle = 4
    case 0xae:
        inst.Inst = LDX
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xbe:
        inst.Inst = LDX
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0xa0:
        inst.Inst = LDY
        inst.Mode = Immediate
        inst.Cycle = 2
    case 0xa4:
        inst.Inst = LDY
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0xb4:
        inst.Inst = LDY
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0xac:
        inst.Inst = LDY
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xbc:
        inst.Inst = LDY
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x85:
        inst.Inst = STA
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x95:
        inst.Inst = STA
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x8d:
        inst.Inst = STA
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x9d:
        inst.Inst = STA
        inst.Mode = AbsoluteX
        inst.Cycle = 4
    case 0x99:
        inst.Inst = STA
        inst.Mode = AbsoluteY
        inst.Cycle = 4
    case 0x81:
        inst.Inst = STA
        inst.Mode = IndirectX
        inst.Cycle = 6
    case 0x91:
        inst.Inst = STA
        inst.Mode = IndirectY
        inst.Cycle = 5
    case 0x86:
        inst.Inst = STX
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x96:
        inst.Inst = STX
        inst.Mode = ZeroPageY
        inst.Cycle = 4
    case 0x8e:
        inst.Inst = STX
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0x84:
        inst.Inst = STY
        inst.Mode = ZeroPage
        inst.Cycle = 3
    case 0x94:
        inst.Inst = STY
        inst.Mode = ZeroPageX
        inst.Cycle = 4
    case 0x8c:
        inst.Inst = STY
        inst.Mode = Absolute
        inst.Cycle = 4
    case 0xaa:
        inst.Inst = TAX
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x8a:
        inst.Inst = TXA
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xa8:
        inst.Inst = TAY
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x98:
        inst.Inst = TYA
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x9a:
        inst.Inst = TXS
        inst.Mode = Implied
        inst.Cycle = 2
    case 0xba:
        inst.Inst = TSX
        inst.Mode = Implied
        inst.Cycle = 2
    case 0x48:
        inst.Inst = PHA
        inst.Mode = Implied
        inst.Cycle = 3
    case 0x68:
        inst.Inst = PLA
        inst.Mode = Implied
        inst.Cycle = 4
    case 0x08:
        inst.Inst = PHP
        inst.Mode = Implied
        inst.Cycle = 3
    case 0x28:
        inst.Inst = PLP
        inst.Mode = Implied
        inst.Cycle = 4
    case 0xea:
        inst.Inst = NOP
        inst.Mode = Implied
        inst.Cycle = 2        
    }
    return inst 
}
