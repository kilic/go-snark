package fields

import (
	"math/big"
	"testing"

	fp "github.com/kilic/fp256"
)

var fq1 *fp.Field
var fq1old *FqOld
var fq2 *Fq2
var fq2old *Fq2Old
var fq6 *Fq6
var fq12 *Fq12
var nBox int

func TestMain(m *testing.M) {
	testmodulus := "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
	testnonresiduefq2 := "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd46"
	nBox = 10000
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 = fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 = &Fq2{fq1, nonResidueFq2}
	nonResidueFq6 := [2]*fp.FieldElement{
		fq1.NewElementFromUint(9),
		fq1.NewElementFromUint(1)}
	fq6 = &Fq6{fq2, nonResidueFq6}
	fq12 = &Fq12{
		fq6,
		fq2,
		nonResidueFq6,
	}
	fq1old = NewFqOld(p)
	m.Run()
}

func Benchmark100MultiplicationFqOld(t *testing.B) {
	a, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	b, _ := new(big.Int).SetString("1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b", 16)
	c := new(big.Int)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 100; i++ {
			c = fq1old.Mul(a, b)
		}
	}
	_ = c
}

func Benchmark100MultiplicationFq(t *testing.B) {
	a, _ := new(fp.FieldElement).SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b, _ := new(fp.FieldElement).SetString(nil, "0x1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(fp.FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fq1.Mont(a, a)
		fq1.Mont(b, b)
		for i := 0; i < 100; i++ {
			fq1.Mul(c, a, b)
		}
		fq1.Demont(c, c)
	}
}

func Benchmark100AdditionFqOld(t *testing.B) {
	a, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	b, _ := new(big.Int).SetString("1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b", 16)
	c := new(big.Int)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 100; i++ {
			c = fq1old.Add(a, b)
		}
	}
	_ = c
}

func Benchmark100AdditionFq(t *testing.B) {
	a, _ := new(fp.FieldElement).SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b, _ := new(fp.FieldElement).SetString(nil, "0x1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	c := new(fp.FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fq1.Mont(a, a)
		fq1.Mont(b, b)
		for i := 0; i < 100; i++ {
			fq1.Add(c, a, b)
		}
		fq1.Demont(c, c)
	}
}

func BenchmarkInverseFqOld(t *testing.B) {
	a, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	c := new(big.Int)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c = fq1old.Inverse(a)
	}
	_ = c
}

func BenchmarkInverseFq(t *testing.B) {
	a, _ := fq1.NewElementFromString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	c := new(fp.FieldElement)
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fq1.InvMontUp(c, a)
	}
}

func Benchmark10MultiplicationFq2Old(t *testing.B) {
	a0, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	a1, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	a := [2]*big.Int{a0, a1}
	b0, _ := new(big.Int).SetString("1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b", 16)
	b1, _ := new(big.Int).SetString("1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b", 16)
	b := [2]*big.Int{b0, b1}
	c := [2]*big.Int{new(big.Int), new(big.Int)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for i := 0; i < 10; i++ {
			c = fq2old.Mul(a, b)
		}
	}
	_ = c
}

func Benchmark10MultiplicationFq2(t *testing.B) {
	a := fq2.NewElement()
	b := fq2.NewElement()
	c := fq2.NewElement()
	a[0].SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	a[1].SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	b[0].SetString(nil, "0x1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	b[1].SetString(nil, "0x1fffffffffffffffdddddddd8888888888888888000000119a9a9a9a8b8b8b8b")
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fq1.Mont(a[0], a[0])
		fq1.Mont(a[1], a[1])
		fq1.Mont(b[0], b[0])
		fq1.Mont(b[1], b[1])
		for i := 0; i < 10; i++ {
			fq2.Mul(c, a, b)
		}
		fq2.Demont(c)
	}
}

func BenchmarkInversionq2Old(t *testing.B) {
	a0, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	a1, _ := new(big.Int).SetString("1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa", 16)
	a := [2]*big.Int{a0, a1}
	c := [2]*big.Int{new(big.Int), new(big.Int)}
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		c = fq2old.Inverse(a)
	}
	_ = c
}

func BenchmarkInversionFq2(t *testing.B) {
	a := fq2.NewElement()
	c := fq2.NewElement()
	a[0].SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	a[1].SetString(nil, "0x1aaaaaaaaaaaaaaa44aa44aa44aa44aa91919191ffffff0000119999ffaa01aa")
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		fq2.Inverse(c, a)
	}
}
