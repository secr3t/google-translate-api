package translation

import (
	"bytes"
	"encoding/json"
	"github.com/secr3t/google-translate-api/ctx"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

const GoogleTranslateRpc = "MkEWBc"

type Translation interface {
	Get() ([]string, error)
}

func NewTranslation(sourceLang, targetLang, text string, ctx ctx.Context) Translation {
	return &translation{
		sourceLang: sourceLang,
		targetLang: targetLang,
		text:       text,

		ctx: ctx,
	}
}

type translation struct {
	ctx ctx.Context

	sourceLang, targetLang string
	text                   string

	once sync.Once

	result    []string
	resultErr error
}

func (t *translation) prepareParameters() (string, error) {
	parameters := []interface{}{
		[]interface{}{t.text, t.sourceLang, t.targetLang, true},
		[]interface{}{nil},
	}

	escaped, err := json.Marshal(parameters)
	if err != nil {
		return "", err
	}

	parameters = []interface{}{
		[]interface{}{
			[]interface{}{
				GoogleTranslateRpc,
				string(escaped),
				nil,
				"generic",
			},
		},
	}

	escaped, err = json.Marshal(parameters)
	if err != nil {
		return "", err
	}

	data := "f.req=" + url.QueryEscape(string(escaped))

	return data, nil
}

func (t *translation) get() ([]string, error) {
	data, err := t.prepareParameters()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		"POST",
		t.ctx.TranslateUrl(),
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		return nil, err
	}

	t.ctx.DecorateHeader(req.Header)

	response, err := t.ctx.HttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	all, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	list, err := Analyze(string(all))
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (t *translation) Get() ([]string, error) {
	t.once.Do(func() {
		t.result, t.resultErr = t.get()
	})
	return t.result, t.resultErr
}
