package glPget

// Accept-Ranges: bytes の有無をhttpリクエストを送って確認
// 構造体に結果をいれておく

import (
	context "context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context/ctxhttp"
)

const timeout = 10

func (glp *glPget) CheckSetupURL() error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	res, err := ctxhttp.Head(ctx, http.DefaultClient, glp.URL)

	if err != nil {
		return errors.New("can't http requests")
	}

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return errors.New("not suport accept-ranges from " + glp.URL)
	}

	return nil
}
