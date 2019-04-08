package main

import (
	"fmt"
	"math/rand"
	"time"
)

// D10Roller は 10面ダイスのインターフェース
type D10Roller interface {
	RoleD10() int
}

// Dice implements dice role
type Dice struct {
	rng *rand.Rand
}

// RoleD10 returns 0-9
func (d *Dice) RoleD10() int {
	return d.rng.Intn(10)
}

// Attacker は攻撃役
type Attacker struct {
	d10 D10Roller
}

// Attack 攻撃結果は10面ダイスを2つ振ってパーセンテージ判定する
func (a *Attacker) Attack() (result string) {
	p := a.d10.RoleD10()*10 + a.d10.RoleD10()
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
	a := &Attacker{
		d10: &Dice{
			rng: rand.New(rand.NewSource(time.Now().UnixNano())),
		},
	}
	fmt.Println("Attack Result:", a.Attack())
}
