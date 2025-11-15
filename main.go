package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fajryhamzah/go-loan-sim/constants"
	"github.com/fajryhamzah/go-loan-sim/handler/cli"
	"github.com/fajryhamzah/go-loan-sim/repository"
	"github.com/fajryhamzah/go-loan-sim/services/loan"
	"github.com/fajryhamzah/go-loan-sim/services/user"
	"github.com/fajryhamzah/go-loan-sim/utils"
)

func main() {
	loanRepo := repository.InitLoanRepoByStorage(constants.IN_MEMORY_STORAGE)
	userRepo := repository.InitUserRepoByStorage(constants.IN_MEMORY_STORAGE)

	loanSrv := loan.NewLoanService(loanRepo)
	userSrv := user.NewUserService(userRepo, loanRepo)

	reader := bufio.NewReader(os.Stdin)

	for {
		utils.PrintInlineGreen("> Enter command: ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)

		input := strings.Split(strings.ToLower(cmd), " ")

		switch input[0] {
		case "help":
			cli.HelpHandler()
		case "user":
			cli.UsersHandler(userSrv, input[1:]...)
		case "loan":
			cli.LoanHandler(loanSrv, input[1:]...)
		case "time":
			cli.TimeHandler(loanSrv, input[1:]...)
		case "exit":
			utils.PrintBlue("Byebyee")
			os.Exit(0)
		default:
			utils.PrintRed("Unsupported command.")
		}
	}

}
