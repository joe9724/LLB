package controllers

import (
	"github.com/astaxie/beego"
	_"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"github.com/astaxie/beego/logs"
)

type SingleQuery struct{
	Status _Status  `json:"status"`
	Result _Result `json:"result"`
}
type _Status struct{
	Code int `json:"code"`
	Message string  `json:"message"`
	Detail string  `json:"detail"`
}

type _Result struct{
	__Status int `json:"status"`
	Plan int `json:"plan"`
	Today int `json:"today"`
	Yesterday int `json:"yesterday"`
	TaskId int `json:"taskid"`
}
//查看任务配置
type _Task struct {
	Status _Status
	Result _TaskResult  `json:"result"`

}
type _TaskResult struct{
	TaskId int `json:"id"`
	__Status int `json:"status"`
        Name string `json:"name"`
	Plan int `json:"plan"`
	Url string `json:"url"`
	OperateCount int `json:"operateCount"`
	RndOperate bool `json:"rndOperate"`
	StayTime int `json:"stayTime"`
	Source []_Source
	Operation _Operation
	//Mdi _Mdi
	StartTime int64 `json:"startTime"`
	EndTime int64 `json:"endTime"`

}
type _Source struct {
	Type int `json:"type"`
	Rate int `json:"rate"`
	Url string `json:"url"`
}
type _Operation struct {
	Type int `json:"type"`
	Rate int `json:"rate"`
}
type _Mdi struct {
	Type int `json:"type"`
	Rate int `json:"rate"`
}

type NewTask struct {
	Status _Status  `json:"status"`
	NewTaskResult _NewTaskResult   `json:"result"`
}

type _NewTaskResult struct {
	TaskId int `json:"taskId"`
}


type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "64060764@qq.com"
	c.TplName = "index.tpl"
}

type category struct {
		Id         string
		IsSelected bool
		Value      string
}

//登录
type LoginController struct{
	beego.Controller
}

func (c *LoginController) Post(){
if(c.GetString("username")=="ty"){
	c.TplName = "dataIndex.html"
	cates := []*category{
		&category{"-1", true, "请选择"},
		&category{"查看任务配置", false, "查看任务详情"},
		&category{"查看任务执行状态", false, "查看任务执行状态"},
		&category{"停止任务", false, "停止任务"},
		&category{"删除任务", false, "删除任务"},
		&category{"恢复任务", false, "恢复任务"},
	}

	c.Data["Cates"] = cates

}else{
	c.TplName = "loginerr.html"
}
}


//查询任务
type QueryController struct {
	beego.Controller
}
func (c *QueryController) Post() {
	if (c.GetString("taskid") != "") {
		c.TplName = "dataIndex.html"
		cates := []*category{
			&category{"-1", true, "请选择"},
			&category{"查看任务配置", false, "查看任务配置"},
			&category{"查看任务执行状态", false, "查看任务执行状态"},
			&category{"暂停任务", false, "暂停任务"},
			&category{"停止任务", false, "停止任务"},
			&category{"删除任务", false, "删除任务"},
			&category{"恢复任务", false, "恢复任务"},
		}

		c.Data["Cates"] = cates
		fmt.Println("选中内容是:"+c.Input().Get("category"))
		//判断如果选择的是查询任务执行状态
		if (c.Input().Get("category") == "查看任务执行状态") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/task.mobile.exec",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result SingleQuery
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			var TaskStatus = result.Result.__Status
			var TaskStatusText string
			switch TaskStatus {
			case 0:
				TaskStatusText = "停止中"
			case 1:
				TaskStatusText = "优化中"
			case 3:
				TaskStatusText = "暂停中"
			case 4:
				TaskStatusText = "定时中"
			}
			c.Data["Website"] = "任务编号:" + c.GetString("taskid") + " " + "计划ip/天:" + strconv.Itoa(result.Result.Plan) + " " + "当前状态:" + TaskStatusText + " " + "昨日优化数量:" + strconv.Itoa(result.Result.Yesterday) + " " + "今日优化数量:" + strconv.Itoa(result.Result.Today)

			logs.Info("查看任务执行状态 任务编号:" + c.GetString("taskid") + " " + "计划ip/天:" + strconv.Itoa(result.Result.Plan) + " " + "当前状态:" + TaskStatusText + " " + "昨日优化数量:" + strconv.Itoa(result.Result.Yesterday) + " " + "今日优化数量:" + strconv.Itoa(result.Result.Today))

			fmt.Println(strconv.Itoa(result.Result.Plan))
		} else if  (c.Input().Get("category") == "查看任务配置") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/task.mobile.view",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result _Task
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			var TaskStatus = result.Result.__Status
			var TaskStatusText string
			switch TaskStatus {
			case 0:
				TaskStatusText = "停止中"
			case 1:
				TaskStatusText = "优化中"
			case 3:
				TaskStatusText = "暂停中"
			case 4:
				TaskStatusText = "定时中"
			}
			var endtimestr string
			if(result.Result.EndTime==0){
				endtimestr = "一直优化"
			}else{
				tm := time.Unix(result.Result.EndTime, 0)
				endtimestr = tm.Format("2006-01-02 03:04:05 PM")
			}
			tm := time.Unix(result.Result.StartTime, 0)

                        refer := result.Result.Source[0].Url

			c.Data["Website"] = "任务编号:" + c.GetString("taskid") + " " + "计划ip/天:" + strconv.Itoa(result.Result.Plan) + " " + "当前状态:" + TaskStatusText + " "+" 任务URL:"+result.Result.Url+" "+"refer来源:"+refer+" 开始时间:"+tm.Format("2006-01-02 03:04:05 PM")+" 结束时间:"+endtimestr

			logs.Info("查看任务配置 任务编号:" + c.GetString("taskid") + " " + "计划ip/天:" + strconv.Itoa(result.Result.Plan) + " " + "当前状态:" + TaskStatusText + " "+" 任务URL:"+result.Result.Url+" "+"refer来源:"+refer+" 开始时间:"+tm.Format("2006-01-02 03:04:05 PM")+" 结束时间:"+endtimestr)
		}else if  (c.Input().Get("category") == "暂停任务") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/service.pause",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result _Status
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			c.Data["Website"] = "任务编号" + c.GetString("taskId") +string(body)
			/*if(result.Code==100&&result.Message=="SUCCESS"){
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"暂停成功!"
			}*/
			logs.Info("暂停任务 任务编号" + c.GetString("taskId") +string(body))

		}else if  (c.Input().Get("category") == "恢复任务") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/service.resume",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result _Status
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			c.Data["Website"] = "任务编号" + c.GetString("taskId") +string(body)
			/*if(result.Code==100&&result.Message=="SUCCESS"){
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"恢复成功!"
			}else{
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"恢复失败,只有暂停的任务才可以执行恢复操作!"
			}*/
			logs.Info("恢复任务 任务编号" + c.GetString("taskId") +string(body))


		}else if  (c.Input().Get("category") == "停止任务") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/service.stop",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result _Status
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			c.Data["Website"] = "任务编号" + c.GetString("taskId") +string(body)
			/*if(result.Code==100&&result.Message=="SUCCESS"){
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"停止成功!"
			}else{
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"停止失败!"
			}*/
			logs.Info("停止任务 任务编号" + c.GetString("taskId") +string(body))


		}else if  (c.Input().Get("category") == "删除任务") {
			//发起请求
			resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/task.mobile.delete",
				url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskid")}})

			if err != nil {
				// handle error
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				// handle error
			}
			fmt.Println(string(body))
			var result _Status
			err2 := json.Unmarshal(body, &result)
			if err2 != nil {
				fmt.Println("error:", err2)
			}
			c.Data["Website"] = "任务编号" + c.GetString("taskId") +string(body)
			/*if(result.Code==100&&result.Message=="SUCCESS"){
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"删除成功!"
			}else{
				c.Data["Website"] = "任务编号:" + c.GetString("taskid") +"删除失败!"
			}*/
			logs.Info("删除任务 任务编号" + c.GetString("taskId") +string(body))

		}



	}
}


