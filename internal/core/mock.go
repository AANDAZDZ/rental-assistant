package core

type CenterInfo struct {
	Latitude  float64
	Longitude float64
}

type RentHouseInfo struct {
	Name        string
	Type        RegionType
	Price       int
	Description string
	Latitude    float64
	Longitude   float64
}

type EntertainmentInfo struct {
	Name        string
	Type        RegionType
	Description string
	Latitude    float64
	Longitude   float64
}

type RegionType int

const (
	RegionTypeAll           RegionType = 0
	RegionTypeRentHouse     RegionType = 1
	RegionTypeEntertainment RegionType = 2
)

var (
	CenterMap = map[string]*CenterInfo{
		"北京市北京市朝阳区360大厦": {
			Latitude:  116.49,
			Longitude: 39.98,
		},
	}
	AreaNearRentHouseMap = map[string][]*RentHouseInfo{
		"北京市北京市朝阳区360大厦": {
			{
				Name:        "青年路22号院",
				Type:        RegionTypeRentHouse,
				Price:       2500,
				Description: "押一付三，民水民电",
				Latitude:    116.52,
				Longitude:   39.93,
			},
			{
				Name:        "望京西园次卧",
				Type:        RegionTypeRentHouse,
				Price:       2000,
				Description: "押一付一，民水商电",
				Latitude:    116.48,
				Longitude:   40.01,
			},
			{
				Name:        "宏源公寓主卧",
				Type:        RegionTypeRentHouse,
				Price:       4000,
				Description: "押一付六，水电减免",
				Latitude:    116.49,
				Longitude:   39.98,
			},
		},
	}
	AreaNearEntertainmentMap = map[string][]*EntertainmentInfo{
		"北京市北京市朝阳区360大厦": {
			{
				Name:        "颐堤港",
				Type:        RegionTypeEntertainment,
				Description: "商场",
				Latitude:    116.49,
				Longitude:   39.97,
			},
			{
				Name:        "阜通",
				Type:        RegionTypeEntertainment,
				Description: "地铁站",
				Latitude:    116.47,
				Longitude:   39.99,
			},
			{
				Name:        "北京奥林匹克公园",
				Type:        RegionTypeEntertainment,
				Description: "公园",
				Latitude:    116.39,
				Longitude:   40.01,
			},
		},
	}
)
