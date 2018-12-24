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

func (r Register)reset(){
    r.S = 0xFD
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
    register.reset()
    return &register
}

// ----------------------------------------
// CPU
// ----------------------------------------

type CPU struct {
    Register *Register
}

func NewCPU() *CPU{
    var cpu CPU
    cpu.Register = NewRegister()

    return &cpu
}

func (cpu CPU) reset(){
    cpu.Register.reset()
}
