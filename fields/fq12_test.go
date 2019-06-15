package fields

import (
	"crypto/rand"
	"testing"
)

func TestFq12AddAssoc(t *testing.T) {
	a, b, c, u, v := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.rand(c, rand.Reader)
		fq12.Add(u, a, b)
		fq12.Add(u, u, c)
		fq12.Add(v, b, c)
		fq12.Add(v, v, a)
		if !fq12.Equal(u, v) {
			t.Errorf("additive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq12AddZero(t *testing.T) {
	a, c := fq12.NewElement(), fq12.NewElement()
	zero := fq12.Zero()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Add(c, a, zero)
		if !fq12.Equal(c, a) {
			t.Errorf("zero addition fails\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq12SubAssoc(t *testing.T) {
	a, b, c, u, v := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.rand(c, rand.Reader)
		fq12.Sub(u, a, c)
		fq12.Sub(u, u, b)
		fq12.Sub(v, a, b)
		fq12.Sub(v, v, c)
		if !fq12.Equal(u, v) {
			t.Errorf("subtrctive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq12MulAssoc(t *testing.T) {
	a, b, c, u, v := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.rand(c, rand.Reader)
		fq12.Mul(u, a, b)
		fq12.Mul(u, u, c)
		fq12.Mul(v, a, c)
		fq12.Mul(v, v, b)
		if !fq12.Equal(u, v) {
			t.Errorf("multiplicative associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq12MulId(t *testing.T) {
	a, c := fq12.NewElement(), fq12.NewElement()
	one := fq12.One()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Mul(c, a, one)
		if !fq12.Equal(c, a) {
			t.Errorf("multiplicative identity does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq12MulZero(t *testing.T) {
	a, c := fq12.NewElement(), fq12.NewElement()
	zero := fq12.Zero()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Mul(c, a, zero)
		if !fq12.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
		fq12.Mul(c, zero, a)
		if !fq12.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq12MulComm(t *testing.T) {
	a, b, u, v := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.Mul(u, a, b)
		fq12.Mul(v, b, a)
		if !fq12.Equal(u, v) {
			t.Errorf("multiplicative commutativity does not hold \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq12MulDist(t *testing.T) {
	a, b, c, d, e := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.rand(c, rand.Reader)
		fq12.Mul(d, a, c)
		fq12.Mul(e, b, c)
		fq12.Add(d, d, e)
		fq12.Add(e, a, b)
		fq12.Mul(e, e, c)
		if !fq12.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}

func TestFq12Square(t *testing.T) {
	a, b, c := fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Mul(b, a, a)
		fq12.Square(c, a)
		if !fq12.Equal(c, b) {
			t.Errorf("squaring failed \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq12Neg(t *testing.T) {
	a, b, u, v := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.Sub(u, a, b)
		fq12.Neg(a, a)
		fq12.Neg(b, b)
		fq12.Sub(v, b, a)
		if !fq12.Equal(u, v) {
			t.Errorf("negation failed \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq12Neg2(t *testing.T) {
	a, b := fq12.NewElement(), fq12.NewElement()
	zero := fq12.Zero()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Neg(b, a)
		fq12.Add(a, b, a)
		if !fq12.Equal(zero, a) {
			t.Errorf("negation failed\na:\n%v\na:\n%v\n", a, b)
			return
		}
	}
}

func TestFq12Inv(t *testing.T) {
	a, b, c := fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.Mul(c, a, b)
		fq12.Inverse(b, b)
		fq12.Mul(c, c, b)
		if !fq12.Equal(c, a) {
			t.Errorf("invertion fails \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq12Div(t *testing.T) {
	a, c := fq12.NewElement(), fq12.NewElement()
	one := fq12.One()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.Div(c, a, a)
		if !fq12.Equal(c, one) {
			t.Errorf("division fails \na:\n%v\nc:\n%v\n", a, c)
			return
		}
	}
}

func TestFq12DivDist(t *testing.T) {
	a, b, c, d, e := fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement(), fq12.NewElement()
	for i := 0; i < nBox; i++ {
		fq12.rand(a, rand.Reader)
		fq12.rand(b, rand.Reader)
		fq12.rand(c, rand.Reader)
		fq12.Div(d, a, c)
		fq12.Div(e, b, c)
		fq12.Add(d, d, e)
		fq12.Add(e, a, b)
		fq12.Div(e, e, c)
		if !fq12.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}
