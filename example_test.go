package timeseq_test

import (
	"fmt"
	"time"

	"github.com/tebeka/timeseq"
)

// Example shows basic usage of package
func Example() {
	timeFmt := "15:04:05"
	ts := timeseq.New(5)
	now, _ := time.Parse(timeFmt, "1:02:03")
	fmt.Printf("now = %s\n", now.Format(timeFmt))

	// Add elements to the sequence
	for i := 0; i < 100; i++ {
		duration := time.Duration(i) * time.Second
		if i%3 == 0 {
			duration = -duration
		}
		ts.Add(now.Add(duration), i)
	}

	fmt.Printf("size = %d\n", ts.Len())

	// Get elements in time slice
	elems := ts.Slice(now, now.Add(10*time.Second))
	for _, e := range elems {
		val := e.Value.(int)
		fmt.Printf("%s -> %d\n", e.Time.Format(timeFmt), val)
	}

	// Drop elements until now
	ts.Drop(now)
	fmt.Printf("size after drop = %d\n", ts.Len())

	// now = 13:02:03
	// Output: size = 100
	// 01:02:04 -> 1
	// 01:02:05 -> 2
	// 01:02:07 -> 4
	// 01:02:08 -> 5
	// 01:02:10 -> 7
	// 01:02:11 -> 8
	// 01:02:13 -> 10
	// size after drop = 66
}
