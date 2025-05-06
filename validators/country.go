package validators

import (
	"errors"
	"strings"
)

var (
	countriesNamesInEnglish = []string{
		"afghanistan",
		"albania",
		"algeria",
		"andorra",
		"angola",
		"antigua and barbuda",
		"argentina",
		"armenia",
		"australia",
		"austria",
		"azerbaijan",
		"bahamas",
		"bahrain",
		"bangladesh",
		"barbados",
		"belarus",
		"belgium",
		"belize",
		"benin",
		"bhutan",
		"bolivia",
		"bosnia and herzegovina",
		"botswana",
		"brazil",
		"brunei",
		"bulgaria",
		"burkina faso",
		"burundi",
		"cabo verde",
		"cambodia",
		"cameroon",
		"canada",
		"central african republic",
		"chad",
		"chile",
		"china",
		"colombia",
		"comoros",
		"congo, democratic republic of the",
		"congo, republic of the",
		"costa rica",
		"côte d'ivoire",
		"croatia",
		"cuba",
		"cyprus",
		"czechia",
		"denmark",
		"djibouti",
		"dominica",
		"dominican republic",
		"ecuador",
		"egypt",
		"el salvador",
		"equatorial guinea",
		"eritrea",
		"estonia",
		"eswatini",
		"ethiopia",
		"fiji",
		"finland",
		"france",
		"gabon",
		"gambia",
		"georgia",
		"germany",
		"ghana",
		"greece",
		"grenada",
		"guatemala",
		"guinea",
		"guinea-bissau",
		"guyana",
		"haiti",
		"honduras",
		"hungary",
		"iceland",
		"india",
		"indonesia",
		"iran",
		"iraq",
		"ireland",
		"israel",
		"italy",
		"jamaica",
		"japan",
		"jordan",
		"kazakhstan",
		"kenya",
		"kiribati",
		"korea, north",
		"korea, south",
		"kuwait",
		"kyrgyzstan",
		"laos",
		"latvia",
		"lebanon",
		"lesotho",
		"liberia",
		"libya",
		"liechtenstein",
		"lithuania",
		"luxembourg",
		"madagascar",
		"malawi",
		"malaysia",
		"maldives",
		"mali",
		"malta",
		"marshall islands",
		"mauritania",
		"mauritius",
		"mexico",
		"micronesia, federated states of",
		"moldova",
		"monaco",
		"mongolia",
		"montenegro",
		"morocco",
		"mozambique",
		"myanmar",
		"namibia",
		"nauru",
		"nepal",
		"netherlands",
		"new zealand",
		"nicaragua",
		"niger",
		"nigeria",
		"north macedonia",
		"norway",
		"oman",
		"pakistan",
		"palau",
		"panama",
		"papua new guinea",
		"paraguay",
		"peru",
		"philippines",
		"poland",
		"portugal",
		"qatar",
		"romania",
		"russia",
		"rwanda",
		"saint kitts and nevis",
		"saint lucia",
		"saint vincent and the grenadines",
		"samoa",
		"san marino",
		"sao tome and principe",
		"saudi arabia",
		"senegal",
		"serbia",
		"seychelles",
		"sierra leone",
		"singapore",
		"slovakia",
		"slovenia",
		"solomon islands",
		"somalia",
		"south africa",
		"south sudan",
		"spain",
		"sri lanka",
		"sudan",
		"suriname",
		"sweden",
		"switzerland",
		"syria",
		"taiwan",
		"tajikistan",
		"tanzania",
		"thailand",
		"timor-leste",
		"togo",
		"tonga",
		"trinidad and tobago",
		"tunisia",
		"turkey",
		"turkmenistan",
		"tuvalu",
		"uganda",
		"ukraine",
		"united arab emirates",
		"united kingdom",
		"united states",
		"uruguay",
		"uzbekistan",
		"vanuatu",
		"vatican city",
		"venezuela",
		"vietnam",
		"yemen",
		"zambia",
		"zimbabwe",
	}

	africaCountries = []string{
		"algeria", "angola", "benin", "botswana", "burkina faso", "burundi", "cabo verde", "cameroon",
		"central african republic", "chad", "comoros", "congo, democratic republic of the",
		"congo, republic of the", "côte d'ivoire", "djibouti", "egypt", "equatorial guinea", "eritrea",
		"eswatini", "ethiopia", "gabon", "gambia", "ghana", "guinea", "guinea-bissau", "kenya",
		"lesotho", "liberia", "libya", "madagascar", "malawi", "mali", "mauritania", "mauritius",
		"morocco", "mozambique", "namibia", "niger", "nigeria", "rwanda", "sao tome and principe",
		"senegal", "seychelles", "sierra leone", "somalia", "south africa", "south sudan", "sudan",
		"tanzania", "togo", "tunisia", "uganda", "zambia", "zimbabwe",
	}

	asiaCountries = []string{
		"afghanistan", "armenia", "azerbaijan", "bahrain", "bangladesh", "bhutan", "brunei", "cambodia",
		"china", "cyprus", "georgia", "india", "indonesia", "iran", "iraq", "israel", "japan", "jordan",
		"kazakhstan", "kuwait", "kyrgyzstan", "laos", "lebanon", "malaysia", "maldives", "mongolia",
		"myanmar", "nepal", "north korea", "oman", "pakistan", "palestine", "philippines", "qatar",
		"saudi arabia", "singapore", "south korea", "sri lanka", "syria", "taiwan", "tajikistan",
		"thailand", "timor-leste", "turkey", "turkmenistan", "united arab emirates", "uzbekistan",
		"vietnam", "yemen",
	}

	europeCountries = []string{
		"albania", "andorra", "austria", "belarus", "belgium", "bosnia and herzegovina", "bulgaria",
		"croatia", "czechia", "denmark", "estonia", "finland", "france", "germany", "greece", "hungary",
		"iceland", "ireland", "italy", "latvia", "liechtenstein", "lithuania", "luxembourg", "malta",
		"moldova", "monaco", "montenegro", "netherlands", "north macedonia", "norway", "poland",
		"portugal", "romania", "russia", "san marino", "serbia", "slovakia", "slovenia", "spain",
		"sweden", "switzerland", "ukraine", "united kingdom", "vatican city",
	}

	northAmericaCountries = []string{
		"antigua and barbuda", "bahamas", "barbados", "belize", "canada", "costa rica", "cuba",
		"dominica", "dominican republic", "el salvador", "grenada", "guatemala", "haiti", "honduras",
		"jamaica", "mexico", "nicaragua", "panama", "saint kitts and nevis", "saint lucia",
		"saint vincent and the grenadines", "trinidad and tobago", "united states",
	}

	southAmericaCountries = []string{
		"argentina", "bolivia", "brazil", "chile", "colombia", "ecuador", "guyana", "paraguay", "peru",
		"suriname", "uruguay", "venezuela",
	}

	oceaniaCountries = []string{
		"australia", "fiji", "kiribati", "marshall islands", "micronesia, federated states of", "nauru",
		"new zealand", "palau", "papua new guinea", "samoa", "solomon islands", "tonga", "tuvalu",
		"vanuatu",
	}
)

func CountryValidator(country string) (bool, error) {
	country = strings.ToLower(country)
	for _, c := range countriesNamesInEnglish {
		if country == c {
			return true, nil
		}
	}
	return false, errors.New("country name is not valid")
}

func CountryValidatorByContinent(country string, continent string) (bool, error) {
	country = strings.ToLower(country)
	continent = strings.ToLower(continent)

	var countries []string

	switch continent {
	case "africa":
		countries = africaCountries
	case "asia":
		countries = asiaCountries
	case "europe":
		countries = europeCountries
	case "north america":
		countries = northAmericaCountries
	case "south america":
		countries = southAmericaCountries
	case "oceania":
		countries = oceaniaCountries
	default:
		return false, errors.New("continent name is not valid")
	}

	for _, c := range countries {
		if country == c {
			return true, nil
		}
	}
	return false, errors.New("country name is not valid in the specified continent")
}
