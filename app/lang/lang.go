package lang

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
)

var LangTypes []*LangType

// LangType represents a language type.
type LangType struct {
	Lang, Name string
}

func LoadLang() {
	langs := strings.Split(beego.AppConfig.DefaultString("lang::types", "ta-LK"), "|")
	names := strings.Split(beego.AppConfig.DefaultString("lang::names", "ta-LK"), "|")
	LangTypes = make([]*LangType, 0, len(langs))
	for i, v := range langs {
		LangTypes = append(LangTypes, &LangType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		logs.Trace("Loading language: " + lang)
		root := "app/lang/" + lang
		files, err := ioutil.ReadDir(root)
		if err != nil {
			logs.Error(err)
		}
		os.Remove(root + ".ini")
		dstFile, err := os.Create(root + ".ini")
		if err != nil {
			logs.Error(err)
			continue
		}
		_, _ = dstFile.WriteString("#GENERATED FILE - DO NOT EDIT\n")
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			srcFile, err := os.Open(root + "/" + file.Name())
			if err != nil {
				logs.Error(err)
				continue
			}

			_, _ = dstFile.WriteString("\n["+file.Name()[0:len(file.Name())-4]+"]\n")
			bytesWritten, err := io.Copy(dstFile, srcFile)
			logs.Info(bytesWritten)
			if err != nil {
				logs.Error(err)
				continue
			}

			fmt.Println(file.Name())
		}

		if err := i18n.SetMessage(lang, root+".ini"); err != nil {
			logs.Error("Fail to set message file: " + err.Error())
		}
	}
}
