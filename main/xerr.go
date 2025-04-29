package main

import (
	"fmt"
	"github.com/aobco/log"
	"xsky.com/ocpf/pkg/error/biz_error"
	"xsky.com/ocpf/pkg/error/xerr"
)

func main() {

	bizErr := fmt.Errorf("test error")
	trace := xerr.Trace(bizErr)
	log.Errorf("%+v", trace)
	//trace4 := getTrace4()
	//log.Errorf("%+v", trace4)
	//println("=====================================================================")
	//log.Errorf("%v", trace4)
	//println("=====================================================================")
	//log.Errorf("error: %s", trace4.Error())
	//httpStatus, code, msg := xerr.ParseBizError(trace4)
	//println("=====================================================================")
	//log.Errorf("httpStatus: %d, code: %s, msg: %s", httpStatus, code, msg)
	//err := fmt.Errorf("test error")
	//httpStatus, code, msg := xerr.ParseBizError(err)
	//println("=====================================================================")
	//log.Infof("httpStatus: %d, code: %s, msg: %s", httpStatus, code, msg)
}

func getTrace4() error {
	trace4 := xerr.Trace(getTrace3())
	return trace4
}

func getTrace3() error {
	trace3 := xerr.Trace(getTrace2())
	return trace3
}

func getTrace2() error {
	trace2 := xerr.Trace(getTrace1())
	return trace2
}

func getTrace1() error {
	trace1 := xerr.Trace(getBizErr())
	return trace1
}

func getBizErr() error {
	//bizErr := xerr.BizWrap(errors.New("postgres io timeout"), biz_error.ExecuteSqlError, "insert users error")
	bizErr := xerr.DefaultBizWrap(biz_error.ResourceNotExists, "aaa")
	return bizErr
}
