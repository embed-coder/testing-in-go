package bench

import (
	"io/ioutil"
	"testing"
)

func TestSolve(t *testing.T) {
	c, _ := ioutil.ReadFile("1.txt")
	res := solve(c)
	expected := 3
	if expected != res {
		t.Fatalf("expected %d, got %d", expected, res)
	}
}

func BenchmarkSolve(b *testing.B) {
	c, _ := ioutil.ReadFile("1.txt")
	for i := 0; i < b.N; i++ {
		solve(c)
	}
}
