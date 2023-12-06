package routes

import (
	"hestia/api/logger"
	"hestia/api/middleware"
	"hestia/api/models"
	"hestia/api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func technicalFilesGet(ctx *gin.Context) {
	type FilterOptions struct {
		//models.TechnicalFile
		models.LimitOffset
	}
	var filterOptions FilterOptions
	if err := ctx.ShouldBindJSON(&filterOptions); err != nil {

		logger.Error.Println("Error while binding JSON: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Input validation failed, check documentation for correct input types",
		})
		return
	}

	db, err := utils.GetDbPoolConn()
	if err != nil {
		logger.Error.Println("Error while connecting to DB: ", err)
	}

	var technicalFiles []models.TechnicalFile
	rows, err := db.Query(
		ctx.Request.Context(),
		// TODO: Query
		"SELECT id FROM itv.technical_file WHERE company_id = $1 LIMIT $2 OFFSET $3",
		ctx.GetHeader("X-Company-Id"),
		filterOptions.Limit,
		filterOptions.Offset,
	)

	if err != nil {
		logger.Error.Println("Error while querying technical files: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while fetching technical files",
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var tf models.TechnicalFile
		err := rows.Scan(&tf.Id)
		if err != nil {
			logger.Error.Println("Error while scanning technical files: ", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: "Error while fetching technical files",
			})
			return
		}
		technicalFiles = append(technicalFiles, tf)
	}

	ctx.JSON(http.StatusOK, technicalFiles)
}

func ITVRoutes(r *gin.Engine) {
	itv := r.Group("/itv", middleware.BearerAuthenticate(), middleware.CompanyId())

	// /itv/technicalfiles
	technicalfiles := itv.Group("/technicalfiles")
	technicalfiles.GET("", technicalFilesGet)

}
