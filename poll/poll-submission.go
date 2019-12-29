package poll

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strconv"

	sdk "inside.omertex.com/bitbucket/scm/mf/blockchain_mediafm.git/types"
)

//MsgPollSubmission is used to register user completing survey
type MsgPollSubmission struct {
	PollID         string `json:"poll_id"`
	SubmissionTime int64  `json:"submission_time"`
	AccAddr        string `json:"address"`
	Signature      string `json:"signature"`
}

type submissionForSigning struct {
	PollID         string `json:"poll_id"`
	SubmissionTime int64  `json:"submission_time"`
	AccAddr        string `json:"address"`
}

func NewSubsubmissionForSigning(PollID string, SubmissionTime int64, AccAddr string) submissionForSigning {
	return submissionForSigning{
		PollID:         PollID,
		SubmissionTime: SubmissionTime,
		AccAddr:        AccAddr,
	}
}

func getKey(PollID string) *rsa.PrivateKey {
	keyName := PollID + ".rsa"
	path := PollKeysPath + "/" + keyName
	boolExists, err := Exists(path)
	if err != nil {
		panic(err)
	}
	if !boolExists {
		panic(fmt.Sprintf("Keypair %s does not exist", PollID))
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(file)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return key
}

func (v submissionForSigning) sign(key *rsa.PrivateKey) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	sorted := sdk.MustSortJSON(bytes)

	hashFunc := crypto.SHA256
	hashBuf := hashFunc.New()
	hashBuf.Write(sorted)
	hashed := hashBuf.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, key, hashFunc, hashed)
	if err != nil {
		panic(err)
	}
	return signature
}

func NewMsgPollSubmission(PollID string, SubmissionTime string, AccAddr string) MsgPollSubmission {
	number, err := strconv.Atoi(SubmissionTime)
	if err != nil {
		panic(err)
	}

	forSig := NewSubsubmissionForSigning(PollID, int64(number), AccAddr)
	key := getKey(PollID)
	signature := forSig.sign(key)
	return MsgPollSubmission{
		PollID:         PollID,
		SubmissionTime: int64(number),
		AccAddr:        AccAddr,
		Signature:      base64.StdEncoding.EncodeToString(signature),
	}
}
