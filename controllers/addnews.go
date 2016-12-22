package controllers

import (
	"github.com/astaxie/beego"
	_"fmt"
	_"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	"github.com/astaxie/beego/orm"



	"math"
)
var News []NewDetail
type AddNewsController struct {
	beego.Controller
}
type NewDetail struct {
	//Id int
	Title string
	Content string
}
func (c *AddNewsController) Get() {
	c.TplName = "addnews.html"
	orm := orm.NewOrm()
	if(orm!=nil){
		fmt.Println("connect mysql ok")
	}else{
		fmt.Println("err")
	}
	orm.Using("ibus")

	//

	var err error
	num,err := orm.Raw("select id,title,content from news").QueryRows(&News)
	if err == nil && num > 0 {
		for i, _ := range News {
		     c.Data["Website1"] = News[i].Title+"<br>"
		}

	}else{

		c.Data["Website1"] = err.Error()
	}

	//
	//count, _ := models.M("logoperation").Alias(`op`).Field(`count(op.id) as count`).Where(where).Count()
	//
	res := Paginator(1, 2, num)
	c.Data["paginator"] = res






}
func(c *AddNewsController) Post(){
	c.TplName = "addnews.html"
	orm := orm.NewOrm()
	if(orm!=nil){
		fmt.Println("connect mysql ok")
	}else{
		fmt.Println("err")
	}
	orm.Using("ibus")
	err :=orm.Raw("insert into news set title=? , content=?","pf","content").QueryRow()
	if(err!=nil){
		fmt.Println("insert err",err.Error())
	}else{
		fmt.Println("insert ok")
	}




}
func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}


