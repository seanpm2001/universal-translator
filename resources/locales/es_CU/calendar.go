package es_CU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, AD: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "jun.", 7: "jul.", 8: "ago.", 10: "oct.", 3: "mar.", 4: "abr.", 5: "may.", 1: "ene.", 2: "feb.", 11: "nov.", 9: "sept.", 12: "dic."}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 10: "O", 4: "A", 6: "J", 11: "N", 3: "M", 8: "A", 12: "D", 1: "E", 2: "F", 5: "M", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "junio", 10: "octubre", 8: "agosto", 1: "enero", 3: "marzo", 7: "julio", 11: "noviembre", 2: "febrero", 4: "abril", 5: "mayo", 9: "septiembre", 12: "diciembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dom.", 1: "lun.", 2: "mar.", 3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb."}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "X", 4: "J", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI"}, Wide: ut.CalendarDayFormatNameValue{3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
