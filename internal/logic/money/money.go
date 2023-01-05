package money

import (
	"context"
	"petStore/internal/model"
	"petStore/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sUserMoney struct {
	moneyRecord     []*userMoneyRecord
	moneyAccountMap userMoneyMap
}

type userMoneyRecord struct {
	UserName           string
	ActionType         string
	ActionTime         string
	DestinationAccount string
	Amount             int
	Comment            string
}
type userMoneyMap map[string]int

func init() {
	service.RegisterUserMoney(New())
}

func New() *sUserMoney {
	return &sUserMoney{moneyRecord: []*userMoneyRecord{}, moneyAccountMap: make(userMoneyMap)}
}
func (s *sUserMoney) QueryMoney(ctx context.Context, in model.MoneyQueryInput) (int, error) {
	for user, money := range s.moneyAccountMap {
		if user == in.UserName {
			return money, nil
		}
	}
	return -1, INVALID_QUERY_USER_ERROR
}

func (s *sUserMoney) CreateMoneyAccount(ctx context.Context, in model.CreateMoneyAccountInput) error {
	for user, _ := range s.moneyAccountMap {
		if user == in.UserName {
			return DUPLICATE_MONEY_ACCOUNT_ERROR
		}
	}
	s.moneyAccountMap[in.UserName] = 0
	return nil
}

func (s *sUserMoney) AddMoney2User(ctx context.Context, in model.MoneyAddInput) error {
	for user, _ := range s.moneyAccountMap {
		if user == in.UserName {
			s.moneyAccountMap[user] += in.Amount
			s.moneyRecord = append(s.moneyRecord, &userMoneyRecord{UserName: in.UserName, ActionType: MONEY_ADD, ActionTime: gtime.TimestampMilliStr(), DestinationAccount: in.DestinationAccount, Amount: in.Amount, Comment: in.Comment})

		}
	}
	// s.moneyAccountMap[in.UserName] += in.Amount
	return nil
}

func (s *sUserMoney) TransferMoney2User(ctx context.Context, in model.MoneyTransferInput) error {
	g.Log().Print(ctx, "from: ", in.UserName)
	g.Log().Print(ctx, "to: ", in.DestinationAccount)
	for user, _ := range s.moneyAccountMap {
		if user == in.UserName {
			if s.moneyAccountMap[user] < in.Amount {
				return INVALID_TRANSFER_AMOUNT_ERROR
			}
			for destUser, _ := range s.moneyAccountMap {
				if destUser == in.DestinationAccount {
					s.moneyRecord = append(s.moneyRecord, &userMoneyRecord{UserName: in.UserName, ActionType: MONEY_TRANSFER, ActionTime: gtime.TimestampMilliStr(), DestinationAccount: in.DestinationAccount, Amount: in.Amount, Comment: in.Comment})
					s.moneyAccountMap[in.UserName] -= in.Amount
					s.moneyAccountMap[in.DestinationAccount] += in.Amount
					g.Log().Printf(ctx, "Transfer Money from user %s to user %s successfully, amount: %d", user, destUser, in.Amount)
					return nil
				}
			}
			return INVALID_TRANSFER_DESTINATION_USER_ERROR
		}
	}
	return INVALID_TRANSFER_SOURCE_USER_ERROR
}
