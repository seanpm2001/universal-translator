package nl_CW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, AD: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan.", 2: "feb.", 5: "mei", 10: "okt.", 4: "apr.", 7: "jul.", 11: "nov.", 12: "dec.", 8: "aug.", 9: "sep.", 3: "mrt.", 6: "jun."}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 11: "N", 12: "D", 8: "A", 1: "J", 4: "A", 5: "M", 6: "J", 3: "M", 9: "S", 10: "O", 2: "F"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mei", 7: "juli", 8: "augustus", 12: "december", 3: "maart", 9: "september", 10: "oktober", 2: "februari", 6: "juni", 11: "november", 1: "januari", 4: "april"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za"}, Narrow: ut.CalendarDayFormatNameValue{2: "D", 3: "W", 4: "D", 5: "V", 6: "Z", 0: "Z", 1: "M"}, Short: ut.CalendarDayFormatNameValue{2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma"}, Wide: ut.CalendarDayFormatNameValue{3: "woensdag", 4: "donderdag", 5: "vrijdag", 6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht", "am": "a.m.", "pm": "p.m."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
