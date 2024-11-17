package core

import (
	"context"
	"rental-assistant/internal/utils"
	"rental-assistant/internal/view"
)

func GetNearRegionInfo(_ context.Context, req *view.GetNearRegionInfoReq) (resp *view.GetNearRegionInfoResp) {
	resp = new(view.GetNearRegionInfoResp)
	if err := view.ParseGetNearRegionInfo(req); err != nil {
		resp.AddErr(err)
		return
	}
	identify := req.Province + req.City + req.District + req.Area
	centerInfo, ok := CenterMap[identify]
	if !ok {
		resp.AddErr(view.ErrMsgCenterRegionNotExist)
		return
	}
	reqType, validRentHouseInfos, validEntertainmentInfos := RegionType(req.Type), make([]*RentHouseInfo, 0), make([]*EntertainmentInfo, 0)
	if reqType == RegionTypeAll || reqType == RegionTypeRentHouse {
		rentHouseInfos, rentExist := AreaNearRentHouseMap[identify]
		if !rentExist {
			resp.AddErr(view.ErrMsgCenterRegionNotExist)
			return
		}
		validRentHouseInfos = make([]*RentHouseInfo, 0)
		for _, value := range rentHouseInfos {
			distance := utils.Haversine(centerInfo.Latitude, centerInfo.Longitude, value.Latitude, value.Longitude)
			if req.Distance != 0 && distance > req.Distance {
				continue
			}
			if req.RentHousePrice != 0 && value.Price > req.RentHousePrice {
				continue
			}
			validRentHouseInfos = append(validRentHouseInfos, value)
		}
	}
	if reqType == RegionTypeAll || reqType == RegionTypeEntertainment {
		entertainmentInfos, entertainmentExist := AreaNearEntertainmentMap[identify]
		if !entertainmentExist {
			resp.AddErr(view.ErrMsgCenterRegionNotExist)
			return
		}
		validEntertainmentInfos = make([]*EntertainmentInfo, 0)
		for _, value := range entertainmentInfos {
			distance := utils.Haversine(centerInfo.Latitude, centerInfo.Longitude, value.Latitude, value.Longitude)
			if req.Distance != 0 && distance > req.Distance {
				continue
			}
			validEntertainmentInfos = append(validEntertainmentInfos, value)
		}
	}
	resp.Data = transferNearRegionInfo(validRentHouseInfos, validEntertainmentInfos)
	return
}

func transferNearRegionInfo(inRentList []*RentHouseInfo, inEntertainmentList []*EntertainmentInfo) (out *view.NearRegionInfo) {
	out = &view.NearRegionInfo{
		RentHouseInfos:     make([]*view.RentHouseInfo, 0, len(inRentList)),
		EntertainmentInfos: make([]*view.EntertainmentInfo, 0, len(inEntertainmentList)),
	}
	for _, value := range inRentList {
		out.RentHouseInfos = append(out.RentHouseInfos, &view.RentHouseInfo{
			Name:        value.Name,
			Type:        int(value.Type),
			Price:       value.Price,
			Description: value.Description,
			Latitude:    value.Latitude,
			Longitude:   value.Longitude,
		})
	}
	for _, value := range inEntertainmentList {
		out.EntertainmentInfos = append(out.EntertainmentInfos, &view.EntertainmentInfo{
			Name:        value.Name,
			Type:        int(value.Type),
			Description: value.Description,
			Latitude:    value.Latitude,
			Longitude:   value.Longitude,
		})
	}
	return
}
