package fields

import (
	"crypto/rand"
	"testing"
)

func TestFq2AddAssoc(t *testing.T) {
	a, b, c, u, v := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.rand(c, rand.Reader)
		fq2.Add(u, a, b)
		fq2.Add(u, u, c)
		fq2.Add(v, b, c)
		fq2.Add(v, v, a)
		if !fq2.Equal(u, v) {
			t.Errorf("additive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq2AddZero(t *testing.T) {
	a, c := fq2.NewElement(), fq2.NewElement()
	zero := fq2.Zero()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Add(c, a, zero)
		if !fq2.Equal(c, a) {
			t.Errorf("zero addition fails\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq2SubAssoc(t *testing.T) {
	a, b, c, u, v := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.rand(c, rand.Reader)
		fq2.Sub(u, a, c)
		fq2.Sub(u, u, b)
		fq2.Sub(v, a, b)
		fq2.Sub(v, v, c)
		if !fq2.Equal(u, v) {
			t.Errorf("subtrctive associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq2MulAssoc(t *testing.T) {
	a, b, c, u, v := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.rand(c, rand.Reader)
		fq2.Mul(u, a, b)
		fq2.Mul(u, u, c)
		fq2.Mul(v, a, c)
		fq2.Mul(v, v, b)
		if !fq2.Equal(u, v) {
			t.Errorf("multiplicative associativity does not hold\na:\n%v\nb:\n%v\nc:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, c, u, v)
			return
		}
	}
}

func TestFq2MulId(t *testing.T) {
	a, c := fq2.NewElement(), fq2.NewElement()
	one := fq2.One()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Mul(c, a, one)
		if !fq2.Equal(c, a) {
			t.Errorf("multiplicative identity does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq2MulZero(t *testing.T) {
	a, c := fq2.NewElement(), fq2.NewElement()
	zero := fq2.Zero()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Mul(c, a, zero)
		if !fq2.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
		fq2.Mul(c, zero, a)
		if !fq2.Equal(c, zero) {
			t.Errorf("multiplicative zero does not hold\na:\n%v\nc:\n%v\n",
				a, c)
			return
		}
	}
}

func TestFq2MulComm(t *testing.T) {
	a, b, u, v := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.Mul(u, a, b)
		fq2.Mul(v, b, a)
		if !fq2.Equal(u, v) {
			t.Errorf("multiplicative commutativity does not hold \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq2MulDist(t *testing.T) {
	a, b, c, d, e := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.rand(c, rand.Reader)
		fq2.Mul(d, a, c)
		fq2.Mul(e, b, c)
		fq2.Add(d, d, e)
		fq2.Add(e, a, b)
		fq2.Mul(e, e, c)
		if !fq2.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}

func TestFq2Square(t *testing.T) {
	a, b, c := fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Mul(b, a, a)
		fq2.Square(c, a)
		if !fq2.Equal(c, b) {
			t.Errorf("squaring failed \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq2Neg(t *testing.T) {
	a, b, u, v := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.Sub(u, a, b)
		fq2.Neg(a, a)
		fq2.Neg(b, b)
		fq2.Sub(v, b, a)
		if !fq2.Equal(u, v) {
			t.Errorf("negation failed \na:\n%v\nb:\n%v\nu:\n%v\nv:\n%v\n",
				a, b, u, v)
			return
		}
	}
}

func TestFq2Neg2(t *testing.T) {
	a, b := fq2.NewElement(), fq2.NewElement()
	zero := fq2.Zero()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Neg(b, a)
		fq2.Add(a, b, a)
		if !fq2.Equal(zero, a) {
			t.Errorf("negation failed\na:\n%v\na:\n%v\n", a, b)
			return
		}
	}
}

func TestFq2Inv(t *testing.T) {
	a, b, c := fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.Mul(c, a, b)
		fq2.Inverse(b, b)
		fq2.Mul(c, c, b)
		if !fq2.Equal(c, a) {
			t.Errorf("invertion fails \na:\n%v\nb:\n%v\nc:\n%v\n", a, b, c)
			return
		}
	}
}

func TestFq2Div(t *testing.T) {
	a, c := fq2.NewElement(), fq2.NewElement()
	one := fq2.One()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.Div(c, a, a)
		if !fq2.Equal(c, one) {
			t.Errorf("division fails \na:\n%v\nc:\n%v\n", a, c)
			return
		}
	}
}

func TestFq2DivDist(t *testing.T) {
	a, b, c, d, e := fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement(), fq2.NewElement()
	for i := 0; i < nBox; i++ {
		fq2.rand(a, rand.Reader)
		fq2.rand(b, rand.Reader)
		fq2.rand(c, rand.Reader)
		fq2.Div(d, a, c)
		fq2.Div(e, b, c)
		fq2.Add(d, d, e)
		fq2.Add(e, a, b)
		fq2.Div(e, e, c)
		if !fq2.Equal(d, e) {
			t.Errorf("division fails \nd:\n%v\nd:\n%v\n", d, e)
			return
		}
	}
}
