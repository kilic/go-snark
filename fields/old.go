package fields

import (
	"bytes"
	"math/big"
)

type FqOld struct {
	Q *big.Int
}

func NewFqOld(q *big.Int) *FqOld {
	return &FqOld{
		q,
	}
}

func (fq FqOld) Zero() *big.Int {
	return big.NewInt(int64(0))
}

// One returns a One value on the Fq
func (fq FqOld) One() *big.Int {
	return big.NewInt(int64(1))
}

func (fq FqOld) Add(a, b *big.Int) *big.Int {
	r := new(big.Int).Add(a, b)
	return new(big.Int).Mod(r, fq.Q)
}

func (fq FqOld) Double(a *big.Int) *big.Int {
	r := new(big.Int).Add(a, a)
	return new(big.Int).Mod(r, fq.Q)
}

func (fq FqOld) Sub(a, b *big.Int) *big.Int {
	r := new(big.Int).Sub(a, b)
	return new(big.Int).Mod(r, fq.Q)
}

func (fq FqOld) Neg(a *big.Int) *big.Int {
	m := new(big.Int).Neg(a)
	return new(big.Int).Mod(m, fq.Q)
}

func (fq FqOld) Mul(a, b *big.Int) *big.Int {
	m := new(big.Int).Mul(a, b)
	return new(big.Int).Mod(m, fq.Q)
}

func (fq FqOld) MulScalar(base, e *big.Int) *big.Int {
	return fq.Mul(base, e)
}

func (fq FqOld) Inverse(a *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, fq.Q)
}

func (fq FqOld) Div(a, b *big.Int) *big.Int {
	d := fq.Mul(a, fq.Inverse(b))
	return new(big.Int).Mod(d, fq.Q)
}

func (fq FqOld) Square(a *big.Int) *big.Int {
	m := new(big.Int).Mul(a, a)
	return new(big.Int).Mod(m, fq.Q)
}

func (fq FqOld) Exp(base *big.Int, e *big.Int) *big.Int {
	res := fq.One()
	rem := fq.Copy(e)
	exp := base

	for !bytes.Equal(rem.Bytes(), big.NewInt(int64(0)).Bytes()) {
		if BigIsOdd(rem) {
			res = fq.Mul(res, exp)
		}
		exp = fq.Square(exp)
		rem = new(big.Int).Rsh(rem, 1)
	}
	return res
}

func (fq FqOld) IsZero(a *big.Int) bool {
	return bytes.Equal(a.Bytes(), fq.Zero().Bytes())
}

func (fq FqOld) Copy(a *big.Int) *big.Int {
	return new(big.Int).SetBytes(a.Bytes())
}
func (fq FqOld) Affine(a *big.Int) *big.Int {
	nq := fq.Neg(fq.Q)

	aux := a
	if aux.Cmp(big.NewInt(int64(0))) == -1 {
		if aux.Cmp(nq) != 1 {
			aux = new(big.Int).Mod(aux, fq.Q)
		}
		if aux.Cmp(big.NewInt(int64(0))) == -1 {
			aux = new(big.Int).Add(aux, fq.Q)
		}
	} else {
		if aux.Cmp(fq.Q) != -1 {
			aux = new(big.Int).Mod(aux, fq.Q)
		}
	}
	return aux
}

func (fq FqOld) Equal(a, b *big.Int) bool {
	aAff := fq.Affine(a)
	bAff := fq.Affine(b)
	return bytes.Equal(aAff.Bytes(), bAff.Bytes())
}

func BigIsOdd(n *big.Int) bool {
	one := big.NewInt(int64(1))
	and := new(big.Int).And(n, one)
	return bytes.Equal(and.Bytes(), big.NewInt(int64(1)).Bytes())
}

type Fq2Old struct {
	F          *FqOld
	NonResidue *big.Int
}

// NewFq2Old generates a new Fq2Old
func NewFq2Old(f *FqOld, nonResidue *big.Int) *Fq2Old {
	return &Fq2Old{
		f,
		nonResidue,
	}
}

// Zero returns a Zero value on the Fq2Old
func (fq2 Fq2Old) Zero() [2]*big.Int {
	return [2]*big.Int{fq2.F.Zero(), fq2.F.Zero()}
}

// One returns a One value on the Fq2Old
func (fq2 Fq2Old) One() [2]*big.Int {
	return [2]*big.Int{fq2.F.One(), fq2.F.Zero()}
}

func (fq2 Fq2Old) mulByNonResidue(a *big.Int) *big.Int {
	return fq2.F.Mul(fq2.NonResidue, a)
}

