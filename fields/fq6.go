package fields

import (
	"io"
	"math/big"

	fp "github.com/kilic/fp256"
)

type Fq6 struct {
	F          *Fq2
	NonResidue [2]*fp.FieldElement
}

func NewFq6(f *Fq2, nonResidue [2]*fp.FieldElement) Fq6 {
	fq6 := Fq6{
		f,
		nonResidue,
	}
	return fq6
}

func (fq6 Fq6) Zero() [3][2]*fp.FieldElement {
	return [3][2]*fp.FieldElement{fq6.F.Zero(), fq6.F.Zero(), fq6.F.Zero()}
}

func (fq6 Fq6) One() [3][2]*fp.FieldElement {
	return [3][2]*fp.FieldElement{fq6.F.One(), fq6.F.Zero(), fq6.F.Zero()}
}

func (fq6 Fq6) Equal(a, b [3][2]*fp.FieldElement) bool {
	return fq6.F.Equal(a[0], b[0]) && fq6.F.Equal(a[1], b[1]) && fq6.F.Equal(a[2], b[2])
}

func (fq6 Fq6) Copy(c, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	fq6.F.Copy(c[0], a[0])
	fq6.F.Copy(c[1], a[1])
	fq6.F.Copy(c[2], a[2])
	return c
}

func (fq6 Fq6) Demont(a [3][2]*fp.FieldElement) {
	fq6.F.Demont(a[0])
	fq6.F.Demont(a[1])
	fq6.F.Demont(a[2])
}

func (fq6 Fq6) NewElement() [3][2]*fp.FieldElement {
	return [3][2]*fp.FieldElement{
		fq6.F.NewElement(), fq6.F.NewElement(), fq6.F.NewElement()}
}

func (fq6 Fq6) rand(a [3][2]*fp.FieldElement, r io.Reader) [3][2]*fp.FieldElement {
	fq6.F.rand(a[0], r)
	fq6.F.rand(a[1], r)
	fq6.F.rand(a[2], r)
	return a
}

func (fq6 Fq6) mulByNonResidue(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	return fq6.F.Mul(c, fq6.NonResidue, a)
}

