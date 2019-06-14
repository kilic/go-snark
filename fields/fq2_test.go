package fields

import (
	"math/big"
	"testing"

	fp "github.com/kilic/fp256"
)

type fe struct {
	fp.FieldElement
}

// todo
// single test func
// boxes

var testmodulus = "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"
var testnonresiduefq2 = "0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd46"

func TestFq2Add(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	b0 := "0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
	b1 := "0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b"
	b, _ := fq2.NewElementFromString(b0, b1)
	e0 := "0x2ac93d94edab8c618d1571e6ec2785094b150a3e03a2ae5893c5895e408d7497"
	e1 := "0x22951d63ea6d38b05eff59649b0790bbadf57a66a65d7605ed9160948015aa1d"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Add(c, a, b)
	if !fq2.Equal(c, e) {
		t.Errorf("bad addition")
		return
	}
}

func TestFq2Double(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	e0 := "0x3001bdde243e3cec84d400ccbcb888f2ce8645a9eebdb5ba8dbd7ab9b325edda"
	e1 := "0x02b8d8b442e8f04b2c7139b7e27561ee4bd327d502e20397f3a87f588569283d"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Double(c, a)
	if !fq2.Equal(c, e) {
		t.Errorf("bad doubling")
		return
	}
}

func TestFq2Sub(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	b0 := "0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
	b1 := "0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b"
	b, _ := fq2.NewElementFromString(b0, b1)
	e0 := "0x053880493692b08af7be8ee5d09103e983713b6beb1b0761f9f7f15b72987943"
	e1 := "0x108809c339ad57c485c22609c8ef2990355f17ffc4f6581f4237aadaddd07b67"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Sub(c, a, b)
	if !fq2.Equal(c, e) {
		t.Errorf("bad subtraction")
		return
	}
}

func TestFq2Neg(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	e0 := "0x18636f83cf1281b375e64550232513e4303e47bc7112efaff541ceb9feea065a"
	e1 := "0x16d5badf4f2457ef45ef85ff4f85fb37a5d7215e32c7e37aa43c065f2989ea85"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Neg(c, a)
	if !fq2.Equal(c, e) {
		t.Errorf("bad negate")
		return
	}
}

func TestFq2Mul(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	b0 := "0x12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa"
	b1 := "0x090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b"
	b, _ := fq2.NewElementFromString(b0, b1)
	e0 := "0x21aa0d94f8b20db74d7e485100e81c1020302d6fe18c2f820eabd11f323bed1e"
	e1 := "0x23ec88f9862cd431afcf4ddc8a6bffa2faea737cb95ddcb46ac28068f81d0b0c"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Mul(c, a, b)
	if !fq2.Equal(c, e) {
		t.Errorf("bad multiplication")
		return
	}
}

func TestFq2Square(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	e0 := "0x3010f575ce4a684cc50cf0ee970804db4f648c358eca242228ef10bdd6158959"
	e1 := "0x040fec49d9509ae2dfb8c2cffdd54ce8d4fbb77bb051b087926cbff20ae6911e"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Square(c, a)
	if !fq2.Equal(c, e) {
		t.Errorf("bad squaring")
		return
	}
}

func TestFq2Inverse(t *testing.T) {
	p, _ := new(big.Int).SetString(testmodulus[2:], 16)
	fq1 := fp.NewField(p)
	nonResidueFq2, _ := fq1.NewElementFromString(testnonresiduefq2)
	fq2 := Fq2{fq1, nonResidueFq2}
	a0 := "0x1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed"
	a1 := "0x198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c2"
	a, _ := fq2.NewElementFromString(a0, a1)
	e0 := "0x0239320d03633f5a0913bdcb1e715ad12764ddf2a62f86b299db1e7ac5ab137d"
	e1 := "0x164693135edb8b716f380259f4f5d3c973b43213dbf1f2ba2d761a61843dab82"
	e, _ := fq2.NewElementFromString(e0, e1)
	c := fq2.NewElement()
	fq2.Inverse(c, a)
	if !fq2.Equal(c, e) {
		t.Errorf("bad inverse")
		return
	}
}
