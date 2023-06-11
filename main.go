package main

import (
	"crypto/sha512"
	"fmt"
)


const (
    hashFunctionAmount = 3
    hashLength = 18
    byteMax = 255
    m = hashLength
    k = hashFunctionAmount
)

var (
    filter = [hashLength]bool{}
)

type hashFunction func(data string)[hashLength]string

var (
    dataSet = []string{
        "goma",
        "anko",
        "syoyu",
    }
)

// 追加された新しい要素 input
// ハッシュテーブルの位置
//byte型は0~255までをとる
//ハッシュ関数の数の分だけ先頭から値を取得し、ハッシュ配列の長さでbyte型の返却値を正規化した時の値を配列で返す
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

