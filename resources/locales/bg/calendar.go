package bg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y 'г'.", Short: "d.MM.yy 'г'."}, AD: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d.MM.y 'г'.", Short: "d.MM.yy 'г'."}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "авг", 9: "сеп", 12: "дек", 1: "яну", 2: "фев", 6: "юни", 7: "юли", 10: "окт", 11: "ное", 3: "март", 4: "апр", 5: "май"}, Narrow: ut.CalendarMonthFormatNameValue{1: "я", 3: "м", 4: "а", 5: "м", 6: "ю", 7: "ю", 9: "с", 10: "о", 11: "н", 12: "д", 2: "ф", 8: "а"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "март", 4: "април", 5: "май", 6: "юни", 7: "юли", 8: "август", 9: "септември", 1: "януари", 12: "декември", 10: "октомври", 11: "ноември", 2: "февруари"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "чт", 5: "пт", 6: "сб", 0: "нд", 1: "пн", 2: "вт", 3: "ср"}, Narrow: ut.CalendarDayFormatNameValue{0: "н", 1: "п", 2: "в", 3: "с", 4: "ч", 5: "п", 6: "с"}, Short: ut.CalendarDayFormatNameValue{0: "нд", 1: "пн", 2: "вт", 3: "ср", 4: "чт", 5: "пт", 6: "сб"}, Wide: ut.CalendarDayFormatNameValue{0: "неделя", 1: "понеделник", 2: "вторник", 3: "сряда", 4: "четвъртък", 5: "петък", 6: "събота"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "полунощ", "am": "пр.об.", "pm": "сл.об.", "morning1": "сутрин", "morning2": "на обед", "afternoon1": "следобед", "evening1": "вечер", "night1": "нощ"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "нощ", "midnight": "полунощ", "am": "пр.об.", "pm": "сл.об.", "morning1": "сутрин", "morning2": "на обед", "afternoon1": "следобед", "evening1": "вечер"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "пр.об.", "pm": "сл.об.", "morning1": "сутринта", "morning2": "на обед", "afternoon1": "следобед", "evening1": "вечерта", "night1": "нощ", "midnight": "полунощ"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "след Христа", Abbrev: "сл.Хр.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "преди Христа", Abbrev: "пр.Хр.", Narrow: ""}}}}
