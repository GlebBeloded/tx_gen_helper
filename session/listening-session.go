package session

import (
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"strconv"

	"github.com/btcsuite/btcutil/bech32"
)

//MsgAd store information about advertisment
type MsgAd struct {
	IntegrationID string `json:"integration_id"`
	AdTime        int64  `json:"integration_time"`
	Fingerprint   string `json:"fingerprint"`
}

//MsgRegisterListeningSession is used to pass params to sendAd transactions
type MsgRegisterListeningSession struct {
	AccAddr      string  `json:"acc_addr"`
	SessionStart int64   `json:"time_start"`
	SessionEnd   int64   `json:"time_end"`
	ChannelID    string  `json:"channel_id"`
	Ads          []MsgAd `json:"ads"`
}

//NewMsgRegisterListeningSession creates new Send ad message
func NewMsgRegisterListeningSession(addr string, start, end string, channel string, Ads ...string) MsgRegisterListeningSession {

	tBegin, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}

	tEnd, err := strconv.Atoi(end)
	if err != nil {
		panic(err)
	}

	return MsgRegisterListeningSession{
		AccAddr:      addr,
		SessionStart: int64(tBegin),
		SessionEnd:   int64(tEnd),
		ChannelID:    channel,
		Ads:          GenerateFingerprints(addr, Ads...),
	}
}

func GenerateFingerprints(user string, ads ...string) (output []MsgAd) {

	for i := 0; i < len(ads); i = i + 2 {
		bytes, err := ioutil.ReadFile(AdBytesPath + "/" + ads[i])
		if err != nil {
			panic(err)
		}

		fprint := base64.StdEncoding.EncodeToString(UserHash(user, string(bytes)))

		time, err := strconv.Atoi(ads[i+1])
		if err != nil {
			panic(err)
		}

		output = append(output, MsgAd{
			IntegrationID: ads[i],
			AdTime:        int64(time),
			Fingerprint:   fprint,
		})
	}
	return
}

func UserHash(addr, adBytes string) []byte {
	randomBytes, err := base64.StdEncoding.DecodeString(adBytes)
	if err != nil {
		panic(err)
	}
	_, addrBytes, err := bech32.Decode(addr)
	if err != nil {
		panic(err)
	}
	converted, err := bech32.ConvertBits(addrBytes, 5, 8, false)
	if err != nil {
		panic(err)
	}

	result := append(randomBytes, ';')
	result = append(result, converted[:20]...)

	hash := sha256.Sum256(result)
	return hash[:]
}
