package utils

import (
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"math/big"
)

func StringToElement[T emulated.FieldParams](s string) emulated.Element[T] {
	var z big.Int
	z.SetString(s, 10)
	return emulated.ValueOf[T](z)
}

func G1MulAffine(P *bn254.G1Affine, s *big.Int) bn254.G1Affine {
	var out bn254.G1Affine
	out.ScalarMultiplication(P, s)
	return out
}

func G1AddAffine(a, b *bn254.G1Affine) bn254.G1Affine {
	var out bn254.G1Affine
	out.Add(a, b)
	return out
}

func GenerateBaseECC() bn254.G1Affine {
	_, _, g1, _ := bn254.Generators()
	return g1
}

func ECCToAffinePoint(p bn254.G1Affine) sw_emulated.AffinePoint[emulated.BN254Fp] {
	var xb, yb big.Int
	p.X.BigInt(&xb)
	p.Y.BigInt(&yb)

	return sw_emulated.AffinePoint[emulated.BN254Fp]{
		X: emulated.ValueOf[emulated.BN254Fp](&xb),
		Y: emulated.ValueOf[emulated.BN254Fp](&yb),
	}
}

func AssertOnCurve(fp *emulated.Field[emulated.BN254Fp], P *sw_emulated.AffinePoint[emulated.BN254Fp]) {
	x2 := fp.Mul(&P.X, &P.X)
	x3 := fp.Mul(x2, &P.X)
	rhs := fp.Add(x3, fp.NewElement(3)) // BN254: y^2 = x^3 + 3
	lhs := fp.Mul(&P.Y, &P.Y)
	fp.AssertIsEqual(lhs, rhs)
}

// Split big.Int x into 4 little-endian 64-bit limbs.
func splitLE64(x *big.Int) [4]*big.Int {
	var two64 = new(big.Int).Lsh(big.NewInt(1), 64)                       // 2^64
	var mask64 = new(big.Int).Sub(new(big.Int).Set(two64), big.NewInt(1)) // 2^64-1

	t := new(big.Int).Set(x) // don't mutate caller's x
	var out [4]*big.Int
	for i := 0; i < 4; i++ {
		out[i] = new(big.Int).And(t, mask64) // limb i
		t.Rsh(t, 64)                         // x >>= 64
	}
	return out
}

// Recompose limbs -> big.Int (useful for sanity checks)
func joinLE64(limbs [4]*big.Int) *big.Int {
	acc := new(big.Int).Set(limbs[3])
	acc.Lsh(acc, 64)
	acc.Add(acc, limbs[2])
	acc.Lsh(acc, 64)
	acc.Add(acc, limbs[1])
	acc.Lsh(acc, 64)
	acc.Add(acc, limbs[0])
	return acc
}
