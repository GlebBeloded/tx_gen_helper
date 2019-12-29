package session

import (
	"crypto/rand"
	"encoding/base64"
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
func NewMsgDistributeRewards(adArray []MsgIntegrationData) MsgDistributeRewards {

	for _, i := range adArray {
		save(i)
	}

	return MsgDistributeRewards{
		Ads: adArray,
	}
}

func MakeIntegrationDataArray(args ...string) []MsgIntegrationData {

	var result []MsgIntegrationData

	for _, i := range args {
		bytes := make([]byte, 32)
		rand.Read(bytes)

		result = append(result, MsgIntegrationData{
			IntegrationID: i,
			AdBytes:       base64.StdEncoding.EncodeToString(bytes),
		})
	}
	return result
}

func save(msg MsgIntegrationData) {
	boolExists, err := Exists(AdBytesPath)
	if err != nil {
		panic(err)
	}

	if !boolExists {
		err = os.Mkdir(AdBytesPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	filepath := AdBytesPath + "/" + msg.IntegrationID

	boolExists, err = Exists(filepath)
	if err != nil {
		panic(err)
	}

	if boolExists {
		panic("IntegrationID with the given name already exists")
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	file.WriteString(msg.AdBytes)
	file.Close()
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
