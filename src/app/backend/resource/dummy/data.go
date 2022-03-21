package dummy

import (
	"github.com/Triticumdico/dashboard/src/app/backend/client"
)

type Accounts struct {
	Table []Account `json:"accounts"`
}

type Account struct {
	Id         int    `json:"id"`
	AccountUid string `json:"account_uid"`
}

var account Accounts

func GetTableRows() (Accounts, error) {

	db := client.ClientDb.GetClientDb()

	db.Find(&account.Table)

	return account, nil
}
