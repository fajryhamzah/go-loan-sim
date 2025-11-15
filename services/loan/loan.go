package loan

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/fajryhamzah/go-loan-sim/constants"
	"github.com/fajryhamzah/go-loan-sim/types"
	"github.com/fajryhamzah/go-loan-sim/utils"
	"github.com/google/uuid"
)

func (l *LoanService) AddLoan(userId string, amount int) (*types.Loan, error) {
	var (
		loanData     *types.Loan
		installments []*types.LoanPaymentSchedule
	)

	loanActive, err := l.loanRepo.CheckActiveLoanByUserId(userId)

	if err != nil {
		return nil, err
	}

	if loanActive != nil {
		return nil, errors.New("can't create new loan, user still has active loan")
	}
	totalLoanAmount := int(math.Round(float64(constants.FLAT_INTEREST)/100*float64(amount) + float64(amount)))
	weeklyPayment := totalLoanAmount / constants.WEEK_LOAN
	remainder := totalLoanAmount % constants.WEEK_LOAN
	now := utils.Now()

	loanData = &types.Loan{
		LoanID:              uuid.NewString(),
		PrincipalLoanAmount: amount,
		TotalLoanAmount:     totalLoanAmount,
		Interest:            constants.FLAT_INTEREST,
		Status:              constants.STATUS_ACTIVE,
		LoanDate:            now,
		WeeklyPaymentAmount: weeklyPayment,
	}

	installments = []*types.LoanPaymentSchedule{}

	for i := 1; i <= constants.WEEK_LOAN; i++ {
		payment := weeklyPayment
		if i <= remainder {
			payment += 1
		}

		installments = append(installments, &types.LoanPaymentSchedule{
			LoanPaymentScheduleID: uuid.NewString(),
			Amount:                payment,
			Status:                constants.STATUS_PENDING,
			DueDate:               now,
		})

		now = now.AddDate(0, 0, 7)
	}

	loanData.LoadEndDate = now
	loanData.LoanPaymentSchedule = installments

	return loanData, l.loanRepo.AddLoanToUser(userId, loanData)
}

func (l *LoanService) MakePayment(loanId string, amount int) (*types.Loan, error) {
	loanActive, err := l.loanRepo.GetLoanById(loanId)

	if err != nil {
		return nil, err
	}

	if loanActive == nil {
		return nil, errors.New("user does not have any loan")
	}

	now := utils.Now()

	listPaidLoans := []*types.LoanPaymentSchedule{}

	for _, loanForPay := range loanActive.LoanPaymentSchedule {
		if loanForPay.Status == constants.STATUS_PAID || !now.After(loanForPay.DueDate) {
			continue
		}

		if amount != loanForPay.Amount {
			return nil, fmt.Errorf("can't pay less or more than %d", loanForPay.Amount)
		}

		loanForPay.PaidAmount = amount
		loanForPay.Status = constants.STATUS_PAID

		listPaidLoans = append(listPaidLoans, loanForPay)
		break
	}

	if len(listPaidLoans) < 1 {
		return nil, errors.New("there is no payment schedule to pay")
	}

	loanActive.PaidAmount += amount

	if loanActive.MissPayment >= 1 {
		loanActive.MissPayment -= 1
	}

	if loanActive.PaidAmount == loanActive.TotalLoanAmount {
		loanActive.Status = constants.STATUS_FINISH
	}

	err = l.loanRepo.UpdateLoanData(loanId, loanActive, listPaidLoans)

	if err != nil {
		return nil, err
	}

	return loanActive, nil
}

func (l *LoanService) GetLoanInfo(loanId string) (*types.Loan, error) {
	loanActive, err := l.loanRepo.GetLoanById(loanId)

	if err != nil {
		return nil, err
	}

	if loanActive == nil {
		return nil, nil
	}

	return loanActive, nil
}

func (l *LoanService) ListLoanPaymentThisWeek() (map[string]*types.Loan, error) {
	return l.loanRepo.GetLoanPaymentByDate(utils.Now())
}

func (l *LoanService) GetOutstanding(loanId string) (int, error) {
	loan, err := l.loanRepo.GetLoanById(loanId)

	if err != nil {
		return 0, err
	}

	return loan.TotalLoanAmount - loan.PaidAmount, nil
}

func (l *LoanService) WeeklyLoanProcess() error {
	overdueDate := utils.Now().Add(-7 * 24 * time.Hour) // for checking overdue

	loanList, err := l.loanRepo.GetLoanPaymentByDate(overdueDate)

	if err != nil {
		return err
	}

	for _, loan := range loanList {
		missPaymentCounter := 0
		updatedSchedule := []*types.LoanPaymentSchedule{}

		for _, loanSchedule := range loan.LoanPaymentSchedule {
			if overdueDate.Before(loanSchedule.DueDate) || loanSchedule.Status == constants.STATUS_PAID {
				continue
			}

			missPaymentCounter += 1
			loanSchedule.Status = constants.STATUS_MISS_PAYMENT
			updatedSchedule = append(updatedSchedule, loanSchedule)
		}

		loan.MissPayment = missPaymentCounter
		err = l.loanRepo.UpdateLoanData(loan.LoanID, loan, updatedSchedule)

		if err != nil {
			return err
		}
	}

	return nil
}

func (l *LoanService) GetActiveLoanInfoByUserId(userId string) (*types.Loan, error) {
	loanActive, err := l.loanRepo.CheckActiveLoanByUserId(userId)

	if err != nil {
		return nil, err
	}

	if loanActive == nil {
		return nil, nil
	}

	return loanActive, nil
}
