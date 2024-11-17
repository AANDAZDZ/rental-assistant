package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rental-assistant/internal/core"
	"rental-assistant/internal/view"
)

func GetNearRegionInfo(ctx *gin.Context) {
	req, resp := new(view.GetNearRegionInfoReq), new(view.GetNearRegionInfoResp)
	if err := ctx.ShouldBind(req); err != nil {
		resp.AddErr(view.ErrMsgParamVerifyFailed)
		ctx.JSON(http.StatusOK, resp)
		return
	}
	resp = core.GetNearRegionInfo(ctx, req)
	ctx.JSON(http.StatusOK, resp)
	return
}
