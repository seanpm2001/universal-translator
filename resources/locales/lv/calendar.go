package lv

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, y. 'gada' d. MMMM", Long: "y. 'gada' d. MMMM", Medium: "y. 'gada' d. MMM", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, y. 'gada' d. MMMM", Long: "y. 'gada' d. MMMM", Medium: "y. 'gada' d. MMM", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Janv.", 3: "Marts", 4: "Apr.", 10: "Okt.", 11: "Nov.", 9: "Sept.", 12: "Dec.", 2: "Febr.", 5: "Maijs", 6: "Jūn.", 7: "Jūl.", 8: "Aug."}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 9: "S", 10: "O", 11: "N", 12: "D", 3: "M", 6: "J", 7: "J", 5: "M", 1: "J", 2: "F", 4: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Februāris", 3: "Marts", 4: "Aprīlis", 7: "Jūlijs", 8: "Augusts", 11: "Novembris", 1: "Janvāris", 5: "Maijs", 6: "Jūnijs", 9: "Septembris", 10: "Oktobris", 12: "Decembris"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sv", 1: "Pr", 2: "Ot", 3: "Tr", 4: "Ce", 5: "Pk", 6: "Se"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "P", 2: "O", 3: "T", 4: "C", 5: "P"}, Short: ut.CalendarDayFormatNameValue{3: "Tr", 4: "Ce", 5: "Pk", 6: "Se", 0: "Sv", 1: "Pr", 2: "Ot"}, Wide: ut.CalendarDayFormatNameValue{6: "Sestdiena", 0: "Svētdiena", 1: "Pirmdiena", 2: "Otrdiena", 3: "Trešdiena", 4: "Ceturtdiena", 5: "Piektdiena"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "rīts", "afternoon1": "pēcpusdiena", "evening1": "vakars", "night1": "nakts", "midnight": "pusnakts", "am": "priekšp.", "pm": "pēcpusd."}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "vakars", "night1": "nakts", "midnight": "pusnakts", "am": "priekšp.", "noon": "pusd.", "pm": "pēcp.", "morning1": "rīts", "afternoon1": "pēcpusdiena"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "pēcpusdiena", "morning1": "rīts", "afternoon1": "pēcpusdiena", "evening1": "vakars", "night1": "nakts", "midnight": "pusnakts", "am": "priekšpusdienā", "noon": "pusdienlaiks"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: "m.ē."}, BC: ut.CalendarEraFormatNames{Full: "pirms mūsu ēras", Abbrev: "p.m.ē.", Narrow: "p.m.ē."}}}}
