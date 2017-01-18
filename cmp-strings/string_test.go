package main

import "testing"

func FixtureString() string {
	return "hello world"
}

func BenchmarkCompareString_Reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_Reflect(FixtureString(), FixtureString())
	}
}
func BenchmarkCompareString_Simple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_Simple(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_General(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_General(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_BCE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_BCE(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_Pointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_Pointer(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_PointerAndBCE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_PointerAndBCE(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_GeneralNoNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_GeneralNoNil(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_BCENoNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_BCENoNil(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_PointerNoNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_PointerNoNil(FixtureString(), FixtureString())
	}
}

func BenchmarkCompareString_PointerAndBCENoNil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompareString_PointerAndBCENoNil(FixtureString(), FixtureString())
	}
}
