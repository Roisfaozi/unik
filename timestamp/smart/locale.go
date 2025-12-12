package smart

// PluralCategory constants based on CLDR (Common Locale Data Repository)
type PluralCategory int

const (
	PluralOther PluralCategory = iota
	PluralOne
	// Can be extended: PluralZero, PluralTwo, PluralFew, PluralMany
)

// PluralRuleFunc defines how to map a number to a category
type PluralRuleFunc func(n int) PluralCategory

// Locale defines the localization data for a language
type Locale struct {
	Code       string
	PluralRule PluralRuleFunc
	Dictionary map[string]string              // For static fixed words (e.g. "just_now")
	Plurals    map[string]map[PluralCategory]string // For words that change with number (e.g. "minute")
}

var registry = map[string]Locale{}

func init() {
	registerID()
	registerEN()
	registerTH()
	registerVN()
	registerJP()
	registerMY()
}

func registerEN() {
	registry["en"] = Locale{
		Code: "en",
		PluralRule: func(n int) PluralCategory {
			if n == 1 {
				return PluralOne
			}
			return PluralOther
		},
		Dictionary: map[string]string{
			"just_now": "just now",
			"ago":      "ago",
			"in":       "in",
			"s":        "s", // Short forms usually don't pluralize in this context (1s, 2s)
			"m":        "m",
			"h":        "h",
			"d":        "d",
			"y":        "y",
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOne: "second", PluralOther: "seconds"},
			"min":  {PluralOne: "minute", PluralOther: "minutes"},
			"hour": {PluralOne: "hour", PluralOther: "hours"},
			"day":  {PluralOne: "day", PluralOther: "days"},
			"year": {PluralOne: "year", PluralOther: "years"},
		},
	}
}

func registerID() {
	registry["id"] = Locale{
		Code: "id",
		PluralRule: func(n int) PluralCategory {
			// Indonesian doesn't strictly change word form for units based on number
			// 1 menit, 2 menit
			return PluralOther
		},
		Dictionary: map[string]string{
			"just_now": "baru saja",
			"ago":      "lalu",
			"in":       "dalam",
			"s":        "dtk",
			"m":        "mnt",
			"h":        "j",
			"d":        "h",
			"y":        "thn",
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOther: "detik"},
			"min":  {PluralOther: "menit"},
			"hour": {PluralOther: "jam"},
			"day":  {PluralOther: "hari"},
			"year": {PluralOther: "tahun"},
		},
	}
}

func registerTH() {
	registry["th"] = Locale{
		Code: "th",
		PluralRule: func(n int) PluralCategory {
			return PluralOther // Thai has no plural inflection
		},
		Dictionary: map[string]string{
			"just_now": "เมื่อสักครู่", // Muea sak khru
			"ago":      "ที่แล้ว",      // Tee laeo
			"in":       "อีก",          // Eek
			"s":        "วิ",           // Short Wi
			"m":        "น.",           // Short N.
			"h":        "ชม.",          // Short Chom.
			"d":        "วัน",          // Short Wan
			"y":        "ปี",           // Short Pee
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOther: "วินาที"}, // Winathi
			"min":  {PluralOther: "นาที"},   // Nathi
			"hour": {PluralOther: "ชั่วโมง"}, // Chua mong
			"day":  {PluralOther: "วัน"},    // Wan
			"year": {PluralOther: "ปี"},     // Pee
		},
	}
}

func registerVN() {
	registry["vi"] = Locale{
		Code: "vi",
		PluralRule: func(n int) PluralCategory {
			return PluralOther
		},
		Dictionary: map[string]string{
			"just_now": "vừa xong",
			"ago":      "trước",
			"in":       "trong",
			"s":        "giây",
			"m":        "phút",
			"h":        "giờ",
			"d":        "ngày",
			"y":        "năm",
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOther: "giây"},
			"min":  {PluralOther: "phút"},
			"hour": {PluralOther: "giờ"},
			"day":  {PluralOther: "ngày"},
			"year": {PluralOther: "năm"},
		},
	}
}

func registerJP() {
	registry["ja"] = Locale{
		Code: "ja",
		PluralRule: func(n int) PluralCategory {
			return PluralOther
		},
		Dictionary: map[string]string{
			"just_now": "たった今", // Tatta ima
			"ago":      "前",    // Mae
			"in":       "後",    // Go (After/In context)
			"s":        "秒",    // Byo
			"m":        "分",    // Fun
			"h":        "時間",   // Jikan
			"d":        "日",    // Nichi
			"y":        "年",    // Nen
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOther: "秒"},
			"min":  {PluralOther: "分"},
			"hour": {PluralOther: "時間"},
			"day":  {PluralOther: "日"}, // Or Nichi-kan for duration? Usually just Nichi + Mae usually suffices
			"year": {PluralOther: "年"},
		},
	}
}

func registerMY() {
	// Malay is very similar to Indonesian but has nuanced differences
	registry["ms"] = Locale{
		Code: "ms",
		PluralRule: func(n int) PluralCategory {
			return PluralOther
		},
		Dictionary: map[string]string{
			"just_now": "baru saja",
			"ago":      "lepas",       // 5 minit lepas (vs lalu)
			"in":       "dalam",
			"s":        "saat",
			"m":        "minit",
			"h":        "jam",
			"d":        "hari",
			"y":        "tahun",
		},
		Plurals: map[string]map[PluralCategory]string{
			"sec":  {PluralOther: "saat"},
			"min":  {PluralOther: "minit"},
			"hour": {PluralOther: "jam"},
			"day":  {PluralOther: "hari"},
			"year": {PluralOther: "tahun"},
		},
	}
}

// GetTrans retrieves a static translation.
func GetTrans(lang, key string) string {
	loc, ok := registry[lang]
	if !ok {
		loc = registry["en"] // Fallback to EN
	}
	
	if val, ok := loc.Dictionary[key]; ok {
		return val
	}
	
	// Fallback to EN dictionary if key missing in target lang
	if fallbackVal, ok := registry["en"].Dictionary[key]; ok {
		return fallbackVal
	}
	
	return key // Return key if absolutely nothing found
}

// GetPlural retrieves a word form based on count.
func GetPlural(lang, key string, count int) string {
	loc, ok := registry[lang]
	if !ok {
		loc = registry["en"]
	}

	category := loc.PluralRule(count)
	
	// Try finding the specific plural form
	if forms, ok := loc.Plurals[key]; ok {
		if val, ok := forms[category]; ok {
			return val
		}
		// Fallback to Other if specific category missing
		if val, ok := forms[PluralOther]; ok {
			return val
		}
	}

	// Fallback to EN logic
	enLoc := registry["en"]
	enCategory := enLoc.PluralRule(count)
	if forms, ok := enLoc.Plurals[key]; ok {
		if val, ok := forms[enCategory]; ok {
			return val
		}
		if val, ok := forms[PluralOther]; ok {
			return val
		}
	}

	return key
}
