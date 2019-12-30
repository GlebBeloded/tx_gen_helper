package session

import (
	"io/ioutil"
	"os"
)

var (
	home, _     = os.UserHomeDir()
	AdBytesPath = home + "/.ads"
)

//MsgIntegrationData represents information about a certain integration
type MsgIntegrationData struct {
	IntegrationID string `json:"integration_id"`
	AdBytes       string `json:"ad_bytes"`
}

//MsgDistributeRewards is an array of MsgReward
type MsgDistributeRewards struct {
	Ads []MsgIntegrationData `json:"ads"`
}

//NewMsgDistributeRewards creates new Send ad message
func NewMsgDistributeRewards(ids []string) MsgDistributeRewards {

	var array []MsgIntegrationData
	for _, i := range ids {
		bytes, err := ioutil.ReadFile(AdBytesPath + "/" + i)
		if err != nil {
			panic(err)
		}
		array = append(array, MsgIntegrationData{IntegrationID: i, AdBytes: string(bytes)})
	}
	return MsgDistributeRewards{
		Ads: array,
	}
}
