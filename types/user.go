package types

type User struct {
	Name        string
	LoanActive  *Loan
	LoanHistory []*Loan
}
