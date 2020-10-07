package main

import (
	"filestore-server/handler"
	"fmt"
	"net/http"
)

// 主方法，用于启动go程序
func main() {
	// 静态资源处理
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(assets.AssetFS())))
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// 转到上传文件的页面
	http.HandleFunc("/file/upload", handler.UploadHandler)
	// 上传文件成功后返回的参数
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	// 查询上传文件的信息
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	// 下载上传的文件
	http.HandleFunc("/file/download", handler.DownloadHandler)
	// 更新上传文件的信息
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler)
	// 删除上传的文件
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)
	//http.HandleFunc("/file/query  ", handler.FileQueryHandler)

	// 用户注册
	http.HandleFunc("/user/signup", handler.SignupHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/info", handler.UserInfoHandler)
	// 监听端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print("Failed to start server,err:%s", err.Error())
	}

}
