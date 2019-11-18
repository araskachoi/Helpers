/*
	Copyright 2019 whiteblock Inc.
	This file is a part of the genesis.
	Genesis is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.
	Genesis is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

//Package ethereum provides functions to assist with Ethereum related functionality
package generate

import (
	"fmt"
	"strings"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Account represents an ethereum account
type Account struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  *ecdsa.PublicKey
	Address    common.Address
}

// HexPrivateKey gets the private key in hex format
func (acc Account) HexPrivateKey() string {
	return hex.EncodeToString(crypto.FromECDSA(acc.PrivateKey))
}

// HexPublicKey gets the public key in hex format
func (acc Account) HexPublicKey() string {
	return hex.EncodeToString(crypto.FromECDSAPub(acc.PublicKey))[2:]
}

// HexAddress gets the address in hex format
func (acc Account) HexAddress() string {
	return strings.ToLower(acc.Address.Hex())
}

// NewAccount creates an account from a SECP256K1 ECDSA private key
func NewAccount(privKey *ecdsa.PrivateKey) *Account {
	pubKey := privKey.Public().(*ecdsa.PublicKey)
	addr := crypto.PubkeyToAddress(*pubKey)
	return &Account{PrivateKey: privKey, PublicKey: pubKey, Address: addr}
}

// GenerateEthereumAddress generates a new, random Ethereum account
func GenerateEthereumAddress() (*Account, error) {
	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, error(err)
	}
	return NewAccount(privKey), nil
}

// GenerateAccounts is a convience function to generate an arbitrary number of accounts
// using GenerateEthereumAddress
func GenerateAccounts(accounts int) ([]*Account, error) {
	out := []*Account{}
	for i := 0; i < accounts; i++ {
		acc, err := GenerateEthereumAddress()
		if err != nil {
			return nil, error(err)
		}
		out = append(out, acc)
	}
	return out, nil
}

// MarshalJSON handles the marshaling of Acount into JSON, so that
// the fields are exposed in their hex encodings
func (acc Account) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		PrivateKey string `json:"privateKey"`
		PublicKey  string `json:"publicKey"`
		Address    string `json:"address"`
	}{
		PrivateKey: acc.HexPrivateKey(),
		PublicKey:  acc.HexPublicKey(),
		Address:    acc.HexAddress(),
	})
}

func Export(acc []*Account) (string, error) {
	out := "{"
	for i:=0;i<len(acc);i++ {
		//fmt.Println(acc[i].HexPrivateKey(), acc[i].HexPublicKey(), acc[i].HexAddress())
		account, err := acc[i].MarshalJSON()
		if err != nil {
			return "", error(err)
		}
		//fmt.Println(fmt.Sprintf("%s",account))
		if i == len(acc)-1 {
			out += fmt.Sprintf("%s", account)
		} else {
			out += fmt.Sprintf("%s", account) + ","
		}
	}
	out += "}"
	//fmt.Println(out)
	return out, nil
}
