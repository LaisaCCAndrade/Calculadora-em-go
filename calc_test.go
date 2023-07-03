package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 150

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	teste := subtrai(10, 5)
	resultado := -5

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubIncorrect(t *testing.T) {
	teste := subtrai(10, 10)
	resultado := 5

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	teste := dividi(10)
	resultado := 10

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivIncorrect(t *testing.T) {
	teste := dividi(10)
	resultado := 15

	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
