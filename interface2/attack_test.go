package main

import (
	"testing"
)

type DiceMock struct {
	Results []int
}

func (rng *DiceMock) RoleD10() int {
	result := rng.Results[0]
	rng.Results = rng.Results[1:]
	return result
}

// Diceをモックにすることで確実に望みの結果に対する挙動をテストできる
func TestAttack_DiceのMockで確実に成功(t *testing.T) {
	// 1 が出るように仕込む
	dice := &DiceMock{Results: []int{0, 1}}
	a := &Attacker{d10: dice}
	got := a.Attack()
	want := "HIT!"

	if got != want {
		t.Errorf("Attack() = %v, want %v", got, want)
	}
}

// 滅多に起きない確率のケースもテスト可能
func TestAttack_ケース網羅(t *testing.T) {
	type fakes struct {
		FakeResults []int
	}
	tests := []struct {
		name       string
		fakes      fakes
		wantResult string
	}{
		{"0", fakes{FakeResults: []int{0, 0}}, "CRITICAL!!"},
		{"55", fakes{FakeResults: []int{5, 5}}, "HIT!"},
		{"99", fakes{FakeResults: []int{9, 9}}, "FUMBLE..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fakes で指定したDiceの結果をここで仕込む
			dice := &DiceMock{Results: tt.fakes.FakeResults}
			a := &Attacker{d10: dice}
			if gotResult := a.Attack(); gotResult != tt.wantResult {
				t.Errorf("Attack() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
