package randomName

import . "gopackages-layout/pkg/helpers/wordz"

func City() string {
	Words = []string{"Moscow", "NY", "Chicago", "Kek"}
	return Random()
}
