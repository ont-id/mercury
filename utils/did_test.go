/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/stretchr/testify/assert"
)

var (
	testOntSdk   *sdk.OntologySdk
	testPasswd   = []byte("123456")
	testDefAcc   *sdk.Account
	testGasPrice = uint64(2500)
	testGasLimit = uint64(20000)
	testRpcUrl   = "http://127.0.0.1:20336"
)

func init() {
	var err error
	var wallet *sdk.Wallet
	if !FileExisted("./wallet.dat") {
		wallet, err = testOntSdk.CreateWallet("./wallet.dat")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	} else {
		wallet, err = testOntSdk.OpenWallet("./wallet.dat")
		if err != nil {
			fmt.Println("[CreateWallet] error:", err)
			return
		}
	}
	_, err = wallet.NewDefaultSettingAccount([]byte("123456"))
	if err != nil {
		fmt.Printf("NewDefaultSettingAccount err:%s\n", err)
		return
	}
	err = wallet.Save()
	if err != nil {
		fmt.Printf("wallet save err:%s\n", err)
		return
	}
	testWallet, err := testOntSdk.OpenWallet("./wallet.dat")
	if err != nil {
		fmt.Printf("account.Open error:%s\n", err)
		return
	}
	testOntSdk = sdk.NewOntologySdk()
	testOntSdk.NewRpcClient().SetAddress(testRpcUrl)
	testDefAcc, err = testWallet.GetDefaultAccount(testPasswd)
	if err != nil {
		fmt.Printf("GetDefaultAccount err: %s\n", err)
		return
	}
}

//./ontology --testmode --testmode-gen-block-time 10
func TestValidateDid(t *testing.T) {
	did := "did:ont:TV6jHmuBW33n4tcrggiLnwoRiZ8zHbJBPa"
	assert.Equal(t, true, ValidateDid(did))
}

func RegIDWithPublicKey() (string, error) {
	did, err := sdk.GenerateID()
	if err != nil {
		return "", err
	}
	fmt.Println("did", did)
	txHash, err := testOntSdk.Native.OntId.RegIDWithPublicKey(testGasPrice, testGasLimit, testDefAcc, did, testDefAcc)
	if err != nil {
		return "", err
	}
	fmt.Println("hash:", txHash.ToHexString())
	_, err = testOntSdk.WaitForGenerateBlock(10*time.Second, 1)
	if err != nil {
		return "", err
	}
	return did, nil
}

func TestGetDidDocByDid(t *testing.T) {
	did, err := RegIDWithPublicKey()
	assert.Nil(t, err)
	t.Logf("did:%s", did)
	didDoc, err := GetDidDocByDid(did, testOntSdk)
	assert.Nil(t, err)
	t.Logf("doc:%s", didDoc)
}

func TestGetServiceEndpointByDid(t *testing.T) {
	did, err := RegIDWithPublicKey()
	assert.Nil(t, err)
	t.Logf("did:%s", did)
	didDoc, err := GetServiceEndpointByDid(did, testOntSdk)
	assert.Nil(t, err)
	t.Logf("addrs:%s", didDoc)
}

func TestGetPubKeyByDid(t *testing.T) {
	did, err := RegIDWithPublicKey()
	assert.Nil(t, err)
	t.Logf("did:%s", did)
	pubKey, err := GetPubKeyByDid(did, testOntSdk)
	assert.Nil(t, err)
	_, err = keypair.DeserializePublicKey(pubKey)
	assert.Nil(t, err)
}
