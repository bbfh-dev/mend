package attrs

import "sort"

// Sorts all attributes based on [AttrSortOrder]
func (attrs *Attributes) Sort() *Attributes {
	order := make(map[string]int)
	for i, s := range AttrSortOrder {
		order[s] = i
	}

	// Ort a while keeping the order for equal elements.
	sort.SliceStable(attrs.order, func(i, j int) bool {
		iIdx, iInB := order[attrs.order[i]]
		jIdx, jInB := order[attrs.order[j]]

		// Both elements are in b: sort by the order in b.
		if iInB && jInB {
			return iIdx < jIdx
		}
		// Only a[i] is in b: it comes first.
		if iInB {
			return true
		}
		// Only a[j] is in b: it comes first.
		if jInB {
			return false
		}
		// Neither element is in b: keep original order.
		return false
	})

	return attrs
}
