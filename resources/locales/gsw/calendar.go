package gsw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "Mai", 6: "Jun", 7: "Jul", 8: "Aug", 11: "Nov", 12: "Dez", 1: "Jan", 3: "Mär", 4: "Apr", 9: "Sep", 10: "Okt", 2: "Feb"}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 7: "J", 8: "A", 9: "S", 1: "J", 3: "M", 4: "A", 11: "N", 12: "D", 2: "F", 5: "M", 10: "O"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "März", 4: "April", 6: "Juni", 8: "Auguscht", 9: "Septämber", 10: "Oktoober", 11: "Novämber", 1: "Januar", 12: "Dezämber", 5: "Mai", 7: "Juli", 2: "Februar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mä.", 2: "Zi.", 3: "Mi.", 4: "Du.", 5: "Fr.", 6: "Sa.", 0: "Su."}, Narrow: ut.CalendarDayFormatNameValue{2: "D", 3: "M", 4: "D", 5: "F", 6: "S", 0: "S", 1: "M"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "Mittwuch", 4: "Dunschtig", 5: "Friitig", 6: "Samschtig", 0: "Sunntig", 1: "Määntig", 2: "Ziischtig"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "n.m.", "am": "v.m."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Vormittag", "pm": "Namittag"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "n. Chr.", Abbrev: "n. Chr.", Narrow: "n. Chr."}, BC: ut.CalendarEraFormatNames{Full: "v. Chr.", Abbrev: "v. Chr.", Narrow: "v. Chr."}}}}
