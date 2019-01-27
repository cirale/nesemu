package nes

import (
    "log"
    "encoding/json"
)

type KeyPad struct {
    index int
    isSet bool
    KeyRegister []bool
    KeyBuffer []bool
}

const (
    A = iota
    B
    SELECT
    START
    UP
    DOWN
    LEFT
    RIGHT
)

func (keypad *KeyPad) Read() byte {
    boolToByte := map[bool]byte{true:1,false:0}
    keypad.index++
    return boolToByte[keypad.KeyRegister[keypad.index-1]]
}

func (keypad *KeyPad) Write(data byte){
    log.Printf("error: %v",data)
    if data == 0x01 {
        keypad.isSet = true
    } else if data == 0x00 && keypad.isSet {
        keypad.isSet = false
        keypad.index = 0
        keypad.KeyRegister = keypad.KeyBuffer
        keypad.KeyBuffer = make([]bool,8)
    }
}

func (keypad *KeyPad) Set(data []byte){
    var msg Message
    json.Unmarshal(data, &msg)
    
    if msg.Msgtype == "keydown" {
        keypad.KeyBuffer[keypad.KeyToNumber(string(msg.Data))] = true

    }else if msg.Msgtype == "keyup" {
        keypad.KeyBuffer[keypad.KeyToNumber(string(msg.Data))] = false

    }else{
        log.Printf("error: mismatch msgtype in keypad: %s", msg.Msgtype)
    }
}

func (keypad *KeyPad) KeyToNumber(key string) int {
    switch(key){
    case "a":
        return A
    case "b":
        return B
    case "select":
        return SELECT
    case "start":
        return START
    case "up":
        return UP
    case "down":
        return DOWN
    case "left":
        return LEFT
    case "right":
        return RIGHT
    default:
        return -1
    }
}

func NewKeyPad() *KeyPad {
    var keypad KeyPad
    keypad.index = 0
    keypad.isSet = false
    keypad.KeyRegister = make([]bool,8)
    keypad.KeyBuffer = make([]bool,8)
    
    return &keypad
}
