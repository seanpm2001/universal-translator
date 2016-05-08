package ln

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats ya Letoni", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guarani", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso ya Kolombi", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Falánga ya Ginɛ", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso ya Mexiko", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Motolé ya Norvej", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Falánga CFA BCEAO", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colon ya Kosta Rika", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinarɛ ya Libí", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas ya Litwani", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Falánga ya Rwanda", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dolarɛ ya Zimbabwɛ", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dolarɛ ya Ositali", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Esikudo ya Kapevɛrɛ", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Falánga ya Komoro", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ugwiya ya Moritani", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupi ya Morisi", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dolarɛ ya Namibi", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupi ya Sɛshɛlɛ", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula ya Botswana", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuanɛ Renminbi ya Sinɛ", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra ya Sao Tomé mpé Presipe", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinarɛ ya Bahrɛnɛ", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Motolé ya Swédi", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Paunɛ ya Ezípitɛ", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Falánga ya Gine", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwasha ya Malawi", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metikali ya Mozambiki", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leonɛ", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilingɛ ya Tanzani", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza ya Angóla", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Motolé Sheki", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwasha ya Zambi", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso ya Kuba", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa ya Elitlɛ", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirihamɛ ya Lémila alabo", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso ya Shili", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birɛ ya Etsiópi", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Ɛlɔ́", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gurde", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyalɛ ya Alabi Sawuditɛ", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Falánga ya Burundi", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinarɛ ya Alizeri", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwasha ya Zambi (1968–2012)", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilingɛ ya Kenya", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilingɛ ya Somali", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Falánga ya Kongó", Symbol: "FC"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso Dominikani", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Falánga ya Madagasikarɛ", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Dinarɛ ya Sudá", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Paunɛ ya Sudá", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilingɛ ya Uganda", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real ya Brazil", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dolarɛ ya Kanadá", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Motolé ya Islandi", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti ya Lesóto", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira ya Nizerya", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Motolé ya Danemark", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Paunɛ ya Angɛlɛtɛ́lɛ", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Falánga CFA BEAC", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Sedi ya Gana", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yeni ya Zapɔ", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi ya Gambi", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dolarɛ ya Ameriki", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Randɛ ya Afríka Súdi", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso y’Argentina", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Falánga ya Dzibuti", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Sol Sika", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirihame ya Marokɛ", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Paunɛ ya Sántu elena", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupi ya Índɛ", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dolarɛ ya Liberya", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Falánga ya Swisɛ", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinarɛ ya Tinizi", Symbol: ""}}
