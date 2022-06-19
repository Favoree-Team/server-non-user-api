package controller

import "github.com/Favoree-Team/server-non-user-api/service"

type logReqTransController struct {
	logReqTransService service.LogReqTransService
}

func NewLogReqTransController(logReqTransService service.LogReqTransService) *logReqTransController {
	return &logReqTransController{logReqTransService: logReqTransService}
}
