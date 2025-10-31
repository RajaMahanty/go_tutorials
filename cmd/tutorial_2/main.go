package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int16 = 32767
	intNum += 1
	fmt.Println(intNum)
	intNum += 2
	fmt.Println(intNum)

	var floatNum float64 = 123456.9876543210
	fmt.Println(floatNum)

	var floatNumPlus float32 = 12.5
	var intNumPlus int = 10
	fmt.Println(floatNumPlus + float32(intNumPlus))

	var intNumDiv1 int = 5
	var intNumDiv2 int = 2
	fmt.Println(intNumDiv1 / intNumDiv2)
	fmt.Println(intNumDiv1 % intNumDiv2)

	var str string = "raja" + " " + "mahanty"
	fmt.Println(str)

	var str2 string = `
Raja 
Mahanty
In Roorkee
UK - 266772
`

	fmt.Println(str2)

	var lenght = len(str)

	fmt.Print("Lenght using len function: ")
	fmt.Println(lenght)

	var strLen string = "Raja©"
	fmt.Println(strLen)
	fmt.Print("Length of Raja© using len(): ")
	fmt.Println(len(strLen))

	fmt.Print("Length of Raja© using RuneCountInString(): ")
	fmt.Println(utf8.RuneCountInString(strLen))

	var runeChar rune = 'a'
	fmt.Println(runeChar)

	var myBool bool = false
	fmt.Println(myBool)

	var myBool2 bool = true
	fmt.Println(myBool2)

	// Default values
	var intNumD int
	fmt.Println(intNumD)

	var floatNumD float64
	fmt.Println(floatNumD)

	var strD string
	fmt.Println(strD)

	var runeD rune
	fmt.Println(runeD)

	myName := "Raja"
	fmt.Println(myName)

	var1, var2 := 12, "raaaaaa"
	fmt.Println(var1, var2)

	const Pi float32 = 3.143
	// Pi = 23.43
	fmt.Println(Pi)

	// const raja int
}
