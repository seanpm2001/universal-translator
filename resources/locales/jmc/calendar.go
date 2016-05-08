package jmc

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sep", 12: "Des", 1: "Jan", 2: "Feb", 6: "Jun", 7: "Jul", 8: "Ago", 10: "Okt", 11: "Nov", 3: "Mac", 4: "Apr", 5: "Mei"}, Narrow: ut.CalendarMonthFormatNameValue{3: "M", 4: "A", 5: "M", 6: "J", 7: "J", 11: "N", 1: "J", 8: "A", 9: "S", 10: "O", 12: "D", 2: "F"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "Oktoba", 12: "Desemba", 1: "Januari", 4: "Aprilyi", 8: "Agusti", 6: "Junyi", 7: "Julyai", 9: "Septemba", 11: "Novemba", 2: "Februari", 3: "Machi", 5: "Mei"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Jtn", 4: "Alh", 5: "Iju", 6: "Jmo", 0: "Jpi", 1: "Jtt", 2: "Jnn"}, Narrow: ut.CalendarDayFormatNameValue{1: "J", 2: "J", 3: "J", 4: "A", 5: "I", 6: "J", 0: "J"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{5: "Ijumaa", 6: "Jumamosi", 0: "Jumapilyi", 1: "Jumatatuu", 2: "Jumanne", 3: "Jumatanu", 4: "Alhamisi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "utuko", "pm": "kyiukonyi"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "utuko", "pm": "kyiukonyi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Kabla ya Kristu", Abbrev: "KK", Narrow: ""}}}}
