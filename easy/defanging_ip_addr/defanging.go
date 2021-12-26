package defanging

import "strings"

func DefangIPAddr(address string) string {
	address = strings.ReplaceAll(address, ".", "[.]")
	return address
}
