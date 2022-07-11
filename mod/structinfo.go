package mod

import "time"

// ServiceInfo 参考 https://github.com/YetClass/QuakeAPI/blob/master/core/quake.go 中的结构
type ServiceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Time      time.Time `json:"time"`
		Transport string    `json:"transport"`
		Service   struct {
			HTTP struct {
				HTMLHash string `json:"html_hash"`
				Favicon  struct {
					Hash     string `json:"hash"`
					Location string `json:"location"`
					Data     string `json:"data"`
				} `json:"favicon"`
				Robots          string `json:"robots"`
				SitemapHash     string `json:"sitemap_hash"`
				Server          string `json:"server"`
				Body            string `json:"body"`
				XPoweredBy      string `json:"x_powered_by"`
				MetaKeywords    string `json:"meta_keywords"`
				RobotsHash      string `json:"robots_hash"`
				Sitemap         string `json:"sitemap"`
				Path            string `json:"path"`
				Title           string `json:"title"`
				Host            string `json:"host"`
				SecurityText    string `json:"security_text"`
				StatusCode      int    `json:"status_code"`
				ResponseHeaders string `json:"response_headers"`
			} `json:"http"`
			Version  string `json:"version"`
			Name     string `json:"name"`
			Product  string `json:"product"`
			Banner   string `json:"banner"`
			Response string `json:"response"`
		} `json:"service"`
		Images     []interface{} `json:"images"`
		OsName     string        `json:"os_name"`
		Components []interface{} `json:"components"`
		Location   struct {
			DistrictCn  string    `json:"district_cn"`
			ProvinceCn  string    `json:"province_cn"`
			Gps         []float64 `json:"gps"`
			ProvinceEn  string    `json:"province_en"`
			CityEn      string    `json:"city_en"`
			CountryCode string    `json:"country_code"`
			CountryEn   string    `json:"country_en"`
			Radius      float64   `json:"radius"`
			DistrictEn  string    `json:"district_en"`
			Isp         string    `json:"isp"`
			StreetEn    string    `json:"street_en"`
			Owner       string    `json:"owner"`
			CityCn      string    `json:"city_cn"`
			CountryCn   string    `json:"country_cn"`
			StreetCn    string    `json:"street_cn"`
		} `json:"location"`
		Asn       int    `json:"asn"`
		Hostname  string `json:"hostname"`
		Org       string `json:"org"`
		OsVersion string `json:"os_version"`
		IsIpv6    bool   `json:"is_ipv6"`
		IP        string `json:"ip"`
		Port      int    `json:"port"`
	} `json:"data"`
	Meta struct {
		Total        int    `json:"total"`
		PaginationID string `json:"pagination_id"`
	} `json:"meta"`
}

