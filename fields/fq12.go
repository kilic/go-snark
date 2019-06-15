package fields

import (
	"io"
	"math/big"

	fp "github.com/kilic/fp256"
)

type Fq12 struct {
	F          *Fq6
	Fq2        *Fq2
	NonResidue [2]*fp.FieldElement
}

func NewFq12(f *Fq6, fq2 *Fq2, nonResidue [2]*fp.FieldElement) Fq12 {
	fq12 := Fq12{
		f,
		fq2,
		nonResidue,
	}
	return fq12
}

func (fq12 Fq12) Zero() [2][3][2]*fp.FieldElement {
	return [2][3][2]*fp.FieldElement{fq12.F.Zero(), fq12.F.Zero()}
}

func (fq12 Fq12) One() [2][3][2]*fp.FieldElement {
	return [2][3][2]*fp.FieldElement{fq12.F.One(), fq12.F.Zero()}
}

func (fq12 Fq12) Equal(a, b [2][3][2]*fp.FieldElement) bool {
	return fq12.F.Equal(a[0], b[0]) && fq12.F.Equal(a[1], b[1])
}

func (fq12 Fq12) Copy(c, a [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	fq12.F.Copy(c[0], a[0])
	fq12.F.Copy(c[1], a[1])
	return c
}

func (fq12 Fq12) Demont(a [2][3][2]*fp.FieldElement) {
	fq12.F.Demont(a[0])
	fq12.F.Demont(a[1])
}

func (fq12 Fq12) NewElement() [2][3][2]*fp.FieldElement {
	return [2][3][2]*fp.FieldElement{
		fq12.F.NewElement(), fq12.F.NewElement()}
}

func (fq12 Fq12) rand(a [2][3][2]*fp.FieldElement, r io.Reader) [2][3][2]*fp.FieldElement {
	fq12.F.rand(a[0], r)
	fq12.F.rand(a[1], r)
	return a
}

func (fq12 Fq12) mulByNonResidue(c [3][2]*fp.FieldElement, a [3][2]*fp.FieldElement) [3][2]*fp.FieldElement {
	t := fq12.Fq2.NewElement()
	fq12.Fq2.Mul(t, fq12.NonResidue, a[2])
	fq12.Fq2.Copy(c[2], a[1])
	fq12.Fq2.Copy(c[1], a[0])
	fq12.Fq2.Copy(c[0], t)
	return c
}

func (fq12 Fq12) Add(c, a, b [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	fq12.F.Add(c[0], a[0], b[0])
	fq12.F.Add(c[1], a[1], b[1])
	return c
}

func (fq12 Fq12) Sub(c, a, b [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	fq12.F.Sub(c[0], a[0], b[0])
	fq12.F.Sub(c[1], a[1], b[1])
	return c
}

func (fq12 Fq12) Double(c, a [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	fq12.Add(c, a, a)
	return c
}

func (fq12 Fq12) Neg(c, a [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	fq12.Sub(c, fq12.Zero(), a)
	return c
}

func (fq12 Fq12) Mul(c, a, b [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	t0 := fq12.F.NewElement()
	t1 := fq12.F.NewElement()
	t2 := fq12.F.NewElement()
	t3 := fq12.F.NewElement()
	fq12.F.Mul(t0, a[0], b[0])
	fq12.F.Mul(t1, a[1], b[1])
	fq12.F.Add(t2, t0, t1)
	fq12.mulByNonResidue(t1, t1)
	fq12.F.Add(t0, t0, t1)
	fq12.F.Add(t1, a[0], a[1])
	fq12.F.Add(t3, b[0], b[1])
	fq12.F.Mul(t1, t1, t3)
	fq12.F.Sub(c[1], t1, t2)
	fq12.F.Copy(c[0], t0)
	return c
}

func (fq12 Fq12) MulScalar(res, base [2][3][2]*fp.FieldElement, e *big.Int) [2][3][2]*fp.FieldElement {

	fq12.Copy(res, fq12.Zero())
	rem := new(big.Int).SetBytes(e.Bytes())
	exp := fq12.NewElement()
	fq12.Copy(exp, base)
	zero := new(big.Int)

	for rem.Cmp(zero) != 0 {
		if rem.Bit(0) == 1 {
			fq12.Add(res, res, exp)
		}
		fq12.Double(exp, exp)
		rem.Rsh(rem, 1)
	}
	return res
}

func (fq12 Fq12) Inverse(c, a [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	t0 := fq12.F.NewElement()
	t1 := fq12.F.NewElement()
	fq12.F.Square(t0, a[0])
	fq12.F.Square(t1, a[1])
	fq12.mulByNonResidue(t1, t1)
	fq12.F.Sub(t0, t0, t1)
	fq12.F.Inverse(t0, t0)
	fq12.F.Mul(c[0], a[0], t0)
	fq12.F.Mul(t0, a[1], t0)
	fq12.F.Neg(c[1], t0)
	return c
}

func (fq12 Fq12) Div(c, a, b [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	t := fq12.NewElement()
	fq12.Inverse(t, b)
	return fq12.Mul(c, a, t)
}

func (fq12 Fq12) Square(c, a [2][3][2]*fp.FieldElement) [2][3][2]*fp.FieldElement {
	t0 := fq12.F.NewElement()
	t1 := fq12.F.NewElement()
	t2 := fq12.F.NewElement()
	fq12.F.Mul(t0, a[0], a[1])
	fq12.mulByNonResidue(t1, a[1])
	fq12.F.Add(t1, t1, a[0])
	fq12.F.Add(t2, a[0], a[1])
	fq12.F.Mul(t1, t1, t2)
	fq12.mulByNonResidue(t2, t0)
	fq12.F.Add(t2, t0, t2)
	fq12.F.Sub(c[0], t1, t2)
	fq12.F.Add(c[1], t0, t0)
	return c
}

func (fq12 Fq12) Exp(res, base [2][3][2]*fp.FieldElement, e *big.Int) [2][3][2]*fp.FieldElement {

	fq12.Copy(res, fq12.One())
	rem := new(big.Int).SetBytes(e.Bytes())
	exp := fq12.NewElement()
	fq12.Copy(exp, base)
	zero := new(big.Int)

	for rem.Cmp(zero) != 0 {
		if rem.Bit(0) == 1 {
			fq12.Mul(res, res, exp)
		}
		fq12.Square(exp, exp)
		rem.Rsh(rem, 1)
	}
	return res
}
