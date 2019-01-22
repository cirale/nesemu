package nes

import (
    "log"
    "encoding/json"
)

type KeyPad struct {
    index int
}

func (keypad KeyPad) Read(){
    
}

func (keypad KeyPad) Write(data []byte){
    var msg Message
    json.Unmarshal(data, &msg)
    //log.Printf("error: keypad:%s", msg.Data)

    if msg.Msgtype == "keydown" {
        
    }else if msg.Msgtype == "keyup" {
        
    }else{
        log.Printf("error: mismatch msgtype in keypad: %s", msg.Msgtype)
    }
    
}

func NewKeyPad() *KeyPad {
    var keypad KeyPad
    keypad.index = 0

    return &keypad
}
