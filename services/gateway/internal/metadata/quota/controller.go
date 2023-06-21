package quota

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/arvanStorage/pkg/grpcext"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
	log "github.com/sirupsen/logrus"
)

var (
	controller *Controller
)

type Controller struct {
	quotaService metadataApi.QuotaServiceClient
	logger       *log.Logger
	*gin.Engine
}

func NewController(logger *log.Logger, engine *gin.Engine, metadataAddress string) *Controller {
	if controller == nil {
		metadataConn := grpcext.NewConnection(metadataAddress)
		controller = &Controller{
			quotaService: metadataApi.NewQuotaServiceClient(metadataConn),
			logger:       logger,
		}
	}
	return controller
}

func (c Controller) RegisterRoutes(router *gin.RouterGroup) {
	_ = router.Group("/quota")
}
