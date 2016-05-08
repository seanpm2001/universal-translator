package tr

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"GWE": ut.Currency{Currency: "GWE", DisplayName: "Portekiz Ginesi Esküdosu", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rus Rublesi (1991–1998)", Symbol: "RUR"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Birleşik Avrupa Birimi", Symbol: ""}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ADB Hesap Birimi", Symbol: "XUA"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danimarka Kronu", Symbol: "DKK"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Yunan Drahmisi", Symbol: "GRD"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuveyt Dinarı", Symbol: "KWD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Lübnan Lirası", Symbol: "LBP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberya Doları", Symbol: "LRD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Sırp Dinarı", Symbol: "RSD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei Doları", Symbol: "BND"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Alman Markı", Symbol: "DEM"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norveç Kronu", Symbol: "NOK"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapur Doları", Symbol: "SGD"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "El Salvador Kolonu", Symbol: "SVC"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad ve Tobago Doları", Symbol: "TTD"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Yeni Angola Kvanzası (1990–2000)", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Avustralya Doları", Symbol: "AU$"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ekvador Unidad de Valor Constante (UVC)", Symbol: "ECV"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letonya Latı", Symbol: "LVL"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Eski Romen Leyi", Symbol: "ROL"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Güney Sudan Lirası", Symbol: "SSP"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba Florini", Symbol: "AWG"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Konvertibl Bosna Hersek Markı", Symbol: "BAM"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Yeni Beyaz Rusya Rublesi (1994–1999)", Symbol: "BYB"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Endonezya Rupiahı", Symbol: "IDR"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rus Rublesi", Symbol: "RUB"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brezilya Kruzadosu", Symbol: "BRC"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Doğu Alman Markı", Symbol: "DDM"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Cezayir Dinarı", Symbol: "DZD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Estonya Krunu", Symbol: "EEK"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litvanya Talonu", Symbol: "LTT"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Solomon Adaları Doları", Symbol: "SBD"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanada Doları", Symbol: "CA$"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Çek Cumhuriyeti Korunası", Symbol: "CZK"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrayna Karbovanetz", Symbol: "UAK"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnam Dongu", Symbol: "₫"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutan Ngultrumu", Symbol: "BTN"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Moritanya Ouguiyası", Symbol: "MRO"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldiv Rufiyaası", Symbol: "MVR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panama Balboası", Symbol: "PAB"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Bilinmeyen Para Birimi", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Çin Halk Cumhuriyeti Merkez Bankası Doları", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritius Rupisi", Symbol: "MUR"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistan Afganisi", Symbol: "AFN"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Arnavutluk Leki", Symbol: "ALL"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldova Leyi", Symbol: "MDL"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Meksika Unidad de Inversion (UDI)", Symbol: "MXV"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "İrlanda Lirası", Symbol: "IEP"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaskar Frangı", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedonya Dinarı", Symbol: "MKD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguay Peso en Unidades Indexadas", Symbol: "UYI"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo Frangı", Symbol: "CDF"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Fas Frangı", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Moğolistan Tugriki", Symbol: "MNT"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Uganda Şilini (1966–1987)", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET Fonları", Symbol: "XRE"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabve Doları (2008)", Symbol: "ZWR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbaycan Manatı", Symbol: "AZN"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kamboçya Rieli", Symbol: "KHR"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Makedonya Dinarı (1992–1993)", Symbol: "MKN"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Makao Patakası", Symbol: "MOP"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinam Doları", Symbol: "SRD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu Vatusu", Symbol: "VUV"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Güney Kore Wonu", Symbol: "₩"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibya Doları", Symbol: "NAD"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Polonya Zlotisi", Symbol: "PLN"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Suriye Lirası", Symbol: "SYP"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tacikistan Rublesi", Symbol: "TJR"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bulgar Levası (Hard)", Symbol: "BGL"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR Avrosu", Symbol: "CHE"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Eski Sırbistan Dinarı", Symbol: "CSD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Mali Frangı", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA Frangı BEAC", Symbol: "FCFA"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Gümüş", Symbol: "XAG"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Çekoslavak Korunası (Hard)", Symbol: "CSK"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzanya Şilini", Symbol: "TZS"}, "USS": ut.Currency{Currency: "USS", DisplayName: "ABD Doları (Aynı gün)", Symbol: "USS"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP Frangı", Symbol: "CFPF"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Kosta Rika Kolonu", Symbol: "CRC"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seyşeller Rupisi", Symbol: "SCR"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguay Pezosu (1975–1993)", Symbol: "UYP"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoa Talası", Symbol: "WST"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Yeni Yugoslav Dinarı", Symbol: "YUM"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libya Dinarı", Symbol: "LYD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua Yeni Gine Kinası", Symbol: "PGK"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistan Rupisi", Symbol: "PKR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguay Guaranisi", Symbol: "PYG"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Arjantin Peso Leyi (1970–1983)", Symbol: "ARL"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Arjantin Pesosu (1881–1970)", Symbol: "ARM"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Konvertibl Küba Pesosu", Symbol: "CUC"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Gine Frangı", Symbol: "GNF"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Fas Dirhemi", Symbol: "MAD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Türkmenistan Manatı", Symbol: "TMT"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burma Kyatı", Symbol: "BUK"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominik Pesosu", Symbol: "DOP"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Hindistan Rupisi", Symbol: "₹"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "İtalyan Lireti", Symbol: "ITL"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika Doları", Symbol: "JMD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland Adaları Lirası", Symbol: "FKP"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Özel Çekme Hakkı (SDR)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Test Para Birimi Kodu", Symbol: "XTS"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Yemen Riyali", Symbol: "YER"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabve Doları (2009)", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haiti Gurdu", Symbol: "HTG"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Slovak Korunası", Symbol: "SKK"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somali Şilini", Symbol: "SOS"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonga Paʻangası", Symbol: "TOP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda Şilini", Symbol: "UGX"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Avusturya Şilini", Symbol: "ATS"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tacikistan Somonisi", Symbol: "TJS"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Türkmenistan Manatı (1993–2009)", Symbol: "TMM"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiya Kvaçası", Symbol: "ZMW"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Fransız Frangı", Symbol: "FRF"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Yeni İsrail Şekeli", Symbol: "₪"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambik Esküdosu", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreyn Dinarı", Symbol: "BHD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahama Doları", Symbol: "BSD"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ekvador Sukresi", Symbol: "ECS"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Yeni Zelanda Doları", Symbol: "NZ$"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Surinam Guldeni", Symbol: "SRG"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Bulgar Levası (1879–1952)", Symbol: "BGO"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Fin Markkası", Symbol: "FIM"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platin", Symbol: "XPT"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botsvana Pulası", Symbol: "BWP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Gine Sylisi", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekvator Ginesi Ekuelesi", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laos Kipi", Symbol: "LAK"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timor Esküdosu", Symbol: "TPE"}, "USN": ut.Currency{Currency: "USN", DisplayName: "ABD Doları (Ertesi gün)", Symbol: "USN"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Arjantin Pesosu", Symbol: "ARS"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Azerbaycan Manatı (1993–2006)", Symbol: "AZM"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivya Bolivyanosu", Symbol: "BOB"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguay Pesosu", Symbol: "UYU"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Umman Riyali", Symbol: "OMR"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belçika Frangı (finansal)", Symbol: "BEL"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolivya Mvdolu", Symbol: "BOV"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Macar Forinti", Symbol: "HUF"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "İzlanda Kronu", Symbol: "ISK"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Ürdün Dinarı", Symbol: "JOD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepal Rupisi", Symbol: "NPR"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Rodezya Doları", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosna Hersek Dinarı", Symbol: "BAD"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Brezilya Kruzeirosu (1942–1967)", Symbol: "BRZ"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Cibuti Frangı", Symbol: "DJF"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Gana Sedisi", Symbol: "GHS"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Peru Solu", Symbol: "PES"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Avrupa Hesap Birimi (XBD)", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Yeni Bosna Hersek Dinarı (1994–1997)", Symbol: "BAN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belize Doları", Symbol: "BZD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemala Quetzalı", Symbol: "GTQ"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Lüksemburg Frangı", Symbol: "LUF"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Eski Sudan Dinarı", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuela Bolivarı (1871–2008)", Symbol: "VEB"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Moldova Kuponu", Symbol: "MDC"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabve Doları", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong Doları", Symbol: "HK$"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komorlar Frangı", Symbol: "KMF"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Monako Frangı", Symbol: "MCF"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malezya Ringgiti", Symbol: "MYR"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Fransız UIC-Frangı", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Yeni Brezilya Kruzeirosu (1967–1986)", Symbol: "BRB"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Hırvatistan Kunası", Symbol: "HRK"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "İsveç Kronu", Symbol: "SEK"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolombiya Pesosu", Symbol: "COP"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Gürcistan Kupon Larisi", Symbol: "GEK"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irak Dinarı", Symbol: "IQD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leone Leonesi", Symbol: "SLL"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Arjantin Australi", Symbol: "ARA"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "İsviçre Frangı", Symbol: "CHF"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Cebelitarık Lirası", Symbol: "GIP"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Cayman Adaları Doları", Symbol: "KYD"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda Frangı", Symbol: "RWF"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrayna Grivnası", Symbol: "UAH"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiya Kvaçası (1968–2012)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Arnavutluk Leki (1946–1965)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brezilya Kruzeirosu (1990–1993)", Symbol: "BRE"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verde Esküdosu", Symbol: "CVE"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaragua Kordobası", Symbol: "NIO"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svaziland Lilangenisi", Symbol: "SZL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Doğu Karayip Doları", Symbol: "EC$"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Güney Kore Wonu (1945–1953)", Symbol: "KRO"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Polonya Zlotisi (1950–1995)", Symbol: "PLZ"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Yeni Tayvan Doları", Symbol: "NT$"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Yemen Dinarı", Symbol: "YDD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belçika Frangı", Symbol: "BEF"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda Doları", Symbol: "BMD"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "İsrail Şekeli (1980–1985)", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Eski Mozambik Metikali", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tayland Bahtı", Symbol: "฿"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Güney Afrika Randı (finansal)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Birleşik Arap Emirlikleri Dirhemi", Symbol: "AED"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angola Kvanzası (1977–1990)", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angola Kvanzası Reajustado (1995–1999)", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Beyaz Rusya Rublesi", Symbol: "BYR"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiyopya Birri", Symbol: "ETB"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Ermenistan Dramı", Symbol: "AMD"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Güney Kıbrıs Lirası", Symbol: "CYP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipinler Pesosu", Symbol: "PHP"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Türk Lirası", Symbol: "₺"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Yugoslav Dinarı (Hard)", Symbol: "YUD"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Güney Afrika Randı", Symbol: "ZAR"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belçika Frangı (konvertibl)", Symbol: "BEC"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi Frangı", Symbol: "BIF"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Çin Yuanı", Symbol: "CN¥"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Gine-Bissau Pezosu", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kırgızistan Somu", Symbol: "KGS"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Malta Sterlini", Symbol: "MTP"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Paladyum", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afganistan Afganisi (1927–2002)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Arjantin Pezosu (1983–1985)", Symbol: "ARP"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brezilya Kruzeirosu", Symbol: "BRR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambik Metikali", Symbol: "MZN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Suudi Arabistan Riyali", Symbol: "SAR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA Frangı BCEAO", Symbol: "CFA"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyana Doları", Symbol: "GYD"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Vietnam Dongu (1978–1985)", Symbol: "VNN"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Yeni Zaire Zairesi", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "İsrail Lirası", Symbol: "ILP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenya Şilini", Symbol: "KES"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Romen Leyi", Symbol: "RON"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezuela Bolivarı", Symbol: "VEF"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Avrupa Para Birimi (EMU)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brezilya Reali", Symbol: "R$"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Şili Pesosu", Symbol: "CLP"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Güney Kore Hwanı (1953–1962)", Symbol: "KRH"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Özbekistan Somu", Symbol: "UZS"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritre Nakfası", Symbol: "ERN"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduras Lempirası", Symbol: "HNL"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Eski Türk Lirası", Symbol: "TRL"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR Frangı", Symbol: "CHW"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Şili Esküdosu", Symbol: "CLE"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Letonya Rublesi", Symbol: "LVR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudan Lirası", Symbol: "SDG"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé ve Príncipe Dobrası", Symbol: "STD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Fransız Altın Frangı", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Bolivya Bolivyanosu (1863–1963)", Symbol: "BOL"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Yeni Brezilya Kruzadosu", Symbol: "BRN"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "İspanyol Pezetası (konvertibl hesap)", Symbol: "ESB"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar Kyatı", Symbol: "MMK"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Gümüş Meksika Pezosu (1861–1992)", Symbol: "MXP"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Saint Helena Lirası", Symbol: "SHP"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: "COU"}, "USD": ut.Currency{Currency: "USD", DisplayName: "ABD Doları", Symbol: "$"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaire Zairesi", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peru Nuevo Solü", Symbol: "PEN"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Hollanda Antilleri Guldeni", Symbol: "ANG"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "İspanyol Pezetası (A hesabı)", Symbol: "ESA"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "İspanyol Pezetası", Symbol: "ESP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lanka Rupisi", Symbol: "LKR"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Konvertibl Lüksemburg Frangı", Symbol: "LUC"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Malta Lirası", Symbol: "MTL"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "İzlanda Kronu (1918–1981)", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Kuzey Kore Wonu", Symbol: "KPW"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Slovenya Toları", Symbol: "SIT"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Avrupa Para Birimi", Symbol: "XEU"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angola Kvanzası", Symbol: "AOA"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Şili Unidades de Fomento", Symbol: "CLF"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Hırvatistan Dinarı", Symbol: "HRD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japon Yeni", Symbol: "¥"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nijerya Nairası", Symbol: "NGN"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nikaragua Kordobası (1988–1991)", Symbol: "NIC"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Sovyet Rublesi", Symbol: "SUR"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgar Levası", Symbol: "BGN"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji Doları", Symbol: "FJD"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiya Dalasisi", Symbol: "GMD"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "İran Riyali", Symbol: "IRR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litvanya Litası", Symbol: "LTL"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskar Ariarisi", Symbol: "MGA"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorra Pezetası", Symbol: "ADP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Gana Sedisi (1979–2007)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Finansal Lüksemburg Frangı", Symbol: "LUL"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Avrupa Hesap Birimi (XBC)", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakistan Tengesi", Symbol: "KZT"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Maldiv Rupisi", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Hollanda Florini", Symbol: "NLG"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Konvertibl Yugoslav Dinarı", Symbol: "YUN"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Sosyalist Bulgaristan Levası", Symbol: "BGM"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Küba Pesosu", Symbol: "CUP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksika Pesosu", Symbol: "MX$"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunus Dinarı", Symbol: "TND"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Gürcistan Larisi", Symbol: "GEL"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portekiz Esküdosu", Symbol: "PTE"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar Riyali", Symbol: "QAR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Eski Sudan Lirası", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "İyileştirilmiş Yugoslav Dinarı (1992–1993)", Symbol: "YUR"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados Doları", Symbol: "BBD"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolivya Pezosu", Symbol: "BOP"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malavi Kvaçası", Symbol: "MWK"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Peru İnti", Symbol: "PEI"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Altın", Symbol: "XAU"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladeş Takası", Symbol: "BDT"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Mısır Lirası", Symbol: "EGP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "İngiliz Sterlini", Symbol: "£"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho Lotisi", Symbol: ""}}
