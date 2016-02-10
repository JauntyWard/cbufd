package cbuf

import "testing"

func TestEnqueueDequeue(t *testing.T) {
	testBuffer := NewCircularBuffer(1000)
	for i := 0; i < 2000; i++ {
		testBuffer.Enqueue(i)
	}

	for i := 1000; i < 2000; i++ {
		result := testBuffer.Dequeue()
		if result != i {
			t.Errorf("Expected %d, got %d", i, result)
		}
	}
}

func EnqueuePeakDequeue(t *testing.T) {
	testBuffer := NewCircularBuffer(1000)
	for i := 0; i < 2000; i++ {
		testBuffer.Enqueue(i)
	}

	for i := 1000; i < 2000; i++ {
		peakResult := testBuffer.Peak()
		if peakResult != i {
			t.Errorf("Expected %d, got %d", i, peakResult)
		}
		dequeueResult := testBuffer.Dequeue()
		if dequeueResult != i {
			t.Errorf("Expected %d, got %d", i, dequeueResult)
		}
	}
}
