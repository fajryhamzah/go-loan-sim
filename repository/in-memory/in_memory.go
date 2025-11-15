package inmemory

import (
	"errors"
	"sync"
	"time"

	"github.com/fajryhamzah/go-loan-sim/constants"
	"github.com/fajryhamzah/go-loan-sim/types"
)

var (
	once            sync.Once
	inMemoryStorage *InMemoryStorage
)

type InMemoryStorage struct {
	UserList        map[string]*types.User
	mapLoanToUserId map[string]string
}

func Init() *InMemoryStorage {
	once.Do(func() {
		inMemoryStorage = &InMemoryStorage{
			UserList:        make(map[string]*types.User),
			mapLoanToUserId: make(map[string]string),
		}
	})

	return inMemoryStorage
}

func (i *InMemoryStorage) AddUser(userId string, name string) error {
	if _, isUserExist := i.UserList[userId]; isUserExist {
		return errors.New("user already exist")
	}

	i.UserList[userId] = &types.User{
		UserID: userId,
		Name:   name,
	}

	return nil
}

func (i *InMemoryStorage) AddLoanToUser(userId string, Loan *types.Loan) error {
	if _, isUserExist := i.UserList[userId]; !isUserExist {
		return errors.New("user does not exist")
	}

	i.UserList[userId].LoanActive = Loan
	i.mapLoanToUserId[Loan.LoanID] = userId

	return nil
}

func (i *InMemoryStorage) GetByUser(userId string) (*types.User, error) {
	if _, isUserExist := i.UserList[userId]; !isUserExist {
		return nil, errors.New("user does not exist")
	}

	return i.UserList[userId], nil
}

func (i *InMemoryStorage) UpdateLoanData(loanId string, loanData *types.Loan, loanSchedule []*types.LoanPaymentSchedule) error {
	userId, isLoanIdExist := i.mapLoanToUserId[loanId]

	if !isLoanIdExist {
		return errors.New("loan does not exist")
	}

	currentLoan, isUserExist := i.UserList[userId]
	if !isUserExist {
		return errors.New("user does not exist")
	}

	if loanData.Status == constants.STATUS_FINISH {
		currentLoan.LoanActive = nil
		currentLoan.LoanHistory = append(currentLoan.LoanHistory, loanData)
		return nil
	}

	i.UserList[userId].LoanActive = loanData

	return nil
}

func (i *InMemoryStorage) GetLoanById(loanId string) (*types.Loan, error) {
	userId, isLoanIdExist := i.mapLoanToUserId[loanId]

	if !isLoanIdExist {
		return nil, errors.New("loan does not exist")
	}

	if _, isUserExist := i.UserList[userId]; !isUserExist {
		return nil, errors.New("user is not exist")
	}

	return i.UserList[userId].LoanActive, nil
}

func (i *InMemoryStorage) GetLoanPaymentByDate(date time.Time) (map[string]*types.Loan, error) {
	listLoan := map[string]*types.Loan{}

	for userId, loan := range i.UserList {
		if loan.LoanActive == nil {
			continue
		}

		for _, loanPayment := range loan.LoanActive.LoanPaymentSchedule {
			if loanPayment.Status == constants.STATUS_PAID || !date.After(loanPayment.DueDate) {
				continue
			}

			listLoan[userId] = loan.LoanActive
		}
	}

	return listLoan, nil
}

func (i *InMemoryStorage) CheckActiveLoanByUserId(userId string) (*types.Loan, error) {
	if _, isUserExist := i.UserList[userId]; !isUserExist {
		return nil, errors.New("user is not exist")
	}

	return i.UserList[userId].LoanActive, nil
}
