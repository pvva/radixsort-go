package radixsort

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

const listSize = 100000

type oint int64

func (i oint) OrderN() int64 {
	return int64(i)
}

type ostring string

func (s ostring) OrderL() []byte {
	return []byte(s)
}

func prepareListNumeric() []NumericOrder {
	rand.Seed(time.Now().UnixNano())
	list := make([]NumericOrder, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = oint(rand.Int63())
	}

	return list
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func prepareListLexicographical() []LexicographicalOrder {
	rand.Seed(time.Now().UnixNano())
	list := make([]LexicographicalOrder, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = ostring(randomString(64))
	}

	return list
}

func TestSortNumeric(t *testing.T) {
	list := prepareListNumeric()
	dup := make([]NumericOrder, len(list))
	copy(dup, list)

	sort.SliceStable(dup, func(i, j int) bool {
		return dup[i].OrderN() < dup[j].OrderN()
	})
	SortNumericOrder(list)

	for i, o := range dup {
		if list[i].OrderN() != o.OrderN() {
			t.Fatalf("SortNumericOrder failed")
		}
	}
}

func TestSortLexicographical(t *testing.T) {
	ostrings := []ostring{"ab", "a", "d", "abc", "ff", "qq", "aaa"}
	list := make([]LexicographicalOrder, len(ostrings))
	for i, v := range ostrings {
		list[i] = v
	}

	expected := []ostring{"a", "aaa", "ab", "abc", "d", "ff", "qq"}
	SortLexicographicalOrder(list)

	for i, v := range list {
		if v != expected[i] {
			t.Fatalf("SortLexicographicalOrder failed")
		}
	}
}

func BenchmarkSortNumeric(b *testing.B) {
	list := prepareListNumeric()

	for i := 0; i < b.N; i++ {
		l := make([]NumericOrder, len(list))
		copy(l, list)
		SortNumericOrder(l)
	}
}

func BenchmarkMergeSortNumeric(b *testing.B) {
	list := prepareListNumeric()

	for i := 0; i < b.N; i++ {
		l := make([]NumericOrder, len(list))
		copy(l, list)
		sort.SliceStable(l, func(i, j int) bool {
			return l[i].OrderN() < l[j].OrderN()
		})
	}
}

func BenchmarkSortLexicographical(b *testing.B) {
	list := prepareListLexicographical()

	for i := 0; i < b.N; i++ {
		l := make([]LexicographicalOrder, len(list))
		copy(l, list)
		SortLexicographicalOrder(l)
	}
}
