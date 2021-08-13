package pkg


func Page(data []int, pageNo int, pageSize, totalCount int) []int {
	if pageNo <= 0 {
		return []int{}
	}

	start := pageNo-1
	idxStart := start * pageSize
	idxEnd := (start+1) * pageSize

	if idxStart >= totalCount {
		return []int{}
	}

	if idxEnd > totalCount {
		idxEnd = totalCount
	}


	return data[idxStart: idxEnd]
}
