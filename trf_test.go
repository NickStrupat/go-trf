package trf_test

import (
	"testing"
	. "trf"
)

func TestTrf(t *testing.T) {
	Try(
		func() {
			println("try body")
			panic("panic")
		},
		Recovers{
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
