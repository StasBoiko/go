package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type jsonData struct {
	json.RawMessage
}

// func ValidateJson(jsonText string) (err error) {
// 	var js jsonData
// 	err = json.Unmarshal([]byte(jsonText), &js)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func Save(flag string, jsonText string) (err error) {

	md5Hash := GetMD5Hash(jsonText)

	err = CheckMd5(md5Hash)
	if err != nil {
		return err
	}
	err = SaveMd5Hash(md5Hash)
	if err != nil {
		return err
	}

	bytes := []byte(jsonText)
	timeNow := time.Now().String()

	// }
	// if strings.Contains(file.Name(), ".json") {
	// 	fileBytes, err := ioutil.ReadFile(file.Name())
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if strings.Compare(string(fileBytes), jsonText) == 0 {
	// 		return fmt.Errorf("already exist")
	// 	}
	// }
	// }

	// for _, file := range files {
	// 	if strings.Contains(file.Name(), ".json") {
	// 		fileBytes, err := ioutil.ReadFile(file.Name())
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if strings.Compare(string(fileBytes), jsonText) == 0 {
	// 			return fmt.Errorf("already exist")
	// 		}
	// 	}
	// }

	// err = ioutil.WriteFile(fmt.Sprintf("./%s.%s", timeNow, flag), bytes, 0644)
	// if err != nil {
	// 	return err
	// }
	return ioutil.WriteFile(fmt.Sprintf("./%s.%s", timeNow, flag), bytes, 0644)
}

// func GetMD5Hash(text string) string {
// 	hash := md5.Sum([]byte(text))
// 	return hex.EncodeToString(hash[:])
// }
