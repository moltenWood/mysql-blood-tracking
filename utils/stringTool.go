package utils

// DeleteExtraSpace 清除多余的空格 换行 制表
func DeleteExtraSpace(raw string) string {
	var byteRaw = []byte(raw)
	var byteOutput []byte
	var isEmpty = false
	for i := range byteRaw {
		var c = byteRaw[i]
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			c = ' '
			if isEmpty {
				continue
			}
			isEmpty = true
		} else {
			isEmpty = false
		}
		byteOutput = append(byteOutput, c)
	}
	return string(byteOutput)
}
