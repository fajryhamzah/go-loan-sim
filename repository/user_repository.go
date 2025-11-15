package repository

import (
	"fmt"

	"github.com/fajryhamzah/go-loan-sim/constants"
	inmemory "github.com/fajryhamzah/go-loan-sim/repository/in-memory"
	"github.com/fajryhamzah/go-loan-sim/types"
)

//go:generate mockgen -source=user_repository.go -destination=../mocks/mock_user_repository.go -package=mocks
type UserRepository interface {
	AddUser(userId string, name string) error
	GetByUser(userId string) (*types.User, error)
}

func InitUserRepoByStorage(storageName string) UserRepository {
	switch storageName {
	case constants.IN_MEMORY_STORAGE:
		return inmemory.Init()
	default:
		panic(fmt.Sprintf("Unsupported storage: %s", storageName))
	}
}
