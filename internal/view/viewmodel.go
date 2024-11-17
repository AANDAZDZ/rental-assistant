package view

type CommonReq struct {
	AuthKey string `form:"auth_key" json:"auth_key"`
}

type CommonResp struct {
	ErrNum int    `json:"err_num"`
	ErrMsg string `json:"err_msg"`
}

type GetNearRegionInfoReq struct {
	CommonReq
	Province       string `form:"province" json:"province"`
	City           string `form:"city" json:"city"`
	District       string `form:"district" json:"district"`
	Area           string `form:"area" json:"area"`
	Type           int    `form:"type" json:"type"`
	Distance       int    `form:"distance" json:"distance"`
	RentHousePrice int    `form:"price" json:"price"`
}

type GetNearRegionInfoResp struct {
	CommonResp
	Data *NearRegionInfo `json:"data,omitempty"`
}

type NearRegionInfo struct {
	RentHouseInfos     []*RentHouseInfo     `json:"rent_house_infos,omitempty"`
	EntertainmentInfos []*EntertainmentInfo `json:"entertainment_infos,omitempty"`
}

type RentHouseInfo struct {
	Name        string
	Type        int
	Price       int
	Description string
	Latitude    float64
	Longitude   float64
}

type EntertainmentInfo struct {
	Name        string
	Type        int
	Description string
	Latitude    float64
	Longitude   float64
}
