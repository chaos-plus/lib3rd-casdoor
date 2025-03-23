package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/chaos-plus/chaos-plus-toolx/xfile"
	"github.com/chaos-plus/chaos-plus-toolx/xpkg"
)

//go:generate go run gen.go

func main() {
	casdir := xpkg.GetPkgPath("github.com/casdoor/casdoor")
	fmt.Println("casdoor dir:", "["+casdir+"]")
	cfgdir := filepath.Join(casdir, "conf")
	cfg := filepath.Join(cfgdir, "app.conf")
	if !xfile.IsExist("../../conf/app.conf") {
		err := xfile.CopyFile(cfg, "../conf/app.conf", true)
		if err != nil {
			panic(err)
		}
	}
	src := "build"
	webdir := filepath.Join(casdir, "web")
	xfile.RemoveAll(src)
	err := xfile.CopyDir(webdir, src, true)
	if err != nil {
		panic(err)
	}
	err = cmd(src, []string{"yarn", "install", "--frozen-lockfile"}, []string{})
	if err != nil {
		panic(err)
	}
	err = cmd(src, []string{"yarn", "build"}, []string{"GENERATE_SOURCEMAP=false"})
	if err != nil {
		panic(err)
	}
	output := filepath.Join(src, "build")
	err = xfile.CopyDir(output, "web", true)
	if err != nil {
		panic(err)
	}

}

func cmd(dir string, args []string, env []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir // Set the working directory
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}
