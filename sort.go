package radixsort

import "math"

const negShift = math.MaxInt64

type NumericOrder interface {
	OrderN() int64
}

type LexicographicalOrder interface {
	OrderL() []byte
}

func getMaxElement(list []NumericOrder) NumericOrder {
	var max NumericOrder

	for _, o := range list {
		if max == nil || o.OrderN() > max.OrderN() {
			max = o
		}
	}

	return max
}

// this is LSD radix sort
func SortNumericOrder(list []NumericOrder) {
	if len(list) == 0 {
		return
	}
	max := uint64(getMaxElement(list).OrderN()+negShift) + 1
	base := uint64(256)
	maxOrder := uint64(math.MaxUint64) / base
	size := len(list)
	order := uint64(1)
	intermediate := make([]NumericOrder, size, size)
	indices := make([]byte, size, size)
	bucket := [256]uint64{}

	for order <= max {
		for i, o := range list {
			ov := uint64(o.OrderN()+negShift) + 1
			idx := byte((ov / order) % base)
			bucket[idx]++
			indices[i] = idx
		}

		for i := uint64(1); i < base; i++ {
			bucket[i] += bucket[i-1]
		}

		for i := size - 1; i >= 0; i-- {
			idx := indices[i]
			bucket[idx]--
			intermediate[bucket[idx]] = list[i]
		}

		copy(list, intermediate)

		if order > maxOrder {
			break
		}
		order *= base
		bucket = [256]uint64{}
	}
}

func sortMSD(list, aux []LexicographicalOrder, lo, hi, pos int) {
	count := [258]int{}

	for i := lo; i <= hi; i++ {
		b := list[i].OrderL()
		iv := int16(1)
		if len(b) > pos {
			iv = 2 + int16(b[pos])
		}
		count[iv]++
	}

	for i := 1; i < 258; i++ {
		count[i] += count[i-1]
	}

	for i := lo; i <= hi; i++ {
		b := list[i].OrderL()
		iv := int16(0)
		if len(b) > pos {
			iv = 1 + int16(b[pos])
		}

		aux[count[iv]] = list[i]
		count[iv]++
	}

	for i := lo; i <= hi; i++ {
		list[i] = aux[i-lo]
	}

	for i := 0; i < 256; i++ {
		h := lo + count[i+1] - 1
		l := lo + count[i]
		if h > l {
			sortMSD(list, aux, l, h, pos+1)
		}
	}
}

// this is MSD radix sort
func SortLexicographicalOrder(list []LexicographicalOrder) {
	if len(list) == 0 {
		return
	}

	aux := make([]LexicographicalOrder, len(list))
	sortMSD(list, aux, 0, len(list)-1, 0)
}
