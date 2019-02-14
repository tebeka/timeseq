package timeseq

import (
	"testing"
	"time"
)

func BenchmarkInsert(b *testing.B) {
	b.StopTimer()
	ts, now := New(1035), time.Now()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		dt := time.Duration(i) * time.Millisecond
		if i%2 == 0 {
			dt = -dt
		}
		ts.Add(now.Add(dt), i)
	}
}

func BenchmarkSlice(b *testing.B) {
	b.StopTimer()
	ts, now := New(230), time.Now()
	for i := 0; i < 10017; i++ {
		dt := time.Duration(i) * time.Millisecond
		if i%2 == 0 {
			dt = -dt
		}
		ts.Add(now.Add(dt), i)
	}
	b.StartTimer()

	end := now.Add(400 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		s := ts.Slice(now, end)
		if len(s) == 0 {
			b.Fatal("empty slice")
		}
	}
}
