package algorithmsingo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	t.Log("when sorting elements that are in a random order")
	{
		elements := []string{"a", "z", "c", "q", "g", "e", "o"}
		sortedElements := QuickSort(elements...)

		t.Log("it sorts the elements correctly")
		{
			expectedSortedElements := []string{"a", "c", "e", "g", "o", "q", "z"}
			assert.Equal(t, expectedSortedElements, sortedElements)
		}
	}

	t.Log("when sorting elements that are in the correct order")
	{
		elements := []string{"a", "c", "e", "g", "o", "q", "z"}
		sortedElements := QuickSort(elements...)

		t.Log("it sorts the elements correctly")
		{
			expectedSortedElements := []string{"a", "c", "e", "g", "o", "q", "z"}
			assert.Equal(t, expectedSortedElements, sortedElements)
		}
	}

	t.Log("when sorting elements that are in the reverse order")
	{
		elements := []string{"z", "q", "o", "g", "e", "c", "a"}
		sortedElements := QuickSort(elements...)

		t.Log("it sorts the elements correctly")
		{
			expectedSortedElements := []string{"a", "c", "e", "g", "o", "q", "z"}
			assert.Equal(t, expectedSortedElements, sortedElements)
		}
	}

	t.Log("when sorting no elements")
	{
		sortedElements := QuickSort()

		t.Log("it returns an empty slice")
		{
			assert.Empty(t, sortedElements)
		}
	}

	t.Log("when sorting identical elements")
	{
		elements := []string{"a", "a", "a", "a", "a"}
		sortedElements := QuickSort(elements...)

		t.Log("it sorts the elements correctly")
		{
			assert.Equal(t, elements, sortedElements)
		}
	}
}
