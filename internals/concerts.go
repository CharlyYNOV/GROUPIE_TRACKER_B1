package internals

import (
	"encoding/json"
	"html/template"
	"strings"
)

// Dictionnaire complet de toutes les localisations de ton JSON
var CityCoords = map[string][2]float64{
	"north_carolina-usa": {35.7596, -79.0193}, "georgia-usa": {33.7490, -84.3880},
	"los_angeles-usa": {34.0522, -118.2437}, "saitama-japan": {35.8617, 139.6455},
	"osaka-japan": {34.6937, 135.5023}, "nagoya-japan": {35.1815, 136.9066},
	"penrose-new_zealand": {-36.9110, 174.8110}, "dunedin-new_zealand": {-45.8788, 170.5028},
	"playa_del_carmen-mexico": {20.6296, -87.0739}, "papeete-french_polynesia": {-17.5333, -149.5667},
	"noumea-new_caledonia": {-22.2711, 166.4411}, "london-uk": {51.5074, -0.1278},
	"lausanne-switzerland": {46.5197, 6.6323}, "lyon-france": {45.7640, 4.8357},
	"victoria-australia": {-37.8368, 144.9280}, "new_south_wales-australia": {-31.2532, 146.9211},
	"queensland-australia": {-20.9176, 142.7028}, "auckland-new_zealand": {-36.8485, 174.7633},
	"yogyakarta-indonesia": {-7.7956, 110.3695}, "bratislava-slovakia": {48.1486, 17.1077},
	"budapest-hungary": {47.4979, 19.0402}, "minsk-belarus": {53.9006, 27.5590},
	"california-usa": {36.7783, -119.4179}, "nevada-usa": {38.8026, -116.4194},
	"sao_paulo-brazil": {-23.5505, -46.6333}, "san_isidro-argentina": {-34.4718, -58.5286},
	"arizona-usa": {34.0489, -111.0937}, "texas-usa": {31.9686, -99.9018},
	"stockholm-sweden": {59.3293, 18.0686}, "werchter-belgium": {50.9694, 4.6936},
	"lisbon-portugal": {38.7223, -9.1393}, "bilbao-spain": {43.2630, -2.9350},
	"bogota-colombia": {4.7110, -74.0721}, "new_york-usa": {40.7128, -74.0060},
	"dusseldorf-germany": {51.2271, 6.7735}, "aarhus-denmark": {56.1518, 10.2039},
	"manchester-uk": {53.4808, -2.2426}, "frankfurt-germany": {50.1109, 8.6821},
	"berlin-germany": {52.5200, 13.4050}, "copenhagen-denmark": {55.6761, 12.5683},
	"doha-qatar": {25.2854, 51.5310}, "minnesota-usa": {46.7296, -94.6859},
	"illinois-usa": {40.6331, -89.3985}, "mumbai-india": {19.0760, 72.8777},
	"abu_dhabi-united_arab_emirates": {24.4539, 54.3773}, "pennsylvania-usa": {41.2033, -77.1945},
	"westcliff_on_sea-uk": {51.5377, 0.6934}, "merkers-germany": {50.8167, 10.1167},
	"maine-usa": {45.2538, -69.4455}, "gothenburg-sweden": {57.7089, 11.9746},
	"florida-usa": {27.6648, -81.5158}, "south_carolina-usa": {33.8361, -81.1637},
	"pagney_derriere_barine-france": {48.6925, 5.8622}, "hamburg-germany": {53.5511, 9.9937},
	"boulogne_billancourt-france": {48.8397, 2.2399}, "sion-switzerland": {46.2331, 7.3606},
	"ostrava-czechia": {49.8209, 18.2625}, "klagenfurt-austria": {46.6247, 14.3053},
	"freyming_merlebach-france": {49.1465, 6.8142}, "zaragoza-spain": {41.6488, -0.8891},
	"madrid-spain": {40.4168, -3.7038}, "barcelona-spain": {41.3851, 2.1734},
	"rio_de_janeiro-brazil": {-22.9068, -43.1729}, "recife-brazil": {-8.0543, -34.8813},
	"leipzig-germany": {51.3397, 12.3731}, "salem-germany": {47.7770, 9.2786},
	"monchengladbach-germany": {51.1854, 6.4417}, "cuxhaven-germany": {53.8593, 8.6877},
	"skanderborg-denmark": {56.0319, 9.9272}, "amsterdam-netherlands": {52.3676, 4.9041},
	"burriana-spain": {39.8896, -0.0847}, "oulu-finland": {65.0121, 25.4651},
	"napoca-romania": {46.7712, 23.6236}, "riyadh-saudi_arabia": {24.7136, 46.6753},
	"canton-usa": {40.7989, -81.3784}, "las_vegas-usa": {36.1716, -115.1391},
	"mexico_city-mexico": {19.4326, -99.1332}, "monterrey-mexico": {25.6866, -100.3161},
	"del_mar-usa": {32.9595, -117.2653}, "paris-france": {48.8566, 2.3522},
	"missouri-usa": {37.9643, -91.8318}, "chicago-usa": {41.8781, -87.6298},
	"birmingham-uk": {52.4862, -1.8904}, "madison-usa": {43.0731, -89.4012},
	"cleveland-usa": {41.4993, -81.6944}, "boston-usa": {42.3601, -71.0589},
	"utah-usa": {39.3210, -111.0937}, "glasgow-uk": {55.8642, -4.2518},
	"dublin-ireland": {53.3498, -6.2603}, "cardiff-uk": {51.4816, -3.1791},
	"aberdeen-uk": {57.1497, -2.0943}, "warsaw-poland": {52.2297, 21.0122},
	"sochaux-france": {47.5143, 6.8322}, "eindhoven-netherlands": {51.4416, 5.4697},
	"colorado-usa": {39.5501, -105.7821}, "jakarta-indonesia": {-6.2088, 106.8456},
	"huizhou-china": {23.1118, 114.4161}, "changzhou-china": {31.8112, 119.9741},
	"hong_kong-china": {22.3193, 114.1694}, "sanya-china": {18.2528, 109.5119},
	"aalborg-denmark": {57.0488, 9.9217}, "seattle-usa": {47.6062, -122.3321},
	"omaha-usa": {41.2565, -95.9345}, "kansas_city-usa": {39.0997, -94.5786},
	"st_louis-usa": {38.6270, -90.1994}, "indianapolis-usa": {39.7684, -86.1581},
	"rosemont-usa": {41.9867, -87.8723}, "grand_rapids-usa": {42.9634, -85.6681},
	"montreal-usa": {44.5000, -72.5667}, "newark-usa": {40.7357, -74.1724},
	"uniondale-usa": {40.7004, -73.5929}, "philadelphia-usa": {39.9526, -75.1652},
	"hershey-usa": {40.2859, -76.6502}, "pittsburgh-usa": {40.4406, -79.9959},
	"columbia-usa": {34.0007, -81.0348}, "santiago-chile": {-33.4489, -70.6693},
	"houston-usa": {29.7604, -95.3698}, "atlanta-usa": {33.7490, -84.3880},
	"new_orleans-usa": {29.9511, -90.0715}, "frauenfeld-switzerland": {47.5586, 8.8968},
	"turku-finland": {60.4518, 22.2666}, "brooklyn-usa": {40.6782, -73.9442},
	"montreal-canada": {45.5017, -73.5673}, "vienna-austria": {48.2082, 16.3738},
	"krakow-poland": {50.0647, 19.9450}, "zurich-switzerland": {47.3769, 8.5417},
	"amityville-usa": {40.6793, -73.4173}, "minneapolis-usa": {44.9778, -93.2650},
	"detroit-usa": {42.3314, -83.0458}, "oakland-usa": {37.8044, -122.2712},
	"charlotte-usa": {35.2271, -80.8431}, "inglewood-usa": {33.9617, -118.3531},
	"windsor-canada": {42.3149, -83.0364}, "cincinnati-usa": {39.1031, -84.5120},
	"anaheim-usa": {33.8366, -117.9143}, "manila-philippines": {14.5995, 120.9842},
	"brisbane-australia": {-27.4705, 153.0260}, "melbourne-australia": {-37.8136, 144.9631},
	"lima-peru": {-12.0464, -77.0428}, "groningen-netherlands": {53.2192, 6.5667},
	"antwerp-belgium": {51.2194, 4.4025}, "pico_rivera-usa": {33.9831, -118.0967},
	"berwyn-usa": {41.8506, -87.7937}, "dallas-usa": {32.7767, -96.7970},
	"brixton-uk": {51.4624, -0.1145}, "rotselaar-belgium": {50.9508, 4.7145},
	"alabama-usa": {32.3182, -86.9023}, "massachusetts-usa": {42.4072, -71.3824},
	"athens-greece": {37.9838, 23.7275}, "florence-italy": {43.7696, 11.2558},
	"landgraaf-netherlands": {50.9126, 6.0315}, "burswood-australia": {-31.9610, 115.8970},
	"wellington-new_zealand": {-41.2865, 174.7762}, "seville-spain": {37.3891, -5.9845},
	"taipei-taiwan": {25.0330, 121.5654}, "seoul-south_korea": {37.5665, 126.9780},
	"munich-germany": {48.1351, 11.5820}, "mannheim-germany": {49.4875, 8.4660},
	"san_francisco-usa": {37.7749, -122.4194}, "buenos_aires-argentina": {-34.6037, -58.3816},
	"porto_alegre-brazil": {-30.0346, -51.2177}, "belo_horizonte-brazil": {-19.9167, -43.9345},
	"la_plata-argentina": {-34.9214, -57.9545}, "dubai-united_arab_emirates": {25.2048, 55.2708},
	"willemstad-netherlands_antilles": {12.1167, -68.9333}, "brasilia-brazil": {-15.7975, -47.8919},
	"oklahoma-usa": {35.4676, -97.5164}, "scheessel-germany": {53.1667, 9.4833},
	"st_gallen-switzerland": {47.4239, 9.3748}, "gdynia-poland": {54.5189, 18.5305},
	"arras-france": {50.2910, 2.7775}, "san_jose-costa_rica": {9.9281, -84.0907},
	"nickelsdorf-austria": {47.9404, 17.0673}, "oregon-usa": {43.8041, -120.5542},
	"prague-czechia": {50.0755, 14.4378},
}

type MapMarker struct {
	Artist string  `json:"artist"`
	City   string  `json:"city"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

func GetMarkersJSON(query string) template.HTML {
	var finalMarkers []MapMarker
	query = strings.ToLower(strings.TrimSpace(query))

	for _, locItem := range Locations {
		var artistName string
		for _, a := range Artists {
			if a.Id == locItem.Id {
				artistName = a.Name
				break
			}
		}

		// Filtre Go : On n'ajoute que si le nom correspond
		if query != "" && !strings.Contains(strings.ToLower(artistName), query) {
			continue
		}

		for _, locName := range locItem.Locations {
			if coords, ok := CityCoords[locName]; ok {
				finalMarkers = append(finalMarkers, MapMarker{
					Artist: artistName,
					City:   locName,
					Lat:    coords[0],
					Lng:    coords[1],
				})
			}
		}
	}

	jsonData, _ := json.Marshal(finalMarkers)
	return template.HTML(jsonData)
}
