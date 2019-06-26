package bn128

import (
	"fmt"
	"math/big"

	"github.com/arnaucube/go-snark/fields"
	fp "github.com/kilic/fp256"
)

type PointG2 [3][2]*fp.FieldElement

type G2 struct {
	f *fields.Fq2
	g PointG2
	b [2]*fp.FieldElement
}

func NewG2(f *fields.Fq2, g PointG2, b [2]*fp.FieldElement) G2 {

	bb := f.Copy(f.NewElement(), b)
	return G2{
		f: f,
		g: PointG2{
			g[0],
			g[1],
			f.One(),
		},
		b: bb,
	}
}

func (g2 G2) NewPoint() PointG2 {

	return PointG2{
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	}
}

func (g2 G2) NewPointFromBytes(b []byte) (PointG2, error) {

	if len(b) < 127 {
		return PointG2{}, fmt.Errorf("")
	}
	x, _ := g2.f.NewElementFromBytes(b[:64])
	y, _ := g2.f.NewElementFromBytes(b[64:])
	return PointG2{x, y, g2.f.One()}, nil
}

func (g2 G2) ToBytes(b []byte, p PointG2) ([]byte, error) {

	if len(b) < 127 {
		return nil, fmt.Errorf("")
	}
	a := g2.NewPoint()
	g2.Affine(a, p)
	g2.f.ToBytes(b[:64], a[0])
	g2.f.ToBytes(b[64:], a[1])
	return b, nil
}

func (g2 G2) G() PointG2 {

	a := g2.NewPoint()
	g2.Copy(a, g2.g)
	return a
}

func (g2 G2) Copy(p1 PointG2, p2 PointG2) PointG2 {

	p1[0][0].Set(p2[0][0])
	p1[1][1].Set(p2[1][1])
	p1[2][0].Set(p2[2][0])
	p1[0][1].Set(p2[0][1])
	p1[1][0].Set(p2[1][0])
	p1[2][1].Set(p2[2][1])
	return p1
}

func (g2 G2) Zero() PointG2 {

	return PointG2{
		g2.f.Zero(),
		g2.f.One(),
		g2.f.Zero(),
	}
}

func (g2 G2) One() PointG2 {

	return PointG2{
		g2.f.Zero(),
		g2.f.One(),
		g2.f.Zero(),
	}
}

func (g2 G2) IsZero(p PointG2) bool {

	return g2.f.IsZero(p[2])
}

func (g2 G2) Equal(p1, p2 PointG2) bool {

	if g2.IsZero(p1) {
		return g2.IsZero(p2)
	}

	if g2.IsZero(p2) {
		return g2.IsZero(p1)
	}

	// X1 * Z2^2 == X2 * Z1^2 and Y1 * Z2^3 == Y2 * Z1^3
	g2.f.Square(tt0, p1[2])
	g2.f.Square(tt1, p2[2])
	g2.f.Mul(tt2, tt0, p2[0])
	g2.f.Mul(tt3, tt1, p1[0])
	g2.f.Mul(tt0, tt0, p1[2])
	g2.f.Mul(tt1, tt1, p2[2])
	g2.f.Mul(tt1, tt1, p1[1])
	g2.f.Mul(tt0, tt0, p2[1])
	return g2.f.Equal(tt1, tt0) && g2.f.Equal(tt2, tt3)
}

func (g2 G2) IsOnCurve(p PointG2) bool {

	if g2.IsZero(p) {
		return true
	}
	// Y^2 = X^3 + b Z^6
	g2.f.Square(tt0, p[1])
	g2.f.Square(tt1, p[0])
	g2.f.Mul(tt1, tt1, p[0])
	g2.f.Sub(tt0, tt0, tt1)
	g2.f.Square(tt1, p[2])
	g2.f.Mul(tt1, tt1, p[2])
	g2.f.Square(tt1, tt1)
	//g2.f.Double(tt2, tt1)
	//g2.f.Add(tt1, tt1, tt2)
	g2.f.Mul(tt1, tt1, g2.b)
	return tt1[0].Eq(tt0[0]) && tt1[1].Eq(tt0[1])
}

func (g2 G2) Affine(r, p PointG2) PointG2 {

	if g2.IsZero(p) {
		g2.Copy(r, g2.Zero())
		return r
	}

	g2.f.Inverse(tt0, p[2])
	g2.f.Square(tt1, tt0)
	g2.f.Mul(r[0], p[0], tt1)
	g2.f.Mul(tt0, tt0, tt1)
	g2.f.Mul(r[1], p[1], tt0)
	g2.f.Copy(r[2], g2.f.One())
	return r
}

