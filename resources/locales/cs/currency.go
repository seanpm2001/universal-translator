package cs

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"PTE": ut.Currency{Currency: "PTE", DisplayName: "portugalské escudo", Symbol: "PTE"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "uruguayské peso (v indexovaných jednotkách)", Symbol: "UYI"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "islandská koruna (1918–1981)", Symbol: "ISJ"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "turecká lira", Symbol: "TRY"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "nizozemskoantilský gulden", Symbol: "ANG"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "běloruský rubl", Symbol: "BYR"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "srbský dinár (2002–2006)", Symbol: "CSD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "hongkongský dolar", Symbol: "HK$"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "malijský frank", Symbol: "MLF"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "myanmarský kyat", Symbol: "MMK"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ruský rubl", Symbol: "RUB"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "belgický finanční frank", Symbol: "BEL"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "chilské escudo", Symbol: "CLE"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "litevský talonas", Symbol: "LTT"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "neznámá měna", Symbol: "XXX"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "chilské peso", Symbol: "CLP"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lotyšský lat", Symbol: "LVL"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mauritánská ouguiya", Symbol: "MRO"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "somálský šilink", Symbol: "SOS"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "sovětský rubl", Symbol: "SUR"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "turkmenský manat (1993–2009)", Symbol: "TMM"}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemenský rijál", Symbol: "YER"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "bolivijský mvdol", Symbol: "BOV"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "tchajwanský dolar", Symbol: "NT$"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "argentinský austral", Symbol: "ARA"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "brazilské cruzeiro (1993–1994)", Symbol: "BRR"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltarská libra", Symbol: "GIP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "americký dolar", Symbol: "US$"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatský vatu", Symbol: "VUV"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "svatohelenská libra", Symbol: "SHP"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albánský lek", Symbol: "ALL"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubský zlatý", Symbol: "AWG"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuvajtský dinár", Symbol: "KWD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mauricijská rupie", Symbol: "MUR"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "maledivská rupie (1947–1981)", Symbol: "MVP"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawijská kwacha", Symbol: "MWK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norská koruna", Symbol: "NOK"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "jihoafrický rand", Symbol: "ZAR"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "španělská peseta (konvertibilní účet)", Symbol: "ESB"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldavský leu", Symbol: "MDL"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "nikaragujská córdoba (1988–1991)", Symbol: "NIC"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "uruguayské peso (1975–1993)", Symbol: "UYP"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "zimbabwský dolar (2009)", Symbol: "ZWL"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "východoněmecká marka", Symbol: "DDM"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "maltská lira", Symbol: "MTL"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "svazijský lilangeni", Symbol: "SZL"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "evropská jednotka účtu 9 (XBC)", Symbol: "XBC"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "chorvatská kuna", Symbol: "HRK"}, "AON": ut.Currency{Currency: "AON", DisplayName: "angolská kwanza (1990–2000)", Symbol: "AON"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "argentinské peso (1881–1970)", Symbol: "ARM"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "syrská libra", Symbol: "SYP"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "jemenský dinár", Symbol: "YDD"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "zimbabwský dolar (1980–2008)", Symbol: "ZWD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brazilský real", Symbol: "R$"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "lotyšský rubl", Symbol: "LVR"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malajsijský ringgit", Symbol: "MYR"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palladium", Symbol: "XPD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "andorrská peseta", Symbol: "ADP"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "monacký frank", Symbol: "MCF"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frank", Symbol: "CFPF"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haitský gourde", Symbol: "HTG"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "marocký dinár", Symbol: "MAD"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "běloruský rubl (1994–1999)", Symbol: "BYB"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belizský dolar", Symbol: "BZD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "filipínské peso", Symbol: "PHP"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "ázerbájdžánský manat (1993–2006)", Symbol: "AZM"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiopský birr", Symbol: "ETB"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "francouzský frank", Symbol: "FRF"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "gruzínské kuponové lari", Symbol: "GEK"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kyrgyzský som", Symbol: "KGS"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "turkmenský manat", Symbol: "TMT"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "bosenský nový dinár (1994–1997)", Symbol: "BAN"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "srílanská rupie", Symbol: "LKR"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "mexické stříbrné peso (1861–1992)", Symbol: "MXP"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "švýcarský frank", Symbol: "CHF"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "dánská koruna", Symbol: "DKK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falklandská libra", Symbol: "FKP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "jihokorejský won", Symbol: "₩"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "marocký frank", Symbol: "MAF"}, "USN": ut.Currency{Currency: "USN", DisplayName: "americký dolar (příští den)", Symbol: "USN"}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnamský dong", Symbol: "VND"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "argentinské peso ley (1970–1983)", Symbol: "ARL"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordánský dinár", Symbol: "JOD"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "madagaskarský frank", Symbol: "MGF"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "turecká lira (1922–2005)", Symbol: "TRL"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indická rupie", Symbol: "INR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "peruánský nový sol", Symbol: "PEN"}, "THB": ut.Currency{Currency: "THB", DisplayName: "thajský baht", Symbol: "THB"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudský dolar", Symbol: "BMD"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "brazilské cruzeiro (1942–1967)", Symbol: "BRZ"}, "COP": ut.Currency{Currency: "COP", DisplayName: "kolumbijské peso", Symbol: "COP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongolský tugrik", Symbol: "MNT"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "novozélandský dolar", Symbol: "NZ$"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "jihosúdánská libra", Symbol: "SSP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandský šilink", Symbol: "UGX"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadský dolar", Symbol: "CA$"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubánské konvertibilní peso", Symbol: "CUC"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "maďarský forint", Symbol: "HUF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "liberijský dolar", Symbol: "LRD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "mexické peso", Symbol: "MX$"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mozambický metical", Symbol: "MZN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "súdánská libra", Symbol: "SDG"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundský frank", Symbol: "BIF"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "španělská peseta", Symbol: "ESP"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "portugalskoguinejské escudo", Symbol: "GWE"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelský bolívar", Symbol: "VEF"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zairský zaire (1971–1993)", Symbol: "ZRZ"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laoský kip", Symbol: "LAK"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ománský rijál", Symbol: "OMR"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "katarský rijál", Symbol: "QAR"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "trinidadský dolar", Symbol: "TTD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambijská kwacha", Symbol: "ZMW"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "belgický frank", Symbol: "BEF"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolivijský boliviano", Symbol: "BOB"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "bolivijské peso", Symbol: "BOP"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "švýcarský WIR-frank", Symbol: "CHW"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "severokorejský won", Symbol: "KPW"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maledivská rupie", Symbol: "MVR"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "evropská smíšená jednotka", Symbol: "XBA"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afghánský afghán (1927–2002)", Symbol: "AFA"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "belgický konvertibilní frank", Symbol: "BEC"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "íránský rijál", Symbol: "IRR"}, "STD": ut.Currency{Currency: "STD", DisplayName: "svatotomášská dobra", Symbol: "STD"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "zlato", Symbol: "XAU"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "zimbabwský dolar (2008)", Symbol: "ZWR"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "bosenský dinár (1992–1994)", Symbol: "BAD"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "nizozemský gulden", Symbol: "NLG"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamská balboa", Symbol: "PAB"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "singapurský dolar", Symbol: "SGD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrajinská hřivna", Symbol: "UAH"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "argentinské peso (1983–1985)", Symbol: "ARP"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "moldavský kupon", Symbol: "MDC"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "saúdský rijál", Symbol: "SAR"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "madagaskarský ariary", Symbol: "MGA"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "sucre", Symbol: "XSU"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "australský dolar", Symbol: "AU$"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubánské peso", Symbol: "CUP"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "španělská peseta („A“ účet)", Symbol: "ESA"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "britská libra", Symbol: "£"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "izraelská libra", Symbol: "ILP"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "lesothský loti", Symbol: "LSL"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "jugoslávský konvertibilní dinár (1990–1992)", Symbol: "YUN"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonéská rupie", Symbol: "IDR"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "lucemburský finanční frank", Symbol: "LUL"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "macajská pataca", Symbol: "MOP"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "mosambický escudo", Symbol: "MZE"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namibijský dolar", Symbol: "NAD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "srbský dinár", Symbol: "RSD"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "kód zvlášť vyhrazený pro testovací účely", Symbol: "XTS"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bulharský lev", Symbol: "BGN"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "brazilské cruzado (1986–1989)", Symbol: "BRC"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "mosambický metical (1980–2006)", Symbol: "MZM"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodéský dolar", Symbol: "RHD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "konžský frank", Symbol: "CDF"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "kostarický colón", Symbol: "CRC"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "ruský rubl (1991–1998)", Symbol: "RUR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "rwandský frank", Symbol: "RWF"}, "TND": ut.Currency{Currency: "TND", DisplayName: "tuniský dinár", Symbol: "TND"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "řecká drachma", Symbol: "GRD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "izraelský nový šekel", Symbol: "ILS"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "seychelská rupie", Symbol: "SCR"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadoský dolar", Symbol: "BBD"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "italská lira", Symbol: "ITL"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "arménský dram", Symbol: "AMD"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "bulharský lev (1879–1952)", Symbol: "BGO"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamský dolar", Symbol: "BSD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswanská pula", Symbol: "BWP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "rovníkovoguinejský ekwele", Symbol: "GQE"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "surinamský zlatý", Symbol: "SRG"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tonžská paanga", Symbol: "TOP"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "jugoslávský reformovaný dinár (1992–1993)", Symbol: "YUR"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentinské peso", Symbol: "ARS"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "rakouský šilink", Symbol: "ATS"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ekvádorská jednotka konstantní hodnoty", Symbol: "ECV"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "estonská koruna", Symbol: "EEK"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "lucemburský konvertibilní frank", Symbol: "LUC"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "vietnamský dong (1978–1985)", Symbol: "VNN"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "francouzský zlatý frank", Symbol: "XFO"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "jihoafrický finanční rand", Symbol: "ZAL"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "guyanský dolar", Symbol: "GYD"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "německá marka", Symbol: "DEM"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "finská marka", Symbol: "FIM"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "komorský frank", Symbol: "KMF"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguajské guarani", Symbol: "PYG"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA/BCEAO frank", Symbol: "CFA"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "guinejský syli", Symbol: "GNS"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "chorvatský dinár", Symbol: "HRD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigerijská naira", Symbol: "NGN"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA/BEAC frank", Symbol: "FCFA"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "evropská jednotka účtu 17 (XBD)", Symbol: "XBD"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "", Symbol: "XUA"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "angolská kwanza (1995–1999)", Symbol: "AOR"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "švýcarské WIR-euro", Symbol: "CHE"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyptská libra", Symbol: "EGP"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "jihokorejský won (1945–1953)", Symbol: "KRO"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "polský zlotý (1950–1995)", Symbol: "PLZ"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "slovenská koruna", Symbol: "SKK"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "kód fondů RINET", Symbol: "XRE"}, "AED": ut.Currency{Currency: "AED", DisplayName: "SAE dirham", Symbol: "AED"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosenská konvertibilní marka", Symbol: "BAM"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodžský riel", Symbol: "KHR"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanonská libra", Symbol: "LBP"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "SDR", Symbol: "XDR"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kazašské tenge", Symbol: "KZT"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "makedonský denár (1992–1993)", Symbol: "MKN"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tádžické somoni", Symbol: "TJS"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "ghanský cedi (1979–2007)", Symbol: "GHC"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "islandská koruna", Symbol: "ISK"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "papuánská nová kina", Symbol: "PGK"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzanský šilink", Symbol: "TZS"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "uzbecký sum", Symbol: "UZS"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "venezuelský bolívar (1871–2008)", Symbol: "VEB"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fidžijský dolar", Symbol: "FJD"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "rumunské leu (1952–2006)", Symbol: "ROL"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "slovinský tolar", Symbol: "SIT"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "salvadorský colón", Symbol: "SVC"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "zambijská kwacha (1968–2012)", Symbol: "ZMK"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "zairský nový zaire (1993–1998)", Symbol: "ZRN"}, "KES": ut.Currency{Currency: "KES", DisplayName: "keňský šilink", Symbol: "KES"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "mexická investiční jednotka", Symbol: "MXV"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "čínský jüan", Symbol: "CN¥"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "alžírský dinár", Symbol: "DZD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pákistánská rupie", Symbol: "PKR"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "", Symbol: "BGM"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepálská rupie", Symbol: "NPR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "peruánský sol (1863–1965)", Symbol: "PES"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "ázerbájdžánský manat", Symbol: "AZN"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladéšská taka", Symbol: "BDT"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "guinejsko-bissauské peso", Symbol: "GWP"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "timorské escudo", Symbol: "TPE"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "chilská účetní jednotka (UF)", Symbol: "CLF"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "honduraská lempira", Symbol: "HNL"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "irácký dinár", Symbol: "IQD"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "východokaribský dolar", Symbol: "EC$"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolská kwanza", Symbol: "AOA"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalský quetzal", Symbol: "GTQ"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "súdánský dinár (1992–2007)", Symbol: "SDD"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "tádžický rubl", Symbol: "TJR"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "jugoslávský dinár (1966–1990)", Symbol: "YUD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "brunejský dolar", Symbol: "BND"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "barmský kyat", Symbol: "BUK"}, "COU": ut.Currency{Currency: "COU", DisplayName: "kolumbijská jednotka reálné hodnoty", Symbol: "COU"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "gruzínské lari", Symbol: "GEL"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "súdánská libra (1957–1998)", Symbol: "SDP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "evropská měnová jednotka", Symbol: "ECU"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "bolivijský boliviano (1863–1963)", Symbol: "BOL"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "izraelský šekel (1980–1985)", Symbol: "ILR"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "džibutský frank", Symbol: "DJF"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ukrajinský karbovanec", Symbol: "UAK"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamajský dolar", Symbol: "JMD"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "jihokorejský hwan (1953–1962)", Symbol: "KRH"}, "WST": ut.Currency{Currency: "WST", DisplayName: "samojská tala", Symbol: "WST"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "brazilské nové cruzeiro (1967–1986)", Symbol: "BRB"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "kyperská libra", Symbol: "CYP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "rumunské leu", Symbol: "RON"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "šalamounský dolar", Symbol: "SBD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "stříbro", Symbol: "XAG"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "francouzský UIC frank", Symbol: "XFU"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "jugoslávský nový dinár (1994–2002)", Symbol: "YUM"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "albánský lek (1946–1965)", Symbol: "ALK"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "brazilské nové cruzado (1989–1990)", Symbol: "BRN"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhútánský ngultrum", Symbol: "BTN"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghanský cedi", Symbol: "GHS"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "makedonský denár", Symbol: "MKD"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "peruánská inti", Symbol: "PEI"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "evropská peněžní jednotka", Symbol: "XBB"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japonský jen", Symbol: "JP¥"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afghánský afghán", Symbol: "AFN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "bahrajnský dinár", Symbol: "BHD"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "československá koruna", Symbol: "Kčs"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "kapverdské escudo", Symbol: "CVE"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikánské peso", Symbol: "DOP"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "guinejský frank", Symbol: "GNF"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "irská libra", Symbol: "IEP"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "kajmanský dolar", Symbol: "KYD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "libyjský dinár", Symbol: "LYD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "polský zlotý", Symbol: "PLN"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "sierro-leonský leone", Symbol: "SLL"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ugandský šilink (1966–1987)", Symbol: "UGS"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: "XPT"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "lucemburský frank", Symbol: "LUF"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "švédská koruna", Symbol: "SEK"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "česká koruna", Symbol: "Kč"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ekvádorský sucre", Symbol: "ECS"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritrejská nakfa", Symbol: "ERN"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambijský dalasi", Symbol: "GMD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litevský litas", Symbol: "LTL"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "maltská libra", Symbol: "MTP"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nikaragujská córdoba", Symbol: "NIO"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "angolská kwanza (1977–1991)", Symbol: "AOK"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "BGL", Symbol: "BGL"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "brazilské cruzeiro (1990–1993)", Symbol: "BRE"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "surinamský dolar", Symbol: "SRD"}, "USS": ut.Currency{Currency: "USS", DisplayName: "americký dolar (týž den)", Symbol: "USS"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguayské peso", Symbol: "UYU"}}
