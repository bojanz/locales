package hr_BA

import (
	"math"
	"strconv"
	"time"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/currency"
)

type hr_BA struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyPositiveSuffix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'hr_BA' locale
func New() locales.Translator {
	return &hr_BA{
		locale:                 "hr_BA",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "KM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UZS", "VEB", "VEF", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "sij", "velj", "ožu", "tra", "svi", "lip", "srp", "kol", "ruj", "lis", "stu", "pro"},
		monthsNarrow:           []string{"", "1.", "2.", "3.", "4.", "5.", "6.", "7.", "8.", "9.", "10.", "11.", "12."},
		monthsWide:             []string{"", "siječnja", "veljače", "ožujka", "travnja", "svibnja", "lipnja", "srpnja", "kolovoza", "rujna", "listopada", "studenoga", "prosinca"},
		daysAbbreviated:        []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysNarrow:             []string{"N", "P", "U", "S", "Č", "P", "S"},
		daysShort:              []string{"ned", "pon", "uto", "sri", "čet", "pet", "sub"},
		daysWide:               []string{"nedjelja", "ponedjeljak", "utorak", "srijeda", "četvrtak", "petak", "subota"},
		periodsAbbreviated:     []string{"AM", "PM"},
		periodsNarrow:          []string{"AM", "PM"},
		periodsWide:            []string{"AM", "PM"},
		erasAbbreviated:        []string{"pr. Kr.", "p. Kr."},
		erasNarrow:             []string{"pr.n.e.", "AD"},
		erasWide:               []string{"prije Krista", "poslije Krista"},
		timezones:              map[string]string{"OESZ": "istočnoeuropsko ljetno vrijeme", "MESZ": "srednjoeuropsko ljetno vrijeme", "WESZ": "zapadnoeuropsko ljetno vrijeme", "ARST": "argentinsko ljetno vrijeme", "SAST": "južnoafričko vrijeme", "∅∅∅": "peruansko ljetno vrijeme", "CLST": "čileansko ljetno vrijeme", "UYST": "urugvajsko ljetno vrijeme", "IST": "indijsko vrijeme", "CHADT": "ljetno vrijeme Chathama", "EST": "istočno standardno vrijeme", "HADT": "havajsko-aleutsko ljetno vrijeme", "ADT": "atlantsko ljetno vrijeme", "BT": "butansko vrijeme", "NZDT": "novozelandsko ljetno vrijeme", "AEST": "istočnoaustralsko standardno vrijeme", "WART": "zapadno-argentinsko standardno vrijeme", "ACST": "srednjoaustralsko standardno vrijeme", "JST": "japansko standardno vrijeme", "CST": "središnje standardno vrijeme", "COST": "kolumbijsko ljetno vrijeme", "WITA": "srednjoindonezijsko vrijeme", "GMT": "univerzalno vrijeme", "ACDT": "srednjoaustralsko ljetno vrijeme", "WIB": "zapadnoindonezijsko vrijeme", "CLT": "čileansko standardno vrijeme", "MST": "planinsko standardno vrijeme", "SRT": "surinamsko vrijeme", "EAT": "istočnoafričko vrijeme", "HKST": "hongkonško ljetno vrijeme", "WAT": "zapadnoafričko standardno vrijeme", "GFT": "vrijeme Francuske Gvajane", "VET": "venezuelsko vrijeme", "PST": "pacifičko standardno vrijeme", "AKST": "aljaško standardno vrijeme", "MEZ": "srednjoeuropsko standardno vrijeme", "ACWST": "australsko središnje zapadno standardno vrijeme", "ACWDT": "australsko središnje zapadno ljetno vrijeme", "HNT": "newfoundlandsko standardno vrijeme", "BOT": "bolivijsko vrijeme", "ECT": "ekvadorsko vrijeme", "MDT": "planinsko ljetno vrijeme", "AEDT": "istočnoaustralsko ljetno vrijeme", "EDT": "istočno ljetno vrijeme", "PDT": "pacifičko ljetno vrijeme", "SGT": "singapursko vrijeme", "OEZ": "istočnoeuropsko standardno vrijeme", "HAT": "newfoundlandsko ljetno vrijeme", "LHST": "standardno vrijeme otoka Lord Howe", "NZST": "novozelandsko standardno vrijeme", "ART": "argentinsko standardno vrijeme", "JDT": "japansko ljetno vrijeme", "HKT": "hongkonško standardno vrijeme", "WAST": "zapadnoafričko ljetno vrijeme", "ChST": "standardno vrijeme Chamorra", "WARST": "zapadno-argentinsko ljetno vrijeme", "WIT": "istočnoindonezijsko vrijeme", "LHDT": "ljetno vrijeme otoka Lord Howe", "CAT": "srednjoafričko vrijeme", "UYT": "urugvajsko standardno vrijeme", "COT": "kolumbijsko standardno vrijeme", "HAST": "havajsko-aleutsko standardno vrijeme", "TMST": "turkmenistansko ljetno vrijeme", "GYT": "gvajansko vrijeme", "TMT": "turkmenistansko standardno vrijeme", "MYT": "malezijsko vrijeme", "AKDT": "aljaško ljetno vrijeme", "AST": "atlantsko standardno vrijeme", "WEZ": "zapadnoeuropsko standardno vrijeme", "CDT": "središnje ljetno vrijeme", "AWST": "zapadnoaustralsko standardno vrijeme", "AWDT": "zapadnoaustralsko ljetno vrijeme", "CHAST": "standardno vrijeme Chathama"},
	}
}

