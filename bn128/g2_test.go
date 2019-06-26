package bn128

import (
	"crypto/rand"
	"io"
	"math/big"
	"testing"
)

func (g2 G2) randomPoint(r io.Reader) (PointG2, error) {

	k := new(big.Int)
	k, err := rand.Int(r, g1.Q)
	if err != nil {
		return PointG2{}, err
	}
	p := g2.G()
	g2.MulScalar(p, p, k)
	return p, nil
}

func TestG2Point(t *testing.T) {

	for i := 0; i < nBox; i++ {

		a, _ := g2.randomPoint(rand.Reader)
		bytes := make([]byte, 128)
		_, err := g2.ToBytes(bytes, a)
		if err != nil {
			t.Errorf("")
			return
		}

		b, err := g2.NewPointFromBytes(bytes)
		if err != nil {
			t.Errorf("")
			return
		}

		if !g2.Equal(a, b) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Gen(t *testing.T) {
	g := g2.G()
	if !g2.IsOnCurve(g) {
		t.Errorf("")
		return
	}
}

func TestG2Rand(t *testing.T) {

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		if !g2.IsOnCurve(a) {
			t.Errorf(":(")
			return
		}
	}
}

func TestG2Add(t *testing.T) {

	var t0, t1 = g2.NewPoint(), g2.NewPoint()

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		b, _ := g2.randomPoint(rand.Reader)
		c, _ := g2.randomPoint(rand.Reader)
		g2.Add(t0, a, b)
		g2.Add(t0, t0, c)
		g2.Add(t1, b, c)
		g2.Add(t1, t1, a)
		if !g2.Equal(t0, t1) || !g2.IsOnCurve(t1) || !g2.IsOnCurve(t0) {
			t.Errorf("")
			return
		}
	}

	zero := g2.Zero()

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		b, _ := g2.randomPoint(rand.Reader)
		g2.Add(b, a, zero)
		if !g2.Equal(a, b) || !g2.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Mul(t *testing.T) {

	var t0 = g2.NewPoint()
	var t1 = g2.NewPoint()

	for i := 0; i < nBox; i++ {
		a, _ := rand.Int(rand.Reader, g1.Q)
		b, _ := rand.Int(rand.Reader, g1.Q)
		c, _ := rand.Int(rand.Reader, g1.Q)
		g := g2.G()
		g2.MulScalar(t0, g, b)
		g2.MulScalar(t0, t0, a)
		c.Mul(a, b)
		g2.MulScalar(t1, g, c)
		if !g2.Equal(t0, t1) || !g2.IsOnCurve(t1) || !g2.IsOnCurve(t0) {
			t.Errorf("")
			return
		}
	}

	zeroPoint := g2.Zero()
	zeroScalar := new(big.Int).SetUint64(0)

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		g2.MulScalar(a, a, zeroScalar)
		if !g2.Equal(a, zeroPoint) || !g2.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}

	oneScalar := new(big.Int).SetUint64(1)

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		b := g2.NewPoint()
		g2.MulScalar(b, a, oneScalar)
		if !g2.Equal(a, b) || !g2.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Eq(t *testing.T) {

	zero := g2.Zero()

	if !g2.Equal(zero, zero) || !g2.IsOnCurve(zero) {
		t.Errorf("")
		return
	}

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		if !g2.Equal(a, a) || g2.Equal(zero, a) || !g2.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Neg(t *testing.T) {

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		b := g2.NewPoint()
		zero := g2.NewPoint()
		g2.Neg(b, a)
		g2.Add(b, a, b)
		if !g2.Equal(b, zero) || !g2.IsOnCurve(zero) {
			t.Errorf("")
			return
		}
	}

	for i := 0; i < nBox; i++ {
		a, _ := g2.randomPoint(rand.Reader)
		b := g2.NewPoint()
		g2.Neg(b, a)
		g2.Neg(b, b)
		if !g2.Equal(b, a) || !g2.IsOnCurve(b) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Inverse(t *testing.T) {

	for i := 0; i < nBox; i++ {
		a, _ := rand.Int(rand.Reader, g1.Q)
		ai := new(big.Int).ModInverse(a, g1.Q)
		p1, _ := g2.randomPoint(rand.Reader)
		p2 := g2.NewPoint()
		g2.MulScalar(p2, p1, ai)
		g2.MulScalar(p2, p2, a)
		if !g2.Equal(p1, p2) || !g2.IsOnCurve(p1) || !g2.IsOnCurve(p2) {
			t.Errorf("")
			return
		}
	}
}

func TestG2Order(t *testing.T) {

	zero := g2.Zero()
	g := g2.G()

	for i := 0; i < nBox; i++ {

		a, _ := g2.randomPoint(rand.Reader)

		g2.MulScalar(a, a, g1.Q)
		if !g2.Equal(zero, a) || !g2.IsOnCurve(a) {
			t.Errorf("")
			return
		}

		g2.MulScalar(a, g, g1.Q)
		if !g2.Equal(zero, a) || !g2.IsOnCurve(a) {
			t.Errorf("")
			return
		}

		g2.MulScalar(a, zero, g1.Q)
		if !g2.Equal(zero, a) || !g2.IsOnCurve(a) {
			t.Errorf("")
			return
		}
	}
}

func BenchmarkG2Add(t *testing.B) {

	var t0 = g2.NewPoint()
	a, _ := g2.randomPoint(rand.Reader)
	b, _ := g2.randomPoint(rand.Reader)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		g2.Add(t0, a, b)
	}
}

func BenchmarkG2Double(t *testing.B) {

	var t0 = g2.NewPoint()
	a, _ := g2.randomPoint(rand.Reader)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		g2.Double(t0, a)
	}
}

func BenchmarkG2Mul(t *testing.B) {

	var t0 = g2.NewPoint()
	a, _ := g2.randomPoint(rand.Reader)
	// s, _ := rand.Int(rand.Reader, g1.Q)
	s := g1.Q
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		g2.MulScalar(t0, a, s)
	}
}

func TestSome(t *testing.T) {

	var t0 = g2.NewPoint()
	a, _ := g2.randomPoint(rand.Reader)
	s := g1.Q
	g2.MulScalar(t0, a, s)
}
