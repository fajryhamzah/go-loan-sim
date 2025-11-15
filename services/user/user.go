package user

import "github.com/fajryhamzah/go-loan-sim/types"

func (u *UserService) AddUser(userId string, name string) error {
	return u.userRepo.AddUser(userId, name)
}

func (u *UserService) GetUserInfo(userId string) (*types.User, error) {
	user, err := u.userRepo.GetByUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) IsDeliquent(userId string) (bool, error) {
	loan, err := u.loanRepo.CheckActiveLoanByUserId(userId)

	if err != nil {
		return false, nil
	}

	if loan == nil {
		return true, nil
	}

	return loan.MissPayment >= 2, nil
}
