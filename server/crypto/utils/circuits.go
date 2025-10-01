package utils

import (
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/std/math/emulated/emparams"
	"math/big"
)

func StringToBLS12377Element(s string) emulated.Element[emulated.BLS12377Fp] {
	var z big.Int
	z.SetString(s, 10)
	return emulated.ValueOf[emulated.BLS12377Fp](z)
}

func GenerateBaseECCAffinePoint() sw_emulated.AffinePoint[emulated.BLS12377Fp] {
	_, _, g1, _ := bls12377.Generators()

	// Converting Generator base point coordinates into big.Int
	var Xb, Yb big.Int
	g1.X.BigInt(&Xb)
	g1.Y.BigInt(&Yb)
	// Constructing AffinePoint based on G coordinates
	return sw_emulated.AffinePoint[emulated.BLS12377Fp]{
		X: emulated.ValueOf[emparams.BLS12377Fp](g1.X.Bits()),
		Y: emulated.ValueOf[emparams.BLS12377Fp](g1.Y.Bits()),
	}
}
