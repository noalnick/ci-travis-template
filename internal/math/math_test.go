package math

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}

func TestSubs(t *testing.T) {
	total := Substract(10, 5)
    if total != 5 {
       t.Errorf("Substract was incorrect, got: %d, want: %d.", total, 5)
    }
}

func TestSquare(t *testing.T) {
	total := Square(5)
    if total != 20 {
       t.Errorf("Square was incorrect, got: %d, want: %d.", total, 25)
    }
}

func TestMulti(t *testing.T) {
	total := Mult(10, 10)
    if total != 100 {
       t.Errorf("Mult was incorrect, got: %d, want: %d.", total, 100)
    }
}