package main

import (
	"flag"
	"fmt"
	"math/big"
	"volte/backend/chain"
)

//func main() {
//	//r1cs := groth16.NewVolteBLS12377R1CS(10 /* nullifierTreeDepth */)
//	//cs := r1cs.Compile()
//	//slog.Info(fmt.Sprintf("Rank 1 Constraint system is created."))
//	//slog.Info(fmt.Sprintf("%-30s %d.", "Number of constraints:", cs.GetNbConstraints()))
//	//slog.Info(fmt.Sprintf("%-30s %d.", "Number of public variables:", cs.GetNbPublicVariables()))
//	//slog.Info(fmt.Sprintf("%-30s %d.", "Number of secret variables:", cs.GetNbSecretVariables()))
//	//slog.Info(fmt.Sprintf("%-30s %d.", "Number of commitments:", cs.GetCommitments()))
//
//}

// ---- parsing helpers ----

//func mustElemFromDec(dec string) fr.Element {
//	bi, ok := new(big.Int).SetString(dec, 10)
//	if !ok {
//		panic("bad decimal string: " + dec)
//	}
//	var e fr.Element
//	// Prefer SetBigInt (handles reduction mod r if needed)
//	e.SetBigInt(bi)
//	return e
//}
//
//func uint64Elem(v uint64) fr.Element {
//	var e fr.Element
//	e.SetUint64(v)
//	return e
//}
//
//// ---- Poseidon2 2→1 Compress (fixed arity) ----
//// Use this if your Merkle parent is defined as H2(left,right) = Compress(left,right).
//func poseidon2Compress(left, right fr.Element) (fr.Element, error) {
//	params := p2.GetDefaultParameters()
//	perm := p2.NewPermutation(2, params.NbFullRounds, params.NbPartialRounds)
//
//	lb := left.Bytes()
//	rb := right.Bytes()
//
//	outBytes, err := perm.Compress(lb[:], rb[:])
//	if err != nil {
//		return fr.Element{}, err
//	}
//	var out fr.Element
//	if err := out.SetBytesCanonical(outBytes); err != nil {
//		return fr.Element{}, err
//	}
//	return out, nil
//}
//
//// ---- Poseidon2 Merkle–Damgård (streaming) ----
//// This matches gnark/std/hash/poseidon2.NewMerkleDamgardHasher(api)
//func poseidon2MD(items ...fr.Element) (fr.Element, error) {
//	h := p2.NewMerkleDamgardHasher()
//
//	for _, it := range items {
//		b := it.Bytes()
//		if _, err := h.Write(b[:]); err != nil {
//			return fr.Element{}, err
//		}
//	}
//	sum := h.Sum(nil)
//
//	var out fr.Element
//	if err := out.SetBytesCanonical(sum); err != nil {
//		return fr.Element{}, err
//	}
//	return out, nil
//}
//
//func main() {
//	// Your inputs:
//	// left  = "2201405337893316815918273758681089615279726740723671154548377099132895554651"
//	// right = "20"
//	left := mustElemFromDec("2992998858742992003866448380394208494829972004809173354017382673461872646273")
//	right := mustElemFromDec("3310810034820363091334307352403815617830879962547922857199997655893198882998")
//
//	// If your circuit uses MD (your code does):
//	md, err := poseidon2MD(left, right)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("MD:", md.String())
//}

//func main() {
//	groth16 := zkproofs.VolteGroth16{}
//	groth16.Setup()
//	fmt.Println(
//		"ProvingKey -> G1 group number: ", groth16.GetProvingKey().NbG1(),
//		" G2 group number: ", groth16.GetProvingKey().NbG2(),
//	)
//	fmt.Println(
//		"VerifyingKey -> G1 group number: ", groth16.GetVerifyingKey().NbG1(),
//		" G2 group number: ", groth16.GetVerifyingKey().NbG2(),
//	)
//
//}

// first define a data structure with Serialize method to be used as data block
//type testData struct {
//	data []byte
//}

//func (t *testData) Serialize() ([]byte, error) {
//	return t.data, nil
//}
//
//func handleError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}
//
//// generate dummy data blocks
//func generateRandBlocks(size int) (blocks []mt.DataBlock) {
//	for i := 0; i < size; i++ {
//		block := &testData{
//			data: make([]byte, 100),
//		}
//		_, err := rand.Read(block.data)
//		handleError(err)
//		blocks = append(blocks, block)
//	}
//	return
//}

func main() {
	flag.Parse()
	h := chain.NewEthereumChainHandler()
	res, err := h.GetVolteContract().GetEventHash(big.NewInt(2))
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
