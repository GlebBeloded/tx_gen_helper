package session

import (
	"os"

	"inside.omertex.com/bitbucket/scm/mf/blockchain_sdk.git/x/melodia"
)

var (
	home, _     = os.UserHomeDir()
	AdBytesPath = home + "/.ads"
)

type (
	Integration          = melodia.Integration
	Pair                 = melodia.Pair
	MsgDistributeRewards = melodia.MsgDistributeRewards
)

//NewMsgDistributeRewards creates new Send ad message
func NewMsgDistributeRewards(i []Integration) MsgDistributeRewards {

	return MsgDistributeRewards{
		Integrations: i,
	}
}
