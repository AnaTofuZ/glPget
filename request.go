package glPget

// Accept-Ranges: bytes の有無をhttpリクエストを送って確認
// 構造体に結果をいれておく

import (
	"net/http"

	"github.com/pkg/errors"
)

const timeout = 10

func (glp *glPget) CheckSetupURL() error {

	client := &http.Client{}
	res, err := client.Head(glp.URL)

	if err != nil {
		return errors.New("can't http requests")
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return errors.New("not suport accept-ranges from " + glp.URL)
	}

	return nil
}
