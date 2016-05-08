package bs_Cyrl

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}, AD: ut.CalendarDateFormat{Full: "EEEE, dd. MMMM y.", Long: "dd. MMMM y.", Medium: "dd.MM.y.", Short: "d.M.yy."}}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'u' {0}", Long: "{1} 'u' {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "окт", 11: "нов", 1: "јан", 3: "мар", 5: "мај", 6: "јун", 7: "јул", 8: "авг", 12: "дец", 2: "феб", 4: "апр", 9: "сеп"}, Narrow: ut.CalendarMonthFormatNameValue{12: "д", 2: "ф", 11: "н", 4: "а", 5: "м", 6: "ј", 7: "ј", 8: "а", 9: "с", 1: "ј", 3: "м", 10: "о"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "април", 5: "мај", 6: "јуни", 7: "јули", 8: "август", 9: "септембар", 1: "јануар", 3: "март", 11: "новембар", 12: "децембар", 2: "фебруар", 10: "октобар"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "пет", 6: "суб", 0: "нед", 1: "пон", 2: "уто", 3: "сри", 4: "чет"}, Narrow: ut.CalendarDayFormatNameValue{2: "у", 3: "с", 4: "ч", 5: "п", 6: "с", 0: "н", 1: "п"}, Short: ut.CalendarDayFormatNameValue{3: "sri", 4: "čet", 5: "pet", 6: "sub", 0: "ned", 1: "pon", 2: "uto"}, Wide: ut.CalendarDayFormatNameValue{0: "недеља", 1: "понедељак", 2: "уторак", 3: "сриједа", 4: "четвртак", 5: "петак", 6: "субота"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "поподне", "midnight": "ponoć", "noon": "podne", "morning1": "jutro", "afternoon1": "poslijepodne", "evening1": "večer", "night1": "noć", "am": "пре подне"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "prijepodne", "noon": "podne", "pm": "popodne"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"midnight": "u ponoć", "noon": "podne", "morning1": "jutro", "afternoon1": "poslijepodne", "evening1": "veče", "night1": "noć", "am": "пре подне", "pm": "поподне"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Пре нове ере", Abbrev: "п. н. е.", Narrow: "п.н.е."}}}}
