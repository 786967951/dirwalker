package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"github.com/gobwas/glob"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var listfile, pathpass, filepass []string

//go run dirwalker.go d:/test d:/test/test.txt pathpass:d:/test/test1 filepass:d:/test/test2/test2.txt
func main() {
	//获取参数
	flag.Parse()
	//所选文件夹    参数1 路径
	listpath := flag.Arg(0)
	//结果存放文件  参数2 文件
	reslutfile := flag.Arg(1)
	//获取过滤条件 参数3，4...
	for i := 2; i < flag.NArg(); i++ {
		s := flag.Arg(i)
		if strings.HasPrefix(s, "pathpass:") {
			//过滤文件夹
			pathpass = strings.Split(strings.Split(s, "pathpass:")[1], ",")
		}
		if strings.HasPrefix(s, "filepass:") {
			//过滤文件
			filepass = strings.Split(strings.Split(s, "filepass:")[1], ",")
		}
	}
	getSha(listpath, pathpass, filepass)
	f, _ := os.Create(reslutfile)
	for _, a := range listfile {
		//比如txt文件里，体现不出换行
		f.WriteString(a + "\n")
	}
	defer f.Close()
}

//遍历所有文件及文件夹
func getSha(listpath string, pathpass []string, filepass []string) {
	infos, _ := ioutil.ReadDir(listpath)
	for _, info := range infos {
		if info.IsDir() {
			//路径
			//检查本路径
			bol := false
			bol2 := false
			//判断是否过滤路径
			for _, pass := range pathpass {
				g := glob.MustCompile(pass)
				bol = g.Match(listpath + "/" + info.Name())
				//是过滤路径，遍历下一目标，否则递归遍历下一级路径
				if bol == true {
					bol2 = true
				}
			}
			if bol2 == true {
				continue
			} else {
				path := listpath + "/" + info.Name()
				getSha(path, pathpass, filepass)
			}

		} else {
			//文件
			bol := false
			bol2 := false
			//判断是否过滤文件
			for _, pass := range filepass {
				g := glob.MustCompile(pass)
				bol = g.Match(listpath + "/" + info.Name())
				if bol == true {
					bol2 = true
				}
			}
			if bol2 == true {
				continue
			} else {
		        	str := listpath + ":  " + info.Name() + "," + getSha1(listpath+"/"+info.Name()) + "," + getSize(info)
				listfile = append(listfile, str)
			}
		}
	}
}

//获取Sha1哈希值
func getSha1(path string) string {
	file, _ := os.Open(path)
	h := sha1.New()
	io.Copy(h, file)
	return fmt.Sprintf("%x", h.Sum(nil))
}

//获取文件大小
func getSize(info os.FileInfo) string {
	return strconv.FormatInt(info.Size(), 10)
}
