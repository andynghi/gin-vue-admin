package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"testing"
)

type PageInfoTest struct {
	PageInfo request.PageInfo
	Name     string
}

func TestVerify(t *testing.T) {
	PageInfoVerify := Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}, "Name": {NotEmpty()}}
	var testInfo PageInfoTest
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 0
	testInfo.PageInfo.PageSize = 0
	err := Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("Verification failed, failed to capture 0 value")
	}
	testInfo.Name = ""
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err == nil {
		t.Error("Verification failed, failed to detect that name is empty")
	}
	testInfo.Name = "test"
	testInfo.PageInfo.Page = 1
	testInfo.PageInfo.PageSize = 10
	err = Verify(testInfo, PageInfoVerify)
	if err != nil {
		t.Error("Verification failed and failed to pass the test normally")
	}
}
