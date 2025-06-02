package passbyvalue

import "fmt"

func modifyValuesForInt(a int) {
	a = 20
	fmt.Println("Inside modifyValuesForInt:", a)
}

func modifyValuesForString(a string) {
	a = "modify"
	fmt.Println("Inside modifyValuesForString:", a)
}

func modifyValuesForPointer(a *int) {
	*a = 20
	fmt.Println("Inside modifyValuesForPointer:", *a)
}

func modifyValuesForSliceInt(a []int) {
	a = append(a, 100)
	fmt.Println("Inside modifyValuesForSliceInt:", a)
}

func modifyValuesForSliceInt2(a []int) {
	a[0] = 100
	fmt.Println("Inside modifyValuesForSliceInt:", a)
}