// Add performs an addition on the Fq2Old
func (fq2 Fq2Old) Add(a, b [2]*big.Int) [2]*big.Int {
	return [2]*big.Int{
		fq2.F.Add(a[0], b[0]),
		fq2.F.Add(a[1], b[1]),
	}
}

// Double performs a doubling on the Fq2Old
func (fq2 Fq2Old) Double(a [2]*big.Int) [2]*big.Int {
	return fq2.Add(a, a)
}

// Sub performs a subtraction on the Fq2Old
func (fq2 Fq2Old) Sub(a, b [2]*big.Int) [2]*big.Int {
	return [2]*big.Int{
		fq2.F.Sub(a[0], b[0]),
		fq2.F.Sub(a[1], b[1]),
	}
}

// Neg performs a negation on the Fq2Old
func (fq2 Fq2Old) Neg(a [2]*big.Int) [2]*big.Int {
	return fq2.Sub(fq2.Zero(), a)
}

// Mul performs a multiplication on the Fq2Old
func (fq2 Fq2Old) Mul(a, b [2]*big.Int) [2]*big.Int {
	// Multiplication and Squaring on Pairing-Friendly.pdf; Section 3 (Karatsuba)
	// https://pdfs.semanticscholar.org/3e01/de88d7428076b2547b60072088507d881bf1.pdf
	v0 := fq2.F.Mul(a[0], b[0])
	v1 := fq2.F.Mul(a[1], b[1])
	return [2]*big.Int{
		fq2.F.Add(v0, fq2.mulByNonResidue(v1)),
		fq2.F.Sub(
			fq2.F.Mul(
				fq2.F.Add(a[0], a[1]),
				fq2.F.Add(b[0], b[1])),
			fq2.F.Add(v0, v1)),
	}
}

func (fq2 Fq2Old) MulScalar(p [2]*big.Int, e *big.Int) [2]*big.Int {
	// for more possible implementations see g2.go file, at the function g2.MulScalar()

	q := fq2.Zero()
	d := fq2.F.Copy(e)
	r := p

	foundone := false
	for i := d.BitLen(); i >= 0; i-- {
		if foundone {
			q = fq2.Double(q)
		}
		if d.Bit(i) == 1 {
			foundone = true
			q = fq2.Add(q, r)
		}
	}
	return q
}

// Inverse returns the inverse on the Fq2Old
func (fq2 Fq2Old) Inverse(a [2]*big.Int) [2]*big.Int {
	// High-Speed Software Implementation of the Optimal Ate Pairing over Barreto–Naehrig Curves .pdf
	// https://eprint.iacr.org/2010/354.pdf , algorithm 8
	t0 := fq2.F.Square(a[0])
	t1 := fq2.F.Square(a[1])
	t2 := fq2.F.Sub(t0, fq2.mulByNonResidue(t1))
	t3 := fq2.F.Inverse(t2)
	return [2]*big.Int{
		fq2.F.Mul(a[0], t3),
		fq2.F.Neg(fq2.F.Mul(a[1], t3)),
	}
}

// Div performs a division on the Fq2Old
func (fq2 Fq2Old) Div(a, b [2]*big.Int) [2]*big.Int {
	return fq2.Mul(a, fq2.Inverse(b))
}

// Square performs a square operation on the Fq2Old
func (fq2 Fq2Old) Square(a [2]*big.Int) [2]*big.Int {
	// https://pdfs.semanticscholar.org/3e01/de88d7428076b2547b60072088507d881bf1.pdf , complex squaring
	ab := fq2.F.Mul(a[0], a[1])
	return [2]*big.Int{
		fq2.F.Sub(
			fq2.F.Mul(
				fq2.F.Add(a[0], a[1]),
				fq2.F.Add(
					a[0],
					fq2.mulByNonResidue(a[1]))),
			fq2.F.Add(
				ab,
				fq2.mulByNonResidue(ab))),
		fq2.F.Add(ab, ab),
	}
}

func (fq2 Fq2Old) IsZero(a [2]*big.Int) bool {
	return fq2.F.IsZero(a[0]) && fq2.F.IsZero(a[1])
}

func (fq2 Fq2Old) Affine(a [2]*big.Int) [2]*big.Int {
	return [2]*big.Int{
		fq2.F.Affine(a[0]),
		fq2.F.Affine(a[1]),
	}
}
func (fq2 Fq2Old) Equal(a, b [2]*big.Int) bool {
	return fq2.F.Equal(a[0], b[0]) && fq2.F.Equal(a[1], b[1])
}

func (fq2 Fq2Old) Copy(a [2]*big.Int) [2]*big.Int {
	return [2]*big.Int{
		fq2.F.Copy(a[0]),
		fq2.F.Copy(a[1]),
	}
}
