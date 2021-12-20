package util

func FilterEmptyString(field1, field2 string) string {
	if field1 == "" {
		return field2
	}
	return field1
}
