package api

import (
	"github.com/secr3t/google-translate-api/ctx"
	"github.com/secr3t/google-translate-api/speech"
	"github.com/secr3t/google-translate-api/translation"
	"strings"
)

type Translator interface {
	Translate() (translation.Translation, error)

	TranslateTo(lang string) (translation.Translation, error)

	Speak() (speech.Speech, error)

	SpeakSlowly(slowly bool) (speech.Speech, error)
}

type translator struct {
	ctx  ctx.Context
	text string
	lang string
}

func (t *translator) Translate() (translation.Translation, error) {
	return t.TranslateTo(t.ctx.DefaultTargetLang())
}

func (t *translator) TranslateTo(lang string) (translation.Translation, error) {
	lang = strings.ToLower(lang)
	if err := t.ctx.IsSupported(lang); err != nil {
		return nil, err
	}
	return translation.NewTranslation(t.lang, lang, t.text, t.ctx), nil
}

func (t *translator) Speak() (speech.Speech, error) {
	return t.SpeakSlowly(t.ctx.DefaultSlowly())
}

func (t *translator) SpeakSlowly(slowly bool) (speech.Speech, error) {
	return speech.NewSpeech(t.ctx, t.text, t.lang, slowly), nil
}
