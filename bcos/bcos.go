package bcos

import (
	"fmt"

	"github.com/yekai1003/gobcos/accounts/keystore"
)

const keyDir = "./ks"

func NewAccount(pass string) (string, error) {
	ks := keystore.NewKeyStore(keyDir, keystore.LightScryptN, keystore.LightScryptP)
	account, err := ks.NewAccount(pass)
	if err != nil {
		fmt.Println("Failed to NewAccount", err)
		return "", err
	}
	return account.Address.Hex(), nil
}
