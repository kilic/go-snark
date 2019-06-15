package fields

import (
	"crypto/rand"
	"math/big"
	"testing"

	fp "github.com/kilic/fp256"
)

func TestFq6AdditiveAssoc(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	u := fq6.NewElement()
	v := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Add(u, a, b)
		fq6.Add(u, u, c)
		fq6.Add(v, b, c)
		fq6.Add(v, v, a)
		if !fq6.Equal(u, v) {
			t.Errorf("additive associativity does not hold a:%x, b:%x, c:%x, u:%x, v:%x",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq6SubtractiveAssoc(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	u := fq6.NewElement()
	v := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Sub(u, a, c)
		fq6.Sub(u, u, b)
		fq6.Sub(v, a, b)
		fq6.Sub(v, v, c)
		if !fq6.Equal(u, v) {
			t.Errorf("additive associativity does not hold a:%x, b:%x, c:%x, u:%x, v:%x",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq6MultiplicativeAssoc(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	u := fq6.NewElement()
	v := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Mul(u, a, b)
		fq6.Mul(u, u, c)
		fq6.Mul(v, a, c)
		fq6.Mul(v, v, b)
		if !fq6.Equal(u, v) {
			t.Errorf("multiplicative associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}
func TestFq6MultiplicativeCommutativity(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	u := fq6.NewElement()
	v := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.Mul(u, a, b)
		fq6.Mul(v, b, a)
		if !fq6.Equal(u, v) {
			t.Errorf("multiplicative commutativity does not hold \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq6Squaring(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Mul(b, a, a)
		fq6.Square(c, a)
		if !fq6.Equal(c, b) {
			t.Errorf("multiplicative commutativity does not hold \na:\n%v\nb:\n%v\nc:\n%v\n",
				a, b, c)
			return
		}
	}
}

func TestFq6Neg(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	u := fq6.NewElement()
	v := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.Sub(u, a, b)
		fq6.Neg(a, a)
		fq6.Neg(b, b)
		fq6.Sub(v, b, a)
		if !fq6.Equal(u, v) {
			t.Errorf("negation failed \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq6Neg2(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	zero := fq6.Zero()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Neg(b, a)
		fq6.Add(a, b, a)
		if !fq6.Equal(zero, a) {
			t.Errorf("negation fails \na:\n%v\n", a)
			return
		}
	}
}

func TestFq6Inv(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.Mul(c, a, b)
		fq6.Inverse(b, b)
		fq6.Mul(c, c, b)
		if !fq6.Equal(c, a) {
			t.Errorf("invertion fails \na:\n%v\nc:\n%v\n", a, c)
			return
		}
	}
}

func TestFq6Div(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	c := fq6.NewElement()
	one := fq6.One()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Div(c, a, a)
		if !fq6.Equal(c, one) {
			t.Errorf("division fails \nc:\n%v\n", c)
			return
		}
	}
}

func TestFq6Div2(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 := Fq6{fq2, nonResidueFq6}
	a := fq6.NewElement()
	b := fq6.NewElement()
	c := fq6.NewElement()
	d := fq6.NewElement()
	e := fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Div(d, a, c)
		fq6.Div(e, b, c)
		fq6.Add(d, d, e)
		fq6.Add(e, a, b)
		fq6.Div(e, e, c)
		if !fq6.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\ne:\n%v\n", d, e)
			return
		}
	}
}
