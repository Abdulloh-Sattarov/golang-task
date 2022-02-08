package main

import "fmt"

func main() {
	ok := AddTwoLinkedNumbers(&LinkedNumbers{2, &LinkedNumbers{4, &LinkedNumbers{3, nil}}}, &LinkedNumbers{5, &LinkedNumbers{6, &LinkedNumbers{4, nil}}})
	for ok != nil {
		fmt.Println(ok.Number)
		ok = ok.NextNumber
	}
}

type LinkedNumbers struct {
	Number     int
	NextNumber *LinkedNumbers
}

func AddTwoLinkedNumbers(firstNumbers *LinkedNumbers, secondNumbers *LinkedNumbers) *LinkedNumbers {
	var n int
	result := new(LinkedNumbers)
	linked := result
	for firstNumbers != nil || secondNumbers != nil || n != 0 {
		var (
			n1, n2 int
		)
		if firstNumbers != nil {
			n1, firstNumbers = firstNumbers.Number, firstNumbers.NextNumber
		}
		if secondNumbers != nil {
			n2, secondNumbers = secondNumbers.Number, secondNumbers.NextNumber
		}
		num := n1 + n2 + n
		n = num / 10
		linked.NextNumber = &LinkedNumbers{num % 10, nil}
		linked = linked.NextNumber
	}
	return result.NextNumber
}
