package service

import (
	"flag"
	"fmt"
	"log/slog"
	"net/smtp"
)

var (
	smtpHost     = flag.String("smtp_host", "smtp.google.com", "SMTP host")
	smtpPort     = flag.String("smtp_port", "587", "SMTP port")
	smtpUsername = flag.String("smtp_user", "", "SMTP username")
	smtpPassword = flag.String("smtp_password", "", "SMTP password")
)

type AuthService struct {
	smtpAuth smtp.Auth
}

func NewAuthService() *AuthService {
	return &AuthService{
		smtpAuth: smtp.PlainAuth("", *smtpUsername, *smtpPassword, *smtpHost),
	}
}

func (s *AuthService) sendEmail(to string, body []byte) error {
	err := smtp.SendMail(*smtpHost+":"+*smtpPort, s.smtpAuth, "volte", []string{to}, body)
	if err != nil {
		slog.Error(fmt.Sprintf("Send email error: %v.", err))
	}
	return nil
}

func (s *AuthService) VerifyEmail() {

}

func (s *AuthService) Register() {

}
