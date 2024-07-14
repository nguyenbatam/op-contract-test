package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nguyenbatam/op-contract-test/backend/serializers"
	"io"
	"net/http"
	"strconv"
)

func (api *apiHandler) UserUploadSample(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, serializers.HttpResponse{
			Status: serializers.HttpResponseFail,
			Error:  fmt.Sprintf("error when read body :%v", err.Error()),
		})
		return
	}
	var info serializers.DocumentUpload
	err = json.Unmarshal(body, &info)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, serializers.HttpResponse{
			Status: serializers.HttpResponseFail,
			Error:  fmt.Sprintf("error when parsing JSON body:%v", err.Error()),
		})
		return
	}
	document := api.s.UserUploadSample(c.Request.Context(), &info)
	c.AbortWithStatusJSON(http.StatusOK, serializers.HttpResponse{
		Status: serializers.HttpResponseSuccess,
		Data:   document,
	})
}

func (api *apiHandler) GetDocumentInfo(c *gin.Context) {
	documentIdText := c.Param("document_id")
	documentId, err := strconv.Atoi(documentIdText)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, serializers.HttpResponse{
			Status: serializers.HttpResponseFail,
			Error:  fmt.Sprintf("error document_id is not number:%v", err.Error()),
		})
		return
	}
	document := api.s.GetDocumentSample(c.Request.Context(), uint(documentId))
	c.AbortWithStatusJSON(http.StatusOK, serializers.HttpResponse{
		Status: serializers.HttpResponseSuccess,
		Data:   document,
	})
}
