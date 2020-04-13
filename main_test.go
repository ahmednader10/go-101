package main_test

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"testing"
)

func BenchmarkSHA1(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha1.Sum(data)
	}
}

func BenchmarkSHA256(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha256.Sum256(data)
	}
}

func BenchmarkSHA512(b *testing.B) {
	data := []byte("Mary had a little lamb")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sha512.Sum512(data)
	}
}

func TestAddition(t *testing.T) {
	got := 2 + 2
	expected := 4

	if got != expected {
		t.Errorf("didn't get expected result. Got %v expected %v", got, expected)
	}
}

func TestSubtraction(t *testing.T) {
	got := 10 - 5
	expected := 5

	if got != expected {
		t.Errorf("didn't get expected result. Got %v expected %v", got, expected)
	}
}