package ctx

import (
	"fmt"
	"github.com/secr3t/google-translate-api/tokenizer"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64) " +
		"AppleWebKit/537.36 (KHTML, like Gecko) " +
		"Chrome/47.0.2526.106 Safari/537.36"

	ContentType = "application/x-www-form-urlencoded;charset=utf-8"

	TranslatePath = "/_/TranslateWebserverUi/data/batchexecute"

	DomainTemplate = "https://translate.google.%s"
)

type Context interface {
	TranslateUrl() string

	CheckLang() bool

	DefaultSourceLang() string
	DefaultTargetLang() string
	DefaultSlowly() bool

	HttpClient() *http.Client
	DecorateHeader(header http.Header)

	IsSupported(lang string) error

	PreProcessors() []tokenizer.PreProcessor
	Tokenize() tokenizer.Tokenize
}

type context struct {
	client *http.Client

	tld string

	checkLang bool

	defaultSourceLang string
	defaultTargetLang string
	defaultSlowly     bool

	translateUrl     string
	refererUrl       string
	translateUrlLock sync.Once

	preProcessors []tokenizer.PreProcessor
	tokenize      tokenizer.Tokenize
}

func createDefaultPreProcessors() ([]tokenizer.PreProcessor, error) {
	defaultPreProcessors := make([]tokenizer.PreProcessor, 0)

	processor, err := tokenizer.CreateToneMarks()
	if err != nil {
		return nil, err
	}
	defaultPreProcessors = append(defaultPreProcessors, processor)

	processor, err = tokenizer.CreateEndOfLine()
	if err != nil {
		return nil, err
	}
	defaultPreProcessors = append(defaultPreProcessors, processor)

	processor, err = tokenizer.CreateAbbreviations()
	if err != nil {
		return nil, err
	}
	defaultPreProcessors = append(defaultPreProcessors, processor)

	processors, err := tokenizer.CreateWorSub()
	if err != nil {
		return nil, err
	}
	for _, processor := range processors {
		defaultPreProcessors = append(defaultPreProcessors, processor)
	}
	return defaultPreProcessors, nil
}

func NewContextWithOptions(tld, srcLang, dstLang string) (Context, error) {
	defaultPreProcessors, err := createDefaultPreProcessors()
	if err != nil {
		return nil, err
	}

	netTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	netClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	ctx := &context{
		tld: tld,

		checkLang: true,

		defaultSourceLang: srcLang,
		defaultTargetLang: dstLang,
		defaultSlowly:     true,

		client: netClient,

		preProcessors: defaultPreProcessors,
		tokenize:      tokenizer.TotalTokenize,
	}

	return ctx, nil
}

func NewContext() (Context, error) {
	defaultPreProcessors, err := createDefaultPreProcessors()
	if err != nil {
		return nil, err
	}

	netTransport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	netClient := &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	ctx := &context{
		tld: "com",

		checkLang: true,

		defaultSourceLang: "zh-cn",
		defaultTargetLang: "ko",
		defaultSlowly:     true,

		client: netClient,

		preProcessors: defaultPreProcessors,
		tokenize:      tokenizer.TotalTokenize,
	}

	return ctx, nil
}

func (c *context) TranslateUrl() string {
	c.translateUrlLock.Do(func() {
		c.refererUrl = fmt.Sprintf(DomainTemplate, c.tld)
		c.translateUrl = c.refererUrl + TranslatePath
	})
	return c.translateUrl
}

func (c *context) DefaultSourceLang() string {
	return c.defaultSourceLang
}

func (c *context) DefaultTargetLang() string {
	return c.defaultTargetLang
}

func (c *context) DefaultSlowly() bool {
	return c.defaultSlowly
}

func (c *context) CheckLang() bool {
	return c.checkLang
}

func (c *context) HttpClient() *http.Client {
	return c.client
}

func (c *context) PreProcessors() []tokenizer.PreProcessor {
	return c.preProcessors
}
func (c *context) Tokenize() tokenizer.Tokenize {
	return c.tokenize
}

func (c *context) DecorateHeader(header http.Header) {
	header.Set("Referer", c.refererUrl)
	header.Set("User-Agent", GetUA())
	header.Set("Content-Type", ContentType)
}

func (c *context) IsSupported(lang string) error {
	lang = strings.ToLower(lang)
	if c.checkLang {
		if _, ok := Langs[lang]; !ok {
			return fmt.Errorf("Language not supported: %s ", lang)
		}
	}
	return nil
}
