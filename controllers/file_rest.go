package controllers

import (
	"github.com/astaxie/beego"
	"path"
	"os"
	"io/ioutil"
)

type FileController struct {
	beego.Controller
}

type FileList struct {
	FileNames string `json:"filename"`
}

func (c *FileController) URLMapping() {
	c.Mapping("GetFileList", c.GetFileList)
	c.Mapping("UploadFile", c.UploadFile)
	c.Mapping("DownloadFile", c.DownloadFile)
	c.Mapping("DeleteFile", c.DeleteFile)
}

// @router /getFileList [get]
func (this *FileController) GetFileList() {
	files, _ := ioutil.ReadDir("files")
	var filelist []FileList
	for _, file := range files {
		filelist = append(filelist, FileList{file.Name()})
	}
	print(filelist)
	this.Data["json"] = filelist
	this.ServeJSON()
}

// @router /uploadFile [post]
func (this *FileController) UploadFile() {
	file, info, err := this.GetFile("newfile")
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	}
	defer file.Close()
	filename := info.Filename
	err = this.SaveToFile("newfile", path.Join("files", filename))
	if err != nil {
		this.Ctx.WriteString("File upload failed！")
	} else {
		this.Ctx.WriteString("File upload succeed!")
	}

}

// @router /downloadFile [get]
func (this *FileController) DownloadFile() {
	filename := this.GetString("filename")
	this.Ctx.Output.Download(path.Join("files", filename), filename)
}

// @router /deleteFile/ [get]
func (this *FileController) DeleteFile() {
	filename := this.GetString("filename")
	os.Remove(path.Join("files", filename))
	files, _ := ioutil.ReadDir("files")
	var filelist []FileList
	for _, file := range files {
		filelist = append(filelist, FileList{file.Name()})
	}
	print(filelist)
	this.Data["json"] = filelist
	this.ServeJSON()
}
