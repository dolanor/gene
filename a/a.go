// Package a is a POC of some async await pattern
// in Go.
package a

type promise[T any] struct {
	response <-chan T
}

func (p *promise[T]) Value() T {
	return <-p.response
}

// Sync is the async part of the pattern.
// It takes a niladic function that returns a T
// and transforms it into a promise of T.
func Sync[T any](f func() T) promise[T] {
	tChan := make(chan T)
	go func() {
		tChan <- f()
	}()

	return promise[T]{
		response: tChan,
	}
}

// Wait is the sync part of the pattern.
// It takes a promise of T and returns the Value.
func Wait[T any](p promise[T]) T {
	return p.Value()
}
