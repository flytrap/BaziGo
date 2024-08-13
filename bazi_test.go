package bazi_test

import (
	"testing"

	bz "BaziGo"
)

func TestMingGong(t *testing.T) {
	b := bz.GetBazi(2020, 3, 3, 3, 0, 0, 1) // 男
	if b.MingGong().String() != "己丑" {
		t.Error(b.MingGong())
	}
}
