package main

import (
	"testing"
)

// 乱数生成結果によって、確率で失敗する
func TestAttack_確率で失敗するし網羅性なし(t *testing.T) {
	got := Attack()
	want := "HIT!"

	if got != want {
		t.Errorf("Attack() = %v, want %v", got, want)
	}
}

// このテストは滅多に成功しない
func TestAttack_ケース網羅(t *testing.T) {
	tests := []struct {
		name       string
		wantResult string
	}{
		{"0", "CRITICAL!!"},
		{"1~98", "HIT!"},
		{"99", "FUMBLE..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := Attack(); gotResult != tt.wantResult {
				t.Errorf("Attack() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
