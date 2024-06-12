package authsvc

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/amnestia/tnderlike/internal/domain/constant"
	authmodel "github.com/amnestia/tnderlike/internal/domain/model/auth"
	"github.com/amnestia/tnderlike/internal/domain/model/common"
	"github.com/amnestia/tnderlike/internal/lib/paseto"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	var (
		ctx     = context.Background()
		mockErr = errors.New("mock error")
		tx      = &sqlx.Tx{}
		mockAcc = &authmodel.Account{
			Email:    "email@email.email",
			Password: "testpassword",
		}

		tmpHandleGenerateHash = handleGenerateHash
	)
	defer func() {
		handleGenerateHash = tmpHandleGenerateHash
	}()
	type args struct {
		ctx context.Context
		acc *authmodel.Account
	}
	type expect struct {
		resp *common.DefaultResponse
		err  bool
	}
	tests := []struct {
		name   string
		args   args
		expect expect
		patch  func(args)
	}{
		{
			name: "Register_GenerateHash_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusInternalServerError,
					Error:    mockErr,
				},
				err: true,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "", mockErr
				}
			},
		},
		{
			name: "Register_NewTransaction_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusInternalServerError,
					Error:    mockErr,
				},
				err: true,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "hashedpassword", nil
				}
				mockRepo.On("NewTransaction", a.ctx).Return(nil, mockErr)
			},
		},
		{
			name: "Register_RegisterNewAccount_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusInternalServerError,
					Error:    mockErr,
				},
				err: true,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "hashedpassword", nil
				}
				mockRepo.On("NewTransaction", a.ctx).Return(tx, nil)
				tmp := mockAcc
				tmp.Password = "hashedpassword"
				mockRepo.On("RegisterNewAccount", a.ctx, tx, tmp).Return(int64(0), mockErr)
				mockRepo.On("RollbackOnError", tx, mockErr).Return(mockErr)
			},
		},
		{
			name: "Register_RegisterNewAccount_ReturnErrorDup",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusBadRequest,
					Error:    errors.New("Email have been registered"),
				},
				err: true,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "hashedpassword", nil
				}
				mockRepo.On("NewTransaction", a.ctx).Return(tx, nil)
				tmp := mockAcc
				tmp.Password = "hashedpassword"
				errorDup := errors.New("duplicated data")
				mockRepo.On("RegisterNewAccount", a.ctx, tx, tmp).Return(int64(0), errorDup)
				mockRepo.On("RollbackOnError", tx, errorDup).Return(errorDup)
			},
		},
		{
			name: "Register_Commit_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusInternalServerError,
					Error:    mockErr,
				},
				err: true,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "hashedpassword", nil
				}
				mockRepo.On("NewTransaction", a.ctx).Return(tx, nil)
				tmp := mockAcc
				tmp.Password = "hashedpassword"
				mockRepo.On("RegisterNewAccount", a.ctx, tx, tmp).Return(int64(1), nil)
				mockRepo.On("RollbackOnError", tx, nil).Return(mockErr)
				mockRepo.On("Commit", tx).Return(mockErr)
			},
		},
		{
			name: "Register_NoError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &common.DefaultResponse{
					HTTPCode: http.StatusCreated,
				},
				err: false,
			},
			patch: func(a args) {
				handleGenerateHash = func(s, pepper string) (string, error) {
					return "hashedpassword", nil
				}
				mockRepo.On("NewTransaction", a.ctx).Return(tx, nil)
				tmp := mockAcc
				tmp.Password = "hashedpassword"
				mockRepo.On("RegisterNewAccount", a.ctx, tx, tmp).Return(int64(1), nil)
				mockRepo.On("RollbackOnError", tx, nil).Return(nil)
				mockRepo.On("Commit", tx).Return(nil)
			},
		},
	}

	for _, test := range tests {
		svc := getMock()
		test.patch(test.args)
		t.Run(test.name, func(t *testing.T) {
			resp := svc.Register(test.args.ctx, test.args.acc)
			assert.Equal(t, test.expect.resp, resp)
			if test.expect.err {
				assert.Error(t, resp.Error)
			} else {
				assert.NoError(t, resp.Error)

			}

		})
	}
}

