package counter

import (
	"fmt"
	"io"
	"sort"
)

type String struct {
	items map[string]int
}

func (s *String) Add(v string) {
	if len(s.items) == 0 {
		s.items = make(map[string]int)
	}
	if _, ok := s.items[v]; !ok {
		s.items[v] = 0
	}
	s.items[v]++
}

type StringCount struct {
	String string
	Count  int
}

func (s *String) Sorted() []*StringCount {
	sortedUrls := make([]*StringCount, 0, len(s.items))
	for url, c := range s.items {
		sortedUrls = append(sortedUrls, &StringCount{String: url, Count: c})
	}

	sort.Slice(sortedUrls, func(i int, j int) bool {
		return sortedUrls[i].Count > sortedUrls[j].Count
	})
	return sortedUrls
}

func (s *String) Print(w io.Writer) {
	for _, kv := range s.Sorted() {
		fmt.Fprintf(w, "%d\t%s\n", kv.Count, kv.String)
	}
}
