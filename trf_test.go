package trf

import (
	"fmt"
	"testing"
)

func TestTrf(t *testing.T) {
	Try(
		func() {
			println("try body")
			panic("panic")
		},
		RecoverBlocks{
			Recover(func(err string) {
				println("caught string: ", err)
			}),
			Recover(func(err int) {
				println("caught int:", err)
			}),
			Recover(func(err any) {
				println("caught any", err)
			}),
		},
		func() {
			println("finally")
		},
	)
}

func TestTrfNilBody(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic, got nil")
		}
	}()

	Try(
		nil,
		RecoverBlocks{
			Recover(func(err string) {
				println("caught string: ", err)
			}),
			Recover(func(err int) {
				println("caught int:", err)
			}),
			Recover(func(err any) {
				println("caught any", err)
			}),
		},
		func() {
			println("finally")
		},
	)
}

func TestTrfNilRecoverBlock(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Expected panic, got nil")
		}
		msg := fmt.Sprintf(recoverBlockNilError, 0)
		if r != msg {
			t.Errorf("Expected panic with message '%v', got '%v'", msg, r)
		}
	}()

	Try(
		func() {
			println("try body")
			panic("panic")
		},
		RecoverBlocks{
			nil,
			Recover(func(err int) {
				println("caught int:", err)
			}),
			Recover(func(err any) {
				println("caught any", err)
			}),
		},
		func() {
			println("finally")
		},
	)
}
