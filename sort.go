package radixsort

import "math"

type NumericOrder interface {
	OrderN() int64
}

type LexicographicalOrder interface {
	OrderL() []byte
}

type negativeWrapper struct {
	origin NumericOrder
	order  int64
}

func (nw negativeWrapper) OrderN() int64 {
	return nw.order
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

func sortLSD(list []NumericOrder) {
	if len(list) == 0 {
		return
	}
	max := getMaxElement(list)
	base := int64(256)
	if max.OrderN() <= 65536 {
		base = 32
	}
	maxOrder := int64(math.MaxInt64) / base
	size := len(list)
	order := int64(1)
	intermediate := make([]NumericOrder, size, size)
	indices := make([]byte, size, size)
	bucket := [256]uint64{}

	for order <= max.OrderN() {
		for i, o := range list {
			idx := byte((o.OrderN() / order) % base)
			bucket[idx]++
			indices[i] = idx
		}

		for i := int64(1); i < base; i++ {
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

// this is LSD radix sort
func SortNumericOrder(list []NumericOrder) {
	listNeg := []NumericOrder{}
	listNNeg := []NumericOrder{}

	for _, v := range list {
		o := v.OrderN()
		if o < 0 {
			listNeg = append(listNeg, negativeWrapper{
				origin: v,
				order:  -o,
			})
		} else {
			listNNeg = append(listNNeg, v)
		}
	}

	sortLSD(listNeg)
	sortLSD(listNNeg)

	idx := 0
	for i := len(listNeg) - 1; i >= 0; i-- {
		list[idx] = (listNeg[i].(negativeWrapper)).origin
		idx++
	}
	copy(list[idx:], listNNeg)
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
