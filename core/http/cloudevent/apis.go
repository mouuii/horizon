package cloudevent

import (
	"fmt"
	"strings"

	"github.com/horizoncd/horizon/core/common"
	"github.com/horizoncd/horizon/core/controller/cloudevent"
	"github.com/horizoncd/horizon/pkg/server/response"
	"github.com/horizoncd/horizon/pkg/util/log"

	"github.com/gin-gonic/gin"
)

type API struct {
	cloudEventCtl cloudevent.Controller
}

func NewAPI(cloudEventCtl cloudevent.Controller) *API {
	return &API{
		cloudEventCtl: cloudEventCtl,
	}
}

func (a *API) CloudEvent(c *gin.Context) {
	// "Ce-Source": "/apis/tekton.dev/v1beta1/namespaces/default/taskruns/curl-run-6gplk"
	// Ce-Source is the selfLink in k8s resource
	// if resource is not a pipelineRun, return
	if !strings.Contains(c.GetHeader("Ce-Source"), "pipelineruns") {
		response.Success(c)
		return
	}

	var wpr *cloudevent.WrappedPipelineRun
	if err := c.ShouldBindJSON(&wpr); err != nil {
		response.AbortWithRequestError(c, common.InvalidRequestBody,
			fmt.Sprintf("request body is invalid, err: %v", err))
		return
	}
	if !wpr.IsFinished() {
		response.Success(c)
		return
	}

	if err := a.cloudEventCtl.CloudEvent(c, wpr); err != nil {
		log.Errorf(c, "failed to handle cloud event, pipelinerun name: %s, err: %v",
			wpr.PipelineRun.Name, err)
		response.AbortWithError(c, err)
		return
	}

	response.Success(c)
}
