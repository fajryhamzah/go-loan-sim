package cli

import (
	"fmt"

	"github.com/fajryhamzah/go-loan-sim/constants"
	"github.com/fajryhamzah/go-loan-sim/services/loan"
	"github.com/fajryhamzah/go-loan-sim/utils"
)

func TimeHandler(loanService loan.LoanServiceInterface, inputs ...string) {
	if len(inputs) < 1 {
		utils.PrintRed("need more args", inputs)
		return
	}

	switch inputs[0] {
	case "nextweek":
		utils.PrintInlineBlue("Simulation Time: ")
		utils.PrintGreen(utils.NowFormatted())
		utils.GetSimulationTime().AddWeek()
		utils.PrintInlineBlue("Simulation Time After: ")
		utils.PrintGreen(utils.NowFormatted())

		err := loanService.WeeklyLoanProcess()

		if err != nil {
			utils.PrintRed("Weekly process: ", err)
		}
		loans, err := loanService.ListLoanPaymentThisWeek()

		if err != nil {
			utils.PrintRed("Weekly procee list due loan err: ", err)
			return
		}

		utils.PrintBlue("--------------------------------------------")
		utils.PrintBlue("NEED Payment")
		utils.PrintBlue("--------------------------------------------")
		for userId, loanActive := range loans {
			utils.PrintInlineBlue("User ID: ")
			utils.PrintGreen(userId)
			utils.PrintInlineBlue("Loan ID: ")
			utils.PrintGreen(loanActive.LoanID)
			utils.PrintInlineBlue("Miss payment: ")
			utils.PrintGreen(fmt.Sprintf("%dx", loanActive.MissPayment))
			utils.PrintInlineBlue("Weekly payment: ")
			utils.PrintGreen(utils.FormatRupiah(loanActive.WeeklyPaymentAmount))
			totalDueAmount := 0
			for _, loanSchedule := range loanActive.LoanPaymentSchedule {
				if loanSchedule.Status != constants.STATUS_MISS_PAYMENT {
					continue
				}

				totalDueAmount += loanSchedule.Amount
			}
			utils.PrintInlineBlue("Total due amount: ")
			utils.PrintRed(utils.FormatRupiah(totalDueAmount))
			utils.PrintBlue("--------------------------------------------")
		}
	default:
		utils.PrintInlineBlue("Simulation Time now: ")
		utils.PrintGreen(utils.Now())
	}
}
