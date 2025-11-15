package types

type User struct {
	UserID      string
	Name        string
	LoanActive  *Loan
	LoanHistory []*Loan
}
