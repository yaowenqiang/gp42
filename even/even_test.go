// Package even provides ...
package even

import (
	"testing"
)

func TestLoop(t *testing.T) {
	t.Log("Loop: ", Loop(uint64(34)))
}

func TestFactorial(t *testing.T) {
	t.Log("Factorial:", Factorial(uint64(32)))
}

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Loop(uint64(40))
	}
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(uint64(40))
	}
}

//go test -test.bench=.*
//go test -x -v -test.cpuprofile=pprof.out
