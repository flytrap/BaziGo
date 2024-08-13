package bazi_test

import (
	"testing"
	"time"

	bz "BaziGo"

	"github.com/stretchr/testify/assert"
)

func TestDayun(t *testing.T) {
	n := time.Now()
	b := bz.GetBazi(n.Year(), int(n.Month()), n.Day(), n.Hour(), n.Minute(), n.Second(), 1)

	dy := bz.NewDaYun(b.SiZhu(), 1)
	assert.NotNil(t, dy.Zhu(0))
	assert.NotNil(t, dy.Zhu(0).CangGan())
	assert.NotNil(t, dy.Zhu(0).Changesheng())
	assert.NotNil(t, dy.Zhu(0).ShiShen())
}
