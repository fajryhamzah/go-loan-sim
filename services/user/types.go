package user

import (
	"github.com/fajryhamzah/go-loan-sim/repository"
	"github.com/fajryhamzah/go-loan-sim/types"
)

type UserServiceInterface interface {
	AddUser(userId string, name string) error
	GetUserInfo(userId string) (*types.User, error)

	IsDeliquent(userId string) (bool, error)
}

type UserService struct {
	userRepo repository.UserRepository
	loanRepo repository.LoanRepository
}

func NewUserService(userRepo repository.UserRepository, loanRepo repository.LoanRepository) UserServiceInterface {
	return &UserService{
		userRepo: userRepo,
		loanRepo: loanRepo,
	}
}
