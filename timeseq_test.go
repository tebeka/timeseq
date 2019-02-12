package timeseq

import (
	"fmt"
	"testing"
	"time"
)

func formatTimes(ts []*Elem) string {
	out := make([]string, len(ts))
	for i, e := range ts {
		out[i] = e.Time.Format("15:04:03")
	}
	return fmt.Sprintf("%v", out)
}

func TestTimeSeq(t *testing.T) {
	now, err := time.Parse("15:04:03", "1:02:03")
	if err != nil {
		t.Fatal(err)
	}

	size := 10
	ts := New(size / 2)
	for i := 0; i < size; i++ {
		dt := time.Duration(i) * time.Minute
		if i%2 == 0 {
			dt = -dt
		}
		ts.Add(now.Add(dt), i)
	}

	if count := ts.Len(); count != size {
		t.Fatalf("size mismatch: got %d, expected %d", count, size)
	}

	start, end := now.Add(-3*time.Minute), now.Add(3*time.Minute)
	out := ts.Slice(start, end)
	if count := len(out); count != 4 {
		t.Fatalf("slice mismatch: got %d, expected 4", count)
	}

	i := 3
	until := (*ts)[i].Time
	ts.Drop(until)
	if count := ts.Len(); count != size-i-1 {
		t.Fatalf("slice mismatch: got %d, expected %d", count, size-i)
	}
}
