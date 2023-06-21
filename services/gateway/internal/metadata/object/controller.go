package object

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/arvanStorage/pkg/grpcext"
	metadataApi "github.com/h-varmazyar/arvanStorage/services/metadata/api"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var (
	controller *Controller
)

type Controller struct {
	objectService metadataApi.ObjectServiceClient
	logger        *log.Logger
	*gin.Engine
}

func NewController(logger *log.Logger, engine *gin.Engine, metadataAddress string) *Controller {
	if controller == nil {
		metadataConn := grpcext.NewConnection(metadataAddress)
		controller = &Controller{
			objectService: metadataApi.NewObjectServiceClient(metadataConn),
			logger:        logger,
			Engine:        engine,
		}
	}
	return controller
}

func (c *Controller) RegisterRoutes(router *gin.RouterGroup) {
	objectsRouter := router.Group("/objects")
	uploadRouter := objectsRouter.Group("/upload")
	uploadRouter.POST("/", c.upload)
	uploadRouter.PUT("/:upload_id", c.part)
	uploadRouter.POST("/:upload_id", c.complete)
	uploadRouter.DELETE("/:upload_id", c.abort)
}

func (c *Controller) upload(ctx *gin.Context) {
	newUpload := new(metadataApi.ObjectNewUploadReq)
	err := ctx.ShouldBind(newUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := c.objectService.NewUpload(ctx, newUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) part(ctx *gin.Context) {
	part := new(metadataApi.ObjectUploadPartReq)
	err := ctx.ShouldBind(part)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	part.UploadID = ctx.Param("upload_id")

	_, err = c.objectService.UploadPart(ctx, part)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Controller) complete(ctx *gin.Context) {
	completeUpload := new(metadataApi.ObjectCompleteUploadReq)
	err := ctx.ShouldBind(completeUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	completeUpload.UploadID = ctx.Param("upload_id")

	_, err = c.objectService.CompleteUpload(ctx, completeUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}

func (c *Controller) abort(ctx *gin.Context) {
	abortUpload := new(metadataApi.ObjectAbortUploadReq)
	err := ctx.ShouldBind(abortUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	abortUpload.UploadID = ctx.Param("upload_id")

	_, err = c.objectService.AbortUpload(ctx, abortUpload)
	if err != nil {
		ctx.String(http.StatusBadRequest, err.Error())
		return
	}
	ctx.Status(http.StatusOK)
}
