package main

import (
	"fmt"
	"gopackages-layout/pkg/helpers/randomName"
	"gopackages-layout/pkg/helpers/wordz"
	"github.com/huandu/xstrings"
	"github.com/JesterForAll/utils"
	utilsV2 "github.com/JesterForAll/utils/v2"
)

func main()  {
	//go 3 1, 1st exercise
	wordz.Prefix = ""

	if utilsV2.InSlice(wordz.Words, "Two"){
		fmt.Println("utilsV2: Slice contains value")
	}

	if utilsV2.InSliceInt([]int{1,2,3,4,5}, 3){
		fmt.Println("utilsV2: Slice int contains value")
		return
	}

	if utils.Contains(wordz.Words, "Two"){
		fmt.Println("Slice contains value")
	}

	if utils.ContainsInt([]int{1,2,3,4,5}, 3){
		fmt.Println("Slice int contains value")
	}

	fmt.Println(randomName.City())
	fmt.Println(randomName.Digit())
	fmt.Println(xstrings.Shuffle(randomName.City()))
}