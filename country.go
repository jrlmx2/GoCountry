package gocountry

import (
	"fmt"
	"strings"
)

// Country defines and stores the known identifiers for a given country.
type Country struct {
	full      string
	codeTwo   string
	codeThree string
	num       string
}

// Full returns the full name of the country.
func (c *Country) Full() string {
	return c.full
}

// CodeTwo returns the 	ISO ALPHA-2 Code Example "US"
func (c *Country) CodeTwo() string {
	return c.codeTwo
}

// CodeThree returns the 	ISO ALPHA-3 Code Example "BOL"
func (c *Country) CodeThree() string {
	return c.codeThree
}

// Number returns the ISO Numeric Code UN M49 Numerical Code
func (c *Country) Number() string {
	return c.num
}

// ExistsFull finds the country full name in the input string
func (c *Country) ExistsFull(text string) bool {
	if strings.Contains(text, strings.ToLower(c.full)) {
		return true
	}

	return false
}

// ExistsTwo finds the country ISO ALPHA-2 identifier in the input string
func (c *Country) ExistsTwo(text string) bool {
	if strings.Contains(text, c.codeTwo) {
		return true
	}

	return false

}

// ExistsThree finds the country ISO ALPHA-3 identifier in the input string
func (c *Country) ExistsThree(text string) bool {
	if strings.Contains(text, c.codeThree) {
		return true
	}
	return false

}

// ExistsNum will find ISO Numeric Code UN M49 Numerical Code in the input text
func (c *Country) ExistsNum(text string) bool {
	if strings.Contains(text, c.num) {
		return true
	}

	return false
}

func (c *Country) String() string {
	return c.CodeTwo()
}

// Options toggle search functions and is required to pass into the Search function
type Options struct {
	full      bool
	codeTwo   bool
	codeThree bool
	number    bool
}

// FindByNumber accepts and input int and returns the country associated or nil
func FindByNumber(num int) *Country {
	search := fmt.Sprintf("%03d", num)
	for _, c := range countries {
		if c.ExistsNum(search) {
			return c
		}
	}
	return nil
}

// Search searches the input string text for occurrences of known countries
func Search(options *Options, text string) []*Country {
	contains := make([]*Country, 0)

	for _, c := range countries {
		if options.full && c.ExistsFull(strings.ToLower(text)) || options.codeTwo && c.ExistsTwo(text) || options.codeThree && c.ExistsThree(text) || options.number && c.ExistsNum(text) {
			contains = append(contains, c)
		}
	}

	return contains
}

