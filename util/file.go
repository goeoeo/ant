package util

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//判定文件是否存在
func IsFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//将数据刷入文件
func Data2File(filePath string, data interface{}) error {

	//文件不存在创建文件
	if !IsFileExist(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	content, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, content, 0644)

}

//将文件数据载入到变量中
func File2Data(filePath string, data interface{}) error {
	if !IsFileExist(filePath) {
		return nil
	}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, data)
}
