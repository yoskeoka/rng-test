package main

import (
	"testing"
)

type RandomNumberGeneratorMock struct {
	IntnResults []int
}

func (rng *RandomNumberGeneratorMock) Intn(n int) int {
	result := rng.IntnResults[0]
	rng.IntnResults = rng.IntnResults[1:]
	return result
}

// 乱数生成をモックにすることで確実に望みの乱数に対する挙動をテストできる
func TestAttack_RNGのMockで確実に成功(t *testing.T) {
	// 1 が出るように仕込む
	RNG = &RandomNumberGeneratorMock{IntnResults: []int{0, 1}}
	got := Attack()
	want := "HIT!"

	if got != want {
		t.Errorf("Attack() = %v, want %v", got, want)
	}
}

// 滅多に起きない確率のケースもテスト可能
func TestAttack_ケース網羅(t *testing.T) {
	type fakes struct {
		FakeIntnResults []int
	}
	tests := []struct {
		name       string
		fakes      fakes
		wantResult string
	}{
		{"0", fakes{FakeIntnResults: []int{0, 0}}, "CRITICAL!!"},
		{"55", fakes{FakeIntnResults: []int{5, 5}}, "HIT!"},
		{"99", fakes{FakeIntnResults: []int{9, 9}}, "FUMBLE..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fakes で指定したRNGの結果をここで仕込む
			RNG = &RandomNumberGeneratorMock{IntnResults: tt.fakes.FakeIntnResults}

			if gotResult := Attack(); gotResult != tt.wantResult {
				t.Errorf("Attack() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
