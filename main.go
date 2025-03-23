package main

import casdoor "github.com/chaos-plus/lib3rd-casdoor/casdoor"

func main() {
	casdoor.ExportWebStaticFile("./casdoor/static/web")
	casdoor.RunServer()
}
