package service

type OnChainHandler interface {
	// methods must obey certain name conventions.
	// the same method names must be defined inside the contract
}

type AuthService struct {
}

func (s *AuthService) sendEmail() {

}

func (s *AuthService) VerifyEmail() {

}

func (s *AuthService) Register() {

}
