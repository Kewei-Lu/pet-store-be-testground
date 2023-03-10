// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"petStore/internal/model"
)

type (
	IUserMoney interface {
		QueryMoney(ctx context.Context, in model.MoneyQueryInput) (int, error)
		CreateMoneyAccount(ctx context.Context, in model.CreateMoneyAccountInput) error
		AddMoney2User(ctx context.Context, in model.MoneyAddInput) error
		TransferMoney2User(ctx context.Context, in model.MoneyTransferInput) error
	}
)

var (
	localUserMoney IUserMoney
)

func UserMoney() IUserMoney {
	if localUserMoney == nil {
		panic("implement not found for interface IUserMoney, forgot register?")
	}
	return localUserMoney
}

func RegisterUserMoney(i IUserMoney) {
	localUserMoney = i
}
