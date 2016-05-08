package sr_Latn

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, AD: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "jun", 10: "okt", 3: "mar", 4: "apr", 5: "maj", 7: "jul", 8: "avg", 9: "sep", 11: "nov", 12: "dec", 1: "jan", 2: "feb"}, Narrow: ut.CalendarMonthFormatNameValue{6: "j", 7: "j", 8: "a", 10: "o", 1: "j", 2: "f", 4: "a", 5: "m", 11: "n", 3: "m", 9: "s", 12: "d"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "jun", 9: "septembar", 11: "novembar", 12: "decembar", 1: "januar", 2: "februar", 3: "mart", 4: "april", 5: "maj", 7: "jul", 8: "avgust", 10: "oktobar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "ned", 1: "pon", 2: "uto", 3: "sre", 4: "čet", 5: "pet", 6: "sub"}, Narrow: ut.CalendarDayFormatNameValue{0: "n", 1: "p", 2: "u", 3: "s", 4: "č", 5: "p", 6: "s"}, Short: ut.CalendarDayFormatNameValue{1: "po", 2: "ut", 3: "sr", 4: "če", 5: "pe", 6: "su", 0: "ne"}, Wide: ut.CalendarDayFormatNameValue{5: "petak", 6: "subota", 0: "nedelja", 1: "ponedeljak", 2: "utorak", 3: "sreda", 4: "četvrtak"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "popodne", "evening1": "veče", "night1": "noć", "midnight": "ponoć", "am": "pre podne", "noon": "podne", "pm": "po podne", "morning1": "jutro"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "popodne", "evening1": "veče", "night1": "noć", "midnight": "ponoć", "am": "pre podne", "noon": "podne", "pm": "po podne", "morning1": "jutro"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "podne", "pm": "po podne", "morning1": "jutro", "afternoon1": "popodne", "evening1": "veče", "night1": "noć", "midnight": "ponoć", "am": "pre podne"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "nove ere", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "pre nove ere", Abbrev: "p. n. e.", Narrow: "p.n.e."}}}}