func (fq6 Fq6) Add(c, a, b [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	fq6.F.Add(c[0], a[0], b[0])
	fq6.F.Add(c[1], a[1], b[1])
	fq6.F.Add(c[2], a[2], b[2])
	return c
}

func (fq6 Fq6) Double(c, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	return fq6.Add(c, a, a)
}

func (fq6 Fq6) Sub(c, a, b [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	fq6.F.Sub(c[0], a[0], b[0])
	fq6.F.Sub(c[1], a[1], b[1])
	fq6.F.Sub(c[2], a[2], b[2])
	return c
}

func (fq6 Fq6) Neg(c, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	return fq6.Sub(c, fq6.Zero(), a)
}

func (fq6 Fq6) Mul(c, a, b [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	v0 := fq6.F.NewElement()
	v1 := fq6.F.NewElement()
	v2 := fq6.F.NewElement()
	v3 := fq6.F.NewElement()
	v4 := fq6.F.NewElement()
	v5 := fq6.F.NewElement()
	//
	fq6.F.Mul(v0, a[0], b[0])
	fq6.F.Mul(v1, a[1], b[1])
	fq6.F.Mul(v2, a[2], b[2])
	//
	fq6.F.Add(v3, a[1], a[2])
	fq6.F.Add(v4, b[1], b[2])
	fq6.F.Mul(v3, v3, v4)
	fq6.F.Add(v4, v1, v2)
	fq6.F.Sub(v3, v3, v4)
	fq6.mulByNonResidue(v3, v3)
	fq6.F.Add(v5, v0, v3)
	//
	fq6.F.Add(v3, a[0], a[1])
	fq6.F.Add(v4, b[0], b[1])
	fq6.F.Mul(v3, v3, v4)
	fq6.F.Add(v4, v0, v1)
	fq6.F.Sub(v3, v3, v4)
	fq6.mulByNonResidue(v4, v2)
	fq6.F.Add(c[1], v3, v4)
	//
	fq6.F.Add(v3, a[0], a[2])
	fq6.F.Add(v4, b[0], b[2])
	fq6.F.Mul(v3, v3, v4)
	fq6.F.Add(v4, v0, v2)
	fq6.F.Sub(v3, v3, v4)
	fq6.F.Add(c[2], v1, v3)
	fq6.F.Copy(c[0], v5)
	return c
}

func (fq6 Fq6) Square(c, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	t0 := fq6.F.NewElement()
	t1 := fq6.F.NewElement()
	t2 := fq6.F.NewElement()
	t3 := fq6.F.NewElement()
	t4 := fq6.F.NewElement()
	t5 := fq6.F.NewElement()
	//
	fq6.F.Square(t0, a[0])
	fq6.F.Mul(t1, a[0], a[1])
	fq6.F.Add(t1, t1, t1)
	fq6.F.Sub(t2, a[0], a[1])
	fq6.F.Add(t2, t2, a[2])
	fq6.F.Square(t2, t2)
	fq6.F.Mul(t3, a[1], a[2])
	fq6.F.Add(t3, t3, t3)
	fq6.F.Square(t4, a[2])
	//
	fq6.mulByNonResidue(t5, t3)
	fq6.F.Add(c[0], t0, t5)
	//
	fq6.mulByNonResidue(t5, t4)
	fq6.F.Add(c[1], t1, t5)
	//
	fq6.F.Add(t1, t1, t2)
	fq6.F.Add(t1, t1, t3)
	fq6.F.Add(t0, t0, t4)
	fq6.F.Sub(c[2], t1, t0)
	//
	return c
}

func (fq6 Fq6) MulScalar(base [3][2]*fp.FieldElement, e *big.Int) [3][2]*fp.FieldElement {

	res := fq6.Zero()
	rem := e
	exp := base
	zero := big.NewInt(int64(0))
	for rem.Cmp(zero) != 0 {
		if rem.Bit(1) == 1 {
			fq6.Add(res, res, exp)
		}
		fq6.Double(exp, exp)
		rem.Rsh(rem, 1)
	}
	return res
}

func (fq6 Fq6) Inverse(c, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	t0 := fq6.F.NewElement()
	t1 := fq6.F.NewElement()
	t2 := fq6.F.NewElement()
	t3 := fq6.F.NewElement()
	t4 := fq6.F.NewElement()
	//
	fq6.F.Square(t0, a[0])
	fq6.F.Mul(t1, a[1], a[2])
	fq6.mulByNonResidue(t1, t1)
	fq6.F.Sub(t0, t0, t1)
	//
	fq6.F.Square(t1, a[1])
	fq6.F.Mul(t2, a[0], a[2])
	fq6.F.Sub(t1, t1, t2)
	//
	fq6.F.Square(t2, a[2])
	fq6.mulByNonResidue(t2, t2)
	fq6.F.Mul(t3, a[0], a[1])
	fq6.F.Sub(t2, t2, t3)
	//
	fq6.F.Mul(t3, a[2], t2)
	fq6.F.Mul(t4, a[1], t1)
	fq6.F.Add(t3, t3, t4)
	fq6.mulByNonResidue(t3, t3)
	fq6.F.Mul(t4, a[0], t0)
	fq6.F.Add(t3, t3, t4)
	fq6.F.Inverse(t3, t3)
	//
	fq6.F.Mul(c[0], t0, t3)
	fq6.F.Mul(c[1], t2, t3)
	fq6.F.Mul(c[2], t1, t3)
	//
	return c
}

func (fq6 Fq6) Div(c, a, b [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	t0 := fq6.NewElement()
	fq6.Inverse(t0, b)
	return fq6.Mul(c, a, t0)
}
