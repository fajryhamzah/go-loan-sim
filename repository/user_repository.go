package repository

import (
	"fmt"

	"github.com/fajryhamzah/go-loan-sim/types"
)

type UserRepository interface {
	AddUser(userId string, name string) error
	GetByUser(userId string) (*types.User, error)
}

func InitUserRepoByStorage(storageName string) LoanRepository {
	switch storageName {
	default:
		panic(fmt.Sprintf("Unsupported storage: %s", storageName))
	}
}
