package main

import "testing"

func TestGetMessage(t *testing.T) {
	msg := GetMessage()
	if len(msg) < 1 {
		t.Errorf("message can not be empty")
	}
}
