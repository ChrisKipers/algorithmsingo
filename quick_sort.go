package algorithmsingo

import "math"

func QuickSort(elements ...string) []string {
	quickSort(elements, 0, len(elements))
	return elements
}

func quickSort(elements []string, startPos int, endPos int) {
	if endPos-startPos < 2 {
		return
	}
	center := int(math.Ceil(float64(endPos-startPos)/2)) + startPos - 1
	pivotValue := elements[center]
	elements[center] = elements[endPos-1]
	elements[endPos-1] = pivotValue
	borderIndex := startPos

	for i := startPos; i < endPos-1; i++ {
		if elements[i] <= pivotValue {
			borderValue := elements[borderIndex]
			elements[borderIndex] = elements[i]
			elements[i] = borderValue
			borderIndex++
		}
	}
	quickSort(elements, startPos, borderIndex)
	quickSort(elements, borderIndex, endPos)
}
