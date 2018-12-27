package nes

import (
    "testing"
    "io/ioutil"
)

func TestFetchAddress(t *testing.T){
    bytes, err := ioutil.ReadFile("../helloworld.nes")
    if err != nil {
        t.Fatal(err)
    }
    n := NewNES(bytes)
    var mode AddressingMode
    var address uint16

    //zeropage test
    t.Log("ZeroPage test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0xff)
    mode = ZeroPage
    address = n.CPU.FetchAddress(mode)
    if address != 0x00ff {
        t.Errorf("Zeropage addressing is worong")
    }
    t.Log("ZeroPage test finish")

    //zeropageX test
    t.Log("ZeroPageX test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.X = 0x20
    mode = ZeroPageX
    address = n.CPU.FetchAddress(mode)
    if address != 0x0070 {
        t.Errorf("ZeropageX addressing is worong:%x",address)
    }
    t.Log("ZeroPageX test finish")
    
    //zeropageY test
    t.Log("ZeroPageY test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.Y = 0x20
    mode = ZeroPageY
    address = n.CPU.FetchAddress(mode)
    if address != 0x0070 {
        t.Errorf("ZeropageY addressing is worong:%x",address)
    }
    t.Log("ZeroPageY test finish")

    //absolute test
    t.Log("Absolute test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Bus.WriteByte(0x0001, 0xf0)
    mode = Absolute
    address = n.CPU.FetchAddress(mode)
    if address != 0xf050 {
        t.Errorf("absolute addressing is worong:%x",address)
    }
    t.Log("absolute test finish")

    //absoluteX test
    t.Log("AbsoluteX test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Bus.WriteByte(0x0001, 0xf0)
    n.CPU.Register.X = 0x40
    mode = AbsoluteX
    address = n.CPU.FetchAddress(mode)
    if address != 0xf090 {
        t.Errorf("absoluteX addressing is worong:%x",address)
    }
    t.Log("absoluteX test finish")

    //absoluteY test
    t.Log("AbsoluteY test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Bus.WriteByte(0x0001, 0xf0)
    n.CPU.Register.Y = 0x40
    mode = AbsoluteY
    address = n.CPU.FetchAddress(mode)
    if address != 0xf090 {
        t.Errorf("absoluteY addressing is worong:%x",address)
    }
    t.Log("absoluteY test finish")

    //Relative test
    t.Log("Relative test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    mode = Relative
    address = n.CPU.FetchAddress(mode)
    if address != 0x0051 {
        t.Errorf("Relative addressing is worong:%x",address)
    }
    t.Log("relative test finish")

    //IndirectX test
    t.Log("IndirectX test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.X = 0x50
    n.CPU.Bus.WriteByte(0x00a0, 0xaa)
    n.CPU.Bus.WriteByte(0x00a1, 0xff)
    mode = IndirectX
    address = n.CPU.FetchAddress(mode)
    if address != 0xffaa {
        t.Errorf("IndirectX addressing is worong:%x",address)
    }
    t.Log("IndirectX test finish")

    //IndirectY test
    t.Log("IndirectY test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.Y = 0x50
    n.CPU.Bus.WriteByte(0x0050, 0xaa)
    n.CPU.Bus.WriteByte(0x0051, 0xff)
    mode = IndirectY
    address = n.CPU.FetchAddress(mode)
    if address != 0xfffa {
        t.Errorf("IndirectY addressing is worong:%x",address)
    }
    t.Log("IndirectY test finish")

    //AbsoluteIndirect test
    t.Log("AbsoluteIndirect test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Bus.WriteByte(0x0001, 0x01)
    mode = AbsoluteIndirect
    address = n.CPU.FetchAddress(mode)
    if address != 0x0150 {
        t.Errorf("AbsoluteIndirect addressing is worong:%x",address)
    }
    t.Log("AbsoluteIndirect test finish")
}

func TestExecInstruction(t *testing.T){
    bytes, err := ioutil.ReadFile("../helloworld.nes")
    if err != nil {
        t.Fatal(err)
    }
    n := NewNES(bytes)

    //ADC Test
    t.Log("ADC test")
    var i InstructionSet
    i.Inst = ADC
    i.Mode = Immediate

    // overflow test,negative test
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000,0x7f)
    n.CPU.Register.A = 0
    n.CPU.Register.P.C = true

    n.CPU.ExecInstruction(i)
    t.Log("overflow test, Negative Test start.")
    if n.CPU.Register.A != 128 {
        t.Errorf("result is wrong:%d",n.CPU.Register.A)
    }
    if !n.CPU.Register.P.N {
        t.Errorf("N is wrong")
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong")
    }
    if !n.CPU.Register.P.V {
        t.Errorf("V is wrong")
    }
    if n.CPU.Register.P.C {
        t.Errorf("C is wrong")
    }
    t.Log("overflow test, Negative Test finished.")

    // carry, zero test
    t.Log("carry test start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000,0xff)
    n.CPU.Register.A = 0x1
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x0 {
        t.Errorf("result is wrong:%d, desire:%d",n.CPU.Register.A,0x0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    if n.CPU.Register.P.V {
        t.Errorf("V is wrong:%v",n.CPU.Register.P.V)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }
    t.Log("carry test start")

    // carry test 2
    t.Log("carry test 2 start")
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000,0xff)
    n.CPU.Register.P.C = true
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x0 {
        t.Errorf("result is wrong:%d, desire:%d",n.CPU.Register.A,0x0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    if n.CPU.Register.P.V {
        t.Errorf("V is wrong:%v",n.CPU.Register.P.V)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }
    t.Log("carry test 2 finished")

    //SBC test
    t.Log("SBC test")
    i.Inst = SBC
    i.Mode = Immediate

    t.Log("overflow carry test start")
    n.Reset()
    n.CPU.Register.A = 0x80
    n.CPU.Bus.WriteByte(0x0000,0x7f)
    n.CPU.Register.P.C = true
    
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x1 {
        t.Errorf("result is wrong:%d, desire:%d",n.CPU.Register.A,0x1)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    if !n.CPU.Register.P.V {
        t.Errorf("V is wrong:%v",n.CPU.Register.P.V)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }
    t.Log("overflow  carry test finished")

    // negative test
    t.Log("negative test start")
    n.Reset()
    n.CPU.Register.A = 0xff
    n.CPU.Bus.WriteByte(0x0000,0x0)    
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0xfe {
        t.Errorf("result is wrong:%d, desire:%d",n.CPU.Register.A,0xfe)
    }
    if !n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    if n.CPU.Register.P.V {
        t.Errorf("V is wrong:%v",n.CPU.Register.P.V)
    }
    if n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }
    t.Log("negative test finished")

    // zero test
    t.Log("zero test start")
    n.Reset()
    n.CPU.Register.A = 0x01
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x0 {
        t.Errorf("result is wrong:%d, desire:%d",n.CPU.Register.A,0x0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    if n.CPU.Register.P.V {
        t.Errorf("V is wrong:%v",n.CPU.Register.P.V)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }
    t.Log("zero test finished")

    // AND test
    i.Inst = AND
    i.Mode = Immediate
    t.Log("AND test")

    // negative test
    t.Log("negative test start")
    n.Reset()
    n.CPU.Register.A = 0xaa
    n.CPU.Bus.WriteByte(0x0000, 0xa0)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0xa0 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0xa0)
    }
    if !n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("negative test finished")

    // zero test
    t.Log("zero test start")
    n.Reset()
    n.CPU.Register.A = 0xaa
    n.CPU.Bus.WriteByte(0x0000, 0x55)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("zero test finished")

    //ORA test
    i.Inst = ORA
    i.Mode = Immediate
    t.Log("ORA test")

    // negative test
    t.Log("negative test start")
    n.Reset()
    n.CPU.Register.A = 0xaa
    n.CPU.Bus.WriteByte(0x0000, 0x55)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0xff {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0xff)
    }
    if !n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("negative test finished")

    // zero test
    t.Log("zero test start")
    n.Reset()
    n.CPU.Register.A = 0x00
    n.CPU.Bus.WriteByte(0x0000, 0x00)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("zero test finished")

    //EOR test
    i.Inst = EOR
    i.Mode = Immediate
    t.Log("EOR test")

    // negative test
    t.Log("negative test start")
    n.Reset()
    n.CPU.Register.A = 0x55
    n.CPU.Bus.WriteByte(0x0000, 0xff)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0xaa {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0xaa)
    }
    if !n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("negative test finished")

    // zero test
    t.Log("zero test start")
    n.Reset()
    n.CPU.Register.A = 0xff
    n.CPU.Bus.WriteByte(0x0000, 0xff)
    n.CPU.ExecInstruction(i)
    
    if n.CPU.Register.A != 0 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0)
    }
    if n.CPU.Register.P.N {
        t.Errorf("N is wrong:%v",n.CPU.Register.P.N)
    }
    if !n.CPU.Register.P.Z {
        t.Errorf("Z is wrong:%v",n.CPU.Register.P.Z)
    }
    t.Log("zero test finished")

    //ASL test
    i.Inst = ASL
    i.Mode = Accumulator
    t.Log("ASL test")

    n.Reset()
    n.CPU.Register.A = 0xC1
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x82 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0x82)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }

    //LSR test
    i.Inst = LSR
    i.Mode = Accumulator
    t.Log("LSR test")

    n.Reset()
    n.CPU.Register.A = 0xC1
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0x60 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0x60)
    }
    if !n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }

    //ROL test
    i.Inst = ROL
    i.Mode = Accumulator
    t.Log("ROL test")

    n.Reset()
    n.CPU.Register.A = 0x60
    n.CPU.Register.P.C = true
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0xc1 {
        t.Errorf("result is wrong:%d, desire:%d", n.CPU.Register.A, 0xc1)
    }
    if n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }

    //ROR test
    i.Inst = ROR
    i.Mode = Accumulator
    t.Log("ROR test")

    n.Reset()
    n.CPU.Register.A = 0x50
    n.CPU.Register.P.C = true
    n.CPU.ExecInstruction(i)

    if n.CPU.Register.A != 0xA8 {
        t.Errorf("result is wrong:%x, desire:%x", n.CPU.Register.A, 0xa8)
    }
    if n.CPU.Register.P.C {
        t.Errorf("C is wrong:%v",n.CPU.Register.P.C)
    }

    i.Inst = BCC
    i.Mode = Relative
    t.Log("BCC start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.C = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.C = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }

    i.Inst = BCS
    i.Mode = Relative
    t.Log("BCS start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.C = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.C = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }
    
    i.Inst = BEQ
    i.Mode = Relative
    t.Log("BEQ start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.Z = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.Z = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }

    i.Inst = BNE
    i.Mode = Relative
    t.Log("BNE start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.Z = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.Z = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }
    
    i.Inst = BVC
    i.Mode = Relative
    t.Log("BCC start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.V = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.V = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }

    i.Inst = BVS
    i.Mode = Relative
    t.Log("BCS start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.V = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.V = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }
    
    i.Inst = BPL
    i.Mode = Relative
    t.Log("BCC start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.N = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.N = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }

    i.Inst = BMI
    i.Mode = Relative
    t.Log("BCS start")

    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.N = true
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x51 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x52)
    }
    
    n.Reset()
    n.CPU.Bus.WriteByte(0x0000, 0x50)
    n.CPU.Register.P.N = false
    n.CPU.ExecInstruction(i)
    if n.CPU.Register.PC != 0x01 {
        t.Errorf("result is wrong:%x, desire:%x",n.CPU.Register.PC,0x01)
    }


}
