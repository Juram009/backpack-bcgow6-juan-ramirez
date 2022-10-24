package sorting

func BubbleSort(items []int) {
	L := len(items)
	for i := 0; i < L; i++ {
		for j := 0; j < (L - 1 - i); j++ {
			if items[j] > items[j+1] {
				items[j], items[j+1] = items[j+1], items[j]
			}
		}
	}
}

func InsertionSort(items []int) {
	L := len(items)
	for i := 1; i < L; i++ {
		j := i
		for j > 0 && items[j] < items[j-1] {
			items[j], items[j-1] = items[j-1], items[j]
			j -= 1
		}
	}
}

func QuickSort(list []int, low int, high int) {
	if low < high {
		pi := partition(list, low, high)
		QuickSort(list, low, pi-1)
		QuickSort(list, pi+1, high)
	}
}

func swap(a *int, b *int) {
	t := *a
	*a = *b
	*b = t
}

func partition(list []int, low int, high int) int {
	pivot := list[high]
	i := (low - 1)

	for j := low; j < high; j++ {
		if list[j] <= pivot {
			i++
			swap(&list[i], &list[j])
		}
	}
	swap(&list[i+1], &list[high])

	return i + 1
}
