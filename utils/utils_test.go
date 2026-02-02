package utils

import (
	"reflect"
	"testing"
)

func TestAbsIntNegative(t *testing.T) {
	inputInt := -5
	want := 5
	outputInt := AbsInt(inputInt)
	if !(want == outputInt) {
		t.Errorf(`AbsInt(-5) = %d, want match for %q`, outputInt, want)
	}
}

func TestAbsIntPositive(t *testing.T) {
	inputInt := 5
	want := 5
	outputInt := AbsInt(inputInt)
	if !(want == outputInt) {
		t.Errorf(`AbsInt(5) = %d, want match for %q`, outputInt, want)
	}
}

func TestAbsDiffIntNegative(t *testing.T) {
	inputInt := -5
	want := 7
	outputInt := AbsDiffInt(inputInt, 2)
	if !(want == outputInt) {
		t.Errorf(`AbsInt(-5) = %d, want match for %q`, outputInt, want)
	}
}

func TestFlattenSlice(t *testing.T) {
	inputSlice := []string{"yea", "g", "hghg"}
	want := "yeaghghg"
	outputString := FlattenSlice(inputSlice)
	if !(want == outputString) {
		t.Errorf(`FlattenSlice([]string{"yea", "g", "hghg"}) = %v, want match for %q`, outputString, want)
	}
}

func TestRangeOverlapsSorted(t *testing.T) {
	rangeA := MinMaxRange{5, 10}
	rangeB := MinMaxRange{7, 12}
	want := true
	outputBool, _ := RangeOverlapsSorted(rangeA, rangeB)
	if !(want == outputBool) {
		t.Errorf(`RangeOverlapsSorted = %v, want match for %t`, outputBool, want)

	}
}

// func TestRangeOverlapsUnSorted(t *testing.T) {
// 	rangeA := MinMaxRange{7, 12}
// 	rangeB := MinMaxRange{5, 10}
// 	want :=
// 	outputBool, _ := RangeOverlapsSorted(rangeA, rangeB)
// 	if !(want == outputBool) {
// 		t.Errorf(`RangeOverlapsSorted = %v, want match for %t`, outputBool, want)

// 	}
// }

func TestSplitStrIntoArbitraryLength(t *testing.T) {
	inputStr := "123123123123"
	want := []string{"123", "123", "123", "123"}
	outputSlice, _ := SplitStrIntoArbitraryLength(inputStr, 3)
	if !reflect.DeepEqual(want, outputSlice) {
		t.Errorf(`RangeOverlapsSorted = %v, want match for %v`, outputSlice, want)
	}
}

func TestSplitStrIntoArbitraryLengthDouble(t *testing.T) {
	inputStr := "11"
	want := []string{"1", "1"}
	outputSlice, _ := SplitStrIntoArbitraryLength(inputStr, 1)
	if !reflect.DeepEqual(want, outputSlice) {
		t.Errorf(`RangeOverlapsSorted = %v, want match for %v`, outputSlice, want)
	}
}
