package bn128

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"

	fp "github.com/kilic/fp256"
)

var g1 G1
var g PointG1
var nBox = 1000

// not secure
func (g1 G1) randomPoint(r io.Reader) (PointG1, error) {

	k := new(big.Int)
	k, err := rand.Int(r, g1.Q)
	if err != nil {
		return PointG1{}, err
	}
	p := g1.NewPoint()
	g1.Copy(p, g1.G)
	g1.MulScalar(p, p, k)
	return p, nil
}

func TestMain(m *testing.M) {
	fieldorder := "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
	curveorder := "0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	p, _ := new(big.Int).SetString(fieldorder[2:], 16)
	q, _ := new(big.Int).SetString(curveorder[2:], 16)
	f := fp.NewField(p)
	g = PointG1{
		f.NewElementFromUint(1),
		f.NewElementFromUint(2),
		f.NewElementFromUint(1)}
	g1 = NewG1(f, g, q)
	m.Run()
}

func TestG1Point(t *testing.T) {

	for i := 0; i < nBox; i++ {

		a, _ := g1.randomPoint(rand.Reader)
		bytes := make([]byte, 64)

		err := g1.ToBytes(bytes, a)
		if err != nil {
			t.Errorf("")
			return
		}

		b, err := g1.NewPointFromBytes(bytes)
		if err != nil {
			t.Errorf("")
			return
		}

		if !g1.Equal(a, b) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Generator(t *testing.T) {

	if !g1.IsOnCurve(g1.G) {
		t.Errorf("")
		return
	}
}

func TestG1Add(t *testing.T) {

	var t0, t1 = g1.NewPoint(), g1.NewPoint()

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		b, _ := g1.randomPoint(rand.Reader)
		c, _ := g1.randomPoint(rand.Reader)
		g1.Add(t0, a, b)
		g1.Add(t0, t0, c)
		g1.Add(t1, b, c)
		g1.Add(t1, t1, a)
		if !g1.Equal(t0, t1) || !g1.IsOnCurve(t1) || !g1.IsOnCurve(t0) {
			t.Errorf("")
			return
		}
	}

	zero := g1.Zero()

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		b, _ := g1.randomPoint(rand.Reader)
		g1.Add(b, a, zero)
		if !g1.Equal(a, b) || !g1.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Sub(t *testing.T) {

	zero := g1.Zero()

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		c, _ := g1.randomPoint(rand.Reader)
		g1.Sub(c, a, a)
		if !g1.Equal(zero, c) || !g1.IsOnCurve(c) {
			t.Errorf("")
		}
	}
}

func TestG1Mul(t *testing.T) {

	var t0 = g1.NewPoint()
	var t1 = g1.NewPoint()

	for i := 0; i < nBox; i++ {
		a, _ := rand.Int(rand.Reader, g1.Q)
		b, _ := rand.Int(rand.Reader, g1.Q)
		c, _ := rand.Int(rand.Reader, g1.Q)
		g1.MulScalar(t0, g1.G, b)
		g1.MulScalar(t0, t0, a)
		c.Mul(a, b)
		g1.MulScalar(t1, g1.G, c)
		if !g1.Equal(t0, t1) || !g1.IsOnCurve(t1) || !g1.IsOnCurve(t0) {
			t.Errorf("")
			return
		}
	}

	zeroPoint := g1.Zero()
	zeroScalar := new(big.Int).SetUint64(0)

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		g1.MulScalar(a, a, zeroScalar)
		if !g1.Equal(a, zeroPoint) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}

	oneScalar := new(big.Int).SetUint64(1)

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		b := g1.NewPoint()
		g1.MulScalar(b, a, oneScalar)
		if !g1.Equal(a, b) || !g1.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Eq(t *testing.T) {

	zero := g1.Zero()

	if !g1.Equal(zero, zero) || !g1.IsOnCurve(zero) {
		t.Errorf("")
		return
	}

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		if !g1.Equal(a, a) || g1.Equal(zero, a) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Neg(t *testing.T) {

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		b := g1.NewPoint()
		zero := g1.NewPoint()
		g1.Neg(b, a)
		g1.Add(b, a, b)
		if !g1.Equal(b, zero) || !g1.IsOnCurve(zero) {
			t.Errorf("")
			return
		}
	}

	for i := 0; i < nBox; i++ {
		a, _ := g1.randomPoint(rand.Reader)
		b := g1.NewPoint()
		g1.Neg(b, a)
		g1.Neg(b, b)
		if !g1.Equal(b, a) || !g1.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Inverse(t *testing.T) {

	for i := 0; i < nBox; i++ {
		a, _ := rand.Int(rand.Reader, g1.Q)
		ai := new(big.Int).ModInverse(a, g1.Q)
		p1, _ := g1.randomPoint(rand.Reader)
		p2 := g1.NewPoint()
		g1.MulScalar(p2, p1, ai)
		g1.MulScalar(p2, p2, a)
		if !g1.Equal(p1, p2) || !g1.IsOnCurve(p1) || !g1.IsOnCurve(p2) {
			t.Errorf("")
			return
		}
	}
}

func TestG1Order(t *testing.T) {

	zero := g1.Zero()
	one := g1.One()

	for i := 0; i < nBox; i++ {

		a, _ := g1.randomPoint(rand.Reader)

		g1.MulScalar(a, a, g1.Q)
		if !g1.Equal(zero, a) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}

		g1.MulScalar(a, one, g1.Q)
		if !g1.Equal(zero, a) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}

		g1.MulScalar(a, zero, g1.Q)
		if !g1.Equal(zero, a) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}

		g1.MulScalar(a, g1.G, g1.Q)
		if !g1.Equal(zero, a) || !g1.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}
}

func BenchmarkG1Add(t *testing.B) {

	var t0 = g1.NewPoint()
	a, _ := g1.randomPoint(rand.Reader)
	b, _ := g1.randomPoint(rand.Reader)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		g1.Add(t0, a, b)
	}
}

func BenchmarkG1Double(t *testing.B) {

	var t0 = g1.NewPoint()
	a, _ := g1.randomPoint(rand.Reader)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		g1.Double(t0, a)
	}
}

func BenchmarkG1Mul(t *testing.B) {

	var t0 = g1.NewPoint()
	a, _ := g1.randomPoint(rand.Reader)
	t.ResetTimer()
	s, _ := rand.Int(rand.Reader, g1.Q)
	for i := 0; i < t.N; i++ {
		g1.MulScalar(t0, a, s)
	}
}
