package helpers

func StringRepeat(s string, count int) string {
	byteArr := make([]byte, count*len(s))
	byteToAdd := s[0]
	for i := 0; i < count; i++ {
		byteArr = append(byteArr, byteToAdd)
	}
	return string(byteArr)
}
