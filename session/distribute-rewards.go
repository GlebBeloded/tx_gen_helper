package session

import (
	"io/ioutil"
	"os"

	sdk "inside.omertex.com/bitbucket/scm/mf/blockchain_mediafm.git/types"
	"inside.omertex.com/bitbucket/scm/mf/blockchain_sdk.git/x/melodia"
)

var (
	home, _     = os.UserHomeDir()
	AdBytesPath = home + "/.ads"
)

type (
	MsgIntegrationData   = melodia.MsgIntegrationData
	MsgDistributeRewards = melodia.MsgDistributeRewards
)

//NewMsgDistributeRewards creates new Send ad message
func NewMsgDistributeRewards(payout sdk.Coin, ids []string) MsgDistributeRewards {

	var array []MsgIntegrationData
	for _, i := range ids {
		bytes, err := ioutil.ReadFile(AdBytesPath + "/" + i)
		if err != nil {
			panic(err)
		}
		array = append(array, MsgIntegrationData{IntegrationID: i, AdBytes: string(bytes)})
	}
	return MsgDistributeRewards{
		Ads:         array,
		TotalPayout: payout,
	}
}
