package main

type MaxSegmentTree SegmentTree[int, int]

// First greater than or equal to
func (st *MaxSegmentTree) FindFirstGTE(target int) (index int, val int) {
	var find func(pos int, left, right int) (int, int)
	find = func(pos int, left, right int) (int, int) {
		if st.nodes[pos] < target {
			return -1, 0
		}
		if left == right {
			return left, st.nodes[pos]
		}

		mid := left + (right-left)/2
		if index, val := find(pos*2+1, left, mid); index != -1 {
			return index, val
		}
		return find(pos*2+2, mid+1, right)
	}

	return find(0, 0, st.size-1)
}