var tt0, tt1, tt2, tt3, tt4, tt5, tt6, tt7, tt8 = [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}, [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}, [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}, [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}, [2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)}

func (g2 G2) Add(r, p1, p2 PointG2) PointG2 {

	if g2.IsZero(p1) {
		g2.Copy(r, p2)
		return r
	}

	if g2.IsZero(p2) {
		g2.Copy(r, p1)
		return r
	}

	g2.f.Square(tt7, p1[2])
	g2.f.Mul(tt1, p2[0], tt7)
	g2.f.Mul(tt2, p1[2], tt7)
	g2.f.Mul(tt0, p2[1], tt2)
	g2.f.Square(tt8, p2[2])
	g2.f.Mul(tt3, p1[0], tt8)
	g2.f.Mul(tt4, p2[2], tt8)
	g2.f.Mul(tt2, p1[1], tt4)
	g2.f.Sub(tt1, tt1, tt3)
	g2.f.Double(tt4, tt1)
	g2.f.Mul(tt4, tt4, tt4)
	g2.f.Mul(tt5, tt1, tt4)
	g2.f.Sub(tt0, tt0, tt2)
	g2.f.Double(tt0, tt0)
	g2.f.Square(tt6, tt0)
	g2.f.Sub(tt6, tt6, tt5)
	g2.f.Mul(tt3, tt3, tt4)
	g2.f.Double(tt4, tt3)
	g2.f.Sub(r[0], tt6, tt4)
	g2.f.Sub(tt4, tt3, r[0])
	g2.f.Mul(tt6, tt2, tt5)
	g2.f.Double(tt6, tt6)
	g2.f.Mul(tt0, tt0, tt4)
	g2.f.Sub(r[1], tt0, tt6)
	g2.f.Add(tt0, p1[2], p2[2])
	g2.f.Square(tt0, tt0)
	g2.f.Sub(tt0, tt0, tt7)
	g2.f.Sub(tt0, tt0, tt8)
	g2.f.Mul(r[2], tt0, tt1)
	return r
}

func (g2 G2) Neg(r PointG2, p PointG2) PointG2 {

	g2.f.Copy(r[0], p[0])
	g2.f.Neg(r[1], p[1])
	g2.f.Copy(r[2], p[2])
	return r
}

func (g2 G2) Sub(c, a, b PointG2) PointG2 {

	g2.Neg(c, b)
	g2.Add(c, a, c)
	return c
}

func (g2 G2) Double(r, p PointG2) PointG2 {

	if g2.IsZero(p) {
		g2.Copy(r, p)
		return r
	}

	g2.f.Square(tt0, p[0])
	g2.f.Square(tt1, p[1])
	g2.f.Square(tt2, tt1)
	g2.f.Add(tt1, p[0], tt1)
	g2.f.Square(tt1, tt1)
	g2.f.Sub(tt1, tt1, tt0)
	g2.f.Sub(tt1, tt1, tt2)
	g2.f.Double(tt1, tt1)
	g2.f.Double(tt3, tt0)
	g2.f.Add(tt0, tt3, tt0)
	g2.f.Square(tt4, tt0)
	g2.f.Double(tt3, tt1)
	g2.f.Sub(r[0], tt4, tt3)
	g2.f.Sub(tt1, tt1, r[0])
	g2.f.Double(tt2, tt2)
	g2.f.Double(tt2, tt2)
	g2.f.Double(tt2, tt2)
	g2.f.Mul(tt0, tt0, tt1)
	g2.f.Sub(tt1, tt0, tt2)
	g2.f.Mul(tt0, p[1], p[2])
	g2.f.Copy(r[1], tt1)
	g2.f.Double(r[2], tt0)
	return r
}

func (g2 G2) MulScalar(c, p PointG2, e *big.Int) PointG2 {

	q := PointG2{
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	}
	n := PointG2{
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
		[2]*fp.FieldElement{new(fp.FieldElement), new(fp.FieldElement)},
	}
	ca := 0
	cd := 0
	g2.Copy(n, p)
	l := e.BitLen()
	for i := 0; i < l; i++ {
		if e.Bit(i) == 1 {
			g2.Add(q, q, n)
			ca++
		}
		g2.Double(n, n)
		cd++
	}
	g2.Copy(c, q)
	return c
}
