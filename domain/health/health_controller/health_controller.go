package health_controller

import (
	"net/http"

	"github.com/cqlcorp/gocms/init/service"
	"github.com/cqlcorp/gocms/routes"
	"github.com/cqlcorp/gocms/utility/errors"
	"github.com/gin-gonic/gin"
)

type HealthController struct {
	routes       *routes.Routes
	serviceGroup *service.ServicesGroup
}

func DefaultHealthController(routes *routes.Routes, serviceGroup *service.ServicesGroup) *HealthController {
	hc := &HealthController{
		routes:       routes,
		serviceGroup: serviceGroup,
	}

	hc.Default()
	return hc
}

func (hc *HealthController) Default() {
	hc.routes.Root.GET("/healthy", hc.health)
}

/**
* @api {get} /healthy Service Health Status
* @apiDescription Used to verify that the services are up and running.
* @apiName GetHealthy
* @apiGroup Utility
 */
func (hc *HealthController) health(c *gin.Context) {

	ok, _ := hc.serviceGroup.HealthService.GetHealthStatus()

	if !ok {

		msg := "Service is having health issues"

		// todo add a admin key to get more detailed reports
		//for i, issue := range context {
		//	if i == 0 {
		//		msg = fmt.Sprintf("%v %v", msg, issue)
		//	} else {
		//		msg = fmt.Sprintf("%v, %v", msg, issue)
		//	}
		//}

		errors.Response(c, http.StatusInternalServerError, msg, nil)
		return
	}

	c.Status(http.StatusOK)
}
