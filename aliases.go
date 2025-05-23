package legalform

import (
	"strings"

	"github.com/tilotech/go-phonetics/diacrit"
)

// Aliases represents a data structure that defines country-specific
// aliases for legal forms.
//
// In most cases, you want to simply use the predefined DefaultAliases instance.
type Aliases map[string]map[string]string

// Find returns the alias for the provided legal form if available.
//
// The returned alias is a cleaned text, e.g. it will not contain special
// characters or white spaces.
//
// If no alias was found, then the cleaned version of the legal form is returned.
func (l Aliases) Find(country, legalForm string) string {
	legalForm = strings.ReplaceAll(clean(legalForm), " ", "")
	normalized := diacrit.Normalize(legalForm)
	country = strings.ToUpper(country)
	alias := l[country][normalized]
	if alias != "" {
		return alias
	}
	alias = l["*"][normalized]
	if alias != "" {
		return alias
	}
	return legalForm
}

// DefaultAliases is a list of country specific legal form aliases.
var DefaultAliases = Aliases{
	"*": map[string]string{
		"limited":              "ltd",
		"incorporated":         "inc",
		"corporation":          "corp",
		"company":              "co",
		"publiclimitedcompany": "plc",
		"pvtlimited":           "pvtltd",
		"privatelimited":       "pvtltd",
		"privateltd":           "pvtltd",
		"companylimited":       "coltd",
		"companyltd":           "coltd",
		"colimited":            "coltd",
	},
	"AT": map[string]string{
		"gesmbh":                             "gmbh",
		"stillegesellschaft":                 "stg",
		"gesellschaftmitbeschrankterhaftung": "gmbh",
		"offenehandelsgesellschaft":          "ohg",
		"europaischegesellschaft":            "se",
		"europaischegenossenschaft":          "sce",
		"kommanditerwerbsgesellschaft":       "keg",
		"aktiengesellschaft":                 "ag",
		"offenegesellschaft":                 "og",
		"offeneerwerbsgesellschaft":          "oeg",
		"europaischegesellschaftse":          "se",
		"europaischegenossenschaftsce":       "sce",
		"einzelunternehmer":                  "eu",
		"kommanditgesellschaft":              "kg",
		"gesellschaftdesburgerlichenrechts":  "gesbr",
	},
	"NO": map[string]string{
		"europeisksamvirkeforetak":         "sce",
		"ansvarligselskap":                 "ans",
		"kommunaltforetak":                 "kf",
		"samvirkeforetak":                  "sa",
		"kommandittselskap":                "ks",
		"norskregistrertutenlandskforetak": "nuf",
		"eofg":                             "eofg",
		"borettslag":                       "brl",
		"europeiskselskap":                 "se",
		"interkommunaltselskap":            "iks",
		"boligbyggelag":                    "bbl",
		"ansvarligselskapdeltansvar":       "da",
		"europeiskokonomiskforetaksgruppe": "eofg",
		"enkeltpersonforetak":              "enk",
		"allmennaksjeselskap":              "asa",
		"selskapmedbegrensetansvar":        "ba",
		"aksjeselskap":                     "as",
		"statsforetak":                     "sf",
		"fylkeskommunaltforetak":           "fkf",
	},
	"KY": map[string]string{
		"companylimitedbyguarantee":      "ltd",
		"limitedpartnership":             "lp",
		"limited":                        "ltd",
		"companylimitedbyshares":         "ltd",
		"exemptedlimiteddurationcompany": "ldc",
		"exemptedlimitedpartnership":     "elp",
		"specialeconomiczonecompany":     "sezc",
		"segregatedportfoliocompany":     "spc",
		"limitedliabilitycompany":        "llc",
	},
	"SE": map[string]string{
		"ideellforening":                       "if",
		"europabolag":                          "se",
		"sparbank":                             "sb",
		"europeiskekonomiskintressegruppering": "eeig",
		"trossamfund":                          "tsf",
		"ekonomiskforening":                    "ekfor",
		"kommanditbolag":                       "kb",
		"kooperativhyresrattsforening":         "khf",
		"medlemsbank":                          "mb",
		"europakooperativ":                     "sce",
		"sambruksforening":                     "sf",
		"bankaktiebolag":                       "bab",
		"forsakringsforening":                  "fof",
		"forsakringsaktiebolag":                "fab",
		"handelsbolag":                         "hb",
		"ekfor":                                "ekfor",
		"omsesidigtforsakringsbolag":           "ofb",
		"europeiskgrupperingforterritorielltsamarbete": "egts",
		"bostadsrattsforening":                         "brf",
		"aktiebolag":                                   "ab",
	},
	"CH": map[string]string{
		"partnershipgeneral":                     "snc",
		"societaagaranzialimitata":               "sarl",
		"societeencommandite":                    "sc",
		"gp":                                     "snc",
		"ag":                                     "sa",
		"sagl":                                   "sarl",
		"socencommandite":                        "sc",
		"societeennomcollectif":                  "snc",
		"societaanonima":                         "sa",
		"societainaccomandita":                   "sc",
		"socnomecollettivo":                      "snc",
		"companylimitedbyshares":                 "sa",
		"societeanonyme":                         "sa",
		"plc":                                    "sa",
		"societaagaranzialimitatasagl":           "sarl",
		"kommanditgesellsch":                     "sc",
		"kollektivgesellsch":                     "snc",
		"aktiengesellschaft":                     "sa",
		"limitedliabilitycompany":                "sarl",
		"limited":                                "sarl",
		"societearesponsabilitelimitee":          "sarl",
		"ltdpartnership":                         "sc",
		"societearesponsabilitelimiteesarl":      "sarl",
		"gesellschaftmitbeschrankterhaftung":     "sarl",
		"gmbh":                                   "sarl",
		"kollektivgesellschaft":                  "snc",
		"gesellschaftmitbeschrankterhaftunggmbh": "sarl",
		"limitedpartnership":                     "sc",
		"kommanditgesellschaft":                  "sc",
		"socinaccomandita":                       "sc",
		"generalpartnership":                     "snc",
		"societainnomecollettivo":                "snc",
		"socnomcollectif":                        "snc",
	},
	"RU": map[string]string{
		"обществосограниченноиответственностью": "ooo",
		"публичноеакционерноеобщество":          "pao",
		"закрытоеакционерноеобщество":           "ao",
		"открытоеакционерноеобщество":           "pao",
		"государственноеунитарноепредприятие":   "gup",
		"муниципальноеунитарноепредприятие":     "mup",
		"ооо":                 "ooo",
		"пао":                 "pao",
		"oao":                 "pao",
		"оао":                 "pao",
		"зао":                 "ao",
		"акционерноеобщество": "ao",
		"ао":                  "ao",
		"zao":                 "ao",
		"полноетоварищество":  "pt",
		"пт":                  "pt",
		"гуп":                 "gup",
		"муп":                 "mup",
		"товариществонавере":  "tv",
		"тв":                  "tv",

		// Already translated
		"limitedliabilitycompany":     "ooo",
		"ltdliabilityco":              "ooo",
		"limited":                     "ooo",
		"ltd":                         "ooo",
		"companylimited":              "ooo",
		"companyltd":                  "ooo",
		"colimited":                   "ooo",
		"coltd":                       "ooo",
		"llc":                         "ooo",
		"stateunitaryenterprise":      "gup",
		"generalpartnership":          "pt",
		"limitedpartnership":          "kt",
		"municipalunitaryenterprise":  "mup",
		"jsc":                         "ao",
		"cjsc":                        "ao", // replaced by jsc
		"pjsc":                        "pao",
		"ojsc":                        "pao", // replaced by pjsc
		"gp":                          "pt",
		"lp":                          "tv",
		"productioncooperative":       "pk",
		"businesspartnership":         "khp",
		"limitedliabilitypartnership": "too",
		"llp":                         "too",
	},
	"SG": map[string]string{
		"limited":                     "ltd",
		"limitedpartnership":          "lp",
		"limitedliabilitypartnership": "llp",
		"pteltd":                      "ltd",
	},
	"CR": map[string]string{
		"sociedadcivil": "ac",
		"sociedadderesponsabilidadlimitadaosociedadlimitada": "sl",
		"srl":                 "sl",
		"sociedadcolectiva":   "sc",
		"sociedadanonima":     "sa",
		"sociedadencomandita": "senc",
		"empresaindividualderesponsabilidadlimitada": "eirl",
	},
	"PH": map[string]string{
		"stockcorporation":     "inc",
		"incorporated":         "inc",
		"onepersoncorporation": "opc",
	},
	"RO": map[string]string{
		"societateincomanditapeactiuni":   "sca",
		"grupuleuropeandeintereseconomic": "geie",
		"organizatienonguvernamentala":    "ong",
		"societateinnumecolectiv":         "snc",
		"asociatiefamiliala":              "af",
		"societateincomanditasimpla":      "scs",
		"societateeuropeana":              "se",
		"regieautonoma":                   "ra",
		"grupuldeintereseconomic":         "gie",
		"societatecooperativaeuropeana":   "sce",
		"societatecuraspunderelimitata":   "srl",
		"societatepeactiuni":              "sa",
		"intreprindereindividuala":        "ii",
		"persoanafizicaindependenta":      "pfi",
		"persoanafizicaautorizata":        "pfa",
	},
	"GR": map[string]string{
		"somatiaidrimata":             "clf",
		"anonimieteria":               "sa",
		"naftikieteria":               "sc",
		"εvropaikiεtaireia":           "se",
		"astikiproswpikieteria":       "sp",
		"nomikaprosopadimosioudikeou": "pl",
		"kinopraxia":                  "jv",
		"omorithmieteria":             "oe",
		"eterorithmieteria":           "ee",
		"eteriaperiorismeniseuthinis": "epe",
		"sinetairismos":               "coop",
		"afaniseteria":                "slp",
		"idiotikikefaleouhikieteria":  "ike",
		"simplioktisia":               "jso",
	},
	"UK": map[string]string{
		"cyf":                                   "ltd",
		"unitedkingdomeconomicinterestgrouping": "ukeig",
		"privatelimitedbyguarantee":             "ltd",
		"limited":                               "ltd",
		"ccc":                                   "plc",
		"privatelimitedcompany":                 "ltd",
		"cwmnicyfyngedigcyhoeddus":              "plc",
		"publiclimitedcompany":                  "plc",
		"societaseuropaea":                      "se",
		"cooperativesociety":                    "coop",
		"communitybenefitsociety":               "bencom",
		"europeaneconomicinterestgrouping":      "eeig",
		"partneriaethatebolrwyddcyfyngedig":     "llp",
		"pac":                                   "llp",
		"communityinterestcompany":              "cic",
		"unitedkingdomsocietas":                 "uks",
		"cyfyngedig":                            "ltd",
		"limitedpartnership":                    "lp",
		"limitedliabilitypartnership":           "llp",
	},
	"IT": map[string]string{
		"societacooperativa":                         "sc",
		"societainnomecollettivo":                    "snc",
		"societaeuropea":                             "se",
		"gruppoeuropeodiinteresseeconomico":          "geie",
		"srlsemplificata":                            "srls",
		"societacooperativaeuropea":                  "sce",
		"scrl":                                       "scarl",
		"societaaresponsabilitalimitata":             "srl",
		"societaperazioni":                           "spa",
		"societasemplice":                            "ss",
		"societaaresponsabilitalimitatasemplificata": "srls",
		"societaconsortilearesponsabilitalimitata":   "scarl",
	},
	"LU": map[string]string{
		"societeencommanditesimple":               "secs",
		"societedepargnepensionacapitalvariable":  "sepcav",
		"societearesponsabilitelimitee":           "sarl",
		"societearesponsabilitelimiteesimplifiee": "sarls",
		"societecivile":                           "sci",
		"groupementeuropeendintereteconomique":    "geie",
		"societedinvestissementacapitalfixe":      "sicaf",
		"fondation":                               "fon",
		"entrepriseindividuelle":                  "ei",
		"associationdassurancesmutuelles":         "aam",
		"societeparactionssimplifiee":             "sas",
		"associationdepargnepension":              "assep",
		"associationmomentanee":                   "am",
		"sarls":                                   "sarls",
		"societeencommanditespeciale":             "secsp",
		"etablissementpublic":                     "ep",
		"associationagricole":                     "aa",
		"fondcommundeplacement":                   "fcp",
		"groupementdintereteconomique":            "gie",
		"societecooperative":                      "sc",
		"societeeuropeenne":                       "se",
		"societeencommanditeparactions":           "seca",
		"societedinvestissementacapitalvariable":  "sicav",
		"associationsansbutlucratif":              "asbl",
		"societeanonyme":                          "sa",
		"societeennomcollectif":                   "senc",
		"societascooperativaeuropaea":             "sce",
	},
	"HK": map[string]string{
		"limited":            "ltd",
		"unlimited":          "ultd",
		"unltd":              "ultd",
		"limitedpartnership": "lp",
	},
	"AE": map[string]string{
		"rakeconomiczonecompany":   "fzllc",
		"soleproprietorship":       "sp",
		"civilcompany":             "cc",
		"limitedliabilitycompany":  "llc",
		"generalpartnership":       "gp",
		"limitedpartnership":       "lp",
		"publicjointstockcompany":  "pjsc",
		"privatejointstockcompany": "prjsc",
		"branchofaforeigncompany":  "bf",
		"representativeoffice":     "bro",
		"freezoneestablishment":    "fze",
		"freezonecompany":          "fzc",
	},
	"VG": map[string]string{
		"incorporated":       "inc",
		"unlimited":          "unltd",
		"corporation":        "corp",
		"limitedpartnership": "lp",
		"unlimitedcompanythatisauthorisedtoissueshares": "unltd",
		"segregatedportfoliocompany":                    "spc",
		"limited":                                       "ltd",
	},
	"MY": map[string]string{
		"publiccompany":               "bhd",
		"berhad":                      "bhd",
		"limitedliabilitypartnership": "llp",
		"privatecompany":              "sdnbhd",
		"pteltd":                      "sdnbhd",
		"sendirianberhad":             "sdnbhd",
		"limited":                     "ltd",
	},
	"EC": map[string]string{
		"companiaennombrecolectivo":                     "cncol",
		"companiaencomanditasimpleydivididaporacciones": "ccsda",
		"companiaanonima":                               "sa",
		"companiadeeconomiamixta":                       "cemix",
		"companiaderesponsabilidadlimitada":             "cialtda",
	},
	"DO": map[string]string{
		"sociedadencomanditasimple":                  "senc",
		"sociedadanonimasimplificada":                "sas",
		"empresaindividualderesponsabilidadlimitada": "eirl",
		"sociedadderesponsaibilidadlimitada":         "srl",
		"sociedadanonima":                            "sa",
	},
	"US": map[string]string{
		"rlllp":                                         "lllp",
		"mutualbenefitenterprise":                       "mbe",
		"cooperative":                                   "coop",
		"professionalassociation":                       "pa",
		"federalsavingsassociation":                     "fsa",
		"limitedliabilitycompany":                       "llc",
		"cooperativecorporation":                        "coop",
		"generalcooperativeassociation":                 "coop",
		"businesstrust":                                 "bt",
		"communitydevelopmentcorporation":               "cdc",
		"realestateinvestmenttrusts":                    "reit",
		"limited":                                       "ltd",
		"limitedliabilitylimitedpartnership":            "lllp",
		"publicbenefitcorporationprofit":                "pbc",
		"professionallimitedliabilitycompany":           "pllc",
		"corporation":                                   "corp",
		"statutorypublicbenefitlimitedliabilitycompany": "spbllc",
		"generalpartnership":                            "gp",
		"rllp":                                          "llp",
		"professionalcorporation":                       "pc",
		"limitedcooperativeassociation":                 "lca",
		"registeredlimitedliabilitypartnership":         "llp",
		"limitedpartnership":                            "lp",
		"limitedliabilitypartnership":                   "llp",
		"savingsloanassociation":                        "sl",
		"registeredlimitedliabilitylimitedpartnership":  "lllp",
		"incorporated":                                  "inc",
		"forprofitcorporationprofessional":              "pc",
		"commercialregisteredagent":                     "cra",
		"statutoryfoundation":                           "sf",
	},
	"ES": map[string]string{
		"sociedadanonimaprofesional":                   "sap",
		"sociedadcooperativaprofesional":               "scoopp",
		"fondodepensiones":                             "fp",
		"sociedadcolectiva":                            "sc",
		"sociedadcomanditariasimple":                   "senc",
		"scom":                                         "senc",
		"srll":                                         "sll",
		"sociedadesdeinversioninmobiliaria":            "sii",
		"fondodecapitalriesgopyme":                     "fcrpyme",
		"agrupaciondeintereseconomico":                 "aie",
		"sociedadlimitadaprofesional":                  "slp",
		"instituciondeinversioncolectiva":              "iic",
		"sociedadlimitadaunipersonal":                  "slu",
		"src":                                          "sc",
		"sociedadcomanditariasimpleprofesional":        "sencp",
		"mutualidaddeprevisionsocial":                  "mps",
		"sociedaddecapitalriesgo":                      "scr",
		"sociedaddeinversiondecapitalvariable":         "sicav",
		"sociedadanonimaeuropea":                       "se",
		"sociedadderesponsabilidadlimitada":            "sl",
		"fondodeinversionenactivosdelmercadomonetario": "fiamm",
		"sociedadcomanditariaporacciones":              "scompa",
		"sociedadcooperativaeuropea":                   "sce",
		"sociedadagrariadetrasformacion":               "sat",
		"sociedadanonima":                              "sa",
		"sociedadlimitadadenuevaempresa":               "slne",
		"emprendedorderesponsabilidadlimitada":         "erl",
		"sociedadcomanditariaporaccionesprofesional":   "scompap",
		"scoope":                                   "sce",
		"srcp":                                     "scp",
		"agrupacioneuropeadeintereseconomico":      "aeie",
		"sociedaddecapitalriesgopyme":              "scrpyme",
		"srlp":                                     "slp",
		"fondosdeinversioninmobiliaria":            "fii",
		"fondodeinversionmobiliaria":               "fim",
		"sociedadcooperativa":                      "scoop",
		"sociedadcolectivaprofesional":             "scp",
		"sociedaddegarantiareciproca":              "sgr",
		"sociedadlimitadalaboral":                  "sll",
		"sociedadderesponsabilidadlimitadalaboral": "sll",
		"fondodecapitalriesgo":                     "fcr",
		"sociedadanonimalaboral":                   "sal",
		"fondodeinversion":                         "fi",
		"scomp":                                    "sencp",
		"sociedadregularcolectivaprofesional":      "scp",
		"srl":                                      "sl",
	},
	"NL": map[string]string{
		"naamlozevennootschap":      "nv",
		"commanditairevennootschap": "cv",
		"cooperatie":                "coop",
		"beslotenvennootschapmetbeperkteaansprakelijkheid": "bv",
		"europeeseconomischsamenwerkingsverband":           "eesv",
		"vennootschaponderfirma":                           "vof",
		"europesenaamlozevennootschapsocietaseuropaea":     "se",
		"onderlingewaarborgmaatschappij":                   "owm",
		"rechtspersooninoprichting":                        "rpio",
		"verenigingvaneigenaars":                           "vve",
		"coop":                                             "coop",
		"societascooperativaeuropaea":                      "sce",
	},
	"MX": map[string]string{
		"sociedadanonimapromotoradeinversiondecapitalvariable":  "sapidecv",
		"sociedadcooperativadeahorroyprestamodecapitalvariable": "scdeapdecv",
		"sociedadcivil":                                                "sc",
		"sociedadanonimabursatil":                                      "sab",
		"sociedadporaccionessimplificada":                              "sas",
		"instituciondeasistenciaprivada":                               "iap",
		"sociedadderesponsabilidadlimitada":                            "sderl",
		"sociedadnacionaldecredito":                                    "snc",
		"sociedaddeproduccionruralderesponsabilidadlimitada":           "spr",
		"sociedadanonimapromotoradeinversionbursatildecapitalvariable": "sapibdecv",
		"sociedadcooperativalimitada":                                  "scl",
		"empresaproductivadelestado":                                   "epe",
		"sociedadderesponsabilidadlimitadadecapitalvariable":           "sderldecv",
		"sociedadanonimapromotoradeinversiondecapitalvariablesociedadfinancieradeobjetomultipleentidadnoregulada": "sapidecv",
		"sociedaddeproduccionruralderesponsabilidadlimitadadecapitalvariable":                                     "sprderldecv",
		"sociedadencomanditaporacciones":                                     "sencpora",
		"sociedadanonimadecapitalvariable":                                   "sadecv",
		"agrupacionesfinancieras":                                            "af",
		"sociedadanonima":                                                    "sa",
		"unidadagricolaindustrialdelamujer":                                  "uaim",
		"sociedadderesponsabilidadlimitadamicroindustrial":                   "sderlmi",
		"sociedadcooperativaderesponsabilidadlimitadadecapitalvariable":      "scderldecv",
		"sociedadanonimapromotoradeinversionbursatil":                        "sapib",
		"sociedadencomanditaporaccionesdecapitalvariable":                    "sencporadecv",
		"sociedadporaccionessimplificadadecapitalvariable":                   "sasdecv",
		"sociedadcooperativadeahorroyprestamo":                               "scdeap",
		"sociedadcooperativalimitadadecapitalvariable":                       "scldecv",
		"sociedadderesponsabilidadlimitadadeinterespublicodecapitalvariable": "sderldeipdecv",
		"sociedadanonimabursatildecapitalvariable":                           "sabdecv",
		"sociedadencomanditasimple":                                          "senc",
		"asociacioncivil":                                                    "ac",
		"sociedadencomanditasimpledecapitalvariable":                         "sencdecv",
		"sociedaddesolidaridadsocial":                                        "sdess",
		"sociedadanonimapromotoradeinversion":                                "sapi",
	},
	"TR": map[string]string{
		"komsti":          "komsti",
		"limitedsirket":   "ltdsti",
		"ltdsti":          "ltdsti",
		"anonimsirket":    "as",
		"as":              "as",
		"kollektifsirket": "kollsti",
		"kollsti":         "kollsti",
		"komanditsirket":  "komsti",
	},
	"BG": map[string]string{
		"ednolichentargovets": "et",
		"sie":                 "sd",
		"druzhestvouchredenopozakonazazadalzheniyataidogovorite": "dzzd",
		"sabiratelnodruzhestvo":                                  "sd",
		"sadruzhie":                                              "sd",
		"aktsionernodruzhestvosasspetsialnainvestitsionnatsel":   "adsic",
		"evropeyskoobedineniepoikonomicheskiinteresi":            "eeig",
		"komanditnodruzhestvo":                                   "kd",
		"evropeyskokooperativnodruzhestvo":                       "ekd",
		"aktsionernodruzhestvo":                                  "ad",
		"druzhestvosogranichenaotgovornost":                      "ood",
		"ednolichnodruzhestvosogranichenaotgovornost":            "eood",
		"ednolichnoaktsionernodruzhestvo":                        "ead",
		"komanditnodruzhestvosaktsii":                            "kda",
		"evropeyskodruzhestvo":                                   "ed",
	},
	"PK": map[string]string{
		"limitedliabilitypartnership": "llp",
		"privatelimitedcompany":       "pvtltd",
		"companylimitedbyguarantee":   "guaranteeltd",
		"unlimitedcompany":            "ultd",
		"unlimited":                   "ultd",
		"unltd":                       "ultd",
		"publiclimitedcompany":        "ltd",
		"limited":                     "ltd",
	},
	"FI": map[string]string{
		"saatio":                               "sr",
		"ab":                                   "oy",
		"eurooppalainentaloudellinenetuyhtyma": "etey",
		"publiktaktiebolag":                    "oyj",
		"europaandelslag":                      "sce",
		"eurooppayhtio":                        "se",
		"aktiebolag":                           "oy",
		"andelslag":                            "osk",
		"kb":                                   "ky",
		"julkinenosakeyhtio":                   "oyj",
		"abp":                                  "oyj",
		"osakeyhtio":                           "oy",
		"anl":                                  "osk",
		"kommandiittiyhtio":                    "ky",
		"kommanditbolag":                       "ky",
		"europabolag":                          "se",
		"osuuskunta":                           "osk",
		"europeiskekonomiskintresseguppering":  "etey",
		"eeig":                                 "etey",
		"eurooppaosuuskunta":                   "sce",
		"stiftelse":                            "sr",
	},
	"CZ": map[string]string{
		"evropskehospodarskezajmovesdruzeni": "ehzs",
		"evropskadruzstevnispolecnost":       "sce",
		"evropskaspolecnost":                 "se",
	},
	"FR": map[string]string{
		"societeparactionssimplifiee":          "sas",
		"groupementeuropeendintereteconomique": "geie",
		"societeanonyme":                       "sa",
		"societeenlibrepartenariat":            "slp",
		"societeeuropeenne":                    "se",
	},
	"PT": map[string]string{
		"societascooperativaeuropaea":            "sce",
		"agrupamentoeuropeudeinteresseeconomico": "aeie",
		"sociedadeanonimaeuropeia":               "se",
	},
	"CO": map[string]string{
		"empresaunipersonal":                 "eu",
		"sociedadporaccionessimplificada":    "sas",
		"sociedadesagrariasdetransformacion": "sat",
		"sociedadanonima":                    "sa",
		"sociedadencomanditasimple":          "ciasenc",
		"ciasenc":                            "ciasenc",
		"sociedadencomanditaporacciones":     "ciasca",
		"ciasca":                             "ciasca",
		"sociedadderesponsabilidadlimitada":  "ltda",
	},
	"ID": map[string]string{
		"perusahaanperseroan":                        "persero",
		"perseroanterbataspenanamanmodalasing":       "ptpma",
		"perseroanterbataspenanamanmodaldalamnegeri": "pt",
		"ptpmdn":                    "pt",
		"perseroanterbuka":          "tbk",
		"commanditairevennootschap": "cv",
		"firma":                     "fa",
		"perusahaanumum":            "perum",
	},
	"PL": map[string]string{
		"spolkakomandytowa":                "spk",
		"spolkaakcyjna":                    "sa",
		"spolkaeuropejska":                 "se",
		"spolkakomandytowoakcyjna":         "ska",
		"otwartyfunduszemerytalny":         "ofe",
		"spolkajawna":                      "sj",
		"spj":                              "sj",
		"prostaspolkaakcyjna":              "psa",
		"towarzystwoubezpieczenwzajemnych": "tuw",
		"spolkacywilna":                    "sc",
		"spolkapartnerska":                 "spp",
		"europejskiezgrupowanieinteresowgospodarczych": "ezig",
		"spolkazograniczonaodpowiedzialnoscia":         "spzoo",
	},
	"NZ": map[string]string{
		"limitedpartnership":        "lp",
		"limitedliabilitycompany":   "ltd",
		"limited":                   "ltd",
		"cooperativecompany":        "coop",
		"unlimitedliabilitycompany": "ultd",
		"unltd":                     "ultd",
	},
	"DE": map[string]string{
		"gemeinnutzigegesellschaftmitbeschrankterhaftung": "ggmbh",
		"societaseuropaea":                      "se",
		"gesellschaftmitbeschrankterhaftung":    "gmbh",
		"kommanditgesellschaft":                 "kg",
		"reitaktiengesellschaft":                "reitag",
		"eingetragenekauffrau":                  "ek",
		"ekfm":                                  "ek",
		"eingetragenerverein":                   "ev",
		"europaischegenossenschaft":             "sce",
		"gemeinnutzigeaktiengesellschaft":       "gag",
		"ughaftungsbeschrankt":                  "ug",
		"aktiengesellschaft":                    "ag",
		"partnerschaftsgesellschaft":            "partg",
		"eingetragenegenossenschaft":            "eg",
		"ekffr":                                 "ek",
		"kommanditgesellschaftaufaktien":        "kgaa",
		"offenehandelsgesellschaft":             "ohg",
		"unternehmergesellschaft":               "ug",
		"gesellschaftburgerlichenrechts":        "gbr",
		"investmentaktiengesellschaft":          "invag",
		"eingetragenerkaufmann":                 "ek",
		"versicherungsvereinaufgegenseitigkeit": "vvag",
		"gesmbh":                                "gmbh",
		"partnerschaftsgesellschaftmitbeschrankterberufshaftung": "partgmbb",
		"europaischewirtschaftlicheinteressenvereinigung":        "ewiv",
		"europaischeaktiengesellschaft":                          "se",
	},
	"CA": map[string]string{
		"associationcooperative":           "coop",
		"societeainteretcommunautaire":     "cic",
		"limitedcorporation":               "ltd",
		"sec":                              "lp",
		"cooperative":                      "coop",
		"societearesponsabilitefiduciaire": "sarf",
		"cp":                               "pc",
		"cooperativeassociation":           "coop",
		"incorporated":                     "inc",
		"limitedliabilitypartnership":      "llp",
		"limited":                          "ltd",
		"limitee":                          "ltd",
		"societeennomcollectifaresponsabilitelimitee": "llp",
		"professionalcorporation":                     "pc",
		"societeencommandite":                         "lp",
		"sic":                                         "cic",
		"unlimitedliabilitycorporation":               "ulc",
		"corporationprofessionnelle":                  "pc",
		"communityinterestcompany":                    "cic",
		"unlimitedcompany":                            "ulc",
		"ltee":                                        "ltd",
		"limitedpartnership":                          "lp",
		"corporation":                                 "corp",
		"societeparactions":                           "sa",
		"sencrl":                                      "llp",
	},
	"CL": map[string]string{
		"sociedadlegalminera":                        "slm",
		"asociaciongremial":                          "ag",
		"sociedadderesponsabilidadlimitada":          "srl",
		"sociedadporacciones":                        "spa",
		"sociedadanonima":                            "sa",
		"sociedadcontractualminera":                  "scm",
		"empresaindividualderesponsabilidadlimitada": "eirl",
	},
	"PA": map[string]string{
		"microempresaderesponsabilidadlimitada": "mrl",
		"sociedadanonima":                       "sa",
		"sociedadcivil":                         "sc",
		"sociedadderesponsabilidadlimitada":     "srl",
		"fundaciondeinteresprivado":             "fip",
	},
	"AU": map[string]string{
		"cooperativeltd":                  "coop",
		"coopltd":                         "coop",
		"noliability":                     "nl",
		"indigenouscorporation":           "rntbc",
		"cooperative":                     "coop",
		"cooperativelimited":              "coop",
		"incorporated":                    "inc",
		"limited":                         "ltd",
		"cooplimited":                     "coop",
		"noliabilitycompany":              "nl",
		"ptyltd":                          "pty",
		"incorporatedassociation":         "inc",
		"limitedpartnership":              "lp",
		"publiccompanylimitedbyguarantee": "ltd",
		"limitedproprietarycompany":       "pty",
	},
	"BR": map[string]string{
		"empresapublica":                 "ep",
		"sociedadeemcomdanditasimples":   "scs",
		"sociedadeanonima":               "sa",
		"companhia":                      "sa",
		"cia":                            "sa",
		"sociedadeemcontadeparticipacao": "scp",
		"sociedadeemnomecoletivo":        "snc",
		"sociedadesimples":               "ss",
		"fundacao":                       "fund",
		"associacao":                     "assoc",
		"empresaindividualderesponsabilidadelimitada": "eireli",
		"sociedadelimitada":                           "ltda",
		"sociedadeemcomanditaporacoes":                "ca",
	},
	"ZA": map[string]string{
		"personalliabilitycompany": "inc",
		"incorporated":             "inc",
		"publiccompany":            "ltd",
		"limited":                  "ltd",
		"closecorporation":         "cc",
		"privatecompany":           "ptyltd",
		"nonprofitcompany":         "npc",
		"stateownedcompany":        "soc",
	},
	"PE": map[string]string{
		"sociedadcomercialderesponsabilidadlimitada": "srl",
		"sociedaanonimaabierta":                      "saa",
		"sociedadcivil":                              "scivil",
		"sociedadencomanditasimple":                  "senc",
		"sociedadporaccionescerradasimplificada":     "sacs",
		"sociedadcivilderesponsabilidadlimitada":     "scivilderl",
		"sociedadanonimacerrada":                     "sac",
		"empresaindividualderesponsabilidadlimitada": "eirl",
		"sociedaanonima":                             "sa",
		"sociedadencomanditaporacciones":             "sencpora",
		"sociedadcolectiva":                          "sc",
	},
	"HU": map[string]string{
		"nyilvanosanmukodoreszvenytarsasag": "nyrt",
		"zartkoruenmukodoreszvenytarsasag":  "zrt",
		"reszvenytarsasag":                  "rt",
		"kozhasznutarsasag":                 "kht",
		"kozkeresetitarsasag":               "kkt",
		"korlatoltfelelossegutarsasag":      "kft",
		"takarekeshitelszovetkezet":         "hsz",
		"tksz":                              "hsz",
		"orszagosbetetbiztositasialap":      "oba",
		"europaigazdasagiegyesules":         "ege",
		"betetitarsasag":                    "bt",
		"szovetsegkivevesportszovetseg":     "kbe",
		"europaireszvenytarsasag":           "se",
	},
	"IE": map[string]string{
		"cn":                                          "ulc",
		"grupaileorpachumleaseacnamaioch":             "eeig",
		"companylimitedbyguarantee":                   "clg",
		"cti":                                         "ilp",
		"irishcollectiveassetmanagementvehicle":       "icav",
		"cuideachtafaoitheorainnrathaiochta":          "clg",
		"privateunlimitedcompany":                     "ulc",
		"limited":                                     "ltd",
		"privatecompanylimitedbyshares":               "ltd",
		"publiclimitedcompany":                        "plc",
		"designatedactivitycompany":                   "dac",
		"publicunlimitedcompanythathasnosharecapital": "pulc",
		"teoranta":                                    "ltd",
		"publicunlimitedcompany":                      "puc",
		"cpt":                                         "plc",
		"europeaneconomicinterestgrouping":            "eeig",
		"societaseuropaea":                            "se",
		"cuideachtaphriobhaideachtheoranta":           "ltd",
		"cuideachtaphoiblitheoranta":                  "plc",
		"gele":                                        "eeig",
		"cga":                                         "dac",
		"cuideachtaneamhtheoranta":                    "ulc",
		"limitedcompany":                              "ltd",
		"teo":                                         "ltd",
		"investmentlimitedpartnership":                "ilp",
		"cuideachtaghniomhaiochtaainmnithe":           "dac",
		"ctr":                                         "clg",
		"comhphairtiochttheorantainfheistiochta":      "ilp",
	},
	"IN": map[string]string{
		"publiclimitedcompany":        "ltd",
		"limited":                     "ltd",
		"limitedliabilitypartnership": "llp",
		"privatelimitedcompany":       "pvtltd",
	},
	"DK": map[string]string{
		"iværksætterselskab":              "ivs",
		"interessentskab":                 "is",
		"anpartsselskab":                  "aps",
		"andelsselskabmedbegrænsetansvar": "amba",
		"foreningmedbegrænsetansvar":      "fmba",
		"aktieselskab":                    "as",
		"europæiskselskab":                "seselskab",
		"partnerselskab":                  "ps",
		"europæiskandelsselskab":          "sceselskab",
		"selskabmedbegrænsetansvar":       "smba",
		"kommanditselskab":                "ks",
		"europæiskokonomiskfirmagruppe":   "eofg",
		"eofg":                            "eofg",
	},
	"AR": map[string]string{
		"sociedadcolectiva":               "soccol",
		"sociedadanonima":                 "sa",
		"sociedadanonimaunipersonal":      "sau",
		"sociedaddegarantiareciproca":     "sgr",
		"sociedadencomanditaporacciones":  "scpa",
		"sociedadderesponsalidedlimitada": "srl",
		"sociedaddelestado":               "se",
		"sociedaddecapitaleindustria":     "scei",
		"sociedadporaccionessimplificada": "sas",
		"sociedadencomanditasimple":       "scs",
		"sociedadcooperativa":             "scoop",
	},
	"VN": map[string]string{
		"congtytrachnhiemhuuhan": "llc",
		"congtytnhh":             "llc",
		"congtnhh":               "llc",
		"cttnhh":                 "llc",
		"ctytnhh":                "llc",
		"congtycophan":           "jsc",
		"congtycp":               "jsc",
		"ctycp":                  "jsc",
		"ctcp":                   "jsc",
		"tongcongty":             "gc",
		"doanhnghieptunhan":      "pe",
		"dntn":                   "pe",
		"congtyhopdanh":          "pc",
		"congtyhd":               "pc",
		"congtyliendoanh":        "jvc",
		"congtyld":               "jvc",
		"congtynhanuoc":          "soc",
		"congtynn":               "soc",

		// Already translated
		"limitedliabilitycompany": "llc",
		"ltdliabilityco":          "llc",
		"limited":                 "llc",
		"ltd":                     "llc",
		"companylimited":          "llc",
		"companyltd":              "llc",
		"colimited":               "llc",
		"coltd":                   "llc",
	},
	"JP": map[string]string{
		"株式会社":  "kk",  // Joint-stock Company(jsc) kabushiki kaisha(kk)
		"株":     "kk",  // Joint-stock Company(jsc) kabushiki kaisha(kk)
		"有限会社":  "kk",  // Limited Company(ltd) - yūgen kaisha(yk) (Old form replaced with kk)
		"有":     "kk",  // Limited Company(ltd) - yūgen kaisha(yk) (Old form replaced with kk)
		"合同会社":  "gk",  // Limited Liability Company(llc) - gōdō kaisha(gk)
		"合":     "gk",  // Limited Liability Company(llc) - gōdō kaisha(gk)
		"合名会社":  "gmk", // General Partnership Company(gp) - gōmei kaisha(gmk)
		"名":     "gmk", // General Partnership Company(gp) - gōmei kaisha(gmk)
		"合資会社":  "gsk", // Limited Partnership Company(lp) - gōshi kaisha(gsk)
		"資":     "gsk", // Limited Partnership Company(lp) - gōshi kaisha(gsk)
		"外国会社等": "外",   // Foreign Company, etc.
		"その他":   "そ",   // Others
		"医療法人":  "医",   // Medical Corporation(mc) - iryo hojin(ih)

		// detected but no alias
		// "その他の設立登記法人": "-", // Other Registered Corporations
		// "地方公共団体":     "-", // Local Public Entity
		// "国の機関":       "-", // National Agency

		// romaji (japanese)
		"kabushikikaisha": "kk",
		"godokaisha":      "gk",
		"yugenkaisha":     "kk",
		"goshikaisha":     "gsk",
		"gomeikaisha":     "gmk",

		// pinyin (chinese)
		"zhushihuishe":  "kk",
		"hetonghuishe":  "gk",
		"youxianhuishe": "kk",
		"hezihuishe":    "gsk",
		"heminghuishe":  "gmk",

		// Already translated
		"limitedliabilitycompany": "gk",
		"ltdliabilityco":          "gk",
		"llc":                     "gk",
		"ltd":                     "kk",
		"companylimited":          "kk",
		"companyltd":              "kk",
		"colimited":               "kk",
		"coltd":                   "kk",
		"limited":                 "kk",
		"incorporated":            "kk",
		"inc":                     "kk",
		"corporation":             "kk",
		"corp":                    "kk",
		"corpinc":                 "kk",
		"corpkk":                  "kk",
		"corpyk":                  "kk",
	},
}
