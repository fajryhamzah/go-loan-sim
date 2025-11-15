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
	utils.PrintBlue("- loan {args} : to use loan service")
}
