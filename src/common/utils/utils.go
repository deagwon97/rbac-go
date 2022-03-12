package utils

func IsIn(item int, list []int) bool {
	for _, validItem := range list {
		if validItem == item {
			return true
		}
	}
	return false
}
