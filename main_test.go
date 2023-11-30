package main

import "testing"

func TestTahmin(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		target := "6453"
		guess := "6453"

		plus, minus := result(guess, target)

		if plus != 4 || minus != 0 {
			t.Errorf("Beklenen değerler: Artı=4, Eksi=0; Alınan değerler: Artı=%d, Eksi=%d", plus, minus)
		}
	})

	t.Run("Test2", func(t *testing.T) {
		target := "6453"
		guess := "1234"

		plus, minus := result(guess, target)

		if plus != 0 || minus != 2 {
			t.Errorf("Beklenen değerler: Artı=0, Eksi=2; Alınan değerler: Artı=%d, Eksi=%d", plus, minus)
		}
	})

}
