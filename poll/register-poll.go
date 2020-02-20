package poll

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"strconv"

	sdk "inside.omertex.com/bitbucket/scm/mf/blockchain_mediafm.git/types"
	"inside.omertex.com/bitbucket/scm/mf/blockchain_sdk.git/x/melodia"
)

var (
	home, _      = os.UserHomeDir()
	PollKeysPath = home + "/.pollKeys"
)

type MsgRegisterPoll = melodia.MsgRegisterPoll

func NewMsgRegisterPoll(PollID string, startTime, endTime string, amount sdk.Coin, limit int) MsgRegisterPoll {

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	var startTimeT, endTimeT int

	startTimeT, err = strconv.Atoi(startTime)
	if err != nil {
		panic(err)
	}

	endTimeT, err = strconv.Atoi(endTime)
	if err != nil {
		panic(err)
	}

	saveKey(key, PollID)

	return MsgRegisterPoll{
		PollID:    PollID,
		StartTime: int64(startTimeT),
		EndTime:   int64(endTimeT),
		PubKey:    base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&key.PublicKey)),
		Amount:    amount,
		Limit:     limit,
	}
}

func saveKey(key *rsa.PrivateKey, name string) {
	boolExists, err := Exists(PollKeysPath)
	if err != nil {
		panic(err)
	}

	if !boolExists {
		err = os.Mkdir(PollKeysPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	keyName := name + ".rsa"
	pubKeyName := keyName + ".pub"

	boolExists, err = Exists(pubKeyName)
	if err != nil {
		panic(err)
	}

	if boolExists {
		panic("KeyPair with the given name exists")
	}

	binaryKey := x509.MarshalPKCS1PrivateKey(key)
	binaryPubKey := x509.MarshalPKCS1PublicKey(&key.PublicKey)

	fpriv, err := os.Create(PollKeysPath + "/" + keyName)
	if err != nil {
		panic(err)
	}
	defer fpriv.Close()

	fpub, err := os.Create(PollKeysPath + "/" + pubKeyName)
	if err != nil {
		panic(err)
	}
	defer fpub.Close()

	privPemBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: binaryKey,
	}

	pubPemBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: binaryPubKey,
	}

	pem.Encode(fpriv, &privPemBlock)
	pem.Encode(fpub, &pubPemBlock)

}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
