package loadavg

import "testing"

func TestGet(t *testing.T) {
	t.Log(Get())
}

func BenchmarkGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Get()
	}
}