type HunterInfo struct {
	Code int `json:"code"`
	Data struct {
		AccountType string `json:"account_type"`
		Total       int    `json:"total"`
		Time        int    `json:"time"`
		Arr         []struct {
			IsRisk         string `json:"is_risk"`
			URL            string `json:"url"`
			IP             string `json:"ip"`
			Port           int    `json:"port"`
			WebTitle       string `json:"web_title"`
			Domain         string `json:"domain"`
			IsRiskProtocol string `json:"is_risk_protocol"`
			Protocol       string `json:"protocol"`
			BaseProtocol   string `json:"base_protocol"`
			StatusCode     int    `json:"status_code"`
			Component      []struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"component"`
			Os        string `json:"os"`
			Company   string `json:"company"`
			Number    string `json:"number"`
			Country   string `json:"country"`
			Province  string `json:"province"`
			City      string `json:"city"`
			UpdatedAt string `json:"updated_at"`
			IsWeb     string `json:"is_web"`
			AsOrg     string `json:"as_org"`
			Isp       string `json:"isp"`
			Banner    string `json:"banner"`
		} `json:"arr"`
		ConsumeQuota string `json:"consume_quota"`
		RestQuota    string `json:"rest_quota"`
		SyntaxPrompt string `json:"syntax_prompt"`
	} `json:"data"`
	Message string `json:"message"`
}

type ZoomInfo struct {
	Code      int `json:"code"`
	Total     int `json:"total"`
	Available int `json:"available"`
	Matches   []struct {
		Rdns string `json:"rdns"`
		Jarm string `json:"jarm"`
		Ico  struct {
			Mmh3 string `json:"mmh3"`
			Md5  string `json:"md5"`
		} `json:"ico"`
		Txtfile struct {
			Robotsmd5   string `json:"robotsmd5"`
			Securitymd5 string `json:"securitymd5"`
		} `json:"txtfile"`
		IP       string `json:"ip"`
		Portinfo struct {
			Hostname  string      `json:"hostname"`
			Os        string      `json:"os"`
			Port      int         `json:"port"`
			Service   string      `json:"service"`
			Title     interface{} `json:"title"`
			Version   string      `json:"version"`
			Device    string      `json:"device"`
			Extrainfo string      `json:"extrainfo"`
			Rdns      string      `json:"rdns"`
			App       string      `json:"app"`
			Banner    string      `json:"banner"`
		} `json:"portinfo"`
		Timestamp string `json:"timestamp"`
		Geoinfo   struct {
			Continent struct {
				Code  string `json:"code"`
				Names struct {
					En   string `json:"en"`
					ZhCN string `json:"zh-CN"`
				} `json:"names"`
				GeonameID interface{} `json:"geoname_id"`
			} `json:"continent"`
			Country struct {
				Code  string `json:"code"`
				Names struct {
					En   string `json:"en"`
					ZhCN string `json:"zh-CN"`
				} `json:"names"`
				GeonameID interface{} `json:"geoname_id"`
			} `json:"country"`
			BaseStation string `json:"base_station"`
			City        struct {
				Names struct {
					En   string `json:"en"`
					ZhCN string `json:"zh-CN"`
				} `json:"names"`
				GeonameID interface{} `json:"geoname_id"`
			} `json:"city"`
			Isp          string `json:"isp"`
			Organization string `json:"organization"`
			Idc          string `json:"idc"`
			Location     struct {
				Lon string `json:"lon"`
				Lat string `json:"lat"`
			} `json:"location"`
			Aso          interface{} `json:"aso"`
			Asn          string      `json:"asn"`
			Subdivisions struct {
				Names struct {
					En   string `json:"en"`
					ZhCN string `json:"zh-CN"`
				} `json:"names"`
				GeonameID interface{} `json:"geoname_id"`
			} `json:"subdivisions"`
			PoweredBy string `json:"PoweredBy"`
			Scene     struct {
				En string `json:"en"`
				Cn string `json:"cn"`
			} `json:"scene"`
			OrganizationCN interface{} `json:"organization_CN"`
		} `json:"geoinfo"`
		Protocol struct {
			Application string `json:"application"`
			Probe       string `json:"probe"`
			Transport   string `json:"transport"`
		} `json:"protocol"`
		Honeypot int `json:"honeypot"`
		Whois    struct {
			Two0161130 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2016-11-30"`
		} `json:"whois,omitempty"`
		Whois0 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois1 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois2 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois3 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois4 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois5 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois6 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois7 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois8 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
		Whois9 struct {
			Two0170518 struct {
				AdminC       string        `json:"admin_c"`
				Country      string        `json:"country"`
				Descr        string        `json:"descr"`
				Inetnum      string        `json:"inetnum"`
				IPEnd        string        `json:"ip_end"`
				IPStart      string        `json:"ip_start"`
				Irt          []interface{} `json:"irt"`
				LastModified string        `json:"last_modified"`
				MntBy        string        `json:"mnt_by"`
				MntIrt       string        `json:"mnt_irt"`
				MntLower     string        `json:"mnt_lower"`
				MntRoutes    string        `json:"mnt_routes"`
				Netname      string        `json:"netname"`
				Notify       string        `json:"notify"`
				Organization struct {
					Address      string `json:"address"`
					AdminC       string `json:"admin_c"`
					Country      string `json:"country"`
					Descr        string `json:"descr"`
					Email        string `json:"email"`
					FaxNo        string `json:"fax_no"`
					LastModified string `json:"last_modified"`
					MntBy        string `json:"mnt_by"`
					MntRef       string `json:"mnt_ref"`
					Notify       string `json:"notify"`
					Org          string `json:"org"`
					OrgName      string `json:"org_name"`
					Organization string `json:"organization"`
					Phone        string `json:"phone"`
					Source       string `json:"source"`
					TechC        string `json:"tech_c"`
				} `json:"organization"`
				Person []interface{} `json:"person"`
				Role   []interface{} `json:"role"`
				Source string        `json:"source"`
				Status string        `json:"status"`
				TechC  string        `json:"tech_c"`
			} `json:"2017-05-18"`
		} `json:"whois,omitempty"`
	} `json:"matches"`
	Facets struct {
	} `json:"facets"`
}

