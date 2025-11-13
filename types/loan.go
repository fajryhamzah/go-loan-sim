package types

import "time"

type LoanPaymentSchedule struct {
	Amount     int
	PaidAmount int
	Status     string
	DueDate    time.Time
}

type Loan struct {
	PrincipalLoanAmount int
	TotalLoanAmount     int
	PaidAmount          int
	PendingAmount       int
	MissPayment         int
	Interest            float32
	LoanDate            time.Time
	LoadEndDate         time.Time
	LoanPaymentSchedule []LoanPaymentSchedule
}
