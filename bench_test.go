package queues_test

import (
	"testing"

	"github.com/twharmon/queues"
)

func BenchmarkEnqueue(b *testing.B) {
	b.Run("enqueue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			q := queues.New[int]()
			for j := 0; j < 1000; j++ {
				q.Enqueue(j)
			}
		}
	})
	b.Run("enqueue dequeue", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			q := queues.New[int]()
			for j := 0; j < 1000; j++ {
				q.Enqueue(j)
				if j%100 == 0 {
					for k := 0; k < 100; k++ {
						q.Dequeue()
					}
				}
			}
		}
	})
}
