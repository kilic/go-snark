package bn128

import (
	"fmt"
	"math/big"

	fp "github.com/kilic/fp256"
)

type PointG1 [3]*fp.FieldElement

type G1 struct {
	F *fp.Field
	G PointG1
	Q *big.Int
}

func NewG1(f *fp.Field, g PointG1, q *big.Int) G1 {

	return G1{
		F: f,
		G: PointG1{
			g[0],
			g[1],
			f.NewElementFromUint(1),
		},
		Q: new(big.Int).Set(q),
	}
}

func (g1 G1) NewPoint() PointG1 {

	return PointG1{
		new(fp.FieldElement),
		new(fp.FieldElement),
		new(fp.FieldElement),
	}
}

// test address ith tis
func (g1 G1) NewPointFromBytes(b []byte) (PointG1, error) {

	if len(b) < 63 {
		return PointG1{}, fmt.Errorf("insufficient space")
	}
	p := PointG1{
		g1.F.NewElementFromBytes(b[:32]),
		g1.F.NewElementFromBytes(b[32:]),
		g1.F.NewElementFromUint(1),
	}
	return p, nil
}

func (g1 G1) ToBytes(b []byte, p PointG1) error {

	if len(b) < 63 {
		return fmt.Errorf("insufficient space")
	}
	a := g1.NewPoint()
	g1.Affine(a, p)
	g1.F.Demont(a[0], a[0])
	g1.F.Demont(a[1], a[1])
	a[0].Marshal(b[:32])
	a[1].Marshal(b[32:])
	return nil
}

func (g1 G1) Zero() PointG1 {

	return PointG1{
		g1.F.NewElementFromUint(0),
		g1.F.NewElementFromUint(1),
		g1.F.NewElementFromUint(0),
	}
}

func (g1 G1) One() PointG1 {

	return PointG1{
		new(fp.FieldElement).SetUint(1),
		new(fp.FieldElement).SetUint(2),
		g1.F.NewElementFromUint(1),
	}
}

func (g1 G1) Copy(p1 PointG1, p2 PointG1) PointG1 {

	p1[0].Set(p2[0])
	p1[1].Set(p2[1])
	p1[2].Set(p2[2])
	return p1
}

func (g1 G1) IsZero(p PointG1) bool {

	return p[2].IsZero()
}

// Intermeriate variables defined outside of arithmetic scope,
// in order to diminish operation cost.
// These field element variables are not used before assigned.
var t0, t1, t2, t3, t4, t5, t6, t7, t8 fp.FieldElement

func (g1 G1) Add(r, p1, p2 PointG1) {

	if g1.IsZero(p1) {
		g1.Copy(r, p2)
		return
	}

	if g1.IsZero(p2) {
		g1.Copy(r, p1)
		return
	}

	g1.F.Square(&t7, p1[2])
	g1.F.Mul(&t1, p2[0], &t7)
	g1.F.Mul(&t2, p1[2], &t7)
	g1.F.Mul(&t0, p2[1], &t2)
	g1.F.Square(&t8, p2[2])
	g1.F.Mul(&t3, p1[0], &t8)
	g1.F.Mul(&t4, p2[2], &t8)
	g1.F.Mul(&t2, p1[1], &t4)
	g1.F.Sub(&t1, &t1, &t3)
	g1.F.Double(&t4, &t1)
	g1.F.Mul(&t4, &t4, &t4)
	g1.F.Mul(&t5, &t1, &t4)
	g1.F.Sub(&t0, &t0, &t2)
	g1.F.Double(&t0, &t0)
	g1.F.Square(&t6, &t0)
	g1.F.Sub(&t6, &t6, &t5)
	g1.F.Mul(&t3, &t3, &t4)
	g1.F.Double(&t4, &t3)
	g1.F.Sub(r[0], &t6, &t4)
	g1.F.Sub(&t4, &t3, r[0])
	g1.F.Mul(&t6, &t2, &t5)
	g1.F.Double(&t6, &t6)
	g1.F.Mul(&t0, &t0, &t4)
	g1.F.Sub(r[1], &t0, &t6)
	g1.F.Add(&t0, p1[2], p2[2])
	g1.F.Square(&t0, &t0)
	g1.F.Sub(&t0, &t0, &t7)
	g1.F.Sub(&t0, &t0, &t8)
	g1.F.Mul(r[2], &t0, &t1)
}

