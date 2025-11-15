package cli

import (
	"github.com/fajryhamzah/go-loan-sim/services/user"
	"github.com/fajryhamzah/go-loan-sim/utils"
)

func UsersHandler(userService user.UserServiceInterface, inputs ...string) {
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

		err := userService.AddUser(inputs[1], inputs[2])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		utils.PrintGreen("done.")
	case "get":
		if len(inputs) != 2 {
			utils.PrintRed("Need more args")
			return
		}

		user, err := userService.GetUserInfo(inputs[1])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		utils.PrintInlineBlue("User ID: ")
		utils.PrintGreen(user.UserID)
		utils.PrintInlineBlue("Name: ")
		utils.PrintGreen(user.Name)
		utils.PrintInlineBlue("Has Active Loan: ")

		if user.LoanActive != nil {
			utils.PrintRed("Yes")
		} else {
			utils.PrintGreen("No")
		}

		utils.PrintInlineBlue("Loan History Count: ")
		utils.PrintGreen(len(user.LoanHistory))
	case "isdeliquent":
		if len(inputs) != 2 {
			utils.PrintRed("Need more args")
			return
		}

		isDeliquent, err := userService.IsDeliquent(inputs[1])

		if err != nil {
			utils.PrintRed(err)
			return
		}

		utils.PrintInlineBlue("Is Deliquent: ")

		if isDeliquent {
			utils.PrintRed("Yes")
		} else {
			utils.PrintGreen("No")
		}
	default:
		utils.PrintRed("Unknown cmd")
	}
}
