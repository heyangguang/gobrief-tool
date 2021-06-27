package codeGen

import (
	"bytes"
	templateEnv "gobrief-tool/template"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var ProjectName string

func GenCodeFiles(folder string) {
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if file.IsDir() {
			GenCodeFiles(folder + "/" + file.Name())
		} else {
			HandleCodeGen(folder + "/" + file.Name(), ProjectName)
		}
	}
}

func HandleCodeGen(filePath, projectName string)  {
	readText := Read2String(filePath)
	tmpl, err := template.New("test").Parse(readText)
	if err != nil {
		panic(err)
	}
	TemExecute(tmpl, projectName, filePath)
}

func TemExecute(tmpl *template.Template, projectName, filePath string)  {
	filePrefix := string([]byte(filePath)[10:])
	filePathGen := DirProjectReplace(filePrefix, projectName)
	buffer := new(bytes.Buffer)
	err := os.MkdirAll("./" + filePathGen, os.ModePerm)
	InitFileDir(projectName)
	file, err := os.Create("./" + NoDirProjectReplace(FileNameSetSuffix(filePrefix), projectName))
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buffer, map[string]interface{}{
		"ProjectName": projectName})
	if err != nil {
		panic(err)
	}
	_, _ = file.Write(buffer.Bytes())
}

func Read2String(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func FileNameSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, ".template")
}

func FileNameSetSuffix(fileNameSuffix string) string {
	fileName := path.Base(fileNameSuffix)
	if fileName == "go.template" {
		return FileNameSuffix(fileNameSuffix) + ".mod"
	}
	return FileNameSuffix(fileNameSuffix) + ".go"
}

func DirProjectReplace(filePath, projectName string) string {
	return strings.Replace(path.Dir(filePath), "project", projectName, -1)
}

func NoDirProjectReplace(filePath, projectName string) string {
	return strings.Replace(filePath, "project", projectName, -1)
}

func InitFileDir(projectName string)  {
	templateEnv.ConfigDev(projectName)
	templateEnv.ConfigProd(projectName)
	templateEnv.AirToml(projectName)
	templateEnv.BuildSh(projectName)
	templateEnv.LogDir(projectName)
	templateEnv.TmpDir(projectName)
	templateEnv.ServiceDir(projectName)
}