package main

import "reflect"

func CompareSlices_Reflect(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}
func CompareSlices_General(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CompareSlices_BCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)] // this line is the key
	for i, v := range a {
		if v != b[i] { // here is no bounds checking for b[i]
			return false
		}
	}

	return true
}

func CompareSlices_Pointer(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	if &a[0] == &b[0] {
		return true // early exit
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CompareSlices_PointerAndBCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) == 0 {
		return true
	}
	b = b[:len(a)]
	if &a[0] == &b[0] {
		return true
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
