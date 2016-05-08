package en_MW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y GGGG", Long: "MMMM d, y GG", Medium: "MMM d, y", Short: "M/d/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "May", 8: "Aug", 1: "Jan", 3: "Mar", 4: "Apr", 7: "Jul", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec", 2: "Feb", 6: "Jun"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 10: "O", 8: "A", 5: "M", 11: "N", 3: "M", 7: "J", 12: "D", 1: "J", 2: "F", 6: "J", 9: "S"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "October", 3: "March", 6: "June", 1: "January", 12: "December", 5: "May", 11: "November", 9: "September", 2: "February", 7: "July", 8: "August", 4: "April"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We"}, Wide: ut.CalendarDayFormatNameValue{4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
