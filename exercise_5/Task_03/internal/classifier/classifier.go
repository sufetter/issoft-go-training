package classifier

import "knn/internal/entities"

func optimizedDistance(a, b entities.Object) float64 {
	return (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
}

func quickSelect(dists []entities.Dist, k int) {
	if k > len(dists) {
		k = len(dists)
	}

	left := 0
	right := len(dists) - 1

	for left < right {
		pivotIndex := partition(dists, left, right)

		if pivotIndex == k {
			break
		} else if pivotIndex > k {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}

func partition(dists []entities.Dist, left int, right int) int {
	pivotIndex := right
	pivotValue := dists[pivotIndex].Dist
	right--

	for {
		for dists[left].Dist < pivotValue && left < right {
			left++
		}

		for dists[right].Dist >= pivotValue && left < right {
			right--
		}

		if left == right {
			break
		}

		dists[left], dists[right] = dists[right], dists[left]
	}

	dists[left], dists[pivotIndex] = dists[pivotIndex], dists[left]

	return left
}

func Classify(objects []entities.Object, unknown entities.Object, k int) string {
	dists := make([]entities.Dist, len(objects))
	for i, obj := range objects {
		dists[i] = entities.Dist{Name: obj.Name, Dist: optimizedDistance(obj, unknown)}
	}

	quickSelect(dists, k)

	freq := make(map[string]int)
	for i := 0; i < k; i++ {
		freq[dists[i].Name]++
	}

	var maxClass string
	maxFreq := 0
	for name, f := range freq {
		if f > maxFreq {
			maxFreq = f
			maxClass = name
		}
	}

	return maxClass
}
