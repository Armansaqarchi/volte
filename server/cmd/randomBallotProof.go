package cmd

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/spf13/cobra"
	"log/slog"
	"math/big"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var randomProof = &cobra.Command{
	Use:   "random-proof",
	Short: "Generates random proof using existing verifying and proving keys",
	Run:   runRandomProof,
}

func init() {
	rootCmd.AddCommand(randomProof)
}

func runRandomProof(_ *cobra.Command, _ []string) {
	runRandomBallotProof()
	runRandomMerklePathProof()
	runRandomNullifierProof()
}

func runRandomBallotProof() {
	slog.Info("Generating random proof for Ballot verification.")

	G := utils.GenerateBaseECC()
	x := big.NewInt(30)
	Y := utils.G1MulAffine(&G, x)

	flag.Set("Gx", G.X.String())
	flag.Set("Gy", G.Y.String())
	flag.Set("Yx", Y.X.String())
	flag.Set("Yy", Y.Y.String())

	flag.Parse()

	k := big.NewInt(13220)
	m := big.NewInt(99)
	mG := utils.G1MulAffine(&G, m)
	kY := utils.G1MulAffine(&Y, k)
	C1 := utils.G1MulAffine(&G, k)
	C2 := utils.G1AddAffine(&mG, &kY)

	assignment := circuits.BallotCircuit{
		C1: utils.ECCToAffinePoint(C1),
		C2: utils.ECCToAffinePoint(C2),
		M:  emulated.ValueOf[emulated.BN254Fr](m),
		K:  emulated.ValueOf[emulated.BN254Fr](k),
	}
	vkFile, pkFile := utils.GetCircuitKeys("../keys/groth16/ballot")
	slog.Info("Successfully read ballot provingKey and verifyingKey from files.")
	g16 := zkproofs.NewBallotGroth16FromExistingKeys(vkFile, pkFile)
	slog.Info("Created ballot groth16.")
	createProof(&assignment, g16)
}

func runRandomMerklePathProof() {
	slog.Info("Generating random proof for membership verification.")
	assignment := circuits.MerkleCircuit{
		MerkleRoot:    "1445849190805089689712507754685671168673091553372230482406710117687961492369",
		LeafValue:     "4137760094704180852789719500758563423980885922685717827383305955441808899436",
		PathPositions: []frontend.Variable{0, 0, 1, 0, 1, 1, 1, 0},
		MerklePath: []frontend.Variable{
			"2201405337893316815918273758681089615279726740723671154548377099132895554651",
			"424860605035859383704348310352427067128543432054056537987894636310336670904",
			"1733733094908890202308645489071232861946837868060311922544482735299421656899",
			"4738184262270190554129232471504937848386298074999300056607424141837299375993",
			"2891247542101801713889729682482526880106482743800010573169732230883085419928",
			"2234929124811023109737676314622609840962882856595367442516400778278767910233",
			"4814974093518456386896925216260606687605659407293043447225930833688370931203",
			"3310810034820363093334307352403815617830879961547922857199997655893198882998",
		},
	}
	vkFile, pkFile := utils.GetCircuitKeys("../keys/groth16/membership")
	slog.Info("Successfully read membership provingKey and verifyingKey from files.")
	g16 := zkproofs.NewMembershipGroth16FromExistingKeys(8, vkFile, pkFile)
	slog.Info("Created membership groth16.")
	createProof(&assignment, g16)
}

func runRandomNullifierProof() {
	assignment := circuits.NullifierCircuit{
		SecretKey: "5483851157728431092195247496931641754059230332466588572732387297395134828078",
		EventID:   "3523196653250260958887739657950671762678466692388251624290163732010351636053",
		Nullifier: "10858051838952645709440492871522823286939885463470476254383445187743108776413",
	}
	vkFile, pkFile := utils.GetCircuitKeys("../keys/groth16/nullifier")
	slog.Info("Successfully read nullifier provingKey and verifyingKey from files.")
	g16 := zkproofs.NewNullifierGroth16FromExistingKeys(vkFile, pkFile)
	slog.Info("Created nullifier groth16.")
	createProof(&assignment, g16)
}

func createProof(assignment frontend.Circuit, g16 *zkproofs.Groth16) {

	fullWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		slog.Error("new witness failed: %v", err)
	}
	slog.Info("Created witness for cs circuit.")
	pubWitness, err := fullWitness.Public()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to extract public inputs from witness, err : %s", err))
	}
	proof, err := groth16.Prove(g16.GetR1CSSystem().GetConstraintSystem(), g16.GetProvingKey(), fullWitness, func(config *backend.ProverConfig) error {
		config.HashToFieldFn = sha256.New()
		return nil
	})
	if err != nil {
		slog.Error(fmt.Sprintf("failed to generate proof: %v", err))
		panic(err)
	}
	slog.Info("Proof created.")
	slog.Info("Testing proof.")
	proofParts, err := utils.ExtractProof(proof)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to extract public inputs from witness, err : %s", err))
	}
	for _, part := range proofParts {
		fmt.Println(part.String())
	}
	fmt.Println("Generated public inputs:")
	fmt.Println(pubWitness.Vector())
}
