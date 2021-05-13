// 1108. Defanging an IP Address
// question
// 替换string中的 . 为 [.]
// example
// Input "1.1.1.1", Output: "1[.]1[.]1[.]1"

package main

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
