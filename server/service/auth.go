package service

import (
	"flag"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log/slog"
	"net/http"
	"net/smtp"
	"volte/backend/databases"
	"volte/backend/models"
)

var (
	// Email confs.
	smtpHost     = flag.String("smtp_host", "smtp.google.com", "SMTP host")
	smtpPort     = flag.String("smtp_port", "587", "SMTP port")
	smtpUsername = flag.String("smtp_user", "", "SMTP username")
	smtpPassword = flag.String("smtp_password", "", "SMTP password")
	// DB confs.
	usersCollection = flag.String("users_collection", "users", "Users collection")
	// Session conf.
	sessionSecret = flag.String("session_secret", "dummy_session", "User's session secret key.")
)

func NewCookieStore() cookie.Store {
	return cookie.NewStore([]byte(*sessionSecret))
}

type AuthService struct {
	smtpAuth    smtp.Auth
	mongoClient *databases.MongoClient
}

func NewAuthService(mongoClient *databases.MongoClient) *AuthService {
	return &AuthService{
		smtpAuth:    smtp.PlainAuth("", *smtpUsername, *smtpPassword, *smtpHost),
		mongoClient: mongoClient,
	}
}

func (s *AuthService) sendEmail(to string, body []byte) error {
	err := smtp.SendMail(*smtpHost+":"+*smtpPort, s.smtpAuth, "volte", []string{to}, body)
	if err != nil {
		slog.Error(fmt.Sprintf("Send email error: %v.", err))
	}
	return nil
}

func (s *AuthService) VerifyEmail() bool {
	return true
}

func (s *AuthService) Register(ctx *gin.Context) {

	usersCollection := s.mongoClient.GetClient().Database(*database).Collection(*usersCollection)
	var user models.User
	slog.Info("Binding and verifying user's info.")
	if err := ctx.ShouldBind(&user); err != nil {
		slog.Error("Failed to bind user info," + err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	slog.Info("User's binding is done.")
	var err error
	user.Password, err = user.HashPassword()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	slog.Info("Saving user.")
	if res, err := usersCollection.UpdateOne(
		ctx, bson.M{"_id": user.Commitment}, bson.M{"$setOnInsert": user}, options.UpdateOne().SetUpsert(true),
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		if res.MatchedCount == 0 {
			if err := s.Session(ctx, user); err != nil {
				slog.Error(fmt.Sprintf("Something went wrong while creating session, err : %s", err.Error()))
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while authenticating."})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"message": "User has been successfully registered.",
				"data":    gin.H{"username": user.Username, "commitment": user.Commitment}},
			)
			return
		} else {
			ctx.JSON(http.StatusAlreadyReported, gin.H{"error": "User already exists."})
		}
	}
}

func (s *AuthService) Login(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.authenticate(ctx, user); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"username": user.Username, "commitment": user.Commitment})
}

func (s *AuthService) authenticate(ctx *gin.Context, user models.User) error {
	usersCollection := s.mongoClient.GetClient().Database(*database).Collection(*usersCollection)
	hashedPassword, err := user.HashPassword()
	if err != nil {
		return err
	}

	if err := usersCollection.FindOne(ctx,
		bson.M{"username": user.Username, "password": hashedPassword, "commitment": user.Commitment},
	).Decode(&user); err != nil {
		slog.Error(fmt.Sprintf("Didnt find any user with %v", user))
		return err
	}

	return s.Session(ctx, user)
}

func (s *AuthService) Session(ctx *gin.Context, user models.User) error {
	store := cookie.NewStore([]byte(*sessionSecret))
	sessionName := "user_session"
	sessionMiddleware := sessions.Sessions(sessionName, store)
	sessionMiddleware(ctx)
	session := sessions.Default(ctx)
	session.Set("user", user.Commitment)
	if err := session.Save(); err != nil {
		return err
	}
	return nil
}
