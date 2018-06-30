package magelib

import "testing"

func TestInfo(t *testing.T) {
	Info("this", "is", "a", "test")
}

func TestWarn(t *testing.T) {
	Warn("this", "is", "a", "test")
}
