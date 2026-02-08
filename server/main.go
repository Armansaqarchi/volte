package main

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
	"volte/backend/databases"
	"volte/backend/service"
	"volte/backend/utils/test"
)

var (
	host        = flag.String("host", "0.0.0.0", "host")
	port        = flag.Int("port", 8000, "port")
	corsOrigins = flag.String("allow_origins", "*", "cors_origins")
)

func main() {
	//cmd.Execute()
	RunServer()
}

func RunServer() {

	flag.Parse()

	mongoClient := databases.NewMongoClient()
	contractHandler := test.NewFakeContractHandler()
	voteService := service.NewVotingService(mongoClient, contractHandler)
	authService := service.NewAuthService(mongoClient)

	data, err := contractHandler.GetVolteContract().GetEventHash("b1213340-2979-465a-8b0c-549dd8e1380e")
	if err != nil {
		fmt.Println("data is : ", data)
	}

	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(*corsOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Use(sessions.Sessions("user_session", service.NewCookieStore()))

	engine.POST("event/:id/members/:commitment", voteService.AddMemberToEvent)
	engine.DELETE("event/:id/members/:commitment", voteService.RemoveMemberFromEvent)
	engine.POST("event/:id/vote", voteService.Vote)
	engine.POST("event/:id/end", voteService.EndEvent)
	engine.GET("event/:id/tally", voteService.GetTallyScore)
	engine.GET("event/:id/membership/merkle", voteService.MembershipDetails)
	engine.GET("users/events", voteService.UserEvents)
	engine.GET("users/event/:event_id", voteService.UserEvent)
	engine.POST("event/:id/start", voteService.StartEvent)
	engine.DELETE("event/:id", voteService.DeleteEvent)
	engine.POST("users/events", voteService.CreateEvent)
	engine.POST("auth/login", authService.Login)
	engine.POST("auth/signup", authService.Register)

	if err := engine.Run("localhost:8000"); err != nil {
		panic(err)
	}
}

//
//package main
//
//import (
//	"fmt"
//	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
//	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
//	"math/big"
//)
//
//// Your function (unchanged logic)
//func MimcHash(inputs ...[]byte) (string, error) {
//	h := mimc.NewMiMC()
//
//	for _, s := range inputs {
//		// interpret the decimal string as a big integer
//		bi, ok := new(big.Int).SetString(string(s), 10)
//		if !ok {
//			return "", fmt.Errorf("invalid integer: %s", s)
//		}
//
//		var fe fr.Element
//		fe.SetBigInt(bi)
//
//		// feed valid 32-byte field element to MiMC
//		if _, err := h.Write(fe.Marshal()); err != nil {
//			return "", err
//		}
//	}
//
//	var outFe fr.Element
//	outFe.Unmarshal(h.Sum(nil))
//
//	return outFe.BigInt(new(big.Int)).String(), nil
//}

//func main() {
//	// ✅ Replace this with your actual leaf value (base-10 integer string)
//	leaf := "9168677387132714528859032584364821542840185104421751184077481129415029533802"
//	// Hardcoded siblings from --membership_merkle_path
//	s0 := "19032173336092932136418734408026294018775387567429336972205773858295624821833"
//	s1 := "18825718635493704650410119359352107559241766236429836968817521873213080856327"
//	s2 := "17572087677047727358357888234032540740236400731572642089246487149926217978451"
//	s3 := "9193494683462176876432874073818710080295318200293268793422234138158243027454"
//	s4 := "13609234847973908052518826904539012241930003978443755833434216376532945050393"
//	s5 := "3224858484404051024982857888695790669797791411616561934612591707458973022269"
//	s6 := "13665800177833662927402950316876809269732144885929355926431148910311192249775"
//	s7 := "13615279995823067091760603498663829773998124104829923453226442674074652310937"
//
//	// Expected root from --membership_merkle_root
//	expectedRoot := "6504025998332187669354915286309691688052979522962330227128133763251080291251"
//
//	// positions are all 0,0,0,0,0,0,0,0
//	// Using the convention:
//	//   pos=0 => sibling is on the RIGHT => H(curr, sibling)
//	// So we do: curr = H(curr, s_i) at each step.
//
//	var err error
//	curr := leaf
//	//curr, _ = MimcHash([]byte(curr))
//	//fmt.Println(curr)
//	// Step 0
//	curr, err = MimcHash([]byte(s0), []byte(curr))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h0 =", curr)
//
//	// Step 1
//	curr, err = MimcHash([]byte(curr), []byte(s1))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h1 =", curr)
//
//	// Step 2
//	curr, err = MimcHash([]byte(curr), []byte(s2))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h2 =", curr)
//
//	// Step 3
//	curr, err = MimcHash([]byte(curr), []byte(s3))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h3 =", curr)
//
//	// Step 4
//	curr, err = MimcHash([]byte(curr), []byte(s4))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h4 =", curr)
//
//	// Step 5
//	curr, err = MimcHash([]byte(curr), []byte(s5))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h5 =", curr)
//
//	// Step 6
//	curr, err = MimcHash([]byte(curr), []byte(s6))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h6 =", curr)
//
//	// Step 7
//	curr, err = MimcHash([]byte(curr), []byte(s7))
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("h7 (root) =", curr)
//
//	// Compare
//	if curr == expectedRoot {
//		fmt.Println("✅ Root matches expected root")
//	} else {
//		fmt.Println("❌ Root does NOT match expected root")
//		fmt.Println("expected =", expectedRoot)
//		fmt.Println("computed =", curr)
//	}
//
//	var commitments []merkletree.DataBlock
//	for _, member := range []string{"19032173336092932136418734408026294018775387567429336972205773858295624821833",
//		"9168677387132714528859032584364821542840185104421751184077481129415029533802"} {
//		commitments = append(commitments, models.Commitment(member))
//	}
//	// define sha256(0) representing empty nodes that fills total members
//	// up to 2^8 to match the setup arguments.
//
//	emptyNodeVal, _ := utils.MimcHash([]byte("0"))
//	for i := 0; i < int(math.Pow(2, 8))-len(commitments); i++ {
//		commitments = append(commitments, models.Commitment(emptyNodeVal))
//	}
//
//	commitmentsTree, err := merkletree.New(
//		&merkletree.Config{Mode: merkletree.ModeProofGenAndTreeBuild, HashFunc: func(args ...[]byte) ([]byte, error) {
//			hash, err := utils.MimcHash(args...)
//			return []byte(hash), err
//		}, DisableLeafHashing: true,
//		},
//		commitments,
//	)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(string(commitmentsTree.Root))
//}
