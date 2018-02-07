package golib

import "testing"

func Test_Xor(t *testing.T) {
	if Xor(true, true) {
		t.Error("xor err")
	}
	if !Xor(false, true) {
		t.Error("xor err")
	}
}

func Test_ThreeInputXOR(t *testing.T) {
	if ThreeInputXOR(true, true, false) {
		t.Error("xor err")
	}
	if !ThreeInputXOR(false, true, false) {
		t.Error("xor err")
	}
}
