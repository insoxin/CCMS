package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
	"os"
	"school/function"
	"school/models"
)

type UploadController struct {
	beego.Controller
}
type Getsheet struct {
	beego.Controller
}
type Getrow struct {
	Base
}
func (c* UploadController) Get() {
	c.TplName = "uploads.html"
}
func (c* UploadController) Post() {
	type Jsons struct {
		Code int
		Info string
	}
	jsontest := &Jsons{200, "上传成功"}
	f, _, _ := c.GetFile("fileInfo") //获取上传的文件
	path := "./filetmp/" + "c48a93c545c647dc9c20a4420e6eee33"
	f.Close()                      //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("fileInfo", path) //存文件
	text, _ := function.ReadAllIntoMemory(path)

	key := []byte("scoresdcet111246141score")
	x1 := function.Encrypt3DES(text, key)
	//x2 := function.Decrypt3DES(x1, key)
	function.WriteWithIoutil(path, string(x1))
	_, err := os.Open(path)
	if err != nil {
		jsontest = &Jsons{100, "上传失败"}
	}
	c.Data["json"] = jsontest
	c.ServeJSON()
}
func (c *Getsheet) Get(){
	var sheets []string
	path := "./filetmp/" + "c48a93c545c647dc9c20a4420e6eee33"
	text, _ := function.ReadAllIntoMemory(path)
	key := []byte("scoresdcet111246141score")
	x2 := function.Decrypt3DES(text, key)
	xlFile, err := xlsx.OpenBinary(x2)
	//fmt.Println(xlFile)
	if err != err{
		fmt.Println("上传失败")
	}
	//var sheetname string
	i := 0
	for _, sheet := range xlFile.Sheets {
		//fmt.Printf("Sheet Name: %s\n", sheet.Name)
		sheets = append(sheets, sheet.Name)
		fmt.Println(sheets[i])
		i++
		//sheetname = sheet.Name
	}
	c.Data["sheets"] = sheets
	c.TplName = "getsheet.html"
}
func (c *Getrow) Get(){
	var title [][]*xlsx.Cell
	var titlesql string
	var cells [][]*xlsx.Cell
	rowid,_ := c.GetInt(":rowid")
	path := "./filetmp/" + "c48a93c545c647dc9c20a4420e6eee33"
	text, _ := function.ReadAllIntoMemory(path)
	key := []byte("scoresdcet111246141score")
	x2 := function.Decrypt3DES(text, key)
	xlFile, err := xlsx.OpenBinary(x2)

	//i := 0
	sheet := xlFile.Sheets[rowid]	//选定Sheet工作表

	if err != nil{
		c.Ctx.WriteString("读取出错")
	}else{
		for k, row := range sheet.Rows {	//遍历工作表中的行
			if k == 0{						//首行为标题，单独拿出
				title = append(title,row.Cells)
				fmt.Println(len(row.Cells))
				for k,v := range row.Cells{
					if k < len(row.Cells) && k > 0{
						titlesql = titlesql + ","
					}
					titlesql = titlesql + v.String()
				}
			}else{							//遍历工作表的每一行
				for k,v := range row.Cells{
					if k == 1 && v.String() == "17372511800289"{
						fmt.Println(row.Cells)
					}
				}
				cells = append(cells,row.Cells)
			}
		}
		o := orm.NewOrm()
		title := models.Dataset{}
		title.Id = 1
		title.Exceltitle = titlesql
		title.Onlytitle = ""
		title.Looktitle = ""
		fmt.Println(o.Update(&title))
	}
	/*c.Data["cells"] = cells
	c.TplName = "getcell.html"*/
	c.Ctx.WriteString(titlesql)
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}