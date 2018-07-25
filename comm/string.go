package comm

//Substring 根据长度截取字符串，超过长度不会报溢出错误
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 {
		start = 0
	} else if start >= length {
		start = length
	}
	if end >= length {
		end = length
	}
	return string(r[start:end])
}
