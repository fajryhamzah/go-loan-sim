package loan

import (
	"errors"
	"testing"
	"time"

	"github.com/fajryhamzah/go-loan-sim/mocks"
	"github.com/fajryhamzah/go-loan-sim/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLoanService_MakePayment(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockLoanRepository(ctrl)
	loanService := &LoanService{
		loanRepo: mockRepo,
	}

	now := time.Now()
	// base loan used in most test cases
	baseLoan := func() *types.Loan {
		return &types.Loan{
			LoanID:              "L1",
			PrincipalLoanAmount: 1000,
			TotalLoanAmount:     1200,
			Status:              "ACTIVE",
			PaidAmount:          0,
			WeeklyPaymentAmount: 100,
			LoanPaymentSchedule: []*types.LoanPaymentSchedule{
				{
					LoanPaymentScheduleID: "P1",
					Amount:                100,
					Status:                "PENDING",
					DueDate:               now.AddDate(0, 0, -1), // due yesterday
				},
			},
		}
	}

	tests := []struct {
		name        string
		loanId      string
		amount      int
		setupMock   func()
		expectError bool
	}{
		{
			name:   "should succeed and pay the first installment",
			loanId: "L1",
			amount: 100,
			setupMock: func() {
				mockLoan := baseLoan()
				mockRepo.EXPECT().GetLoanById("L1").Return(mockLoan, nil)
				mockRepo.EXPECT().UpdateLoanData("L1", gomock.Any(), gomock.Any()).Return(nil)
			},
			expectError: false,
		},
		{
			name:   "should fail if repo GetLoanById returns error",
			loanId: "L1",
			amount: 100,
			setupMock: func() {
				mockRepo.EXPECT().GetLoanById("L1").Return(nil, errors.New("db error"))
			},
			expectError: true,
		},
		{
			name:   "should fail if loan status is not active",
			loanId: "L1",
			amount: 100,
			setupMock: func() {
				mockLoan := baseLoan()
				mockLoan.Status = "FINISHED"
				mockRepo.EXPECT().GetLoanById("L1").Return(mockLoan, nil)
			},
			expectError: true,
		},
		{
			name:   "should fail if amount is not equal to weekly payment",
			loanId: "L1",
			amount: 50,
			setupMock: func() {
				mockLoan := baseLoan()
				mockRepo.EXPECT().GetLoanById("L1").Return(mockLoan, nil)
			},
			expectError: true,
		},
		{
			name:   "should fail if UpdateLoanData returns error",
			loanId: "L1",
			amount: 100,
			setupMock: func() {
				mockLoan := baseLoan()
				mockRepo.EXPECT().GetLoanById("L1").Return(mockLoan, nil)
				mockRepo.EXPECT().UpdateLoanData("L1", gomock.Any(), gomock.Any()).Return(errors.New("db error"))
			},
			expectError: true,
		},
		{
			name:   "should fail if no payment schedule to pay",
			loanId: "L1",
			amount: 100,
			setupMock: func() {
				mockLoan := baseLoan()
				// make the due date in the future so no installment is ready
				mockLoan.LoanPaymentSchedule[0].DueDate = now.AddDate(0, 0, 1)
				mockRepo.EXPECT().GetLoanById("L1").Return(mockLoan, nil)
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock()
			loan, err := loanService.MakePayment(tc.loanId, tc.amount)
			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, loan)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, loan)
				assert.Equal(t, tc.amount, loan.PaidAmount)
				assert.Equal(t, "PAID", loan.LoanPaymentSchedule[0].Status)
				assert.Equal(t, tc.amount, loan.LoanPaymentSchedule[0].PaidAmount)
			}
		})
	}
}
