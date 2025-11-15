package types

import "time"

type LoanPaymentSchedule struct {
	LoanPaymentScheduleID string
	Amount                int
	PaidAmount            int
	Status                string
	DueDate               time.Time
}

type Loan struct {
	LoanID              string
	PrincipalLoanAmount int
	TotalLoanAmount     int
	PaidAmount          int
	WeeklyPaymentAmount int
	MissPayment         int
	Interest            float32
	Status              string
	LoanDate            time.Time
	LoadEndDate         time.Time
	LoanPaymentSchedule []*LoanPaymentSchedule
}
