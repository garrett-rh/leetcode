package deciBinary

func MinPartitions(n string) int {
	var ret byte
	byteArray := []byte(n)

	for _, char := range byteArray {
		if char > ret {
			ret = char
		}
	}

	return int(ret - 48)
}
