package fields

import (
	"crypto/rand"
	"testing"
)

func TestFq6AddAssoc(t *testing.T) {
	a, b, c, u, v := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Add(u, a, b)
		fq6.Add(u, u, c)
		fq6.Add(v, b, c)
		fq6.Add(v, v, a)
		if !fq6.Equal(u, v) {
			t.Errorf("additive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq6AddZero(t *testing.T) {
	a, c := fq6.NewElement(), fq6.NewElement()
	zero := fq6.Zero()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Add(c, a, zero)
		if !fq6.Equal(c, a) {
			t.Errorf("zero addition fails\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq6SubAssoc(t *testing.T) {
	a, b, c, u, v := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Sub(u, a, c)
		fq6.Sub(u, u, b)
		fq6.Sub(v, a, b)
		fq6.Sub(v, v, c)
		if !fq6.Equal(u, v) {
			t.Errorf("subtractive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq6MulAssoc(t *testing.T) {
	a, b, c, u, v := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
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

func TestFq6MulId(t *testing.T) {
	a, c := fq6.NewElement(), fq6.NewElement()
	one := fq6.One()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Mul(c, a, one)
		if !fq6.Equal(c, a) {
			t.Errorf("multiplicative identity does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq6MulZero(t *testing.T) {
	a, c := fq6.NewElement(), fq6.NewElement()
	zero := fq6.Zero()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Mul(c, a, zero)
		if !fq6.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
		fq6.Mul(c, zero, a)
		if !fq6.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}
func TestFq6MulComm(t *testing.T) {
	a, b, u, v := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
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

func TestFq6MulDist(t *testing.T) {
	a, b, c, d, e := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.rand(c, rand.Reader)
		fq6.Mul(d, a, c)
		fq6.Mul(e, b, c)
		fq6.Add(d, d, e)
		fq6.Add(e, a, b)
		fq6.Mul(e, e, c)
		if !fq6.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}

func TestFq6Square(t *testing.T) {
	a, b, c := fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Mul(b, a, a)
		fq6.Square(c, a)
		if !fq6.Equal(c, b) {
			t.Errorf("squaring failed \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq6Neg(t *testing.T) {
	a, b, u, v := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
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
	a, b := fq6.NewElement(), fq6.NewElement()
	zero := fq6.Zero()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Neg(b, a)
		fq6.Add(a, b, a)
		if !fq6.Equal(zero, a) {
			t.Errorf("negation failed\na:\n%v\na:\n%v\n", a, b)
			return
		}
	}
}

func TestFq6Inv(t *testing.T) {
	a, b, c := fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.rand(b, rand.Reader)
		fq6.Mul(c, a, b)
		fq6.Inverse(b, b)
		fq6.Mul(c, c, b)
		if !fq6.Equal(c, a) {
			t.Errorf("invertion fails \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq6Div(t *testing.T) {
	a, c := fq6.NewElement(), fq6.NewElement()
	one := fq6.One()
	for i := 0; i < nBox; i++ {
		fq6.rand(a, rand.Reader)
		fq6.Div(c, a, a)
		if !fq6.Equal(c, one) {
			t.Errorf("division fails \na:\n%v\nc:\n%v\n", a, c)
			return
		}
	}
}

func TestFq6DivDist(t *testing.T) {
	a, b, c, d, e := fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement(), fq6.NewElement()
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
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}
