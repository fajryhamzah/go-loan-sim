package repository

import (
	"fmt"

	"github.com/fajryhamzah/go-loan-sim/types"
)

// all storage struct must implement this interface
type RepoInterface interface {
	AddUser(userId string) error
	AddLoanToUser(userId string, amount int, interest int) error
	GetByUser(userId string) (*types.User, error)
}

func InitRepoByStorage(storageName string) RepoInterface {
	switch storageName {
	default:
		panic(fmt.Sprintf("Unsupported storage: %s", storageName))
	}
}
