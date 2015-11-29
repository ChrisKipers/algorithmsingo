package algorithmsingo
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	t.Log("When I have a populated heap")
	{
		myHeap := NewHeap()
		myHeap.Add(10)
		myHeap.Add(2)
		myHeap.Add(4)
		myHeap.Add(8)
		myHeap.Add(14)
		myHeap.Add(-1)
		myHeap.Add(-10)
		myHeap.Add(100)

		t.Log("peek returns the min item")
		{
			assert.Equal(t,myHeap.Peak(), -10)
		}

		t.Log("Size returns the size")
		{
			assert.Equal(t, myHeap.Size(), 8)
		}

		t.Log("it returns the items in order")
		{
			assert.Equal(t,myHeap.Get(), -10)
			assert.Equal(t,myHeap.Get(), -1)
			assert.Equal(t,myHeap.Get(), 2)
			assert.Equal(t,myHeap.Get(), 4)
			assert.Equal(t,myHeap.Get(), 8)
			assert.Equal(t,myHeap.Get(), 10)
			assert.Equal(t,myHeap.Get(), 14)
			assert.Equal(t,myHeap.Get(), 100)
		}
	}
}
