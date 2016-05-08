package ky

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BYR": ut.Currency{Currency: "BYR", DisplayName: "беларусь рублу", Symbol: "BYR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руанда франкы", Symbol: "RWF"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "тринидад жана тобаго доллары", Symbol: "TTD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Гонконг доллары", Symbol: "HKD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "никарагуа кордобасы", Symbol: "NIO"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистан руписи", Symbol: "PKR"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Йемен риалы", Symbol: "YER"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "гаити гурдусу", Symbol: "HTG"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Макау патакасы", Symbol: "MOP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филиппин песосу", Symbol: "PHP"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Сао Томе жана Принсипе добрасы", Symbol: "STD"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Афганстан афганиси", Symbol: "AFN"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Жибути франкы", Symbol: "DJF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Ыйык Елена фунту", Symbol: "SHP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Тай баты", Symbol: "฿"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "албан леги", Symbol: "ALL"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Армения драмы", Symbol: "AMD"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Гана седиси", Symbol: "GHS"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Мадагаскар ариариси", Symbol: "MGA"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мйанмар кйаты", Symbol: "MMK"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Түштүк Африка ранды", Symbol: "ZAR"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Ирак динары", Symbol: "IQD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "орус рублу", Symbol: "RUB"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "КФП франкы", Symbol: "CFPF"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "бразилия реалы", Symbol: "BRL"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "гуйана доллары", Symbol: "GYD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Намибия доллары", Symbol: "NAD"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "куба песосу", Symbol: "CUP"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "гибралтар фунту", Symbol: "GIP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаос киби", Symbol: "LAK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "норвегия крону", Symbol: "NOK"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "аргентина песосу", Symbol: "ARS"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладеш такасы", Symbol: "BDT"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурунди франкы", Symbol: "BIF"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Коморос франкы", Symbol: "KMF"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монгол тугриги", Symbol: "MNT"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свази лилангени", Symbol: "SZL"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вьетнам доӊу", Symbol: "₫"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Түркмөнстан манаты", Symbol: "TMT"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "багама доллары", Symbol: "BSD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гине франкы", Symbol: "GNF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ямайка доллары", Symbol: "JMD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шри Ланка руписи", Symbol: "LKR"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуа Жаӊы Гине кинасы", Symbol: "PGK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "румын лейи", Symbol: "RON"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "босния-герцоговина жүгүртөлмөлүү маркасы", Symbol: "BAM"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонезия рупийасы", Symbol: "IDR"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Иран риалы", Symbol: "IRR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Түштүк Корея уону", Symbol: "KRW"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "марокко дирхамы", Symbol: "MAD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Мавританий руписи", Symbol: "MUR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "панама балбоасы", Symbol: "PAB"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Бириккен Араб Эмираттары дирхамы", Symbol: "AED"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "белиз доллары", Symbol: "BZD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "молдован лейи", Symbol: "MDL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "польша злотыйы", Symbol: "PLN"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Соломон доллары", Symbol: "SBD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "тунис динары", Symbol: "TND"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "уругвай песосу", Symbol: "UYU"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "бермуд доллары", Symbol: "BMD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конго франкы", Symbol: "CDF"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Кытай юаны", Symbol: "CN¥"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "ливия динары", Symbol: "LYD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзания шиллинги", Symbol: "TZS"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Ангола кванзасы", Symbol: "AOA"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Бруней доллары", Symbol: "BND"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "чех кронасы", Symbol: "CZK"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сирия фунту", Symbol: "SYP"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "венесуэла боливары", Symbol: "VEF"}, "USD": ut.Currency{Currency: "USD", DisplayName: "АКШ доллары", Symbol: "USD"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "нидерланд-антил гулдени", Symbol: "ANG"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Эритреа накфасы", Symbol: "ERN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Эфиопия бирри", Symbol: "ETB"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузия лариси", Symbol: "GEL"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "хорват кунасы", Symbol: "HRK"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израил жаӊы шегели", Symbol: "ILS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индия руписи", Symbol: "INR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "швейцария франкы", Symbol: "CHF"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "фолкленд аралдарынын фунту", Symbol: "FKP"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "судан фунту", Symbol: "SDG"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "суринам доллары", Symbol: "SRD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Түркия лирасы", Symbol: "TRY"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Белгисиз акча", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Оман риалы", Symbol: "OMR"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапур доллары", Symbol: "SGD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Уганда шиллинги", Symbol: "UGX"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "гондурас лемпирасы", Symbol: "HNL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малави квачасы", Symbol: "MWK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непал руписи", Symbol: "NPR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомали шиллинги", Symbol: "SOS"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Тайвань жаӊы доллары", Symbol: "TWD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "КФА ВЕАС франкы", Symbol: "FCFA"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "чыгыш кариб доллары", Symbol: "XCD"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Кабо-Верде эскудосу", Symbol: "CVE"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "мексика песосу", Symbol: "MXN"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "парагвай гуараниси", Symbol: "PYG"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоа таласы", Symbol: "WST"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "барбадос доллары", Symbol: "BBD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Кыргызстан сому", Symbol: "сом"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувейт динары", Symbol: "KWD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Ливан фунту", Symbol: "LBP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либерия доллары", Symbol: "LRD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "украин гривени", Symbol: "UAH"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Австралия доллары", Symbol: "AUD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "аруба флорини", Symbol: "AWG"}, "COP": ut.Currency{Currency: "COP", DisplayName: "колумбия песосу", Symbol: "COP"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "исландия крону", Symbol: "ISK"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Түндүк Корея уону", Symbol: "KPW"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "перу нуэво солу", Symbol: "PEN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Сауд риалы", Symbol: "SAR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербайжан манаты", Symbol: "AZN"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "куба жүгүртүлмөлүү песосу", Symbol: "CUC"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Жапан йени", Symbol: "JP¥"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кения шиллинги", Symbol: "KES"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мавритания угиясы", Symbol: "MRO"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Өзбекстан сому", Symbol: "UZS"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фижи доллары", Symbol: "FJD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "кайман доллары", Symbol: "KYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигерия найрасы", Symbol: "NGN"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катар риалы", Symbol: "QAR"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбия даласиси", Symbol: "GMD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбик метикалы", Symbol: "MZN"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "швеция крону", Symbol: "SEK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Түштүк Судан фунту", Symbol: "SSP"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "доминикан песосу", Symbol: "DOP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "алжир динары", Symbol: "DZD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Йордания динары", Symbol: "JOD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Малдив руфийасы", Symbol: "MVR"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбия квачасы", Symbol: "ZMW"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахрейн динары", Symbol: "BHD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "чили песосу", Symbol: "CLP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "дания крону", Symbol: "DKK"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малайзия ринггити", Symbol: "MYR"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "венгр форинти", Symbol: "HUF"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутан нгултруму", Symbol: "BTN"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "британия фунт стерлинги", Symbol: "GBP"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казакстан теӊгеси", Symbol: "KZT"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сейшел руписи", Symbol: "SCR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Тажикстан сомониси", Symbol: "TJS"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "боливия боливианосу", Symbol: "BOB"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "латвия латы", Symbol: "LVL"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Жаӊы Зеландия доллары", Symbol: "NZD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "серб динары", Symbol: "RSD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сиерра-Леоне леонеси", Symbol: "SLL"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Ботсвана пуласы", Symbol: "BWP"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "канада доллары", Symbol: "CAD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "египет фунту", Symbol: "EGP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "гватемала кетсалы", Symbol: "GTQ"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "македон денары", Symbol: "MKD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонга паангасы", Symbol: "TOP"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "болгар левиси", Symbol: "BGN"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "коста-рика колону", Symbol: "CRC"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Камбожа риели", Symbol: "KHR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "литва литасы", Symbol: "LTL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануату ватусу", Symbol: "VUV"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "КФА ВСЕАО франкы", Symbol: "CFA"}}
