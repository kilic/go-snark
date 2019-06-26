package bn128

import (
	"math/big"
	"testing"

	"github.com/arnaucube/go-snark/fields"
	fp "github.com/kilic/fp256"
)

func TestMain(m *testing.M) {
	fieldorder := "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
	curveorder := "0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"
	//
	p, _ := new(big.Int).SetString(fieldorder[2:], 16)
	q, _ := new(big.Int).SetString(curveorder[2:], 16)
	f := fp.NewField(p)
	gg1 := PointG1{
		f.NewElementFromUint(1),
		f.NewElementFromUint(2),
		f.NewElementFromUint(1)}
	b1 := f.NewElementFromUint(3)
	g1 = NewG1(f, gg1, q, b1)
	//
	//
	//
	nonresiduef2 := "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd46"
	nonResidueFq2, _ := f.NewElementFromString(nonresiduef2)
	f2 := fields.NewFq2(f, nonResidueFq2)
	g2oo, _ := f.NewElementFromString("0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed")
	g2o1, _ := f.NewElementFromString("0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2")
	g21o, _ := f.NewElementFromString("0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa")
	g211, _ := f.NewElementFromString("0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b")
	gg2 := PointG2{
		[2]*fp.FieldElement{g2oo, g2o1},
		[2]*fp.FieldElement{g21o, g211},
	}
	twist := [2]*fp.FieldElement{f.NewElementFromUint(9), f.NewElementFromUint(1)}
	invtwist, b2 := f2.NewElement(), f2.NewElement()
	f2.Inverse(invtwist, twist)
	// bts := [64]byte{}
	// b1.Marshal(bts[:])
	// new(big.Int).SetBytes(bts[:])
	b1Big := new(big.Int).SetUint64(3)
	b2 = f2.MulScalar(invtwist, b1Big)
	g2 = NewG2(f2, gg2, b2)
	///
	//
	m.Run()
	//
}
