package gobar

import (
	"testing"
	"time"
)

func TestBar_NewOption(t *testing.T) {
	bar := Init()
	bar.NewOption(0,8)

	for i:=1;i<=8;i++{
		bar.Play(int64(i))
		time.Sleep(100*time.Millisecond)
	}

	bar.Finish()
}