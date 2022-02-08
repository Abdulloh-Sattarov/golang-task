package main

import (
	"reflect"
	"testing"
)

func Test_AddTwoLinkedNumbers(t *testing.T) {

	tests := []struct {
		testName     string
		firstLinked  *LinkedNumbers
		secondLinked *LinkedNumbers
		want         []int
	}{
		{
			testName:     "first",
			firstLinked:  &LinkedNumbers{2, &LinkedNumbers{4, &LinkedNumbers{3, nil}}},
			secondLinked: &LinkedNumbers{5, &LinkedNumbers{6, &LinkedNumbers{4, nil}}},
			want:         []int{7, 0, 8},
		},
		{
			testName:     "second",
			firstLinked:  &LinkedNumbers{0, nil},
			secondLinked: &LinkedNumbers{0, nil},
			want:         []int{0},
		},
		{
			testName:     "third",
			firstLinked:  &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, nil}}}}}}},
			secondLinked: &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, &LinkedNumbers{9, nil}}}},
			want:         []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tc := range tests {
		t.Run(tc.testName, func(t *testing.T) {
			got := AddTwoLinkedNumbers(tc.firstLinked, tc.secondLinked)
			var result []int
			for got != nil {
				result = append(result, got.Number)
				got = got.NextNumber
			}
			if !reflect.DeepEqual(tc.want, result) {
				t.Fatalf("%s: expected: %v, got: %v", tc.testName, tc.want, result)
			}
		})
	}
}
