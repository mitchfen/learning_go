package main

import (
	"fmt"

	"github.com/mitchfen/learning_go/pkg/outputHelpers"
)

func main() {
	age := 29
	agePointer := &age // Age pointer is the address of age
	getAdultYears(agePointer)
	outputHelpers.PrettyPrint(fmt.Sprintf("Age: %d", age))
}

func getAdultYears(age *int) {
	fmt.Printf("Age address: %v\n", age)
	fmt.Printf("Age value: %v\n", *age) // Pointer dereference
}
