package gd

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d'mh' MMMM y", Long: "d'mh' MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, AD: ut.CalendarDateFormat{Full: "EEEE, d'mh' MMMM y", Long: "d'mh' MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "Màrt", 4: "Gibl", 5: "Cèit", 6: "Ògmh", 7: "Iuch", 8: "Lùna", 10: "Dàmh", 1: "Faoi", 9: "Sult", 11: "Samh", 12: "Dùbh", 2: "Gearr"}, Narrow: ut.CalendarMonthFormatNameValue{5: "C", 6: "Ò", 7: "I", 1: "F", 2: "G", 3: "M", 4: "G", 12: "D", 8: "L", 9: "S", 10: "D", 11: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "An Gearran", 3: "Am Màrt", 5: "An Cèitean", 6: "An t-Ògmhios", 9: "An t-Sultain", 12: "An Dùbhlachd", 1: "Am Faoilleach", 7: "An t-Iuchar", 8: "An Lùnastal", 10: "An Dàmhair", 11: "An t-Samhain", 4: "An Giblean"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "DiC", 4: "Dia", 5: "Dih", 6: "DiS", 0: "DiD", 1: "DiL", 2: "DiM"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "C", 4: "A", 5: "H"}, Short: ut.CalendarDayFormatNameValue{4: "Da", 5: "hA", 6: "Sa", 0: "Dò", 1: "Lu", 2: "Mà", 3: "Ci"}, Wide: ut.CalendarDayFormatNameValue{0: "DiDòmhnaich", 1: "DiLuain", 2: "DiMàirt", 3: "DiCiadain", 4: "DiarDaoin", 5: "DihAoine", 6: "DiSathairne"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "m", "pm": "f"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "m", "pm": "f"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "m", "pm": "f"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Ro Chrìosta", Abbrev: "RC", Narrow: "R"}}}}
