//go:build linux
// +build linux

package linux

import "github.com/pkg/browser"

// BrowserOpenURL 使用默认浏览器打开指定的url

// ff:默认浏览器打开url
// url:
func (f *Frontend) BrowserOpenURL(url string) {
	// 特定方法实现
	_ = browser.OpenURL(url)
}
