package lingxing

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hiscaler/gox/filex"
	"io/ioutil"
	"os"
	"path"
)

type TokenWriterReader interface {
	Read() (Token, error)
	Write(token Token) (bool, error)
}

type FileToken struct {
	Path  string
	Token Token
}

func tokenFilePath() string {
	fmt.Println(path.Join(os.TempDir(), "tong_tool_token.json"))
	return path.Join(os.TempDir(), "tong_tool_token.json")
}

func (ft FileToken) Read() (Token, error) {
	token := Token{}
	var err error
	file := tokenFilePath()
	if filex.Exists(tokenFilePath()) {
		var b []byte
		if b, err = ioutil.ReadFile(file); err == nil {
			err = json.Unmarshal(b, &token)
		}
	} else {
		err = errors.New("tongtool token file is not exists")
	}
	return token, err
}

func (ft FileToken) Write(token Token) (bool, error) {
	b, err := json.Marshal(token)
	err = ioutil.WriteFile(tokenFilePath(), b, 0777)
	if err != nil {
		return false, err
	}
	return true, nil
}
