package manager

import "testing"

func TestManager(t *testing.T) {
	wenxiang := &CommonManager{
		Name: "文祥",
	}
	xiaoyun := &MajorManager{
		Name: "晓云",
	}
	yuanlei := &GeneralManager{
		Name: "苑磊",
	}
	wenxiang.SetNext(xiaoyun)
	xiaoyun.SetNext(yuanlei)

	request := Request{
		RequestType:    "请假",
		RequestContent: "小菜请假",
		Number:         2,
	}
	wenxiang.RequestHandler(request)

	request = Request{
		RequestType:    "请假",
		RequestContent: "小菜请假",
		Number:         5,
	}
	wenxiang.RequestHandler(request)

	request = Request{
		RequestType:    "加薪",
		RequestContent: "小菜请求加薪",
		Number:         500,
	}
	wenxiang.RequestHandler(request)

	request = Request{
		RequestType:    "加薪",
		RequestContent: "小菜请求加薪",
		Number:         1000,
	}
	wenxiang.RequestHandler(request)
}
