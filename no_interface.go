package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RoleD10 returns 0-9
func RoleD10() int {
	return rand.Intn(10)
}

// Attack 攻撃結果は10面ダイスを2つ振ってパーセンテージ判定する
func Attack() (result string) {
	p := RoleD10()*10 + RoleD10()
	fmt.Println("2D10 Result:", p)

	if p == 0 {
		return "CRITICAL!!"
	}

	if p == 99 {
		return "FUMBLE..."
	}

	return "HIT!"
}

func main() {
	fmt.Println("Attack Result:", Attack())
}
