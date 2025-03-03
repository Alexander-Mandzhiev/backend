package service

func convertToInt(ids []int32) []int {
	var result []int
	for _, id := range ids {
		result = append(result, int(id))
	}
	return result
}
