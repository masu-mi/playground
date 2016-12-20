package main

import "reflect"

func CompareString_Reflect(a, b string) bool {
	return reflect.DeepEqual(a, b)
}
func CompareString_Simple(a, b string) bool {
	return a == b
}
func CompareString_General(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
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

func CompareString_BCE(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
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

func CompareString_Pointer(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
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

func CompareString_PointerAndBCE(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
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

func CompareString_GeneralNoNil(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func CompareString_BCENoNil(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
	if len(a) != len(b) {
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

func CompareString_PointerNoNil(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
	if len(a) != len(b) {
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

func CompareString_PointerAndBCENoNil(as, bs string) bool {
	a, b := []byte(as), []byte(bs)
	if len(a) != len(b) {
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
