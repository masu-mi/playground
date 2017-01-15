package main

import "testing"

func Fixture() []int {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
}

func BenchmarkCompareSlices_Reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareSlices_Reflect(Fixture(), Fixture())
	}
}

func BenchmarkCompareSlices_General(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareSlices_General(Fixture(), Fixture())
	}
}

func BenchmarkCompareSlices_BCE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareSlices_BCE(Fixture(), Fixture())
	}
}

func BenchmarkCompareSlices_Pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareSlices_Pointer(Fixture(), Fixture())
	}
}

func BenchmarkCompareSlices_PointerAndBCE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareSlices_PointerAndBCE(Fixture(), Fixture())
	}
}
