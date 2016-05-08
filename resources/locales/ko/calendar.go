package ko

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "y년 M월 d일 EEEE", Long: "y년 M월 d일", Medium: "y. M. d.", Short: "yy. M. d."}, AD: ut.CalendarDateFormat{Full: "y년 M월 d일 EEEE", Long: "y년 M월 d일", Medium: "y. M. d.", Short: "yy. M. d."}}, Time: ut.CalendarDateFormat{Full: "a h시 m분 s초 zzzz", Long: "a h시 m분 s초 z", Medium: "a h:mm:ss", Short: "a h:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "3월", 5: "5월", 6: "6월", 7: "7월", 8: "8월", 12: "12월", 1: "1월", 4: "4월", 9: "9월", 10: "10월", 11: "11월", 2: "2월"}, Narrow: ut.CalendarMonthFormatNameValue{7: "7월", 8: "8월", 9: "9월", 11: "11월", 12: "12월", 2: "2월", 6: "6월", 4: "4월", 5: "5월", 10: "10월", 1: "1월", 3: "3월"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "10월", 12: "12월", 1: "1월", 4: "4월", 9: "9월", 6: "6월", 7: "7월", 8: "8월", 11: "11월", 2: "2월", 3: "3월", 5: "5월"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "토", 0: "일", 1: "월", 2: "화", 3: "수", 4: "목", 5: "금"}, Narrow: ut.CalendarDayFormatNameValue{4: "목", 5: "금", 6: "토", 0: "일", 1: "월", 2: "화", 3: "수"}, Short: ut.CalendarDayFormatNameValue{0: "일", 1: "월", 2: "화", 3: "수", 4: "목", 5: "금", 6: "토"}, Wide: ut.CalendarDayFormatNameValue{6: "토요일", 0: "일요일", 1: "월요일", 2: "화요일", 3: "수요일", 4: "목요일", 5: "금요일"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning2": "오전", "evening1": "저녁", "noon": "정오", "am": "오전", "pm": "오후", "morning1": "새벽", "afternoon1": "오후", "night1": "밤", "midnight": "자정"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "오후", "night1": "밤", "midnight": "자정", "am": "AM", "pm": "PM", "morning2": "오전", "noon": "정오", "morning1": "새벽", "evening1": "저녁"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "오전", "noon": "정오", "pm": "오후", "morning1": "새벽", "evening1": "저녁", "midnight": "자정", "afternoon1": "오후", "night1": "밤", "morning2": "오전"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "서기", Abbrev: "AD", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "기원전", Abbrev: "BC", Narrow: ""}}}}
