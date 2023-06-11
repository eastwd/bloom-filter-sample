package main

import (
	"crypto/sha512"
	"fmt"
)


const (
    hashFunctionAmount = 3
    hashLength = 18
    byteMax = 255
)

var (
    filter = [hashLength]bool{}
)

var (
    dataSet = []string{
        "goma",
        "anko",
        "syoyu",
    }
)

//byte型は0~255までをとる
func UpdateFilter(input string) {
    h := sha512.Sum512([]byte(input))
    for i:=0; i< hashFunctionAmount; i++ {
        a := int(h[i]) * hashLength / byteMax
        filter[a] = true
    }
}

func Validate(input string) bool {
    h := sha512.Sum512([]byte(input))
    for i:=0; i< hashFunctionAmount; i++ {
        a := int(h[i]) * hashLength / byteMax
        if !filter[a] {
            return false
        }
    }
    return true
}

func main(){
    for _, v := range dataSet {
        UpdateFilter(v)
    }
    if Validate("goma") {
        fmt.Println("goma exists.")
    }
    if !Validate("apple") {
        fmt.Println("apple doesn't exist.")
    }
}