// Locale returns the current translators string locale
func (hr *hr_BA) Locale() string {
	return hr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hr_BA'
func (hr *hr_BA) PluralsCardinal() []locales.PluralRule {
	return hr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hr_BA'
func (hr *hr_BA) PluralsOrdinal() []locales.PluralRule {
	return hr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'hr_BA'
func (hr *hr_BA) PluralsRange() []locales.PluralRule {
	return hr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hr_BA'
func (hr *hr_BA) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod10 := i % 10
	iMod100 := i % 100
	fMod10 := f % 10
	fMod100 := f % 100

	if (v == 0 && iMod10 == 1 && iMod100 != 11) || (fMod10 == 1 && fMod100 != 11) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod10 >= 2 && iMod10 <= 4 && (iMod100 < 12 || iMod100 > 14)) || (fMod10 >= 2 && fMod10 <= 4 && (fMod100 < 12 || fMod100 > 14)) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hr_BA'
func (hr *hr_BA) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hr_BA'
func (hr *hr_BA) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := hr.CardinalPluralRule(num1, v1)
	end := hr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hr *hr_BA) MonthAbbreviated(month time.Month) string {
	return hr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hr *hr_BA) MonthsAbbreviated() []string {
	return hr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hr *hr_BA) MonthNarrow(month time.Month) string {
	return hr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hr *hr_BA) MonthsNarrow() []string {
	return hr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hr *hr_BA) MonthWide(month time.Month) string {
	return hr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hr *hr_BA) MonthsWide() []string {
	return hr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hr *hr_BA) WeekdayAbbreviated(weekday time.Weekday) string {
	return hr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hr *hr_BA) WeekdaysAbbreviated() []string {
	return hr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hr *hr_BA) WeekdayNarrow(weekday time.Weekday) string {
	return hr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hr *hr_BA) WeekdaysNarrow() []string {
	return hr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hr *hr_BA) WeekdayShort(weekday time.Weekday) string {
	return hr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hr *hr_BA) WeekdaysShort() []string {
	return hr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hr *hr_BA) WeekdayWide(weekday time.Weekday) string {
	return hr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hr *hr_BA) WeekdaysWide() []string {
	return hr.daysWide
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hr_BA' and handles both Whole and Real numbers based on 'v'
func (hr *hr_BA) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hr_BA' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hr *hr_BA) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hr.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hr_BA'
func (hr *hr_BA) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hr.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hr_BA'
// in accounting notation.
func (hr *hr_BA) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, hr.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hr.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hr.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, hr.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Year()), 10)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'hr_BA'
func (hr *hr_BA) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
