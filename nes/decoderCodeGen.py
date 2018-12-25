ops = """
ADC,Absolute,6D
ADC,AbsoluteX,7D
ADC,AbsoluteY,79
ADC,Immediate,69
ADC,IndirectX,61
ADC,IndirectY,71
ADC,ZeroPage,65
ADC,ZeroPageX,75
AND,Absolute,2D
AND,AbsoluteX,3D
AND,AbsoluteY,39
AND,Immediate,29
AND,IndirectX,21
AND,IndirectY,31
AND,ZeroPage,25
AND,ZeroPageX,35
ASL,Absolute,0E
ASL,AbsoluteX,1E
ASL,Accumulator,0A
ASL,ZeroPage,6
ASL,ZeroPageX,16
BCC,Relative,90
BCS,Relative,B0
BEQ,Relative,F0
BIT,Absolute,2C
BIT,ZeroPage,24
BMI,Relative,30
BNE,Relative,D0
BPL,Relative,10
BRK,Implied,0
BVC,Relative,50
BVS,Relative,70
CLC,Implied,18
CLD,Implied,D8
CLI,Implied,58
CLV,Implied,B8
CMP,Absolute,CD
CMP,AbsoluteX,DD
CMP,AbsoluteY,D9
CMP,Immediate,C9
CMP,IndirectX,C1
CMP,IndirectY,D1
CMP,ZeroPage,C5
CMP,ZeroPageX,D5
CPX,Absolute,EC
CPX,Immediate,E0
CPX,ZeroPage,E4
CPY,Absolute,CC
CPY,Immediate,C0
CPY,ZeroPage,C4
DEC,Absolute,CE
DEC,AbsoluteX,DE
DEC,ZeroPage,C6
DEC,ZeroPageX,D6
DEX,Implied,CA
DEY,Implied,88
EOR,Absolute,4D
EOR,AbsoluteX,5D
EOR,AbsoluteY,59
EOR,Immediate,49
EOR,IndirectX,41
EOR,IndirectY,51
EOR,ZeroPage,45
EOR,ZeroPageX,55
INC,Absolute,EE
INC,AbsoluteX,FE
INC,ZeroPage,E6
INC,ZeroPageX,F6
INX,Implied,E8
INY,Implied,C8
JMP,Absolute,4C
JMP,Indirect,6C
JSR,Absolute,20
LDA,Absolute,AD
LDA,AbsoluteX,BD
LDA,AbsoluteY,B9
LDA,Immediate,A9
LDA,IndirectX,A1
LDA,IndirectY,B1
LDA,ZeroPage,A5
LDA,ZeroPageX,B5
LDX,Absolute,AE
LDX,AbsoluteY,BE
LDX,Immediate,A2
LDX,ZeroPage,A6
LDX,ZeroPageY,B6
LDY,Absolute,AC
LDY,AbsoluteX,BC
LDY,Immediate,A0
LDY,ZeroPage,A4
LDY,ZeroPageX,B4
LSR,Absolute,4E
LSR,AbsoluteX,5E
LSR,Accumulator,4A
LSR,ZeroPage,46
LSR,ZeroPageX,56
NOP,Implied,EA
ORA,Absolute,0D
ORA,AbsoluteX,1D
ORA,AbsoluteY,19
ORA,Immediate,9
ORA,IndirectX,1
ORA,IndirectY,11
ORA,ZeroPage,5
ORA,ZeroPageX,15
PHA,Implied,48
PHP,Implied,8
PLA,Implied,68
PLP,Implied,28
ROL,Absolute,2E
ROL,AbsoluteX,3E
ROL,Accumulator,2A
ROL,ZeroPage,26
ROL,ZeroPageX,36
ROR,Absolute,6E
ROR,AbsoluteX,7E
ROR,Accumulator,6A
ROR,ZeroPage,66
ROR,ZeroPageX,76
RTI,Implied,40
RTS,Implied,60
SBC,Absolute,ED
SBC,AbsoluteX,FD
SBC,AbsoluteY,F9
SBC,Immediate,E9
SBC,IndirectX,E1
SBC,IndirectY,F1
SBC,ZeroPage,E5
SBC,ZeroPageX,F5
SEC,Implied,38
SED,Implied,F8
SEI,Implied,78
STA,Absolute,8D
STA,AbsoluteX,9D
STA,AbsoluteY,99
STA,IndirectX,81
STA,IndirectY,91
STA,ZeroPage,85
STA,ZeroPageX,95
STX,Absolute,8E
STX,ZeroPage,86
STX,ZeroPageY,96
STY,Absolute,8C
STY,ZeroPage,84
STY,ZeroPageX,94
TAX,Implied,AA
TAY,Implied,A8
TSX,Implied,BA
TXA,Implied,8A
TXS,Implied,9A
TYA,Implied,98
"""

template = """
else if opcode == {0} {{
    inst.Inst = {1}
    inst.Mode = {2}
}}
"""

for op in ops.strip().split("\n"):
    inst,mode,code = op.split(",")
    print(template.strip().format("0x{:02x}".format(int(code,16)), inst, mode),end="")
