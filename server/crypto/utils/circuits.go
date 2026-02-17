package utils

import (
	"errors"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/consensys/gnark/backend/groth16"
	bn254_groth16 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"github.com/iden3/go-iden3-crypto/v2/mimc7"
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
func MimcHash(inputs ...[]byte) (string, error) {
	h := mimc.NewMiMC()

	for _, s := range inputs {
		// interpret the decimal string as a big integer
		bi, ok := new(big.Int).SetString(string(s), 10)
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

	fmt.Println("Result : ", outFe.BigInt(new(big.Int)).String())

	return outFe.BigInt(new(big.Int)).String(), nil
}

var frModulus, _ = new(big.Int).SetString(
	"21888242871839275222246405745257275088548364400416034343698204186575808495617",
	10,
)

func ReduceToFr(x *big.Int) *big.Int {
	if x == nil {
		return new(big.Int)
	}
	z := new(big.Int).Mod(new(big.Int).Set(x), frModulus)
	if z.Sign() < 0 {
		z.Add(z, frModulus)
	}
	return z
}

func NewBigIntFromString(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
}

func MiMC7MultiHashCircomFr(inputs []*big.Int) (*big.Int, error) {
	if len(inputs) == 0 {
		return nil, errors.New("inputs must be non-empty")
	}

	reduced := make([]*big.Int, len(inputs))
	for i := range inputs {
		reduced[i] = ReduceToFr(inputs[i])
	}
	// nil key => k = 0
	return mimc7.Hash(reduced, nil)

}

func DecryptM_BSGS(C1, C2 *babyjub.Point, privX *big.Int, max uint64) (uint64, error) {
	if max == 0 {
		return 0, errors.New("max must be > 0")
	}

	// Generator used by iden3 babyjub: B8 (base point * 8, subgroup generator). :contentReference[oaicite:4]{index=4}
	G := babyjub.B8

	// BN254 Fr modulus (used by BabyJubJub in circomlib/iden3). :contentReference[oaicite:5]{index=5}
	q, _ := new(big.Int).SetString(
		"21888242871839275222246405745257275088548364400416034343698204186575808495617",
		10,
	)

	// Helper: affine add via projective (since affine Point has no Add method). :contentReference[oaicite:6]{index=6}
	addAffine := func(a, b *babyjub.Point) *babyjub.Point {
		pa := a.Projective()
		pb := b.Projective()
		sum := babyjub.NewPointProjective().Add(pa, pb).Affine()
		return sum
	}

	// Helper: negate a point on twisted Edwards: -(x,y) = (-x mod q, y)
	negAffine := func(p *babyjub.Point) *babyjub.Point {
		nx := new(big.Int).Neg(p.X)
		nx.Mod(nx, q)
		return &babyjub.Point{X: nx, Y: new(big.Int).Set(p.Y)}
	}

	// 1) H = C2 - x*C1
	S := babyjub.NewPoint().Mul(privX, C1) // S = x*C1 :contentReference[oaicite:7]{index=7}
	H := addAffine(C2, negAffine(S))       // H = C2 - S

	// 2) BSGS parameters
	mStep := uint64(math.Ceil(math.Sqrt(float64(max))))
	if mStep == 0 {
		return 0, errors.New("invalid step size")
	}

	// 3) Baby steps: store j*G for j in [0, mStep)
	babies := make(map[[32]byte]uint64, mStep)

	acc := babyjub.NewPoint().Mul(big.NewInt(0), G) // 0*G (identity)
	for j := uint64(0); j < mStep; j++ {
		babies[acc.Compress()] = j // canonical key :contentReference[oaicite:8]{index=8}
		acc = addAffine(acc, G)
	}

	// 4) Giant steps: look for H - i*(mStep*G) in baby table
	mG := babyjub.NewPoint().Mul(new(big.Int).SetUint64(mStep), G) // mStep*G
	negmG := negAffine(mG)

	limit := uint64(math.Ceil(float64(max) / float64(mStep)))
	cur := H

	for i := uint64(0); i <= limit; i++ {
		if j, ok := babies[cur.Compress()]; ok {
			m := i*mStep + j
			if m < max {
				return m, nil
			}
			return 0, errors.New("match found but m out of range (increase max?)")
		}
		cur = addAffine(cur, negmG) // cur -= mG
	}

	return 0, errors.New("m not found in range; check ciphertext encoding or max")
}
