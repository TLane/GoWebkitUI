package main

import (
    "github.com/mattn/go-gtk/gtk"
    "github.com/mattn/go-webkit/webkit"
	"os"
    )


func main() {
	initPage("", "https://github.com/TLane/GoWebkitUI", 500, 500)
}

func initPage(title string, uri string, size_x int, size_y int) {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle(title)
	window.Connect("destroy", gtk.MainQuit)

	vbox := gtk.NewVBox(false, 1)

	swin := gtk.NewScrolledWindow(nil, nil)
	swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	swin.SetShadowType(gtk.SHADOW_IN)

	webview := webkit.NewWebView()
	webview.LoadUri(uri)
	swin.Add(webview)

	vbox.Add(swin)

	window.Add(vbox)
	window.SetSizeRequest(size_x, size_y)
	window.ShowAll()

	proxy := os.Getenv("HTTP_PROXY")
	if len(proxy) > 0 {
		soup_uri := webkit.SoupUri(proxy)
		webkit.GetDefaultSession().Set("proxy-uri", soup_uri)
		soup_uri.Free()
	}
	gtk.Main()
}