type ShodanInfo struct {
	Matches []struct {
		Hash   int `json:"hash"`
		Shodan struct {
			Region  string `json:"region"`
			Crawler string `json:"crawler"`
			Options struct {
			} `json:"options"`
			Module string `json:"module"`
			ID     string `json:"id"`
		} `json:"_shodan,omitempty"`
		Product string `json:"product"`
		HTTP    struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       interface{}   `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			Robots      interface{}   `json:"robots"`
			Server      string        `json:"server"`
			HeadersHash int           `json:"headers_hash"`
			Host        string        `json:"host"`
			HTML        string        `json:"html"`
			Location    string        `json:"location"`
			Components  struct {
			} `json:"components"`
			HTMLHash        int         `json:"html_hash"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Os        interface{} `json:"os"`
		Timestamp string      `json:"timestamp"`
		Isp       string      `json:"isp"`
		Cpe23     []string    `json:"cpe23,omitempty"`
		Cpe       []string    `json:"cpe,omitempty"`
		Transport string      `json:"transport"`
		Asn       string      `json:"asn"`
		Hostnames []string    `json:"hostnames"`
		Location  struct {
			City        string      `json:"city"`
			RegionCode  string      `json:"region_code"`
			AreaCode    interface{} `json:"area_code"`
			Longitude   float64     `json:"longitude"`
			Latitude    float64     `json:"latitude"`
			CountryCode string      `json:"country_code"`
			CountryName string      `json:"country_name"`
		} `json:"location"`
		Version string   `json:"version,omitempty"`
		IP      int      `json:"ip"`
		Domains []string `json:"domains"`
		Org     string   `json:"org"`
		Data    string   `json:"data"`
		Port    int      `json:"port"`
		IPStr   string   `json:"ip_str"`
		Cloud   struct {
			Region   string `json:"region"`
			Service  string `json:"service"`
			Provider string `json:"provider"`
		} `json:"cloud,omitempty"`
		Tags    []string `json:"tags,omitempty"`
		Info    string   `json:"info,omitempty"`
		Shodan0 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan1 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP0 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP1 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan2 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan3 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan4 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP2 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP3 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan5 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan6 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP4 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan7 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan8 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan9 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan10 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan11 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP5 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan12 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP6 struct {
			Status      int           `json:"status"`
			RobotsHash  int           `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash int           `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      string        `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				JQuery struct {
					Categories []string `json:"categories"`
				} `json:"jQuery"`
				Cdnjs struct {
					Categories []string `json:"categories"`
				} `json:"cdnjs"`
				Bootstrap struct {
					Categories []string `json:"categories"`
				} `json:"Bootstrap"`
				Popper struct {
					Categories []string `json:"categories"`
				} `json:"Popper"`
				JQueryCDN struct {
					Categories []string `json:"categories"`
				} `json:"jQuery CDN"`
				GoogleFontAPI struct {
					Categories []string `json:"categories"`
				} `json:"Google Font API"`
				Cloudflare struct {
					Categories []string `json:"categories"`
				} `json:"Cloudflare"`
				Liveinternet struct {
					Categories []string `json:"categories"`
				} `json:"Liveinternet"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         string      `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Ssl struct {
			ChainSha256 []string `json:"chain_sha256"`
			Jarm        string   `json:"jarm"`
			Chain       []string `json:"chain"`
			Dhparams    struct {
				Prime     string `json:"prime"`
				PublicKey string `json:"public_key"`
				Bits      int    `json:"bits"`
				Generator int    `json:"generator"`
			} `json:"dhparams"`
			Versions      []string      `json:"versions"`
			AcceptableCas []interface{} `json:"acceptable_cas"`
			Tlsext        []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"tlsext"`
			Ja3S interface{} `json:"ja3s"`
			Cert struct {
				SigAlg     string `json:"sig_alg"`
				Issued     string `json:"issued"`
				Expires    string `json:"expires"`
				Expired    bool   `json:"expired"`
				Version    int    `json:"version"`
				Extensions []struct {
					Critical bool   `json:"critical,omitempty"`
					Data     string `json:"data"`
					Name     string `json:"name"`
				} `json:"extensions"`
				Fingerprint struct {
					Sha256 string `json:"sha256"`
					Sha1   string `json:"sha1"`
				} `json:"fingerprint"`
				Serial  int64 `json:"serial"`
				Subject struct {
					Cn string `json:"CN"`
				} `json:"subject"`
				Pubkey struct {
					Type string `json:"type"`
					Bits int    `json:"bits"`
				} `json:"pubkey"`
				Issuer struct {
					C  string `json:"C"`
					Cn string `json:"CN"`
					O  string `json:"O"`
				} `json:"issuer"`
			} `json:"cert"`
			Cipher struct {
				Version string `json:"version"`
				Bits    int    `json:"bits"`
				Name    string `json:"name"`
			} `json:"cipher"`
			Trust struct {
				Revoked bool `json:"revoked"`
				Browser struct {
					Mozilla   bool `json:"mozilla"`
					Apple     bool `json:"apple"`
					Microsoft bool `json:"microsoft"`
				} `json:"browser"`
			} `json:"trust"`
			HandshakeStates []string `json:"handshake_states"`
			Alpn            []string `json:"alpn"`
			Ocsp            struct {
			} `json:"ocsp"`
		} `json:"ssl,omitempty"`
		Shodan13 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan14 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP7 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan15 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan16 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP8 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan17 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan18 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan19 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP9 struct {
			Status          int           `json:"status"`
			RobotsHash      interface{}   `json:"robots_hash"`
			Redirects       []interface{} `json:"redirects"`
			Securitytxt     interface{}   `json:"securitytxt"`
			Title           interface{}   `json:"title"`
			SitemapHash     interface{}   `json:"sitemap_hash"`
			Robots          interface{}   `json:"robots"`
			Server          string        `json:"server"`
			HeadersHash     int           `json:"headers_hash"`
			Host            string        `json:"host"`
			HTML            string        `json:"html"`
			Location        string        `json:"location"`
			HTMLHash        int           `json:"html_hash"`
			Sitemap         interface{}   `json:"sitemap"`
			SecuritytxtHash interface{}   `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan20 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan21 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan22 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan23 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan24 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan25 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan26 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan27 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP10 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP11 struct {
			Status     int         `json:"status"`
			RobotsHash interface{} `json:"robots_hash"`
			Redirects  []struct {
				Host     string `json:"host"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"redirects"`
			Securitytxt     interface{} `json:"securitytxt"`
			Title           interface{} `json:"title"`
			SitemapHash     interface{} `json:"sitemap_hash"`
			Robots          interface{} `json:"robots"`
			Server          string      `json:"server"`
			HeadersHash     int         `json:"headers_hash"`
			Host            string      `json:"host"`
			HTML            string      `json:"html"`
			Location        string      `json:"location"`
			HTMLHash        int         `json:"html_hash"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan28 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Hostname string `json:"hostname"`
				Scan     string `json:"scan"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan29 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP12 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			Waf         string        `json:"waf"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int       `json:"headers_hash"`
			Host        string    `json:"host"`
			HTML        time.Time `json:"html"`
			Location    string    `json:"location"`
			Components  struct {
				GoogleTagManager struct {
					Categories []string `json:"categories"`
				} `json:"Google Tag Manager"`
				Pingdom struct {
					Categories []string `json:"categories"`
				} `json:"Pingdom"`
			} `json:"components"`
			HTMLHash        int         `json:"html_hash"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan30 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Hostname string `json:"hostname"`
				Scan     string `json:"scan"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan31 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan32 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP13 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan33 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan34 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan35 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan36 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP14 struct {
			Status          int           `json:"status"`
			RobotsHash      interface{}   `json:"robots_hash"`
			Redirects       []interface{} `json:"redirects"`
			Securitytxt     interface{}   `json:"securitytxt"`
			Title           interface{}   `json:"title"`
			SitemapHash     interface{}   `json:"sitemap_hash"`
			Robots          interface{}   `json:"robots"`
			Server          string        `json:"server"`
			HeadersHash     int           `json:"headers_hash"`
			Host            string        `json:"host"`
			HTML            string        `json:"html"`
			Location        string        `json:"location"`
			HTMLHash        int           `json:"html_hash"`
			Sitemap         interface{}   `json:"sitemap"`
			SecuritytxtHash interface{}   `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP15 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan37 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP16 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP17 struct {
			Status      int           `json:"status"`
			RobotsHash  int           `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      string        `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan38 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Referrer string `json:"referrer"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan39 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan40 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan41 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP18 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				Prototype struct {
					Categories []string `json:"categories"`
				} `json:"Prototype"`
				ExtJS struct {
					Categories []string `json:"categories"`
				} `json:"ExtJS"`
				SynologyDiskStation struct {
					Categories []string `json:"categories"`
				} `json:"Synology DiskStation"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		SynologyDsm struct {
			Hostname string `json:"hostname"`
		} `json:"synology_dsm,omitempty"`
		Shodan42 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP19 struct {
			Status      int           `json:"status"`
			RobotsHash  int           `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      string        `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				Modernizr struct {
					Categories []string `json:"categories"`
				} `json:"Modernizr"`
				WordPress struct {
					Categories []string `json:"categories"`
				} `json:"WordPress"`
				MySQL struct {
					Categories []string `json:"categories"`
				} `json:"MySQL"`
				W3TotalCache struct {
					Categories []string `json:"categories"`
				} `json:"W3 Total Cache"`
				YoastSEO struct {
					Categories []string `json:"categories"`
				} `json:"Yoast SEO"`
				Php struct {
					Categories []string `json:"categories"`
				} `json:"PHP"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan43 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP20 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				Prototype struct {
					Categories []string `json:"categories"`
				} `json:"Prototype"`
				ExtJS struct {
					Categories []string `json:"categories"`
				} `json:"ExtJS"`
				SynologyDiskStation struct {
					Categories []string `json:"categories"`
				} `json:"Synology DiskStation"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan44 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP21 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				AnimateCSS struct {
					Categories []string `json:"categories"`
				} `json:"animate.css"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP22 struct {
			Status      int           `json:"status"`
			RobotsHash  int           `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      string        `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan45 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan46 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan47 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Referrer string `json:"referrer"`
				Hostname string `json:"hostname"`
				Scan     string `json:"scan"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan48 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan49 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP23 struct {
			Status          int           `json:"status"`
			RobotsHash      interface{}   `json:"robots_hash"`
			Redirects       []interface{} `json:"redirects"`
			Securitytxt     interface{}   `json:"securitytxt"`
			Title           interface{}   `json:"title"`
			SitemapHash     interface{}   `json:"sitemap_hash"`
			Robots          interface{}   `json:"robots"`
			Server          string        `json:"server"`
			HeadersHash     int           `json:"headers_hash"`
			Host            string        `json:"host"`
			HTML            string        `json:"html"`
			Location        string        `json:"location"`
			HTMLHash        int           `json:"html_hash"`
			Sitemap         interface{}   `json:"sitemap"`
			SecuritytxtHash interface{}   `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan50 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Vulns struct {
			CVE20198943 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-8943"`
			CVE202028039 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28039"`
			CVE202028038 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28038"`
			CVE202036326 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-36326"`
			CVE201917673 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17673"`
			CVE201917672 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17672"`
			CVE201917671 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17671"`
			CVE20181000773 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-1000773"`
			CVE202028037 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28037"`
			CVE202028036 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28036"`
			CVE201917675 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17675"`
			CVE201917674 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17674"`
			CVE201819296 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-19296"`
			CVE20199787 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-9787"`
			CVE20198942 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-8942"`
			CVE201820152 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20152"`
			CVE201917669 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17669"`
			CVE201812895 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-12895"`
			CVE20204050 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-4050"`
			CVE202144223 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-44223"`
			CVE202028033 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28033"`
			CVE202011028 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11028"`
			CVE202011029 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11029"`
			CVE201916218 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16218"`
			CVE201916219 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16219"`
			CVE201820151 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20151"`
			CVE202011025 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11025"`
			CVE202011026 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11026"`
			CVE202011027 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11027"`
			CVE201920043 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-20043"`
			CVE201917670 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-17670"`
			CVE201810101 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-10101"`
			CVE201810100 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-10100"`
			CVE201920042 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-20042"`
			CVE201810102 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-10102"`
			CVE201717094 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2017-17094"`
			CVE201920041 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-20041"`
			CVE201717091 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2017-17091"`
			CVE201717093 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2017-17093"`
			CVE201717092 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2017-17092"`
			CVE202025286 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-25286"`
			CVE202028035 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28035"`
			CVE201916217 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16217"`
			CVE202221663 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-21663"`
			CVE202221662 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-21662"`
			CVE202221661 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-21661"`
			CVE202028034 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28034"`
			CVE201820150 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20150"`
			CVE202028040 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28040"`
			CVE202221664 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-21664"`
			CVE202011030 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-11030"`
			CVE202028032 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-28032"`
			CVE20204049 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-4049"`
			CVE20204048 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-4048"`
			CVE20204047 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-4047"`
			CVE20204046 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2020-4046"`
			CVE201820148 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20148"`
			CVE201820149 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20149"`
			CVE201916223 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16223"`
			CVE201916222 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16222"`
			CVE201916221 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16221"`
			CVE201916220 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16220"`
			CVE201916780 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16780"`
			CVE201916781 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2019-16781"`
			CVE20171000600 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2017-1000600"`
			CVE201820147 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20147"`
			CVE201820153 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2018-20153"`
			CVE202129450 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-29450"`
		} `json:"vulns,omitempty"`
		HTTP24 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				JQuery struct {
					Categories []string `json:"categories"`
				} `json:"jQuery"`
				AllInOneSEOPack struct {
					Categories []string `json:"categories"`
				} `json:"All in One SEO Pack"`
				SWFObject struct {
					Categories []string `json:"categories"`
				} `json:"SWFObject"`
				Twitter struct {
					Categories []string `json:"categories"`
				} `json:"Twitter"`
				JQueryMigrate struct {
					Categories []string `json:"categories"`
				} `json:"jQuery Migrate"`
				YouTube struct {
					Categories []string `json:"categories"`
				} `json:"YouTube"`
				WordPress struct {
					Categories []string `json:"categories"`
				} `json:"WordPress"`
				MySQL struct {
					Categories []string `json:"categories"`
				} `json:"MySQL"`
				Php struct {
					Categories []string `json:"categories"`
				} `json:"PHP"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan51 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Hostname string `json:"hostname"`
				Scan     string `json:"scan"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan52 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Vulns0 struct {
			CVE202231626 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-31626"`
			CVE202231625 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2022-31625"`
			CVE202121707 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21707"`
			CVE202121706 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21706"`
			CVE202121705 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21705"`
			CVE202121704 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21704"`
			CVE202121703 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21703"`
			CVE202121708 struct {
				Verified   bool     `json:"verified"`
				References []string `json:"references"`
				Cvss       float64  `json:"cvss"`
				Summary    string   `json:"summary"`
			} `json:"CVE-2021-21708"`
		} `json:"vulns,omitempty"`
		HTTP25 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP26 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan53 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan54 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan55 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan56 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan57 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP27 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan58 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan59 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		Shodan60 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
				Hostname string `json:"hostname"`
				Scan     string `json:"scan"`
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
		HTTP28 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
				NodeJs struct {
					Categories []string `json:"categories"`
				} `json:"Node.js"`
				NuxtJs struct {
					Categories []string `json:"categories"`
				} `json:"Nuxt.js"`
				VueJs struct {
					Categories []string `json:"categories"`
				} `json:"Vue.js"`
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		HTTP29 struct {
			Status      int           `json:"status"`
			RobotsHash  interface{}   `json:"robots_hash"`
			Redirects   []interface{} `json:"redirects"`
			Securitytxt interface{}   `json:"securitytxt"`
			Title       string        `json:"title"`
			SitemapHash interface{}   `json:"sitemap_hash"`
			HTMLHash    int           `json:"html_hash"`
			Robots      interface{}   `json:"robots"`
			Favicon     struct {
				Hash     int    `json:"hash"`
				Data     string `json:"data"`
				Location string `json:"location"`
			} `json:"favicon"`
			HeadersHash int    `json:"headers_hash"`
			Host        string `json:"host"`
			HTML        string `json:"html"`
			Location    string `json:"location"`
			Components  struct {
			} `json:"components"`
			Server          string      `json:"server"`
			Sitemap         interface{} `json:"sitemap"`
			SecuritytxtHash interface{} `json:"securitytxt_hash"`
		} `json:"http,omitempty"`
		Shodan61 struct {
			Region  string `json:"region"`
			Ptr     bool   `json:"ptr"`
			Module  string `json:"module"`
			ID      string `json:"id"`
			Options struct {
			} `json:"options"`
			Crawler string `json:"crawler"`
		} `json:"_shodan,omitempty"`
	} `json:"matches"`
	Total int `json:"total"`
}
