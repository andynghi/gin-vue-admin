package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FileUploadAndDownloadApi struct{}

// UploadFile
// @Tags      ExaFileUploadAndDownload
// @Summary Upload file example
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param file formData file true "Upload file example"
// @Success 200 {object} response.Response{data=exampleRes.ExaFileResponse,msg=string} "Upload file example, return including file details"
// @Router    /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("Failed to receive file!", zap.Error(err))
		response.FailWithMessage("Failed to receive file", c)
		return
	}
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave) // Get the file path after uploading the file
	if err != nil {
		global.GVA_LOG.Error("Failed to modify database link!", zap.Error(err))
		response.FailWithMessage("Failed to modify the database link", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "Upload successful", c)
}

// EditFileName edit file name or remarks
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = fileUploadAndDownloadService.EditFileName(file)
	if err != nil {
		global.GVA_LOG.Error("Editing failed!", zap.Error(err))
		response.FailWithMessage("Editing failed", c)
		return
	}
	response.OkWithMessage("Editing successfully", c)
}

// DeleteFile
// @Tags      ExaFileUploadAndDownload
// @Summary delete file
// @Security  ApiKeyAuth
// @Produce   application/json
// @Param data body example.ExaFileUploadAndDownload true "Just pass in the id in the file"
// @Success 200 {object} response.Response{msg=string} "Delete file"
// @Router    /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	err := c.ShouldBindJSON(&file)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// GetFileList
// @Tags      ExaFileUploadAndDownload
// @Summary Paginated file list
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body request.PageInfo true "Page number, size of each page"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Paging file list, return includes list, total number, page number, number of pages per page"
// @Router    /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "Get successful", c)
}
