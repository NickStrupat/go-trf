package trf

import (
	"fmt"
	"reflect"
)

const (
	bodyNilError            = "Argument `body` is nil. It must be a function."
	recoverBlocksNilError   = "Argument `recovers` is nil. It must be a slice of RecoverBlock items."
	recoverBlocksEmptyError = "Argument `recovers` is empty. It must contain at least one recover block."
	recoverBlockNilError    = "Argument `recovers` contains a nil at index %d. All recover blocks must each be a function."
)

func Try(
	body func(),
	recoverBlocks []RecoverBlock,
	finally func(),
) {
	// Check the arguments
	if body == nil {
		panic(bodyNilError)
	}
	if recoverBlocks == nil {
		panic(recoverBlocksNilError)
	}
	if len(recoverBlocks) == 0 {
		panic(recoverBlocksEmptyError)
	}
	for i, c := range recoverBlocks {
		if c == nil {
			panic(fmt.Sprintf(recoverBlockNilError, i))
		}
	}

	// Defer the finally and recovery funcs
	if finally != nil {
		defer finally()
	}

	defer func() {
		err := recover()
		if err == nil {
			return
		}
		t := reflect.TypeOf(err)
		for _, c := range recoverBlocks {
			if c.tryInvoke(t, err) {
				return
			}
		}
		panic(err)
	}()

	// Call the body
	body()
}

type RecoverBlocks []RecoverBlock

func Recover[TEx any](handler func(ex TEx)) RecoverBlock {
	return recoverBlock[TEx]{
		t: reflect.TypeFor[TEx](),
		h: handler,
	}
}

type RecoverBlock interface {
	tryInvoke(t reflect.Type, ex any) bool
}

type recoverBlock[T any] struct {
	t reflect.Type
	h func(ex T)
}

func (r recoverBlock[T]) tryInvoke(t reflect.Type, ex any) bool {
	invoke := t.AssignableTo(r.t)
	if invoke {
		r.h(ex.(T))
	}
	return invoke
}
