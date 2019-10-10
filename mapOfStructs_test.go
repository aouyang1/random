package random

import (
	"testing"
)

func TestMapOfStructs(t *testing.T) {
	m := initMapOfStructs()
	fillMapOfStructs(m)
	if len(m) != 1000 {
		t.Fatal("Expected 1000 keys")
	}

	for _, v := range m {
		if len(v.Items) != 4800 {
			t.Fatal("Expected 480 items")
		}
		for i := 0; i < len(v.Items); i++ {
			if i != int(v.Items[i]) {
				t.Fatalf("Expected %d, but got %.3f", i, v.Items[i])
			}
		}
	}
}

func TestMapOfStructsReferenceKey(t *testing.T) {
	m := initMapOfStructs()
	fillMapOfStructsReferenceKey(m)
	if len(m) != 1000 {
		t.Fatal("Expected 1000 keys")
	}

	for _, v := range m {
		if len(v.Items) != 4800 {
			t.Fatal("Expected 480 items")
		}
		for i := 0; i < len(v.Items); i++ {
			if i != int(v.Items[i]) {
				t.Fatalf("Expected %d, but got %.3f", i, v.Items[i])
			}
		}
	}
}

func TestMapOfStructsPtr(t *testing.T) {
	m := initMapOfStructsPtr()
	fillMapOfStructsPtr(m)
	if len(m) != 1000 {
		t.Fatal("Expected 1000 keys")
	}

	for _, v := range m {
		if len(v.Items) != 4800 {
			t.Fatal("Expected 480 items")
		}
		expectedSum := float64((len(v.Items) - 1)) * float64(len(v.Items)) / 2.0
		if v.Sum != expectedSum {
			t.Fatalf("Expected sum of %f, but got %f", expectedSum, v.Sum)
		}

		for i := 0; i < len(v.Items); i++ {
			if i != int(v.Items[i]) {
				t.Fatalf("Expected %d, but got %.3f", i, v.Items[i])
			}
		}
	}
}

func TestMapOfStructsPtrReferenceKey(t *testing.T) {
	m := initMapOfStructsPtr()
	fillMapOfStructsPtrReferenceKey(m)
	if len(m) != 1000 {
		t.Fatal("Expected 1000 keys")
	}

	for _, v := range m {
		if len(v.Items) != 4800 {
			t.Fatal("Expected 480 items")
		}
		expectedSum := float64((len(v.Items) - 1)) * float64(len(v.Items)) / 2.0
		if v.Sum != expectedSum {
			t.Fatalf("Expected sum of %f, but got %f", expectedSum, v.Sum)
		}

		for i := 0; i < len(v.Items); i++ {
			if i != int(v.Items[i]) {
				t.Fatalf("Expected %d, but got %.3f", i, v.Items[i])
			}
		}
	}
}

func BenchmarkMapOfStructsSimple(b *testing.B) {
	m := initMapOfStructsSimple()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsSimple(m)
	}
}

func BenchmarkMapOfStructsSimpleReferenceKey(b *testing.B) {
	m := initMapOfStructsSimple()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsSimpleReferenceKey(m)
	}
}

func BenchmarkMapOfStructsSimplePtr(b *testing.B) {
	m := initMapOfStructsSimplePtr()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsSimplePtr(m)
	}
}

func BenchmarkMapOfStructsSimplePtrReferenceKey(b *testing.B) {
	m := initMapOfStructsSimplePtr()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsSimplePtrReferenceKey(m)
	}
}

func BenchmarkMapOfStructs(b *testing.B) {
	m := initMapOfStructs()
	for i := 0; i < b.N; i++ {
		fillMapOfStructs(m)
	}
}

func BenchmarkMapOfStructsReferenceKey(b *testing.B) {
	m := initMapOfStructs()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsReferenceKey(m)
	}
}

func BenchmarkMapOfStructsPtr(b *testing.B) {
	m := initMapOfStructsPtr()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsPtr(m)
	}
}

func BenchmarkMapOfStructsPtrReferenceKey(b *testing.B) {
	m := initMapOfStructsPtr()
	for i := 0; i < b.N; i++ {
		fillMapOfStructsPtrReferenceKey(m)
	}
}
