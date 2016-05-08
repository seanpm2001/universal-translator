package pt_PT

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"THB": ut.Currency{Currency: "THB", DisplayName: "Baht da Tailândia", Symbol: "฿"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som do Uzbequistão", Symbol: "UZS"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florim de Aruba", Symbol: "AWG"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Coroa checa", Symbol: "CZK"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franco guineense", Symbol: "GNF"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial iraniano", Symbol: "IRR"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu moldavo", Symbol: "MDL"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dólar do Suriname", Symbol: "SRD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka de Bangladesh", Symbol: "BDT"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso Plata mexicano (1861–1992)", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Zloti polaco (1950–1995)", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "direito especial de saque", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunisino", Symbol: "TND"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dólar de Trindade e Tobago", Symbol: "TTD"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Libra de Chipre", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi do Gana", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dólar da Guiana", Symbol: "GYD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge do Cazaquistão", Symbol: "KZT"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham marroquino", Symbol: "MAD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina da Papua-Nova Guiné", Symbol: "PGK"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rublo novo bielorusso (1994–1999)", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu romeno", Symbol: "RON"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dólar de Singapura", Symbol: "SGD"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franco CFA (BCEAO)", Symbol: "CFA"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari de Madagáscar", Symbol: "MGA"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba nicaraguano", Symbol: "NIO"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelim somali", Symbol: "SOS"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni do Tajaquistão", Symbol: "TJS"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat do Turquemenistão", Symbol: "TMT"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha zambiano (1968–2012)", Symbol: "ZMK"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham dos Emirados Árabes Unidos", Symbol: "AED"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Marco bósnio-herzegóvino conversível", Symbol: "BAM"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dólar belizense", Symbol: "BZD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi da Gâmbia", Symbol: "GMD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Dinar macedónio", Symbol: "MKD"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Unidad de Inversion (UDI) mexicana", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum do Butão", Symbol: "BTN"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som do Quirguistão", Symbol: "KGS"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupia do Sri Lanka", Symbol: "LKR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia da Ucrânia", Symbol: "UAH"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Unidade da Moeda Europeia", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar forte jugoslavo", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dólar das Bahamas", Symbol: "BSD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira das Honduras", Symbol: "HNL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha do Malawi", Symbol: "MWK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical de Moçambique", Symbol: "MZN"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloti polaco", Symbol: "PLN"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial iemenita", Symbol: "YER"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Novo sol peruano", Symbol: "PEN"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram arménio", Symbol: "AMD"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franco belga (convertível)", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franco jibutiano", Symbol: "DJF"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Libra das Ilhas Falkland", Symbol: "FKP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip de Laos", Symbol: "LAK"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rupia das Ilhas Maldivas", Symbol: "MVR"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar da Bósnia-Herzegóvina", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra esterlina britânica", Symbol: "£"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas da Lituânia", Symbol: "LTL"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba nicaraguano (1988–1991)", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial de Omã", Symbol: "OMR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dólar dos Estados Unidos", Symbol: "US$"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat do Azerbaijão", Symbol: "AZN"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon costa-riquenho", Symbol: "CRC"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Unidad de Valor Constante (UVC) do Equador", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial do Catar", Symbol: "QAR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Rial saudita", Symbol: "SAR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu de Vanuatu", Symbol: "VUV"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afeghani (1927–2002)", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats da Letónia", Symbol: "LVL"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franco do Mali", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo português", Symbol: "\u200b"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni da Suazilândia", Symbol: "SZL"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dólar do Zimbabwe", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula de Botswana", Symbol: "BWP"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dólar de Fiji", Symbol: "FJD"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar conversível jugoslavo", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal da Guatemala", Symbol: "GTQ"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel novo israelita", Symbol: "₪"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa do Panamá", Symbol: "PAB"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga de Tonga", Symbol: "TOP"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franco CFA (BEAC)", Symbol: "FCFA"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Super Dinar jugoslavo", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dólar da Namíbia", Symbol: "NAD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar baremita", Symbol: "BHD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dólar bruneíno", Symbol: "BND"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dólar canadiano", Symbol: "CA$"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franco comoriano", Symbol: "KMF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat de Mianmar", Symbol: "MMK"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya da Mauritânia", Symbol: "MRO"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afegani do Afeganistão", Symbol: "AFN"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi de Gana", Symbol: "GHS"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dólar das Ilhas Caimão", Symbol: "KYD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik da Mongólia", Symbol: "MNT"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca de Macau", Symbol: "MOP"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dólar das Caraíbas Orientais", Symbol: "EC$"}}
