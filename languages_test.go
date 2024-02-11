package gofakeit

import (
	"fmt"
	"testing"
)

func ExampleLanguage() {
	Seed(11)
	fmt.Println(Language())

	// Output: Kazakh
}

func ExampleFaker_Language() {
	f := New(11)
	fmt.Println(f.Language())

	// Output: Kazakh
}

func BenchmarkLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Language()
	}
}

func ExampleLanguageAbbreviation() {
	Seed(11)
	fmt.Println(LanguageAbbreviation())

	// Output: kk
}

func ExampleFaker_LanguageAbbreviation() {
	f := New(11)
	fmt.Println(f.LanguageAbbreviation())

	// Output: kk
}

func BenchmarkLanguageAbbreviation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageAbbreviation()
	}
}

func ExampleLanguageBCP() {
	Seed(11)
	fmt.Println(LanguageBCP())

	// Output: de-DE
}

func ExampleFaker_LanguageBCP() {
	f := New(11)
	fmt.Println(f.LanguageBCP())

	// Output: de-DE
}

func BenchmarkLanguageBCP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LanguageBCP()
	}
}

func ExampleProgrammingLanguage() {
	Seed(464)
	fmt.Println(ProgrammingLanguage())

	// Output: Go
}

func ExampleFaker_ProgrammingLanguage() {
	f := New(464)
	fmt.Println(f.ProgrammingLanguage())

	// Output: Go
}

func BenchmarkProgrammingLanguage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProgrammingLanguage()
	}
}

func ExampleProgrammingLanguageBest() {
	Seed(11)
	fmt.Println(ProgrammingLanguageBest())

	// Output: Go
}

func ExampleFaker_ProgrammingLanguageBest() {
	f := New(11)
	fmt.Println(f.ProgrammingLanguageBest())

	// Output: Go
}

func BenchmarkProgrammingLanguageBest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProgrammingLanguageBest()
	}
}
