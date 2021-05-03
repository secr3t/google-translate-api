package tokenizer

const (
	ToneMarks = "[?!？！]"

	AllPunc = "?!？！.,¡()[]¿…‥،;:—。，、： \t\n\r\v\f"

	OtherPunc = `¡()[]¿…‥،;—。，、：`
)

var (
	ABBREVIATIONS = []string{
		"dr", "jr", "mr",
		"mrs", "ms", "msgr",
		"prof", "sr", "st",
	}

	SUB_PAIRS = map[string]string{
		"Esq.": "Esquire",
	}
)
