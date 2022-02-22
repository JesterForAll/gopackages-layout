package randomName

import . "gopackages-layout/pkg/helpers/wordz"

func Digit() string {
	Words = []string{"one", "two", "three", "four", "five", "six"}
	return Random()
}
