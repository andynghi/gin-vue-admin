import service from '@/utils/request'
// @Tags FileUploadAndDownload
// @Summary Paginated file list
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "Get the file list by page"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Get successful"}"
// @Router /fileUploadAndDownload/getFileList [post]
export const getFileList = (data) => {
  return service({
    url: '/fileUploadAndDownload/getFileList',
    method: 'post',
    data
  })
}

// @Tags FileUploadAndDownload
// @Summary delete file
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body dbModel.FileUploadAndDownload true "Just pass in the id in the file"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"Return successful"}"
// @Router /fileUploadAndDownload/deleteFile [post]
export const deleteFile = (data) => {
  return service({
    url: '/fileUploadAndDownload/deleteFile',
    method: 'post',
    data
  })
}

/**
 * Edit file name or remarks
 * @param data
 * @returns {*}
 */
export const editFileName = (data) => {
  return service({
    url: '/fileUploadAndDownload/editFileName',
    method: 'post',
    data
  })
}
