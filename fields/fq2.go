package fields

import (
	"fmt"
	"io"
	"math/big"

	fp "github.com/kilic/fp256"
)

type Fq2 struct {
	F          *fp.Field
	NonResidue *fp.FieldElement
}

func NewFq2(f *fp.Field, nonResidue *fp.FieldElement) *Fq2 {
	return &Fq2{
		f,
		nonResidue,
	}
}

func (fq2 Fq2) NewElement() [2]*fp.FieldElement {
	return [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}
}

func (fq2 Fq2) NewElementFromBytes(b []byte) ([2]*fp.FieldElement, error) {

	if len(b) < 63 {
		return [2]*fp.FieldElement{}, fmt.Errorf("")
	}

	return [2]*fp.FieldElement{
		fq2.F.NewElementFromBytes(b[:32]),
		fq2.F.NewElementFromBytes(b[32:]),
	}, nil
}

func (fq2 Fq2) rand(a [2]*fp.FieldElement, r io.Reader) [2]*fp.FieldElement {
	fq2.F.RandElement(a[0], r)
	fq2.F.RandElement(a[1], r)
	return a
}

func (fq2 Fq2) ToBytes(b []byte, a [2]*fp.FieldElement) ([]byte, error) {
	if len(b) < 63 {
		return nil, fmt.Errorf("")
	}
	t := fq2.NewElement()
	fq2.Demont(t, a)
	t[0].Marshal(b[:32])
	t[1].Marshal(b[32:])
	return b, nil
}

func (fq2 Fq2) Zero() [2]*fp.FieldElement {
	return [2]*fp.FieldElement{fq2.F.NewElementFromUint(0), fq2.F.NewElementFromUint(0)}
}

func (fq2 Fq2) One() [2]*fp.FieldElement {
	return [2]*fp.FieldElement{fq2.F.NewElementFromUint(1), fq2.F.NewElementFromUint(0)}
}

func (fq2 Fq2) IsZero(a [2]*fp.FieldElement) bool {
	return a[0].IsZero() && a[1].IsZero()
}

func (fq2 Fq2) Equal(a, b [2]*fp.FieldElement) bool {
	return a[0].Eq(b[0]) && a[1].Eq(b[1])
}

func (fq2 Fq2) Copy(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	c[0].Set(a[0])
	c[1].Set(a[1])
	return c
}

func (fq2 Fq2) Demont(c [2]*fp.FieldElement, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Demont(c[0], a[0])
	fq2.F.Demont(c[1], a[1])
	return c
}

func (fq2 Fq2) mulByNonResidue(c, a *fp.FieldElement) *fp.FieldElement {
	fq2.F.Mul(c, fq2.NonResidue, a)
	return c
}

func (fq2 Fq2) Add(c, a, b [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Add(c[0], a[0], b[0])
	fq2.F.Add(c[1], a[1], b[1])
	return c
}

func (fq2 Fq2) Double(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Double(c[0], a[0])
	fq2.F.Double(c[1], a[1])
	return c
}

func (fq2 Fq2) Sub(c, a, b [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Sub(c[0], a[0], b[0])
	fq2.F.Sub(c[1], a[1], b[1])
	return c
}

func (fq2 Fq2) Neg(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Neg(c[0], a[0])
	fq2.F.Neg(c[1], a[1])
	return c
}

var t1, t2, t0, t3 fp.FieldElement

func (fq2 Fq2) Mul(c, a, b [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Mul(&t1, a[0], b[0])
	fq2.F.Mul(&t2, a[1], b[1])
	fq2.F.Add(&t0, &t1, &t2)
	fq2.mulByNonResidue(&t2, &t2)
	fq2.F.Add(&t3, &t1, &t2)
	fq2.F.Add(&t1, a[0], a[1])
	fq2.F.Add(&t2, b[0], b[1])
	fq2.F.Mul(&t1, &t1, &t2)
	c[0].Set(&t3)
	fq2.F.Sub(c[1], &t1, &t0)
	return c
}

func (fq2 Fq2) MulScalar(p [2]*fp.FieldElement, e *big.Int) [2]*fp.FieldElement {
	q := fq2.Zero()
	d := new(big.Int).SetBytes(e.Bytes())
	r := p
	foundone := false
	for i := d.BitLen(); i >= 0; i-- {
		if foundone {
			fq2.Double(q, q)
		}
		if d.Bit(i) == 1 {
			foundone = true
			fq2.Add(q, q, r)
		}
	}
	return q
}

func (fq2 Fq2) Square2(c, a [2]fp.FieldElement) {
	fq2.F.Mul(&t0, &a[0], &a[1])
	fq2.F.Double(&t3, &t0)
	fq2.mulByNonResidue(&t1, &t0)
	fq2.F.Add(&t0, &t1, &t0)
	fq2.mulByNonResidue(&t1, &a[1])
	fq2.F.Add(&t1, &t1, &a[0])
	fq2.F.Add(&t2, &a[0], &a[1])
	fq2.F.Mul(&t2, &t1, &t2)
	fq2.F.Sub(&c[0], &t2, &t0)
	c[1].Set(&t3)
}

func (fq2 Fq2) Square(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Mul(&t0, a[0], a[1])
	fq2.F.Double(&t3, &t0)
	fq2.mulByNonResidue(&t1, &t0)
	fq2.F.Add(&t0, &t1, &t0)
	fq2.mulByNonResidue(&t1, a[1])
	fq2.F.Add(&t1, &t1, a[0])
	fq2.F.Add(&t2, a[0], a[1])
	fq2.F.Mul(&t2, &t1, &t2)
	fq2.F.Sub(c[0], &t2, &t0)
	c[1].Set(&t3)
	return c
}

func (fq2 Fq2) Inverse(c, a [2]*fp.FieldElement) [2]*fp.FieldElement {
	fq2.F.Square(&t0, a[0])
	fq2.F.Square(&t1, a[1])
	fq2.mulByNonResidue(&t1, &t1)
	fq2.F.Sub(&t1, &t0, &t1)
	fq2.F.InvMontUp(&t0, &t1)
	fq2.F.Mul(c[0], a[0], &t0)
	fq2.F.Mul(&t0, a[1], &t0)
	fq2.F.Neg(c[1], &t0)
	return c
}

func (fq2 Fq2) Div(c, a, b [2]*fp.FieldElement) [2]*fp.FieldElement {
	t0 := fq2.NewElement()
	fq2.Inverse(t0, b)
	return fq2.Mul(c, a, t0)
}
