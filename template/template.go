package templateEnv

import (
	"os"
)

func ConfigDev(projectName string) {
	file, err := os.Create("./" + projectName + "/config_dev.yaml")
	if err != nil {
		panic(err)
	}
	_, _ = file.WriteString(`LogPath: "./logs/go-brief.log"
LogLevel: "INFO"
DBConnect: "root:123456@(127.0.0.1:3306)/go_brief"
Port: 8080`)
}

func ConfigProd(projectName string) {
	file, err := os.Create("./" + projectName + "/config_prod.yaml")
	if err != nil {
		panic(err)
	}
	_, _ = file.WriteString(`LogPath: "./logs/go-brief.log"
LogLevel: "INFO"
DBConnect: "root:123456@(127.0.0.1:3306)/go_brief"
Port: 8070`)
}

func AirToml(projectName string) {
	file, err := os.Create("./" + projectName + "/.air.toml")
	if err != nil {
		panic(err)
	}
	_, _ = file.WriteString(`# Config file for [Air](https://github.com/cosmtrek/air) in TOML format
# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp"

[build]
cmd = "go build -o ./tmp/main ."
bin = "tmp/main"
# Customize binary.
# full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
full_bin = "./tmp/main"
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = ["logs"]
# This log file places in your tmp_dir.
log = "./tmp/air.log"
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop running old binary when build errors occur.
stop_on_error = true
# Send Interrupt signal before killing process (windows does not support this feature)
send_interrupt = false
# Delay after sending Interrupt signal
kill_delay = 500 # ms

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true`)
}

func BuildSh(projectName string) {
	file, err := os.Create("./" + projectName + "/build.sh")
	if err != nil {
		panic(err)
	}
	_, _ = file.WriteString(`#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-brief main.go
CGO_ENABLED=0 go build -o go-brief main.go`)
}

func LogDir(projectName string) {
	_ = os.Mkdir("./"+projectName+"/logs", os.ModePerm)
}

func TmpDir(projectName string) {
	_ = os.Mkdir("./"+projectName+"/tmp", os.ModePerm)
}

func ServiceDir(projectName string) {
	_ = os.Mkdir("./"+projectName+"/app/service", os.ModePerm)
}
