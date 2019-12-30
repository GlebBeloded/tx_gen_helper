package session

import (
	"crypto/rand"
	"encoding/base64"
	"os"
)

func SaveAdBytesToOS(args ...string) {

	for _, i := range args {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		checkExists(i)

		save(i, base64.StdEncoding.EncodeToString(bytes))
	}
}

func checkExists(id string) {
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

	filepath := AdBytesPath + "/" + id

	boolExists, err = Exists(filepath)
	if err != nil {
		panic(err)
	}

	if boolExists {
		panic("IntegrationID with the given name already exists")
	}

}

func save(id, adBytes string) {

	filepath := AdBytesPath + "/" + id

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}

	file.WriteString(adBytes)
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
