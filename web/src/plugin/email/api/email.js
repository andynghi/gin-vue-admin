import service from '@/utils/request'
// @Tags System
// @Summary Send test email
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Sent successfully"}"
// @Router /email/emailTest [post]
export const emailTest = (data) => {
  return service({
    url: '/email/emailTest',
    method: 'post',
    data
  })
}

// @Tags System
// @Summary Send email
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body email_response.Email true "Required parameters for sending emails"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Sent successfully"}"
// @Router /email/sendEmail [post]
export const sendEmail = (data) => {
  return service({
    url: '/email/sendEmail',
    method: 'post',
    data
  })
}

