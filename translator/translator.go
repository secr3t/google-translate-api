package translator

import (
	"github.com/secr3t/google-translate-api/api"
	"github.com/secr3t/google-translate-api/ctx"
	"strings"
)

var (
	c, _ = ctx.NewContext()
	a    = api.NewTranslateApi(c)
)

func Translate(text string) string {
	t, e := a.CreateTranslator(text)

	if e != nil {
		return ""
	}

	tt, e := t.Translate()
	if e != nil {
		return ""
	}

	s, _ := tt.Get()
	return strings.Join(s, ",")
}
