package design_patterns

import "testing"

func TestGetNon(t *testing.T) {
	t.Log(GetNon().Value)
}

func BenchmarkGetNon(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.Log(GetNon().Value)
	}
}
