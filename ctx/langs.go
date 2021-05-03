package ctx

var (
	MainLangs = map[string]string{
		"af":    "Afrikaans",
		"ar":    "Arabic",
		"bn":    "Bengali",
		"bs":    "Bosnian",
		"ca":    "Catalan",
		"cs":    "Czech",
		"cy":    "Welsh",
		"da":    "Danish",
		"de":    "German",
		"el":    "Greek",
		"en":    "English",
		"eo":    "Esperanto",
		"es":    "Spanish",
		"et":    "Estonian",
		"fi":    "Finnish",
		"fr":    "French",
		"gu":    "Gujarati",
		"hi":    "Hindi",
		"hr":    "Croatian",
		"hu":    "Hungarian",
		"hy":    "Armenian",
		"id":    "Indonesian",
		"is":    "Icelandic",
		"it":    "Italian",
		"ja":    "Japanese",
		"jw":    "Javanese",
		"km":    "Khmer",
		"kn":    "Kannada",
		"ko":    "Korean",
		"la":    "Latin",
		"lv":    "Latvian",
		"mk":    "Macedonian",
		"ml":    "Malayalam",
		"mr":    "Marathi",
		"my":    "Myanmar (Burmese)",
		"ne":    "Nepali",
		"nl":    "Dutch",
		"no":    "Norwegian",
		"pl":    "Polish",
		"pt":    "Portuguese",
		"ro":    "Romanian",
		"ru":    "Russian",
		"si":    "Sinhala",
		"sk":    "Slovak",
		"sq":    "Albanian",
		"sr":    "Serbian",
		"su":    "Sundanese",
		"sv":    "Swedish",
		"sw":    "Swahili",
		"ta":    "Tamil",
		"te":    "Telugu",
		"th":    "Thai",
		"tl":    "Filipino",
		"tr":    "Turkish",
		"uk":    "Ukrainian",
		"ur":    "Urdu",
		"vi":    "Vietnamese",
		"zh-CN": "Chinese",
	}

	ExtraLangs = map[string]string{
		// Chinese
		"zh-cn": "Chinese (Mandarin/China)",
		"zh-tw": "Chinese (Mandarin/Taiwan)",
		// English
		"en-us": "English (US)",
		"en-ca": "English (Canada)",
		"en-uk": "English (UK)",
		"en-gb": "English (UK)",
		"en-au": "English (Australia)",
		"en-gh": "English (Ghana)",
		"en-in": "English (India)",
		"en-ie": "English (Ireland)",
		"en-nz": "English (New Zealand)",
		"en-ng": "English (Nigeria)",
		"en-ph": "English (Philippines)",
		"en-za": "English (South Africa)",
		"en-tz": "English (Tanzania)",
		// French
		"fr-ca": "French (Canada)",
		"fr-fr": "French (France)",
		// Portuguese
		"pt-br": "Portuguese (Brazil)",
		"pt-pt": "Portuguese (Portugal)",
		// Spanish
		"es-es": "Spanish (Spain)",
		"es-us": "Spanish (United States)",
	}

	Langs = make(map[string]string, len(MainLangs)+len(ExtraLangs))
)

func init() {
	for k, v := range MainLangs {
		Langs[k] = v
	}

	for k, v := range ExtraLangs {
		Langs[k] = v
	}
}
