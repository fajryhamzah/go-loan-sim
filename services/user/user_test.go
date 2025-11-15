package user

import (
	"fmt"
	"testing"

	"github.com/fajryhamzah/go-loan-sim/mocks"
	"github.com/fajryhamzah/go-loan-sim/types"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		userName    string
		mockError   error
		expectError bool
	}{
		{
			name:        "should call AddUser and return no error",
			userID:      "123",
			userName:    "Fajry",
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "should call AddUser and return error",
			userID:      "999",
			userName:    "ErrorCase",
			mockError:   fmt.Errorf("repo error"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			mockLoanRepo := mocks.NewMockLoanRepository(ctrl)
			svc := NewUserService(mockUserRepo, mockLoanRepo)

			mockUserRepo.
				EXPECT().
				AddUser(tc.userID, tc.userName).
				Return(tc.mockError).
				Times(1)

			err := svc.AddUser(tc.userID, tc.userName)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetUserInfo(t *testing.T) {
	tests := []struct {
		name        string
		userID      string
		mockUser    *types.User
		mockError   error
		expectError bool
	}{
		{
			name:   "should return user info",
			userID: "123",
			mockUser: &types.User{
				UserID: "123",
				Name:   "Boss",
			},
			mockError:   nil,
			expectError: false,
		},
		{
			name:        "should return error from repo",
			userID:      "123",
			mockUser:    nil,
			mockError:   fmt.Errorf("not found"),
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			mockLoanRepo := mocks.NewMockLoanRepository(ctrl)
			svc := NewUserService(mockUserRepo, mockLoanRepo)

			mockUserRepo.
				EXPECT().
				GetByUser(tc.userID).
				Return(tc.mockUser, tc.mockError).
				Times(1)

			userInfo, err := svc.GetUserInfo(tc.userID)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockUser, userInfo)
			}
		})
	}
}

func TestIsDeliquent(t *testing.T) {
	tests := []struct {
		name           string
		mockLoan       *types.Loan
		mockErr        error
		expectedResult bool
	}{
		{
			name:           "no active loan → delinquent",
			mockLoan:       nil,
			mockErr:        nil,
			expectedResult: true,
		},
		{
			name: "miss payment < 2 → not delinquent",
			mockLoan: &types.Loan{
				MissPayment: 1,
			},
			mockErr:        nil,
			expectedResult: false,
		},
		{
			name: "miss payment >= 2 → delinquent",
			mockLoan: &types.Loan{
				MissPayment: 2,
			},
			mockErr:        nil,
			expectedResult: true,
		},
		{
			name:           "repo returns error → treat as not delinquent per business rule",
			mockLoan:       nil,
			mockErr:        fmt.Errorf("db error"),
			expectedResult: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockUserRepo := mocks.NewMockUserRepository(ctrl)
			mockLoanRepo := mocks.NewMockLoanRepository(ctrl)
			svc := NewUserService(mockUserRepo, mockLoanRepo)

			mockLoanRepo.
				EXPECT().
				CheckActiveLoanByUserId("123").
				Return(tc.mockLoan, tc.mockErr).
				Times(1)

			result, err := svc.IsDeliquent("123")

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
