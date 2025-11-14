package repository

import (
	"fmt"

	"github.com/fajryhamzah/go-loan-sim/types"
)

// all storage struct must implement this interface
type LoanRepository interface {
	AddLoanToUser(userId string, Loan *types.Loan) error
	UpdateLoanData(userId string, loanData *types.Loan) error
}

func InitLoanRepoByStorage(storageName string) LoanRepository {
	switch storageName {
	default:
		panic(fmt.Sprintf("Unsupported storage: %s", storageName))
	}
}
