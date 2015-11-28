package algorithmsingo

import "math"

func MergeSort(elements ...string) []string {
	if len(elements) < 2 {
		return elements
	}

	center := int(math.Ceil(float64(len(elements)) / 2))
	sortedFirstHalf := MergeSort(elements[:center]...)
	sortedLastHalf := MergeSort(elements[center:]...)

	sortedElements := make([]string, 0, len(elements))
	sortedFirstElementsIndex := 0
	sortedLastElementsIndex := 0

	for len(sortedElements) != len(elements) {
		if sortedFirstElementsIndex < len(sortedFirstHalf) && sortedLastElementsIndex < len(sortedLastHalf) {
			if sortedFirstHalf[sortedFirstElementsIndex] < sortedLastHalf[sortedLastElementsIndex] {
				sortedElements = append(sortedElements, sortedFirstHalf[sortedFirstElementsIndex])
				sortedFirstElementsIndex++
			} else {
				sortedElements = append(sortedElements, sortedLastHalf[sortedLastElementsIndex])
				sortedLastElementsIndex++
			}
		} else if sortedFirstElementsIndex < len(sortedFirstHalf) {
			sortedElements = append(sortedElements, sortedFirstHalf[sortedFirstElementsIndex])
			sortedFirstElementsIndex++
		} else {
			sortedElements = append(sortedElements, sortedLastHalf[sortedLastElementsIndex])
			sortedLastElementsIndex++
		}
	}

	return sortedElements
}
