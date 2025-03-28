package trf

import "reflect"

func Try(
	body func(),
	catches []RecoverBlock,
	finally func(),
) {
	if body == nil {
		panic("body argument is nil")
	}
	if finally != nil {
		defer finally()
	}
	if catches != nil && len(catches) != 0 {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			t := reflect.TypeOf(err)
			for _, c := range catches {
				if c == nil {
					continue
				}
				if c.tryInvoke(t, err) {
					return
				}
			}
			panic(err)
		}()
	}
	body()
}

type Recovers []RecoverBlock

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
