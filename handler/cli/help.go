package cli

import "github.com/fajryhamzah/go-loan-sim/utils"

func HelpHandler() {
	utils.PrintBlue("Supported Command:")
	utils.PrintBlue("- help : to get list of the command")
	utils.PrintBlue("- time : control simulation time")
	utils.PrintBlue("	- time now : current time in simulation")
	utils.PrintBlue("	- time nextweek : skip time in simulation to next week (adding 7 days)")
	utils.PrintBlue("- users {args} : to use user service")
	utils.PrintBlue("	- users add {userID} {name} : add new user")
	utils.PrintBlue("	- users get {userID} : to get user info by id")
	utils.PrintBlue("	- users isdeliquent {userID} : check if the user is deliquent or not")
	utils.PrintBlue("- loan {args} : to use loan service")
	utils.PrintBlue("	- loan add {userID} {amount} : add new loan amount to the user")
	utils.PrintBlue("	- loan pay {loanID} {amount} : user pay the loan")
	utils.PrintBlue("	- loan info {loanID} : complete information about the loan")
	utils.PrintBlue("	- loan infobyuser {userID} : complete information about the active loan by user id")
	utils.PrintBlue("	- loan outstanding {loanID} : get current outstanding of the loan")
	utils.PrintBlue("	- loan due : list of loan that need to have payment that week")

}