func TestAuth(t *testing.T) {
	var (
		ctx     = context.Background()
		mockErr = errors.New("mock error")
		mockAcc = &authmodel.Account{
			ID:       1,
			Email:    "email@email.email",
			Password: "testpassword",
		}

		tmpHandleVerifyHash = handleVerifyHash
	)
	defer func() {
		handleVerifyHash = tmpHandleVerifyHash
	}()
	type args struct {
		ctx context.Context
		acc *authmodel.Account
	}
	type expect struct {
		resp *authmodel.LoginResponse
		err  bool
	}
	tests := []struct {
		name   string
		args   args
		expect expect
		patch  func(args)
	}{
		{
			name: "Auth_AuthError_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "",
						RefreshToken: "",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusInternalServerError,
						Error:    mockErr,
					},
				},
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(nil, mockErr)
			},
		},
		{
			name: "Auth_VerifyHash_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "",
						RefreshToken: "",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusNotFound,
						Error:    constant.LoginFailedError{},
					},
				},
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(mockAcc, nil)
				handleVerifyHash = func(s, p string) (bool, error) {
					return false, mockErr
				}
			},
		},
		{
			name: "Auth_VerifyHashInvalid_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "",
						RefreshToken: "",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusNotFound,
						Error:    constant.LoginFailedError{},
					},
				},
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(mockAcc, nil)
				handleVerifyHash = func(s, p string) (bool, error) {
					return false, nil
				}
			},
		},
		{
			name: "Auth_GenerateAccessToken_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "",
						RefreshToken: "",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusInternalServerError,
						Error:    mockErr,
					},
				},
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(mockAcc, nil)
				handleVerifyHash = func(s, p string) (bool, error) {
					return true, nil
				}
				mockPASTHandler.On("Generate", paseto.Payload{
					ID:        mockAcc.ID,
					Email:     mockAcc.Email,
					TokenType: paseto.AccessToken,
				}).Return("access-token", mockErr)
			},
		},
		{
			name: "Auth_GenerateRefreshToken_ReturnError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "",
						RefreshToken: "",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusInternalServerError,
						Error:    mockErr,
					},
				},
				err: true,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(mockAcc, nil)
				handleVerifyHash = func(s, p string) (bool, error) {
					return true, nil
				}
				mockPASTHandler.On("Generate", paseto.Payload{
					ID:        mockAcc.ID,
					Email:     mockAcc.Email,
					TokenType: paseto.AccessToken,
				}).Return("access-token", nil)
				mockPASTHandler.On("Generate", paseto.Payload{
					ID:        mockAcc.ID,
					Email:     mockAcc.Email,
					TokenType: paseto.RefreshToken,
				}).Return("", mockErr)
			},
		},
		{
			name: "Auth_NoError",
			args: args{
				ctx: ctx,
				acc: mockAcc,
			},
			expect: expect{
				resp: &authmodel.LoginResponse{
					TokenData: authmodel.TokenData{
						AccessToken:  "access-token",
						RefreshToken: "refresh-token",
					},
					DefaultResponse: common.DefaultResponse{
						HTTPCode: http.StatusOK,
						Error:    nil,
					},
				},
				err: false,
			},
			patch: func(a args) {
				mockRepo.On("Auth", a.ctx, a.acc.Email).Return(mockAcc, nil)
				handleVerifyHash = func(s, p string) (bool, error) {
					return true, nil
				}
				mockPASTHandler.On("Generate", paseto.Payload{
					ID:        mockAcc.ID,
					Email:     mockAcc.Email,
					TokenType: paseto.AccessToken,
				}).Return("access-token", nil)
				mockPASTHandler.On("Generate", paseto.Payload{
					ID:        mockAcc.ID,
					Email:     mockAcc.Email,
					TokenType: paseto.RefreshToken,
				}).Return("refresh-token", nil)
			},
		},
	}

	for _, test := range tests {
		svc := getMock()
		test.patch(test.args)
		t.Run(test.name, func(t *testing.T) {
			resp := svc.Auth(test.args.ctx, test.args.acc)
			assert.Equal(t, test.expect.resp, resp)
			if test.expect.err {
				assert.Error(t, resp.Error)
			} else {
				assert.NoError(t, resp.Error)

			}

		})
	}
}
