package poll

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"os"
	"strconv"
)

var (
	home, _      = os.UserHomeDir()
	PollKeysPath = home + "/.pollKeys"
)

type MsgRegisterPoll struct {
	PollID      string `json:"poll_id"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
	PubKey      string `json:"public_key"`
	Coefficient uint8  `json:"coefficient"`
}

func NewMsgRegisterPoll(PollID string, startTime, endTime string, Coefficient string) MsgRegisterPoll {

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	var startTimeT, endTimeT, coef int

	startTimeT, err = strconv.Atoi(startTime)
	if err != nil {
		panic(err)
	}

	endTimeT, err = strconv.Atoi(endTime)
	if err != nil {
		panic(err)
	}

	coef, err = strconv.Atoi(Coefficient)
	if err != nil {
		panic(err)
	}

	saveKey(key, PollID)

	return MsgRegisterPoll{
		PollID:      PollID,
		StartTime:   int64(startTimeT),
		EndTime:     int64(endTimeT),
		PubKey:      base64.StdEncoding.EncodeToString(x509.MarshalPKCS1PublicKey(&key.PublicKey)),
		Coefficient: uint8(coef),
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
