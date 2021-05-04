package api

import (
	"github.com/secr3t/google-translate-api/ctx"
	"strings"
)

type TranslateApi interface {
	CreateTranslator(text string) (Translator, error)

	CreateLangTranslator(text, lang string) (Translator, error)

	GetCtx() ctx.Context
}

func NewTranslateApi(ctx ctx.Context) TranslateApi {
	return &translateApi{
		Ctx: ctx,
	}
}

type translateApi struct {
	Ctx ctx.Context
}

func (a *translateApi) GetCtx() ctx.Context {
	return a.Ctx
}

func (a *translateApi) CreateTranslator(text string) (Translator, error) {
	return a.CreateLangTranslator(text, a.Ctx.DefaultSourceLang())
}

func (a *translateApi) CreateLangTranslator(text, lang string) (Translator, error) {
	lang = strings.ToLower(lang)
	if err := a.Ctx.IsSupported(lang); err != nil {
		return nil, err
	}

	return &translator{
		ctx:  a.Ctx,
		text: text,
		lang: lang,
	}, nil
}
