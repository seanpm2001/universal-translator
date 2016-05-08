package ar_EH

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, AD: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "ديسمبر", 6: "يونيو", 7: "يوليو", 8: "أغسطس", 1: "يناير", 2: "فبراير", 3: "مارس", 9: "سبتمبر", 10: "أكتوبر", 11: "نوفمبر", 4: "أبريل", 5: "مايو"}, Narrow: ut.CalendarMonthFormatNameValue{5: "و", 10: "ك", 11: "ب", 8: "غ", 9: "س", 1: "ي", 3: "م", 4: "أ", 6: "ن", 7: "ل", 2: "ف", 12: "د"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "ديسمبر", 6: "يونيو", 10: "أكتوبر", 8: "أغسطس", 5: "مايو", 7: "يوليو", 9: "سبتمبر", 11: "نوفمبر", 1: "يناير", 3: "مارس", 4: "أبريل", 2: "فبراير"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}, Narrow: ut.CalendarDayFormatNameValue{5: "ج", 6: "س", 0: "ح", 1: "ن", 2: "ث", 3: "ر", 4: "خ"}, Short: ut.CalendarDayFormatNameValue{5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس"}, Wide: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "مساءً", "morning1": "فجرا", "afternoon1": "ظهرًا", "night1": "منتصف الليل", "night2": "ليلاً", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "am": "ص", "pm": "م"}, Narrow: ut.CalendarPeriodFormatNameValue{"night2": "ليلاً", "pm": "م", "morning1": "فجرا", "afternoon1": "ظهرًا", "afternoon2": "بعد الظهر", "am": "ص", "night1": "منتصف الليل", "morning2": "صباحًا", "evening1": "مساءً"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "صباحًا", "pm": "مساءً", "night1": "منتصف الليل", "morning1": "فجرا", "afternoon1": "ظهرًا", "night2": "ليلاً", "afternoon2": "بعد الظهر", "morning2": "صباحًا", "evening1": "مساءً"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