func (g1 G1) Neg(r PointG1, p PointG1) PointG1 {

	r[0].Set(p[0])
	g1.F.Neg(r[1], p[1])
	r[2].Set(p[2])
	return r
}

func (g1 G1) Sub(c, a, b PointG1) PointG1 {

	g1.Neg(c, b)
	g1.Add(c, a, c)
	return c
}

func (g1 G1) Double(r, p PointG1) PointG1 {

	if g1.IsZero(p) {
		g1.Copy(r, p)
		return r
	}

	g1.F.Square(&t0, p[0])
	g1.F.Square(&t1, p[1])
	g1.F.Square(&t2, &t1)
	g1.F.Add(&t1, p[0], &t1)
	g1.F.Square(&t1, &t1)
	g1.F.Sub(&t1, &t1, &t0)
	g1.F.Sub(&t1, &t1, &t2)
	g1.F.Double(&t1, &t1)
	g1.F.Double(&t3, &t0)
	g1.F.Add(&t0, &t3, &t0)
	g1.F.Square(&t4, &t0)
	g1.F.Double(&t3, &t1)
	g1.F.Sub(r[0], &t4, &t3)
	g1.F.Sub(&t1, &t1, r[0])
	g1.F.Double(&t2, &t2)
	g1.F.Double(&t2, &t2)
	g1.F.Double(&t2, &t2)
	g1.F.Mul(&t0, &t0, &t1)
	g1.F.Sub(&t1, &t0, &t2)
	g1.F.Mul(&t0, p[1], p[2])
	r[1].Set(&t1)
	g1.F.Double(r[2], &t0)
	return r
}

func (g1 G1) MulScalar(c, p PointG1, e *big.Int) PointG1 {

	q := PointG1{new(fp.FieldElement), new(fp.FieldElement), new(fp.FieldElement)}
	n := PointG1{new(fp.FieldElement), new(fp.FieldElement), new(fp.FieldElement)}
	g1.Copy(n, p)
	l := e.BitLen()
	for i := 0; i < l; i++ {
		if e.Bit(i) == 1 {
			g1.Add(q, q, n)
		}
		g1.Double(n, n)
	}
	g1.Copy(c, q)
	return c
}

func (g1 G1) IsOnCurve(p PointG1) bool {

	if g1.IsZero(p) {
		return true
	}
	// Y^2 = X^3 + b Z^6
	g1.F.Square(&t0, p[1])
	g1.F.Square(&t1, p[0])
	g1.F.Mul(&t1, &t1, p[0])
	g1.F.Sub(&t0, &t0, &t1)
	g1.F.Square(&t1, p[2])
	g1.F.Mul(&t1, &t1, p[2])
	g1.F.Square(&t1, &t1)
	g1.F.Double(&t2, &t1)
	g1.F.Add(&t1, &t1, &t2)
	return t1.Eq(&t0)
}

func (g1 G1) Affine(r, p PointG1) PointG1 {

	if g1.IsZero(p) {
		g1.Copy(r, g1.Zero())
		return r
	}

	g1.F.InvMontUp(&t0, p[2])
	g1.F.Square(&t1, &t0)
	g1.F.Mul(r[0], p[0], &t1)
	g1.F.Mul(&t0, &t0, &t1)
	g1.F.Mul(r[1], p[1], &t0)
	r[2].Set(g1.F.One())
	return r
}

func (g1 G1) Equal(p1, p2 PointG1) bool {

	if g1.IsZero(p1) {
		return g1.IsZero(p2)
	}

	if g1.IsZero(p2) {
		return g1.IsZero(p1)
	}

	// X1 * Z2^2 == X2 * Z1^2 and Y1 * Z2^3 == Y2 * Z1^3
	g1.F.Square(&t0, p1[2])
	g1.F.Square(&t1, p2[2])
	g1.F.Mul(&t2, &t0, p2[0])
	g1.F.Mul(&t3, &t1, p1[0])
	g1.F.Mul(&t0, &t0, p1[2])
	g1.F.Mul(&t1, &t1, p2[2])
	g1.F.Mul(&t1, &t1, p1[1])
	g1.F.Mul(&t0, &t0, p2[1])
	return t1.Eq(&t0) && t2.Eq(&t3)
}
