package cli

import (
	"fmt"
	"strconv"

	"github.com/fajryhamzah/go-loan-sim/constants"
	"github.com/fajryhamzah/go-loan-sim/services/loan"
	"github.com/fajryhamzah/go-loan-sim/utils"
)

func LoanHandler(loanService loan.LoanServiceInterface, inputs ...string) {
	if len(inputs) < 1 {
		utils.PrintRed("need more args", inputs)
		return
	}

	switch inputs[0] {
	case "add":
		if len(inputs) != 3 {
			utils.PrintRed("Need more args")
			return
		}

		amount, err := strconv.Atoi(inputs[2])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		loan, err := loanService.AddLoan(inputs[1], amount)

		if err != nil {
			utils.PrintRed("Add loan err: ", err)
			return
		}

		utils.PrintInlineBlue("Loan ID: ")
		utils.PrintGreen(loan.LoanID)

		utils.PrintGreen("done.")
	case "infobyuser":
		if len(inputs) != 2 {
			utils.PrintRed("Need more args")
			return
		}
		utils.PrintInlineBlue("Simulation Time: ")
		utils.PrintGreen(utils.NowFormatted())

		loanActive, err := loanService.GetActiveLoanInfoByUserId(inputs[1])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		if loanActive == nil {
			utils.PrintGreen("User does not have active loan.")
			return
		}

		utils.PrintInlineBlue("Loan ID: ")
		utils.PrintGreen(loanActive.LoanID)
		utils.PrintInlineBlue("Principal: ")
		utils.PrintGreen(utils.FormatRupiah(loanActive.PrincipalLoanAmount))
		utils.PrintInlineBlue("Total loan amount + interest: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.TotalLoanAmount))
		utils.PrintInlineBlue("Loan At: ")
		utils.PrintGreen(utils.Format(loanActive.LoanDate))
		utils.PrintInlineBlue("Expected Loan Finish At: ")
		utils.PrintGreen(utils.Format(loanActive.LoadEndDate))
		utils.PrintInlineBlue("Paid amount: ")
		utils.PrintGreen(utils.FormatRupiah(loanActive.PaidAmount))
		utils.PrintInlineBlue("Remaining loan: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.TotalLoanAmount - loanActive.PaidAmount))
		utils.PrintInlineBlue("Weekly Payment: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.WeeklyPaymentAmount))
		utils.PrintInlineBlue("Installment: ")
		utils.PrintRed(fmt.Sprintf("%dx", len(loanActive.LoanPaymentSchedule)))
		utils.PrintInlineBlue("Miss Payment: ")
		utils.PrintRed(fmt.Sprintf("%dx", loanActive.MissPayment))

		if loanActive.MissPayment > 0 {
			utils.PrintBlue("Missing installment: ")

			for idx, loanSchedule := range loanActive.LoanPaymentSchedule {
				if loanSchedule.Status != constants.STATUS_MISS_PAYMENT {
					continue
				}

				utils.PrintInlineBlue("	- Installment: ")
				utils.PrintGreen(idx + 1)
				utils.PrintInlineBlue("	  Due Date: ")
				utils.PrintGreen(utils.Format(loanSchedule.DueDate))
			}
		}
	case "info":
		if len(inputs) != 2 {
			utils.PrintRed("Need more args")
			return
		}

		utils.PrintInlineBlue("Simulation Time: ")
		utils.PrintGreen(utils.NowFormatted())

		loanActive, err := loanService.GetLoanInfo(inputs[1])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		if loanActive == nil {
			utils.PrintGreen("User does not have loan.")
			return
		}

		utils.PrintInlineBlue("Principal: ")
		utils.PrintGreen(utils.FormatRupiah(loanActive.PrincipalLoanAmount))
		utils.PrintInlineBlue("Total loan amount + interest: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.TotalLoanAmount))
		utils.PrintInlineBlue("Loan At: ")
		utils.PrintGreen(utils.Format(loanActive.LoanDate))
		utils.PrintInlineBlue("Expected Loan Finish At: ")
		utils.PrintGreen(utils.Format(loanActive.LoadEndDate))
		utils.PrintInlineBlue("Paid amount: ")
		utils.PrintGreen(utils.FormatRupiah(loanActive.PaidAmount))
		utils.PrintInlineBlue("Remaining loan: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.TotalLoanAmount - loanActive.PaidAmount))
		utils.PrintInlineBlue("Weekly Payment: ")
		utils.PrintRed(utils.FormatRupiah(loanActive.WeeklyPaymentAmount))
		utils.PrintInlineBlue("Installment: ")
		utils.PrintRed(fmt.Sprintf("%dx", len(loanActive.LoanPaymentSchedule)))
		utils.PrintInlineBlue("Miss Payment: ")
		utils.PrintRed(fmt.Sprintf("%dx", loanActive.MissPayment))

		if loanActive.MissPayment > 0 {
			utils.PrintBlue("Missing installment: ")

			for idx, loanSchedule := range loanActive.LoanPaymentSchedule {
				if loanSchedule.Status != constants.STATUS_MISS_PAYMENT {
					continue
				}

				utils.PrintInlineBlue("	- Installment: ")
				utils.PrintGreen(idx + 1)
				utils.PrintInlineBlue("	  Due Date: ")
				utils.PrintGreen(utils.Format(loanSchedule.DueDate))
			}
		}
	case "pay":
		if len(inputs) != 3 {
			utils.PrintRed("Need more args")
			return
		}

		amount, err := strconv.Atoi(inputs[2])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		_, err = loanService.MakePayment(inputs[1], amount)

		if err != nil {
			utils.PrintRed("Pay loan err: ", err)
			return
		}

		utils.PrintGreen("done.")
	case "due":
		loans, err := loanService.ListLoanPaymentThisWeek()

		if err != nil {
			utils.PrintRed("List due loan err: ", err)
			return
		}

		utils.PrintBlue("--------------------------------------------")
		utils.PrintBlue("NEED Payment This Week")
		utils.PrintBlue("--------------------------------------------")
		for userId, loanActive := range loans {
			utils.PrintInlineBlue("User ID: ")
			utils.PrintGreen(userId)
			utils.PrintInlineBlue("Weekly payment: ")
			utils.PrintGreen(utils.FormatRupiah(loanActive.WeeklyPaymentAmount))
			utils.PrintInlineBlue("Total due amount: ")
			payment := loanActive.WeeklyPaymentAmount + loanActive.MissPayment*loanActive.WeeklyPaymentAmount
			utils.PrintRed(utils.FormatRupiah(payment))
			utils.PrintBlue("--------------------------------------------")
		}
	case "outstanding":
		if len(inputs) != 2 {
			utils.PrintRed("Need more args")
			return
		}

		outstanding, err := loanService.GetOutstanding(inputs[1])

		if err != nil {
			utils.PrintRed("Get outstanding loan err: ", err)
			return
		}

		utils.PrintInlineBlue("Current Outstanding: ")
		utils.PrintRed(utils.FormatRupiah(outstanding))
	default:
		utils.PrintRed("Unknown cmd")
	}
}
