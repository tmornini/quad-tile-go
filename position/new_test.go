package position

import "testing"

func TestNew(t *testing.T) {
	_, err := New(RandomData())

	if err != nil {
		t.Error("valid position should not return an error")
	}

	_, err = New(-91, 0, 0)

	if err == nil {
		t.Error("latitude less than -90 should return an error")
	}

	_, err = New(91, 0, 0)

	if err == nil {
		t.Error("latitude greater than 90 should return an error")
	}

	_, err = New(0, -180, 0)

	if err == nil {
		t.Error("longitude less than or equal to -180 should return an error")
	}

	_, err = New(0, 181, 0)

	if err == nil {
		t.Error("longitude greater than 180 should return an error")
	}
}

func BenchmarkNew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		New(RandomData())
	}
}
