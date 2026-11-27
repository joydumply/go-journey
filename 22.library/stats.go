package main

func CountByDecade(books []Book) map[int]int {
	m := make(map[int]int)

	for _, b := range books {
		d := b.Year / 10 * 10
		m[d]++
	}

	return m
}

func AgeCategory(year int) string {
	switch {
	case year < 1950:
		return "classic"
	case year < 2000:
		return "modern"
	default:
		return "contemporary"
	}
}

func LibraryStats(books ...Book) (total int, available int) {
	total = len(books)
	for _, b := range books {
		if b.Available {
			available++
		}
	}
	return total, available
}
