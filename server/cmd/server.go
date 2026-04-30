package cmd

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"strings"
	"time"
	"volte/backend/chain"
	"volte/backend/databases"
	"volte/backend/service"
	"volte/backend/utils/test"
)

var (
	host        = flag.String("host", "0.0.0.0", "host")
	port        = flag.Int("port", 8000, "port")
	corsOrigins = flag.String("allow_origins", "*", "cors_origins")
	testEnv     = flag.Bool("test", true, "test indicates whether contract must be in a test env")
)

var serverCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run:   runServer,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) {

	flag.Parse()

	mongoClient := databases.NewMongoClient()
	redisClient := databases.NewRedisClientProvider()
	var contractHandler chain.ContractHandler
	if *testEnv {
		contractHandler = test.NewFakeContractHandler()
	} else {
		contractHandler = chain.NewEthereumChainHandler()
	}
	voteService := service.NewVotingService(mongoClient, contractHandler, redisClient)
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

	if err := engine.Run(fmt.Sprintf("%s:%d", *host, *port)); err != nil {
		panic(err)
	}
}