//新建任务
type NewTaskController struct {
	beego.Controller
}
func (c *NewTaskController)Post(){
        c.TplName = "dataIndex.html"
	cates := []*category{
		&category{"-1", true, "请选择"},
		&category{"查看任务配置", false, "查看任务详情"},
		&category{"查看任务执行状态", false, "查看任务执行状态"},
		&category{"停止任务", false, "停止任务"},
		&category{"删除任务", false, "删除任务"},
		&category{"恢复任务", false, "恢复任务"},
	}

	c.Data["Cates"] = cates
	//c.Data["Website1"] = "new task ok!"
	sourcetxt :=c.GetString("source")
	//发起请求
	 source := "[{\"type\":1,\"url\":\""+sourcetxt+"\",\"rate\":100}]"
	fmt.Println("source="+source)
	resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/task.mobile.create",
		url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "url": {c.GetString("url")}, "plan": {c.GetString("plan")}, "name": {c.GetString("name")}, "source": {source}})

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	var result NewTask
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	c.Data["Website1"] = "添加任务成功,任务编号请记住"+strconv.Itoa(result.NewTaskResult.TaskId)
	logs.Info("添加任务,任务编号请记住"+strconv.Itoa(result.NewTaskResult.TaskId)+" "+string(body))


}

//开始任务
type StartTaskController struct {
	beego.Controller
}
func (c *StartTaskController)Post() {
	c.TplName = "dataIndex.html"
	cates := []*category{
		&category{"-1", true, "请选择"},
		&category{"查看任务配置", false, "查看任务详情"},
		&category{"查看任务执行状态", false, "查看任务执行状态"},
		&category{"停止任务", false, "停止任务"},
		&category{"删除任务", false, "删除任务"},
		&category{"恢复任务", false, "恢复任务"},
	}

	c.Data["Cates"] = cates
	//c.GetString("taskId")
	var starttime int = 0
	//var endtime int  = 0
	var params url.Values
	if(c.GetString("endtime")==""){
                params = url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskId")}, "startTime": {strconv.Itoa(starttime)}}
	}else{
		//将时间转换成时间戳作为endTime参数值
		toBeCharge := c.GetString("endtime")
		timeLayout := "2006-01-02 15:04:05"
		loc, _ := time.LoadLocation("Local")
		theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
		sr:=strconv.FormatInt(theTime.Unix(),10)
		fmt.Println("endtime.timestamp="+sr)
		params = url.Values{"apiKey": {"00d62d2f527037989bfd768f218da74e"}, "taskId": {c.GetString("taskId")}, "startTime": {strconv.Itoa(starttime)}, "endTime": {sr}}
	}
	resp, err := http.PostForm("http://service.liuliangbao.cn/api/v2/service.start",params)

	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	fmt.Println(string(body))
	var result _Status
	err2 := json.Unmarshal(body, &result)
	if err2 != nil {
		fmt.Println("error:", err2)
	}
	c.Data["Website2"] = "任务编号" + c.GetString("taskId") +string(body)
	logs.Info("开始任务 任务编号" + c.GetString("taskId") +string(body))


}


