package main

import (
	"log"
	"net/http"
)

func main() {
	// http.Dir 传递的是 .（当前目录），所以会将当前目录作为根目录，假如 . 目录下有 1.txt, 2.exe, 3.dmg
	// 这 3 个文件，那么可以通过诸如 http://localhost:8080/1.txt 的方式访问，如果访问
	// http://localhost:8080/，会在网页上显示 . 目录下的所有文件，同时点击可以查看文件内容
	http.Handle("/", http.FileServer(http.Dir(".")))

	// 访问 http://localhost:8080/static/ 为什么会 404？
	http.Handle("/static/", http.FileServer(http.Dir(".")))
	// 这个同样也是 404
	http.Handle("/a/", http.StripPrefix("/http/", http.FileServer(http.Dir("."))))
	// 解决：如果路由路径不是 /，那么必须要调用 http.StripPrefix，同时 StripPrefix 的第一个参数必须
	// 和路由路径相同，比如下面这个示例
	// 可以通过 http://localhost:8080/http/ 访问
	http.Handle("/http/", http.StripPrefix("/http/", http.FileServer(http.Dir("."))))

	// 注意这里的路由是 /b/ 如果是 /b 会 404
	http.HandleFunc("/b/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./demo.lubenwei")
	})

	log.Println("listen in 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
