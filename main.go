package main

import (
	"embed"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// go:embed all:thirdparty/ocr.py
// go:embed all:thirdparty/convert.py
// go:embed all:thirdparty/dist/pdf.exe
// var thirdpartyAsset embed.FS

var (
	log    *logrus.Logger
	logger *logrus.Entry
	logdir string
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// init logger
	log = logrus.New()
	if runtime.GOOS == "windows" {
		logdir = filepath.Join(os.Getenv("USERPROFILE"), ".pdf_ryze")
	} else {
		logdir = filepath.Join(os.Getenv("HOME"), ".pdf_ryze")
	}
	err := os.MkdirAll(logdir, 0755)
	if err != nil {
		err = errors.Wrap(err, "failed to create log directory")
		log.Fatal(err)
	}
	logpath := filepath.Join(logdir, "access.log")
	file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		err = errors.Wrap(err, "failed to create log file")
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		DisableColors:   true,
	})
	logger = log.WithFields(logrus.Fields{
		"service": "pdf-ryze",
	})
	logger.Info("starting pdf-ryze")

	// Create an instance of the app structure
	app := NewApp()

	port := "50051"
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		// net_state = false
		fmt.Println("端口", port, "已被占用:", err)
		// listener.Close()
		var args = []string{"openServer", "open"}
		go_client(args)
	} else {
		// 如果没有错误发生，关闭监听以释放端口
		listener.Close()
		fmt.Println("端口", port, "可用")

		// app.cmdRunner(nil, "pdf")
		// config, err := app.LoadConfig()
		path, err := os.Executable()
		fmt.Printf("Error os.Executable(): %s\n", err)
		// fmt.Println("greeter_server运行=%s", config.PdfPath)
		pdfPath := filepath.Join(filepath.Dir(path), "pdf.exe")
		// fmt.Println("os.Executable()=%s", pdfPath)
		cmd := exec.Command(pdfPath)
		if err := cmd.Start(); err != nil {
			fmt.Printf("Error starting command: %s\n", err)
		}
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "PDF Ryze V1.0.1",
		Width:     1140, /*1280*/
		Height:    542,  /*700*/
		MinWidth:  1040, /*1000*/
		MinHeight: 500,  /*600*/
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		err = errors.Wrap(err, "run wails app failed")
		fmt.Println("Error:", err.Error())
	}

	for i := 0; i < 50; i++ {
		time.Sleep(500 * time.Millisecond)
		listener_, err_ := net.Listen("tcp", ":"+port)
		if err_ != nil {
			fmt.Println("关闭时，端口", port, "已被占用:", err_)
			var args = []string{"closeServer", "close"}
			go_client(args)
			break
		}
		listener_.Close()
	}

	logger.Info("exiting pdf-ryze")
}
