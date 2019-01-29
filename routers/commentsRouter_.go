package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-web-example/controllers:FileController"] = append(beego.GlobalControllerRouter["go-web-example/controllers:FileController"],
        beego.ControllerComments{
            Method: "DeleteFile",
            Router: `/deleteFile/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-web-example/controllers:FileController"] = append(beego.GlobalControllerRouter["go-web-example/controllers:FileController"],
        beego.ControllerComments{
            Method: "DownloadFile",
            Router: `/downloadFile`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-web-example/controllers:FileController"] = append(beego.GlobalControllerRouter["go-web-example/controllers:FileController"],
        beego.ControllerComments{
            Method: "GetFileList",
            Router: `/getFileList`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-web-example/controllers:FileController"] = append(beego.GlobalControllerRouter["go-web-example/controllers:FileController"],
        beego.ControllerComments{
            Method: "UploadFile",
            Router: `/uploadFile`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
