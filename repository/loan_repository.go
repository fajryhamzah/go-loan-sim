package repository

import (
	"fmt"
	"time"

	"github.com/fajryhamzah/go-loan-sim/types"
)

// all storage struct must implement this interface
//
//go:generate mockgen -source=loan_repository.go -destination=../mocks/mock_loan_repository.go -package=mocks
type LoanRepository interface {
	AddLoanToUser(userId string, Loan *types.Loan) error
	UpdateLoanData(loanId string, loanData *types.Loan, loanSchedule []*types.LoanPaymentSchedule) error
	GetLoanById(loanId string) (*types.Loan, error)
	GetLoanPaymentByDate(date time.Time) (map[string]*types.Loan, error)
	CheckActiveLoanByUserId(userId string) (*types.Loan, error)
}

func InitLoanRepoByStorage(storageName string) LoanRepository {
	switch storageName {
	default:
		panic(fmt.Sprintf("Unsupported storage: %s", storageName))
	}
}
