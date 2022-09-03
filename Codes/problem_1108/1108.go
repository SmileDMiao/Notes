package leetcode

func defangIPaddr(address string) string {
	ip := ""

	for _, char := range address {
		if string(char) == "." {
			ip += "[.]"
		} else {
			ip += string(char)
		}
	}
	return ip
}
