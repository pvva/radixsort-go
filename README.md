This repository contains [radix sort algorithm](https://en.wikipedia.org/wiki/Radix_sort)  implementation in Go.

LSD sort is a bit tuned in terms of performance.

Numeric sort.
```
type MyType uint64

func (t MyType) OrderN() uint64 {
	return uint64(t)
}

...

var list []radixsort.NumericOrder

radixsort.SortNumericOrder(list)
// 12, 2, 4, 1, 5, 7 => 1, 2, 4, 5, 7, 12
```

Lexicographical sort.
```
type MyType string

func (t MyType) OrderL() []byte {
	return []byte(t)
}

...

var list []radixsort.LexicographicalOrder

radixsort.SortLexicographicalOrder(list)
// "c", "a", "ba", "ab", "bb", "aaa" => "a", "aaa", "ab", "ba", "bb", "c"
```

Usage of, for example, insertion sort for lexicographical order for small lists is not done. It is actually not planned to be used for small arrays. Feel free to fork and adapt the code.
