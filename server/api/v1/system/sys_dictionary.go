package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictionaryApi struct{}

// CreateSysDictionary
// @Tags      SysDictionary
// @Summary Create SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionary true "SysDictionary model"
// @Success 200 {object} response.Response{msg=string} "Create SysDictionary"
// @Router    /sysDictionary/createSysDictionary [post]
func (s *DictionaryApi) CreateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.CreateSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("Creation failed!", zap.Error(err))
		response.FailWithMessage("Creation failed", c)
		return
	}
	response.OkWithMessage("Created successfully", c)
}

// DeleteSysDictionary
// @Tags      SysDictionary
// @Summary delete SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionary true "SysDictionary model"
// @Success 200 {object} response.Response{msg=string} "Delete SysDictionary"
// @Router    /sysDictionary/deleteSysDictionary [delete]
func (s *DictionaryApi) DeleteSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.DeleteSysDictionary(dictionary)
	if err != nil {
		global.GVA_LOG.Error("Deletion failed!", zap.Error(err))
		response.FailWithMessage("Deletion failed", c)
		return
	}
	response.OkWithMessage("Deletion successful", c)
}

// UpdateSysDictionary
// @Tags      SysDictionary
// @Summary Update SysDictionary
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data body system.SysDictionary true "SysDictionary model"
// @Success 200 {object} response.Response{msg=string} "Update SysDictionary"
// @Router    /sysDictionary/updateSysDictionary [put]
func (s *DictionaryApi) UpdateSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindJSON(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = dictionaryService.UpdateSysDictionary(&dictionary)
	if err != nil {
		global.GVA_LOG.Error("Update failed!", zap.Error(err))
		response.FailWithMessage("Update failed", c)
		return
	}
	response.OkWithMessage("Update successful", c)
}

// FindSysDictionary
// @Tags      SysDictionary
// @Summary Query SysDictionary with id
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param data query system.SysDictionary true "ID or dictionary name"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "Query SysDictionary with id"
// @Router    /sysDictionary/findSysDictionary [get]
func (s *DictionaryApi) FindSysDictionary(c *gin.Context) {
	var dictionary system.SysDictionary
	err := c.ShouldBindQuery(&dictionary)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	sysDictionary, err := dictionaryService.GetSysDictionary(dictionary.Type, dictionary.ID, dictionary.Status)
	if err != nil {
		global.GVA_LOG.Error("The dictionary was not created or opened!", zap.Error(err))
		response.FailWithMessage("The dictionary was not created or opened", c)
		return
	}
	response.OkWithDetailed(gin.H{"resysDictionary": sysDictionary}, "Query successful", c)
}

// GetSysDictionaryList
// @Tags      SysDictionary
// @Summary Get the SysDictionary list by pagination
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "Get the SysDictionary list by paging, and return the list, total number, page number, and number per page"
// @Router    /sysDictionary/getSysDictionaryList [get]
func (s *DictionaryApi) GetSysDictionaryList(c *gin.Context) {
	list, err := dictionaryService.GetSysDictionaryInfoList()
	if err != nil {
		global.GVA_LOG.Error("Acquisition failed!", zap.Error(err))
		response.FailWithMessage("Failed to obtain", c)
		return
	}
	response.OkWithDetailed(list, "Get successful", c)
}
