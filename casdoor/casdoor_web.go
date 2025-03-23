package casdoor

import (
	"embed"

	"github.com/chaos-plus/chaos-plus-toolx/xfile"
	"github.com/chaos-plus/chaos-plus-toolx/xres"
)

//go:embed all:static/web
var fs embed.FS

func ExportWebStaticFile(path string) {
	res := xres.New(fs)
	if !xfile.IsExist("./static/web/index.html") {
		if path != "" {
			res.Export("static/web", path, true)
		} else {
			res.Export("static/web", "web/build", true)
		}
	}
}
