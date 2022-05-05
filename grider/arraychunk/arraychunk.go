package arrayChunk

func Chunk(data []int, subSliceSize int) [][]int {
	result := [][]int{}
	lenData := len(data)

	i := 0
	for i < len(data) {
		endIdx := i + subSliceSize
		if endIdx > lenData {
			endIdx = lenData
		}

		tmpSlice := data[i:endIdx]
		result = append(result, tmpSlice)
		i = i + subSliceSize
	}

	return result
}
