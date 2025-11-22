package utils

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend/groth16"
	bn254_groth16 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"io"
	"log/slog"
	"math"
	"math/big"
	"os"
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

func MakeG1Affine(xStr, yStr string) (bn254.G1Affine, error) {
	// Convert decimal strings to field elements
	var X, Y fp.Element
	if _, err := X.SetString(xStr); err != nil {
		return bn254.G1Affine{}, fmt.Errorf("bad x")
	}
	if _, err := Y.SetString(yStr); err != nil {
		return bn254.G1Affine{}, fmt.Errorf("bad y")
	}

	return bn254.G1Affine{X: X, Y: Y}, nil
}

// BSGS finds integer m in [0, B) such that M = m*G.
// Returns m and a boolean flag indicating success.
func BSGS(M, G bn254.G1Affine, B uint64) (uint64, bool) {
	if B == 0 {
		return 0, false
	}

	n := uint64(math.Ceil(math.Sqrt(float64(B)))) // √B

	H := make(map[string]uint64, n)
	var jG bn254.G1Jac
	jG.FromAffine(&bn254.G1Affine{}) // start from identity
	var step bn254.G1Jac
	step.FromAffine(&G)
	var tmpAff bn254.G1Affine

	tmpAff.FromJacobian(&jG)
	H[string(tmpAff.Marshal())] = 0

	for j := uint64(1); j < n; j++ {
		jG.AddAssign(&step)
		tmpAff.FromJacobian(&jG)
		H[string(tmpAff.Marshal())] = j
	}

	var nG bn254.G1Affine
	nG.ScalarMultiplication(&G, new(big.Int).SetUint64(n))

	var Q bn254.G1Jac
	Q.FromAffine(&M)
	var minusnG bn254.G1Affine
	minusnG.Neg(&nG)

	for i := uint64(0); i <= n; i++ {
		tmpAff.FromJacobian(&Q)
		if j, ok := H[string(tmpAff.Marshal())]; ok {
			m := i*n + j
			if m < B {
				return m, true
			}
		}
		Q.AddMixed(&minusnG)
	}

	return 0, false
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

func GetCircuitKeys(basePath string) (io.Reader, io.Reader) {
	pkFile, err := os.Open(fmt.Sprintf("%s/provingKey", basePath))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to generate pkFile, err : %s", err))
		panic(err)
	}
	vkFile, err := os.Open(fmt.Sprintf("%s/verifyingKey", basePath))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to generate vkFile, err : %s", err))
		panic(err)
	}

	return vkFile, pkFile
}

func ExtractProof(proof groth16.Proof) ([]*big.Int, error) {
	bn254Proof, _ := proof.(*bn254_groth16.Proof)

	ax := new(big.Int).SetBytes(bn254Proof.Ar.X.Marshal())
	ay := new(big.Int).SetBytes(bn254Proof.Ar.Y.Marshal())

	// B (G2) in gnark order: (real, imag)
	bx0 := new(big.Int).SetBytes(bn254Proof.Bs.X.A0.Marshal()) // x.real (A0)
	bx1 := new(big.Int).SetBytes(bn254Proof.Bs.X.A1.Marshal()) // x.imag (A1)
	by0 := new(big.Int).SetBytes(bn254Proof.Bs.Y.A0.Marshal()) // y.real (A0)
	by1 := new(big.Int).SetBytes(bn254Proof.Bs.Y.A1.Marshal()) // y.imag (A1)

	// C/Krs (G1)
	cx := new(big.Int).SetBytes(bn254Proof.Krs.X.Marshal())
	cy := new(big.Int).SetBytes(bn254Proof.Krs.Y.Marshal())

	if len(bn254Proof.Commitments) != 0 {
		commitmentx := new(big.Int).SetBytes(bn254Proof.Commitments[0].X.Marshal())
		commitmenty := new(big.Int).SetBytes(bn254Proof.Commitments[0].Y.Marshal())
		PokCommitmentx := new(big.Int).SetBytes(bn254Proof.CommitmentPok.X.Marshal())
		PokCommitmenty := new(big.Int).SetBytes(bn254Proof.CommitmentPok.Y.Marshal())

		return []*big.Int{
			ax, ay, bx1, bx0, by1, by0, cx, cy, commitmentx, commitmenty, PokCommitmentx, PokCommitmenty,
		}, nil
	}

	// Return in EIP-197 order for your solidity verifier:
	// [A.x, A.y, B.x.imag, B.x.real, B.y.imag, B.y.real, C.x, C.y]
	return []*big.Int{ax, ay, bx1, bx0, by1, by0, cx, cy}, nil
}
func MimcHash(inputs ...string) (string, error) {
	h := mimc.NewMiMC()

	for _, s := range inputs {
		// interpret the decimal string as a big integer
		bi, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return "", fmt.Errorf("invalid integer: %s", s)
		}

		var fe fr.Element
		fe.SetBigInt(bi)

		// feed valid 32-byte field element to MiMC
		if _, err := h.Write(fe.Marshal()); err != nil {
			return "", err
		}
	}

	var outFe fr.Element
	outFe.Unmarshal(h.Sum(nil))
	return outFe.BigInt(new(big.Int)).String(), nil
}
