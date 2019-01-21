package nes

type KeyPad struct {
    index int
}

func (keypad KeyPad) read(){
    
}

func (keypad KeyPad) write(data string){
    
}

func NewKeyPad() *KeyPad {
    var keypad KeyPad
    keypad.index = 0

    return &keypad
}
