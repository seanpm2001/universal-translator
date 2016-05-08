package es_PA

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "MM/dd/y", Short: "MM/dd/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "MM/dd/y", Short: "MM/dd/yy"}}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "nov.", 1: "ene.", 2: "feb.", 8: "ago.", 6: "jun.", 7: "jul.", 9: "sept.", 10: "oct.", 12: "dic.", 3: "mar.", 4: "abr.", 5: "may."}, Narrow: ut.CalendarMonthFormatNameValue{1: "E", 4: "A", 7: "J", 8: "A", 9: "S", 2: "F", 3: "M", 5: "M", 6: "J", 10: "O", 11: "N", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "julio", 9: "septiembre", 1: "enero", 8: "agosto", 10: "octubre", 4: "abril", 12: "diciembre", 3: "marzo", 6: "junio", 11: "noviembre", 2: "febrero", 5: "mayo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "X", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{4: "JU", 5: "VI", 6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI"}, Wide: ut.CalendarDayFormatNameValue{0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "de la madrugada", "morning2": "de la mañana", "evening1": "de la tarde", "night1": "de la noche", "am": "a. m.", "pm": "p. m.", "noon": "mediodía"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning2": "de la mañana", "evening1": "de la tarde", "night1": "de la noche", "am": "a. m.", "pm": "p. m.", "noon": "mediodía", "morning1": "de la madrugada"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
