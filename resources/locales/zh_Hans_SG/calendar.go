package zh_Hans_SG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{DateEra: ut.DateEra{BC: ut.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "dd/MM/yy"}, AD: ut.CalendarDateFormat{Full: "y年M月d日EEEE", Long: "y年M月d日", Medium: "y年M月d日", Short: "dd/MM/yy"}}, Time: ut.CalendarDateFormat{Full: "zzzz ah:mm:ss", Long: "z ah:mm:ss", Medium: "ah:mm:ss", Short: "ah:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "11月", 6: "6月", 10: "10月", 12: "12月", 5: "5月", 7: "7月", 8: "8月", 1: "1月", 9: "9月", 2: "2月", 3: "3月", 4: "4月"}, Narrow: ut.CalendarMonthFormatNameValue{10: "10", 12: "12", 1: "1", 3: "3", 4: "4", 6: "6", 7: "7", 9: "9", 11: "11", 2: "2", 8: "8", 5: "5"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{4: "四月", 8: "八月", 11: "十一月", 2: "二月", 7: "七月", 12: "十二月", 1: "一月", 3: "三月", 5: "五月", 9: "九月", 6: "六月", 10: "十月"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "周三", 4: "周四", 5: "周五", 6: "周六", 0: "周日", 1: "周一", 2: "周二"}, Narrow: ut.CalendarDayFormatNameValue{5: "五", 6: "六", 0: "日", 1: "一", 2: "二", 3: "三", 4: "四"}, Short: ut.CalendarDayFormatNameValue{5: "周五", 6: "周六", 0: "周日", 1: "周一", 2: "周二", 3: "周三", 4: "周四"}, Wide: ut.CalendarDayFormatNameValue{5: "星期五", 6: "星期六", 0: "星期日", 1: "星期一", 2: "星期二", 3: "星期三", 4: "星期四"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "下午", "morning2": "上午", "night1": "凌晨", "afternoon1": "中午", "afternoon2": "下午", "evening1": "晚上", "midnight": "午夜", "am": "上午", "morning1": "早上"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "晚上", "night1": "凌晨", "morning2": "上午", "afternoon2": "下午", "pm": "下午", "am": "上午", "afternoon1": "中午", "midnight": "午夜", "morning1": "早上"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "晚上", "pm": "下午", "morning1": "早上", "am": "上午", "afternoon1": "中午", "afternoon2": "下午", "midnight": "午夜", "night1": "凌晨", "morning2": "上午"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
