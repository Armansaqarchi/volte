package utils

import (
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/std/math/emulated/emparams"
	"math/big"
)

func StringToBLS12377Element(s string) emulated.Element[emulated.BN254Fp] {
	var z big.Int
	z.SetString(s, 10)
	return emulated.ValueOf[emulated.BN254Fp](z)
}

func G1MulAffine(P *bls12377.G1Affine, s *big.Int) bls12377.G1Affine {
	var out bls12377.G1Affine
	out.ScalarMultiplication(P, s)
	return out
}

func G1AddAffine(a, b *bls12377.G1Affine) bls12377.G1Affine {
	var out bls12377.G1Affine
	out.Add(a, b)
	return out
}

func GenerateBaseECC() bls12377.G1Affine {
	_, _, g1, _ := bls12377.Generators()
	return g1
}

func ECCToAffinePoint(p bls12377.G1Affine) *sw_emulated.AffinePoint[emulated.BN254Fp] {

	// Converting Generator base point coordinates into big.Int
	var Xb, Yb big.Int
	p.X.BigInt(&Xb)
	p.Y.BigInt(&Yb)
	// Constructing AffinePoint based on G coordinates
	return &sw_emulated.AffinePoint[emulated.BN254Fp]{
		X: emulated.ValueOf[emparams.BN254Fp](p.X),
		Y: emulated.ValueOf[emparams.BN254Fp](p.Y),
	}
}

func ValFr(bi *big.Int) *emulated.Element[emulated.BN254Fp] {
	fr := emulated.ValueOf[emparams.BN254Fp](bi)
	return &fr
}
