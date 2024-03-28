package clilogger

import (
	"fmt"
	"io"
	"os"

	"github.com/wailsapp/wails/v2/internal/colour"
)

// CLILogger 是被 cli 使用的
type CLILogger struct {
	Writer io.Writer
	mute   bool
}

// New cli logger

// ff:创建
// writer:
func New(writer io.Writer) *CLILogger {
	return &CLILogger{
		Writer: writer,
	}
}

// Mute 设置是否应该禁用日志器

// ff:禁用日志
// value:禁用
func (c *CLILogger) Mute(value bool) {
	c.mute = value
}

// Print 函数类似于 Printf 函数

// ff:日志输出
// args:
// message:消息
func (c *CLILogger) Print(message string, args ...interface{}) {
	if c.mute {
		return
	}

	_, err := fmt.Fprintf(c.Writer, message, args...)
	if err != nil {
		c.Fatal("FATAL: " + err.Error())
	}
}

// Println 工作方式类似于 Printf，但在末尾添加换行符

// ff:日志输出并换行
// args:
// message:消息
func (c *CLILogger) Println(message string, args ...interface{}) {
	if c.mute {
		return
	}
	temp := fmt.Sprintf(message, args...)
	_, err := fmt.Fprintln(c.Writer, temp)
	if err != nil {
		c.Fatal("FATAL: " + err.Error())
	}
}

// Fatal 打印给定的消息，然后中止程序

// ff:日志输出并停止
// args:
// message:消息
func (c *CLILogger) Fatal(message string, args ...interface{}) {
	temp := fmt.Sprintf(message, args...)
	_, err := fmt.Fprintln(c.Writer, colour.Red("FATAL: "+temp))
	if err != nil {
		println(colour.Red("FATAL: " + err.Error()))
	}
	os.Exit(1)
}
