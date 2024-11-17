package view

import (
	"strings"
)

func parseCommon(common CommonReq) (err error) {
	common.AuthKey = strings.TrimSpace(common.AuthKey)
	if common.AuthKey != "rental-assistant" {
		return ErrMsgAuthVerifyFailed
	}
	return
}

func ParseGetNearRegionInfo(req *GetNearRegionInfoReq) (err error) {
	if err = parseCommon(req.CommonReq); err != nil {
		return
	}
	req.Province = strings.TrimSpace(req.Province)
	req.City = strings.TrimSpace(req.City)
	req.District = strings.TrimSpace(req.District)
	req.Area = strings.TrimSpace(req.Area)
	if req.Province == "" {
		return ErrMsgParamVerifyFailed.WithCustomMsg("required province param")
	}
	if req.City == "" {
		return ErrMsgParamVerifyFailed.WithCustomMsg("required city param")
	}
	if req.District == "" {
		return ErrMsgParamVerifyFailed.WithCustomMsg("required district param")
	}
	if req.Area == "" {
		return ErrMsgParamVerifyFailed.WithCustomMsg("required area param")
	}
	if req.Type < 0 {
		return ErrMsgParamVerifyFailed.WithCustomMsg("required type param")
	}
	if req.Distance < 0 {
		return ErrMsgParamVerifyFailed.WithCustomMsg("illegal distance param")
	}
	if req.RentHousePrice < 0 {
		return ErrMsgParamVerifyFailed.WithCustomMsg("illegal price param")
	}
	return
}
