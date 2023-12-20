package a_test

import (
	"log"
	"testing"
	"time"

	"github.com/dolanor/gene/a"
)

func ExampleWait() {
	b := 1
	c := 5
	slowAdd := func() int {
		time.Sleep(1 * time.Second)
		return b + c
	}
	promise := a.Sync(slowAdd)

	sum := a.Wait(promise)

	log.Println("1 + 5 =", sum)
	// 1 + 5 = 6
}

func ExampleSync() {
	b := 1
	c := 5
	slowAdd := func() int {
		time.Sleep(1 * time.Second)
		return b + c
	}
	promise := a.Sync(slowAdd)

	log.Println(promise.Value())
	// 6
}

func TestAwaitAsync(t *testing.T) {
	b := 1
	c := 5
	slowAdd := func() int {
		time.Sleep(1 * time.Second)
		return b + c
	}
	promise := a.Sync(slowAdd)
	start := time.Now()

	sum := a.Wait(promise)

	dur := time.Now().Sub(start)

	t.Log(dur)
	if sum != 6 {
		t.Fatal("sum wrongly computed")
	}

	rounded := dur.Round(1 * time.Millisecond)
	t.Log(rounded)
	if rounded != time.Second {
		t.Fatal("the promise didn't wait correctly")
	}
}
