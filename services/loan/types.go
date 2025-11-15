package loan

import (
	"github.com/fajryhamzah/go-loan-sim/repository"
	"github.com/fajryhamzah/go-loan-sim/types"
)

type LoanServiceInterface interface {
	AddLoan(userId string, amount int) (*types.Loan, error)
	MakePayment(loanId string, amount int) (*types.Loan, error)
	GetLoanInfo(loanId string) (*types.Loan, error)
	ListLoanPaymentThisWeek() (map[string]*types.Loan, error)
	GetOutstanding(loanId string) (int, error)
	GetActiveLoanInfoByUserId(userId string) (*types.Loan, error)

	WeeklyLoanProcess() error
}

type LoanService struct {
	loanRepo repository.LoanRepository
}

func NewLoanService(loanRepo repository.LoanRepository) LoanServiceInterface {
	return &LoanService{
		loanRepo: loanRepo,
	}
}