var countries = [...]*Country{
	&Country{full: "Afghanistan", codeTwo: "AF", codeThree: "AFG", num: "004"},
	&Country{full: "ALA	Aland Islands", codeTwo: "AX", codeThree: "ALA", num: "248"},
	&Country{full: "Albania", codeTwo: "AL", codeThree: "ALB", num: "008"},
	&Country{full: "Algeria", codeTwo: "DZ", codeThree: "DZA", num: "012"},
	&Country{full: "American Samoa", codeTwo: "AS", codeThree: "ASM", num: "016"},
	&Country{full: "Andorra", codeTwo: "AD", codeThree: "AND", num: "020"},
	&Country{full: "Angola", codeTwo: "AO", codeThree: "AGO", num: "024"},
	&Country{full: "Anguilla", codeTwo: "AI", codeThree: "AIA", num: "660"},
	&Country{full: "Antarctica", codeTwo: "AQ", codeThree: "ATA", num: "010"},
	&Country{full: "Antigua and Barbuda", codeTwo: "AG", codeThree: "ATG", num: "028"},
	&Country{full: "Argentina", codeTwo: "AR", codeThree: "ARG", num: "032"},
	&Country{full: "Armenia", codeTwo: "AM", codeThree: "ARM", num: "051"},
	&Country{full: "Aruba", codeTwo: "AW", codeThree: "ABW", num: "533"},
	&Country{full: "Australia", codeTwo: "AU", codeThree: "AUS", num: "036"},
	&Country{full: "Austria", codeTwo: "AT", codeThree: "AUT", num: "040"},
	&Country{full: "Azerbaijan", codeTwo: "AZ", codeThree: "AZE", num: "031"},
	&Country{full: "Bahamas", codeTwo: "BS", codeThree: "BHS", num: "044"},
	&Country{full: "Bahrain", codeTwo: "BH", codeThree: "BHR", num: "048"},
	&Country{full: "Bangladesh", codeTwo: "BD", codeThree: "BGD", num: "050"},
	&Country{full: "Barbados", codeTwo: "BB", codeThree: "BRB", num: "052"},
	&Country{full: "Belarus", codeTwo: "BY", codeThree: "BLR", num: "112"},
	&Country{full: "Belgium", codeTwo: "BE", codeThree: "BEL", num: "056"},
	&Country{full: "Belize", codeTwo: "BZ", codeThree: "BLZ", num: "084"},
	&Country{full: "Benin", codeTwo: "BJ", codeThree: "BEN", num: "204"},
	&Country{full: "Bermuda", codeTwo: "BM", codeThree: "BMU", num: "060"},
	&Country{full: "Bhutan", codeTwo: "BT", codeThree: "BTN", num: "064"},
	&Country{full: "Bolivia", codeTwo: "BO", codeThree: "BOL", num: "068"},
	&Country{full: "Bosnia and Herzegovina", codeTwo: "BA", codeThree: "BIH", num: "070"},
	&Country{full: "Botswana", codeTwo: "BW", codeThree: "BWA", num: "072"},
	&Country{full: "Bouvet Island", codeTwo: "BV", codeThree: "BVT", num: "074"},
	&Country{full: "Brazil", codeTwo: "BR", codeThree: "BRA", num: "076"},
	&Country{full: "British Virgin Islands", codeTwo: "VG", codeThree: "VGB", num: "092"},
	&Country{full: "British Indian Ocean Territory", codeTwo: "IO", codeThree: "IOT", num: "086"},
	&Country{full: "Brunei Darussalam", codeTwo: "BN", codeThree: "BRN", num: "096"},
	&Country{full: "Bulgaria", codeTwo: "BG", codeThree: "BGR", num: "100"},
	&Country{full: "Burkina Faso", codeTwo: "BF", codeThree: "BFA", num: "854"},
	&Country{full: "Burundi", codeTwo: "BI", codeThree: "BDI", num: "108"},
	&Country{full: "Cambodia", codeTwo: "KH", codeThree: "KHM", num: "116"},
	&Country{full: "Cameroon", codeTwo: "CM", codeThree: "CMR", num: "120"},
	&Country{full: "Canada", codeTwo: "CA", codeThree: "CAN", num: "124"},
	&Country{full: "Cape Verde", codeTwo: "CV", codeThree: "CPV", num: "132"},
	&Country{full: "Cayman Islands", codeTwo: "KY", codeThree: "CYM", num: "136"},
	&Country{full: "Central African Republic", codeTwo: "CF", codeThree: "CAF", num: "140"},
	&Country{full: "Chad", codeTwo: "TD", codeThree: "TCD", num: "148"},
	&Country{full: "Chile", codeTwo: "CL", codeThree: "CHL", num: "152"},
	&Country{full: "China", codeTwo: "CN", codeThree: "CHN", num: "156"},
	&Country{full: "Hong Kong, Special Administrative Region of China", codeTwo: "HK", codeThree: "HKG", num: "344"},
	&Country{full: "Macao, Special Administrative Region of China", codeTwo: "MO", codeThree: "MAC", num: "446"},
	&Country{full: "Christmas Island", codeTwo: "CX", codeThree: "CXR", num: "162"},
	&Country{full: "Cocos Islands", codeTwo: "CC", codeThree: "CCK", num: "166"},
	&Country{full: "Colombia", codeTwo: "CO", codeThree: "COL", num: "170"},
	&Country{full: "Comoros", codeTwo: "KM", codeThree: "COM", num: "174"},
	&Country{full: "Congo", codeTwo: "CG", codeThree: "COG", num: "178"},
	&Country{full: "Congo, Democratic Republic of the", codeTwo: "CD", codeThree: "COD", num: "180"},
	&Country{full: "Cook Islands", codeTwo: "CK", codeThree: "COK", num: "184"},
	&Country{full: "Costa Rica", codeTwo: "CR", codeThree: "CRI", num: "188"},
	&Country{full: "Côte d'Ivoire", codeTwo: "CI", codeThree: "CIV", num: "384"},
	&Country{full: "Croatia", codeTwo: "HR", codeThree: "HRV", num: "191"},
	&Country{full: "Cuba", codeTwo: "CU", codeThree: "CUB", num: "192"},
	&Country{full: "Cyprus", codeTwo: "CY", codeThree: "CYP", num: "196"},
	&Country{full: "Czech Republic", codeTwo: "CZ", codeThree: "CZE", num: "203"},
	&Country{full: "Denmark", codeTwo: "DK", codeThree: "DNK", num: "208"},
	&Country{full: "Djibouti", codeTwo: "DJ", codeThree: "DJI", num: "262"},
	&Country{full: "Dominica", codeTwo: "DM", codeThree: "DMA", num: "212"},
	&Country{full: "Dominican Republic", codeTwo: "DO", codeThree: "DOM", num: "214"},
	&Country{full: "Ecuador", codeTwo: "EC", codeThree: "ECU", num: "218"},
	&Country{full: "Egypt", codeTwo: "EG", codeThree: "EGY", num: "818"},
	&Country{full: "El Salvador", codeTwo: "SV", codeThree: "SLV", num: "222"},
	&Country{full: "Equatorial Guinea", codeTwo: "GQ", codeThree: "GNQ", num: "226"},
	&Country{full: "Eritrea", codeTwo: "ER", codeThree: "ERI", num: "232"},
	&Country{full: "Estonia", codeTwo: "EE", codeThree: "EST", num: "233"},
	&Country{full: "Ethiopia", codeTwo: "ET", codeThree: "ETH", num: "231"},
	&Country{full: "Falkland Islands", codeTwo: "FK", codeThree: "FLK", num: "238"},
	&Country{full: "Faroe Islands", codeTwo: "FO", codeThree: "FRO", num: "234"},
	&Country{full: "Fiji", codeTwo: "FJ", codeThree: "FJI", num: "242"},
	&Country{full: "Finland", codeTwo: "FI", codeThree: "FIN", num: "246"},
	&Country{full: "France", codeTwo: "FR", codeThree: "FRA", num: "250"},
	&Country{full: "French Guiana", codeTwo: "GF", codeThree: "GUF", num: "254"},
	&Country{full: "French Polynesia", codeTwo: "PF", codeThree: "PYF", num: "258"},
	&Country{full: "French Southern Territories", codeTwo: "TF", codeThree: "ATF", num: "260"},
	&Country{full: "Gabon", codeTwo: "GA", codeThree: "GAB", num: "266"},
	&Country{full: "Gambia", codeTwo: "GM", codeThree: "GMB", num: "270"},
	&Country{full: "Georgia", codeTwo: "GE", codeThree: "GEO", num: "268"},
	&Country{full: "Germany", codeTwo: "DE", codeThree: "DEU", num: "276"},
	&Country{full: "Ghana", codeTwo: "GH", codeThree: "GHA", num: "288"},
	&Country{full: "Gibraltar", codeTwo: "GI", codeThree: "GIB", num: "292"},
	&Country{full: "Greece", codeTwo: "GR", codeThree: "GRC", num: "300"},
	&Country{full: "Greenland", codeTwo: "GL", codeThree: "GRL", num: "304"},
	&Country{full: "Grenada", codeTwo: "GD", codeThree: "GRD", num: "308"},
	&Country{full: "Guadeloupe", codeTwo: "GP", codeThree: "GLP", num: "312"},
	&Country{full: "Guam", codeTwo: "GU", codeThree: "GUM", num: "316"},
	&Country{full: "Guatemala", codeTwo: "GT", codeThree: "GTM", num: "320"},
	&Country{full: "Guernsey", codeTwo: "GG", codeThree: "GGY", num: "831"},
	&Country{full: "Guinea", codeTwo: "GN", codeThree: "GIN", num: "324"},
	&Country{full: "Guinea-Bissau", codeTwo: "GW", codeThree: "GNB", num: "624"},
	&Country{full: "Guyana", codeTwo: "GY", codeThree: "GUY", num: "328"},
	&Country{full: "Haiti", codeTwo: "HT", codeThree: "HTI", num: "332"},
	&Country{full: "Heard Island and Mcdonald Islands", codeTwo: "HM", codeThree: "HMD", num: "334"},
	&Country{full: "Holy See", codeTwo: "VA", codeThree: "VAT", num: "336"},
	&Country{full: "Honduras", codeTwo: "HN", codeThree: "HND", num: "340"},
	&Country{full: "Hungary", codeTwo: "HU", codeThree: "HUN", num: "348"},
	&Country{full: "Iceland", codeTwo: "IS", codeThree: "ISL", num: "352"},
	&Country{full: "India", codeTwo: "IN", codeThree: "IND", num: "356"},
	&Country{full: "Indonesia", codeTwo: "ID", codeThree: "IDN", num: "360"},
	&Country{full: "Iran, Islamic Republic of", codeTwo: "IR", codeThree: "IRN", num: "364"},
	&Country{full: "Iraq", codeTwo: "IQ", codeThree: "IRQ", num: "368"},
	&Country{full: "Ireland", codeTwo: "IE", codeThree: "IRL", num: "372"},
	&Country{full: "Isle of Man", codeTwo: "IM", codeThree: "IMN", num: "833"},
	&Country{full: "Israel", codeTwo: "IL", codeThree: "ISR", num: "376"},
	&Country{full: "Italy", codeTwo: "IT", codeThree: "ITA", num: "380"},
	&Country{full: "Jamaica", codeTwo: "JM", codeThree: "JAM", num: "388"},
	&Country{full: "Japan", codeTwo: "JP", codeThree: "JPN", num: "392"},
	&Country{full: "Jersey", codeTwo: "JE", codeThree: "JEY", num: "832"},
	&Country{full: "Jordan", codeTwo: "JO", codeThree: "JOR", num: "400"},
	&Country{full: "Kazakhstan", codeTwo: "KZ", codeThree: "KAZ", num: "398"},
	&Country{full: "Kenya", codeTwo: "KE", codeThree: "KEN", num: "404"},
	&Country{full: "Kiribati", codeTwo: "KI", codeThree: "KIR", num: "296"},
	&Country{full: "Korea, Democratic People's Republic of", codeTwo: "KP", codeThree: "PRK", num: "408"},
	&Country{full: "Korea, Republic of", codeTwo: "KR", codeThree: "KOR", num: "410"},
	&Country{full: "Kuwait", codeTwo: "KW", codeThree: "KWT", num: "414"},
	&Country{full: "Kyrgyzstan", codeTwo: "KG", codeThree: "KGZ", num: "417"},
	&Country{full: "Lao PDR", codeTwo: "LA", codeThree: "LAO", num: "418"},
	&Country{full: "Latvia", codeTwo: "LV", codeThree: "LVA", num: "428"},
	&Country{full: "Lebanon", codeTwo: "LB", codeThree: "LBN", num: "422"},
	&Country{full: "Lesotho", codeTwo: "LS", codeThree: "LSO", num: "426"},
	&Country{full: "Liberia", codeTwo: "LR", codeThree: "LBR", num: "430"},
	&Country{full: "Libya", codeTwo: "LY", codeThree: "LBY", num: "434"},
	&Country{full: "Liechtenstein", codeTwo: "LI", codeThree: "LIE", num: "438"},
	&Country{full: "Lithuania", codeTwo: "LT", codeThree: "LTU", num: "440"},
	&Country{full: "Luxembourg", codeTwo: "LU", codeThree: "LUX", num: "442"},
	&Country{full: "Macedonia", codeTwo: "MK", codeThree: "MKD", num: "807"},
	&Country{full: "Madagascar", codeTwo: "MG", codeThree: "MDG", num: "450"},
	&Country{full: "Malawi", codeTwo: "MW", codeThree: "MWI", num: "454"},
	&Country{full: "Malaysia", codeTwo: "MY", codeThree: "MYS", num: "458"},
	&Country{full: "Maldives", codeTwo: "MV", codeThree: "MDV", num: "462"},
	&Country{full: "Mali", codeTwo: "ML", codeThree: "MLI", num: "466"},
	&Country{full: "Malta", codeTwo: "MT", codeThree: "MLT", num: "470"},
	&Country{full: "Marshall Islands", codeTwo: "MH", codeThree: "MHL", num: "584"},
	&Country{full: "Martinique", codeTwo: "MQ", codeThree: "MTQ", num: "474"},
	&Country{full: "Mauritania", codeTwo: "MR", codeThree: "MRT", num: "478"},
	&Country{full: "Mauritius", codeTwo: "MU", codeThree: "MUS", num: "480"},
	&Country{full: "Mayotte", codeTwo: "YT", codeThree: "MYT", num: "175"},
	&Country{full: "Mexico", codeTwo: "MX", codeThree: "MEX", num: "484"},
	&Country{full: "Micronesia", codeTwo: "FM", codeThree: "FSM", num: "583"},
	&Country{full: "Moldova", codeTwo: "MD", codeThree: "MDA", num: "498"},
	&Country{full: "Monaco", codeTwo: "MC", codeThree: "MCO", num: "492"},
	&Country{full: "Mongolia", codeTwo: "MN", codeThree: "MNG", num: "496"},
	&Country{full: "Montenegro", codeTwo: "ME", codeThree: "MNE", num: "499"},
	&Country{full: "Montserrat", codeTwo: "MS", codeThree: "MSR", num: "500"},
	&Country{full: "Morocco", codeTwo: "MA", codeThree: "MAR", num: "504"},
	&Country{full: "Mozambique", codeTwo: "MZ", codeThree: "MOZ", num: "508"},
	&Country{full: "Myanmar", codeTwo: "MM", codeThree: "MMR", num: "104"},
	&Country{full: "Namibia", codeTwo: "NA", codeThree: "NAM", num: "516"},
	&Country{full: "Nauru", codeTwo: "NR", codeThree: "NRU", num: "520"},
	&Country{full: "Nepal", codeTwo: "NP", codeThree: "NPL", num: "524"},
	&Country{full: "Netherlands", codeTwo: "NL", codeThree: "NLD", num: "528"},
	&Country{full: "Netherlands Antilles", codeTwo: "AN", codeThree: "ANT", num: "530"},
	&Country{full: "New Caledonia", codeTwo: "NC", codeThree: "NCL", num: "540"},
	&Country{full: "New Zealand", codeTwo: "NZ", codeThree: "NZL", num: "554"},
	&Country{full: "Nicaragua", codeTwo: "NI", codeThree: "NIC", num: "558"},
	&Country{full: "Niger", codeTwo: "NE", codeThree: "NER", num: "562"},
	&Country{full: "Nigeria", codeTwo: "NG", codeThree: "NGA", num: "566"},
	&Country{full: "Niue", codeTwo: "NU", codeThree: "NIU", num: "570"},
	&Country{full: "Norfolk Island", codeTwo: "NF", codeThree: "NFK", num: "574"},
	&Country{full: "Northern Mariana Islands", codeTwo: "MP", codeThree: "MNP", num: "580"},
	&Country{full: "Norway", codeTwo: "NO", codeThree: "NOR", num: "578"},
	&Country{full: "Oman", codeTwo: "OM", codeThree: "OMN", num: "512"},
	&Country{full: "Pakistan", codeTwo: "PK", codeThree: "PAK", num: "586"},
	&Country{full: "Palau", codeTwo: "PW", codeThree: "PLW", num: "585"},
	&Country{full: "Palestine", codeTwo: "PS", codeThree: "PSE", num: "275"},
	&Country{full: "Panama", codeTwo: "PA", codeThree: "PAN", num: "591"},
	&Country{full: "Papua New Guinea", codeTwo: "PG", codeThree: "PNG", num: "598"},
	&Country{full: "Paraguay", codeTwo: "PY", codeThree: "PRY", num: "600"},
	&Country{full: "Peru", codeTwo: "PE", codeThree: "PER", num: "604"},
	&Country{full: "Philippines", codeTwo: "PH", codeThree: "PHL", num: "608"},
	&Country{full: "Pitcairn", codeTwo: "PN", codeThree: "PCN", num: "612"},
	&Country{full: "Poland", codeTwo: "PL", codeThree: "POL", num: "616"},
	&Country{full: "Portugal", codeTwo: "PT", codeThree: "PRT", num: "620"},
	&Country{full: "Puerto Rico", codeTwo: "PR", codeThree: "PRI", num: "630"},
	&Country{full: "Qatar", codeTwo: "QA", codeThree: "QAT", num: "634"},
	&Country{full: "Réunion", codeTwo: "RE", codeThree: "REU", num: "638"},
	&Country{full: "Romania", codeTwo: "RO", codeThree: "ROU", num: "642"},
	&Country{full: "Russian Federation", codeTwo: "RU", codeThree: "RUS", num: "643"},
	&Country{full: "Rwanda", codeTwo: "RW", codeThree: "RWA", num: "646"},
	&Country{full: "Saint-Barthélemy", codeTwo: "BL", codeThree: "BLM", num: "652"},
	&Country{full: "Saint Helena", codeTwo: "SH", codeThree: "SHN", num: "654"},
	&Country{full: "Saint Kitts and Nevis", codeTwo: "KN", codeThree: "KNA", num: "659"},
	&Country{full: "Saint Lucia", codeTwo: "LC", codeThree: "LCA", num: "662"},
	&Country{full: "Saint-Martin", codeTwo: "MF", codeThree: "MAF", num: "663"},
	&Country{full: "Saint Pierre and Miquelon", codeTwo: "PM", codeThree: "SPM", num: "666"},
	&Country{full: "Saint Vincent and Grenadines", codeTwo: "VC", codeThree: "VCT", num: "670"},
	&Country{full: "Samoa", codeTwo: "WS", codeThree: "WSM", num: "882"},
	&Country{full: "San Marino", codeTwo: "SM", codeThree: "SMR", num: "674"},
	&Country{full: "Sao Tome and Principe", codeTwo: "ST", codeThree: "STP", num: "678"},
	&Country{full: "Saudi Arabia", codeTwo: "SA", codeThree: "SAU", num: "682"},
	&Country{full: "Senegal", codeTwo: "SN", codeThree: "SEN", num: "686"},
	&Country{full: "Serbia", codeTwo: "RS", codeThree: "SRB", num: "688"},
	&Country{full: "Seychelles", codeTwo: "SC", codeThree: "SYC", num: "690"},
	&Country{full: "Sierra Leone", codeTwo: "SL", codeThree: "SLE", num: "694"},
	&Country{full: "Singapore", codeTwo: "SG", codeThree: "SGP", num: "702"},
	&Country{full: "Slovakia", codeTwo: "SK", codeThree: "SVK", num: "703"},
	&Country{full: "Slovenia", codeTwo: "SI", codeThree: "SVN", num: "705"},
	&Country{full: "Solomon Islands", codeTwo: "SB", codeThree: "SLB", num: "090"},
	&Country{full: "Somalia", codeTwo: "SO", codeThree: "SOM", num: "706"},
	&Country{full: "South Africa", codeTwo: "ZA", codeThree: "ZAF", num: "710"},
	&Country{full: "South Georgia and the South Sandwich Islands", codeTwo: "GS", codeThree: "SGS", num: "239"},
	&Country{full: "South Sudan", codeTwo: "SS", codeThree: "SSD", num: "728"},
	&Country{full: "Spain", codeTwo: "ES", codeThree: "ESP", num: "724"},
	&Country{full: "Sri Lanka", codeTwo: "LK", codeThree: "LKA", num: "144"},
	&Country{full: "Sudan", codeTwo: "SD", codeThree: "SDN", num: "736"},
	&Country{full: "Suriname", codeTwo: "SR", codeThree: "SUR", num: "740"},
	&Country{full: "Svalbard and Jan Mayen Islands", codeTwo: "SJ", codeThree: "SJM", num: "744"},
	&Country{full: "Swaziland", codeTwo: "SZ", codeThree: "SWZ", num: "748"},
	&Country{full: "Sweden", codeTwo: "SE", codeThree: "SWE", num: "752"},
	&Country{full: "Switzerland", codeTwo: "CH", codeThree: "CHE", num: "756"},
	&Country{full: "Syrian Arab Republic (Syria)", codeTwo: "SY", codeThree: "SYR", num: "760"},
	&Country{full: "Taiwan, Republic of China", codeTwo: "TW", codeThree: "TWN", num: "158"},
	&Country{full: "Tajikistan", codeTwo: "TJ", codeThree: "TJK", num: "762"},
	&Country{full: "Tanzania *, United Republic of", codeTwo: "TZ", codeThree: "TZA", num: "834"},
	&Country{full: "Thailand", codeTwo: "TH", codeThree: "THA", num: "764"},
	&Country{full: "Timor-Leste", codeTwo: "TL", codeThree: "TLS", num: "626"},
	&Country{full: "Togo", codeTwo: "TG", codeThree: "TGO", num: "768"},
	&Country{full: "Tokelau", codeTwo: "TK", codeThree: "TKL", num: "772"},
	&Country{full: "Tonga", codeTwo: "TO", codeThree: "TON", num: "776"},
	&Country{full: "Trinidad and Tobago", codeTwo: "TT", codeThree: "TTO", num: "780"},
	&Country{full: "Tunisia", codeTwo: "TN", codeThree: "TUN", num: "788"},
	&Country{full: "Turkey", codeTwo: "TR", codeThree: "TUR", num: "792"},
	&Country{full: "Turkmenistan", codeTwo: "TM", codeThree: "TKM", num: "795"},
	&Country{full: "Turks and Caicos Islands", codeTwo: "TC", codeThree: "TCA", num: "796"},
	&Country{full: "Tuvalu", codeTwo: "TV", codeThree: "TUV", num: "798"},
	&Country{full: "Uganda", codeTwo: "UG", codeThree: "UGA", num: "800"},
	&Country{full: "Ukraine", codeTwo: "UA", codeThree: "UKR", num: "804"},
	&Country{full: "United Arab Emirates", codeTwo: "AE", codeThree: "ARE", num: "784"},
	&Country{full: "United Kingdom", codeTwo: "GB", codeThree: "GBR", num: "826"},
	&Country{full: "United States of America", codeTwo: "US", codeThree: "USA", num: "840"},
	&Country{full: "United States Minor Outlying Islands", codeTwo: "UM", codeThree: "UMI", num: "581"},
	&Country{full: "Uruguay", codeTwo: "UY", codeThree: "URY", num: "858"},
	&Country{full: "Uzbekistan", codeTwo: "UZ", codeThree: "UZB", num: "860"},
	&Country{full: "Vanuatu", codeTwo: "VU", codeThree: "VUT", num: "548"},
	&Country{full: "Venezuela", codeTwo: "VE", codeThree: "VEN", num: "862"},
	&Country{full: "Viet Nam", codeTwo: "VN", codeThree: "VNM", num: "704"},
	&Country{full: "Virgin Islands, US", codeTwo: "VI", codeThree: "VIR", num: "850"},
	&Country{full: "Wallis and Futuna Islands", codeTwo: "WF", codeThree: "WLF", num: "876"},
	&Country{full: "Western Sahara", codeTwo: "EH", codeThree: "ESH", num: "732"},
	&Country{full: "Yemen", codeTwo: "YE", codeThree: "YEM", num: "887"},
	&Country{full: "Zambia", codeTwo: "ZM", codeThree: "ZMB", num: "894"},
	&Country{full: "Zimbabwe", codeTwo: "ZW", codeThree: "ZWE", num: "716"},
}
