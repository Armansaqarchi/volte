package cmd

import (
	"flag"
	"fmt"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"github.com/spf13/cobra"
	"log/slog"
	"math/big"
)

var (
	privKey32 = flag.String("privkey32", "20", "32 byte private key")
)

var randomBasePoint = &cobra.Command{
	Use:   "random-basepoint",
	Short: "Generate verifying and proving keys.",
	Run:   runRandomBasePoint,
}

func init() {
	rootCmd.AddCommand(randomBasePoint)
}

func runRandomBasePoint(_ *cobra.Command, _ []string) {
	// Pick any valid point on BabyJubJub. Using B8 is standard and in the right subgroup.
	G := babyjub.B8

	// Convert private key to scalar (big-endian) and reduce to subgroup order.
	sk, ok := new(big.Int).SetString(*privKey32, 10)
	if !ok {
		panic("invalid private key")
	}
	sk.Mod(sk, babyjub.SubOrder) // ensure scalar in [0, SubOrder-1]
	if sk.Sign() == 0 {
		panic("invalid private key: scalar is 0 mod SubOrder")
	}

	// Y = sk * G
	Y := babyjub.NewPoint().Mul(sk, G)

	// Safety checks.
	if !Y.InCurve() {
		panic("derived Y is not on curve (should never happen)")
	}
	if !Y.InSubGroup() {
		panic("derived Y is not in subgroup (should never happen)")
	}

	slog.Info("Successfully generated elgamal base points.")
	slog.Info(fmt.Sprintf("Gx: %s", G.X))
	slog.Info(fmt.Sprintf("Gy: %s", G.Y))
	slog.Info(fmt.Sprintf("Yx: %s", Y.X))
	slog.Info(fmt.Sprintf("Yy: %s", Y.Y))
}
