package xrpc

import (
	"github.com/coder2m/component/pkg/xcode"
)

const (
	//system error
	MysqlErr = iota + 100

	//业务错误 高于 10000
	EmptyData = iota + 10000
	ValidationErrCode
)

var (
	systemErrDepot = []xcode.CodeInfo{
		{xcode.SystemType, MysqlErr, "mysql 错误"},
	}
	businessErrDepot = []xcode.CodeInfo{
		{xcode.BusinessType, EmptyData, "Empty Data"},
        {xcode.BusinessType, ValidationErrCode, "validation error"},
	}
)

func init() {
	xcode.CodeAdds(systemErrDepot)
	xcode.CodeAdds(businessErrDepot)
}
