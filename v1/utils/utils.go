package utils

func Contains(slice []interface{}, item interface{}) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
