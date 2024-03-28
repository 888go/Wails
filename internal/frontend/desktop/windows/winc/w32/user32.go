//go:build windows

/*
 * Copyright (C) 2019 Tad Vizbaras. All Rights Reserved.
 * Copyright (C) 2010-2012 The W32 Authors. All Rights Reserved.
 */

package w32

import (
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procRegisterClassEx               = moduser32.NewProc("RegisterClassExW")
	procLoadIcon                      = moduser32.NewProc("LoadIconW")
	procLoadCursor                    = moduser32.NewProc("LoadCursorW")
	procShowWindow                    = moduser32.NewProc("ShowWindow")
	procShowWindowAsync               = moduser32.NewProc("ShowWindowAsync")
	procUpdateWindow                  = moduser32.NewProc("UpdateWindow")
	procCreateWindowEx                = moduser32.NewProc("CreateWindowExW")
	procFindWindowW                   = moduser32.NewProc("FindWindowW")
	procAdjustWindowRect              = moduser32.NewProc("AdjustWindowRect")
	procAdjustWindowRectEx            = moduser32.NewProc("AdjustWindowRectEx")
	procDestroyWindow                 = moduser32.NewProc("DestroyWindow")
	procDefWindowProc                 = moduser32.NewProc("DefWindowProcW")
	procDefDlgProc                    = moduser32.NewProc("DefDlgProcW")
	procPostQuitMessage               = moduser32.NewProc("PostQuitMessage")
	procGetMessage                    = moduser32.NewProc("GetMessageW")
	procTranslateMessage              = moduser32.NewProc("TranslateMessage")
	procDispatchMessage               = moduser32.NewProc("DispatchMessageW")
	procSendMessage                   = moduser32.NewProc("SendMessageW")
	procPostMessage                   = moduser32.NewProc("PostMessageW")
	procWaitMessage                   = moduser32.NewProc("WaitMessage")
	procSetWindowText                 = moduser32.NewProc("SetWindowTextW")
	procGetWindowTextLength           = moduser32.NewProc("GetWindowTextLengthW")
	procGetWindowText                 = moduser32.NewProc("GetWindowTextW")
	procGetWindowRect                 = moduser32.NewProc("GetWindowRect")
	procGetWindowInfo                 = moduser32.NewProc("GetWindowInfo")
	procSetWindowCompositionAttribute = moduser32.NewProc("SetWindowCompositionAttribute")
	procMoveWindow                    = moduser32.NewProc("MoveWindow")
	procScreenToClient                = moduser32.NewProc("ScreenToClient")
	procCallWindowProc                = moduser32.NewProc("CallWindowProcW")
	procSetWindowLong                 = moduser32.NewProc("SetWindowLongW")
	procSetWindowLongPtr              = moduser32.NewProc("SetWindowLongW")
	procGetWindowLong                 = moduser32.NewProc("GetWindowLongW")
	procGetWindowLongPtr              = moduser32.NewProc("GetWindowLongW")
	procEnableWindow                  = moduser32.NewProc("EnableWindow")
	procIsWindowEnabled               = moduser32.NewProc("IsWindowEnabled")
	procIsWindowVisible               = moduser32.NewProc("IsWindowVisible")
	procSetFocus                      = moduser32.NewProc("SetFocus")
	procGetFocus                      = moduser32.NewProc("GetFocus")
	procSetActiveWindow               = moduser32.NewProc("SetActiveWindow")
	procSetForegroundWindow           = moduser32.NewProc("SetForegroundWindow")
	procBringWindowToTop              = moduser32.NewProc("BringWindowToTop")
	procInvalidateRect                = moduser32.NewProc("InvalidateRect")
	procGetClientRect                 = moduser32.NewProc("GetClientRect")
	procGetDC                         = moduser32.NewProc("GetDC")
	procReleaseDC                     = moduser32.NewProc("ReleaseDC")
	procSetCapture                    = moduser32.NewProc("SetCapture")
	procReleaseCapture                = moduser32.NewProc("ReleaseCapture")
	procGetWindowThreadProcessId      = moduser32.NewProc("GetWindowThreadProcessId")
	procMessageBox                    = moduser32.NewProc("MessageBoxW")
	procGetSystemMetrics              = moduser32.NewProc("GetSystemMetrics")
	procPostThreadMessageW            = moduser32.NewProc("PostThreadMessageW")
	procRegisterWindowMessageA        = moduser32.NewProc("RegisterWindowMessageA")
	// procSysColorBrush 是一个指向用户32模块中名为"GetSysColorBrush"的新进程的指针
	procCopyRect          = moduser32.NewProc("CopyRect")
	procEqualRect         = moduser32.NewProc("EqualRect")
	procInflateRect       = moduser32.NewProc("InflateRect")
	procIntersectRect     = moduser32.NewProc("IntersectRect")
	procIsRectEmpty       = moduser32.NewProc("IsRectEmpty")
	procOffsetRect        = moduser32.NewProc("OffsetRect")
	procPtInRect          = moduser32.NewProc("PtInRect")
	procSetRect           = moduser32.NewProc("SetRect")
	procSetRectEmpty      = moduser32.NewProc("SetRectEmpty")
	procSubtractRect      = moduser32.NewProc("SubtractRect")
	procUnionRect         = moduser32.NewProc("UnionRect")
	procCreateDialogParam = moduser32.NewProc("CreateDialogParamW")
	procDialogBoxParam    = moduser32.NewProc("DialogBoxParamW")
	procGetDlgItem        = moduser32.NewProc("GetDlgItem")
	procDrawIcon          = moduser32.NewProc("DrawIcon")
	procCreateMenu        = moduser32.NewProc("CreateMenu")
	// procSetMenu 是一个指向 "SetMenu" 函数的指针，该函数位于 moduser32 包中
// moduser32.NewProc 用于从用户32模块动态加载并创建指向 "SetMenu" API 函数的指针
// procSetMenu                  = moduser32.NewProc("SetMenu") 
// 这一行代码定义了一个名为 procSetMenu 的变量，它通过 moduser32 包获取并指向 Windows API 中的 "SetMenu" 函数。
	procDestroyMenu        = moduser32.NewProc("DestroyMenu")
	procCreatePopupMenu    = moduser32.NewProc("CreatePopupMenu")
	procCheckMenuRadioItem = moduser32.NewProc("CheckMenuRadioItem")
// procDrawMenuBar = moduser32.NewProc("DrawMenuBar") // 创建并获取名为"DrawMenuBar"的用户32模块（moduser32）中的新过程
// procInsertMenuItem = moduser32.NewProc("InsertMenuItemW") // 创建并获取名为"InsertMenuItemW"的用户32模块中的新过程 // 注意：待修复或进一步处理

	procClientToScreen                = moduser32.NewProc("ClientToScreen")
	procIsDialogMessage               = moduser32.NewProc("IsDialogMessageW")
	procIsWindow                      = moduser32.NewProc("IsWindow")
	procEndDialog                     = moduser32.NewProc("EndDialog")
	procPeekMessage                   = moduser32.NewProc("PeekMessageW")
	procTranslateAccelerator          = moduser32.NewProc("TranslateAcceleratorW")
	procSetWindowPos                  = moduser32.NewProc("SetWindowPos")
	procFillRect                      = moduser32.NewProc("FillRect")
	procDrawText                      = moduser32.NewProc("DrawTextW")
	procAddClipboardFormatListener    = moduser32.NewProc("AddClipboardFormatListener")
	procRemoveClipboardFormatListener = moduser32.NewProc("RemoveClipboardFormatListener")
	procOpenClipboard                 = moduser32.NewProc("OpenClipboard")
	procCloseClipboard                = moduser32.NewProc("CloseClipboard")
	procEnumClipboardFormats          = moduser32.NewProc("EnumClipboardFormats")
	procGetClipboardData              = moduser32.NewProc("GetClipboardData")
	procSetClipboardData              = moduser32.NewProc("SetClipboardData")
	procEmptyClipboard                = moduser32.NewProc("EmptyClipboard")
	procGetClipboardFormatName        = moduser32.NewProc("GetClipboardFormatNameW")
	procIsClipboardFormatAvailable    = moduser32.NewProc("IsClipboardFormatAvailable")
	procBeginPaint                    = moduser32.NewProc("BeginPaint")
	procEndPaint                      = moduser32.NewProc("EndPaint")
	procGetKeyboardState              = moduser32.NewProc("GetKeyboardState")
	procMapVirtualKey                 = moduser32.NewProc("MapVirtualKeyExW")
	procGetAsyncKeyState              = moduser32.NewProc("GetAsyncKeyState")
	procToAscii                       = moduser32.NewProc("ToAscii")
	procSwapMouseButton               = moduser32.NewProc("SwapMouseButton")
	procGetCursorPos                  = moduser32.NewProc("GetCursorPos")
	procSetCursorPos                  = moduser32.NewProc("SetCursorPos")
	procSetCursor                     = moduser32.NewProc("SetCursor")
	procCreateIcon                    = moduser32.NewProc("CreateIcon")
	procDestroyIcon                   = moduser32.NewProc("DestroyIcon")
	procMonitorFromPoint              = moduser32.NewProc("MonitorFromPoint")
	procMonitorFromRect               = moduser32.NewProc("MonitorFromRect")
	procMonitorFromWindow             = moduser32.NewProc("MonitorFromWindow")
	procGetMonitorInfo                = moduser32.NewProc("GetMonitorInfoW")
	procGetDpiForSystem               = moduser32.NewProc("GetDpiForSystem")
	procGetDpiForWindow               = moduser32.NewProc("GetDpiForWindow")
	procEnumDisplayMonitors           = moduser32.NewProc("EnumDisplayMonitors")
	procEnumDisplaySettingsEx         = moduser32.NewProc("EnumDisplaySettingsExW")
	procChangeDisplaySettingsEx       = moduser32.NewProc("ChangeDisplaySettingsExW")
	procSendInput                     = moduser32.NewProc("SendInput")
	procSetWindowsHookEx              = moduser32.NewProc("SetWindowsHookExW")
	procUnhookWindowsHookEx           = moduser32.NewProc("UnhookWindowsHookEx")
	procCallNextHookEx                = moduser32.NewProc("CallNextHookEx")

	libuser32, _        = syscall.LoadLibrary("user32.dll")
	insertMenuItem, _   = syscall.GetProcAddress(libuser32, "InsertMenuItemW")
	setMenuItemInfo, _  = syscall.GetProcAddress(libuser32, "SetMenuItemInfoW")
	setMenu, _          = syscall.GetProcAddress(libuser32, "SetMenu")
	drawMenuBar, _      = syscall.GetProcAddress(libuser32, "DrawMenuBar")
	trackPopupMenuEx, _ = syscall.GetProcAddress(libuser32, "TrackPopupMenuEx")
	getKeyState, _      = syscall.GetProcAddress(libuser32, "GetKeyState")
	getSysColorBrush, _ = syscall.GetProcAddress(libuser32, "GetSysColorBrush")

	getWindowPlacement, _ = syscall.GetProcAddress(libuser32, "GetWindowPlacement")
	setWindowPlacement, _ = syscall.GetProcAddress(libuser32, "SetWindowPlacement")

	setScrollInfo, _ = syscall.GetProcAddress(libuser32, "SetScrollInfo")
	getScrollInfo, _ = syscall.GetProcAddress(libuser32, "GetScrollInfo")

	mainThread HANDLE
)

func init() {
	runtime.LockOSThread()
	mainThread = GetCurrentThreadId()
}


// ff:
// lp:
func GET_X_LPARAM(lp uintptr) int32 {
	return int32(int16(LOWORD(uint32(lp))))
}


// ff:
// lp:
func GET_Y_LPARAM(lp uintptr) int32 {
	return int32(int16(HIWORD(uint32(lp))))
}


// ff:
// wndClassEx:
func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
	ret, _, _ := procRegisterClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ATOM(ret)
}


// ff:
// iconName:
// instance:
func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
	ret, _, _ := procLoadIcon.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(iconName)))

	return HICON(ret)
}


// ff:
// res:
// instance:
func LoadIconWithResourceID(instance HINSTANCE, res uint16) HICON {
	ret, _, _ := procLoadIcon.Call(
		uintptr(instance),
		uintptr(res))

	return HICON(ret)
}


// ff:
// cursorName:
// instance:
func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
	ret, _, _ := procLoadCursor.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(cursorName)))

	return HCURSOR(ret)
}


// ff:
// res:
// instance:
func LoadCursorWithResourceID(instance HINSTANCE, res uint16) HCURSOR {
	ret, _, _ := procLoadCursor.Call(
		uintptr(instance),
		uintptr(res))

	return HCURSOR(ret)
}


// ff:
// cmdshow:
// hwnd:
func ShowWindow(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0
}


// ff:
// cmdshow:
// hwnd:
func ShowWindowAsync(hwnd HWND, cmdshow int) bool {
	ret, _, _ := procShowWindowAsync.Call(
		uintptr(hwnd),
		uintptr(cmdshow))

	return ret != 0
}


// ff:
// hwnd:
func UpdateWindow(hwnd HWND) bool {
	ret, _, _ := procUpdateWindow.Call(
		uintptr(hwnd))
	return ret != 0
}


// ff:
// lp:
// wp:
// msg:
// threadID:
func PostThreadMessage(threadID HANDLE, msg int, wp, lp uintptr) {
	procPostThreadMessageW.Call(threadID, uintptr(msg), wp, lp)
}


// ff:
// name:
func RegisterWindowMessage(name *uint16) uint32 {
	ret, _, _ := procRegisterWindowMessageA.Call(
		uintptr(unsafe.Pointer(name)))

	return uint32(ret)
}


// ff:
// lp:
// wp:
// msg:
func PostMainThreadMessage(msg uint32, wp, lp uintptr) bool {
	ret, _, _ := procPostThreadMessageW.Call(mainThread, uintptr(msg), wp, lp)
	return ret != 0
}


// ff:
func CreateWindowEx(exStyle uint, className, windowName *uint16,
	style uint, x, y, width, height int, parent HWND, menu HMENU,
	instance HINSTANCE, param unsafe.Pointer) HWND {
	ret, _, _ := procCreateWindowEx.Call(
		uintptr(exStyle),
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(param))

	return HWND(ret)
}


// ff:
// windowName:
// className:
func FindWindowW(className, windowName *uint16) HWND {
	ret, _, _ := procFindWindowW.Call(
		uintptr(unsafe.Pointer(className)),
		uintptr(unsafe.Pointer(windowName)))

	return HWND(ret)
}


// ff:
// exStyle:
// menu:
// style:
// rect:
func AdjustWindowRectEx(rect *RECT, style uint, menu bool, exStyle uint) bool {
	ret, _, _ := procAdjustWindowRectEx.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)),
		uintptr(exStyle))

	return ret != 0
}


// ff:
// menu:
// style:
// rect:
func AdjustWindowRect(rect *RECT, style uint, menu bool) bool {
	ret, _, _ := procAdjustWindowRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(style),
		uintptr(BoolToBOOL(menu)))

	return ret != 0
}


// ff:
// hwnd:
func DestroyWindow(hwnd HWND) bool {
	ret, _, _ := procDestroyWindow.Call(hwnd)
	return ret != 0
}


// ff:
func HasGetDpiForWindowFunc() bool {
	err := procGetDpiForWindow.Find()
	return err == nil
}


// ff:
// hwnd:
func GetDpiForWindow(hwnd HWND) UINT {
	dpi, _, _ := procGetDpiForWindow.Call(hwnd)
	return uint(dpi)
}


// ff:
// data:
// hwnd:
func SetWindowCompositionAttribute(hwnd HWND, data *WINDOWCOMPOSITIONATTRIBDATA) bool {
	if procSetWindowCompositionAttribute != nil {
		ret, _, _ := procSetWindowCompositionAttribute.Call(
			hwnd,
			uintptr(unsafe.Pointer(data)),
		)
		return ret != 0
	}
	return false
}


// ff:
// lParam:
// wParam:
// msg:
// hwnd:
func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}


// ff:
// lParam:
// wParam:
// msg:
// hwnd:
func DefDlgProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procDefDlgProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}


// ff:
// exitCode:
func PostQuitMessage(exitCode int) {
	procPostQuitMessage.Call(
		uintptr(exitCode))
}


// ff:
// msgFilterMax:
// msgFilterMin:
// hwnd:
// msg:
func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
	ret, _, _ := procGetMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax))

	return int(ret)
}


// ff:
// msg:
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret != 0

}


// ff:
// msg:
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := procDispatchMessage.Call(
		uintptr(unsafe.Pointer(msg)))

	return ret

}


// ff:
// lParam:
// wParam:
// msg:
// hwnd:
func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procSendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}


// ff:
// lParam:
// wParam:
// msg:
// hwnd:
func PostMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) bool {
	ret, _, _ := procPostMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret != 0
}


// ff:
func WaitMessage() bool {
	ret, _, _ := procWaitMessage.Call()
	return ret != 0
}


// ff:
// text:
// hwnd:
func SetWindowText(hwnd HWND, text string) {
	procSetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}


// ff:
// hwnd:
func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(
		uintptr(hwnd))

	return int(ret)
}


// ff:
// info:
// hwnd:
func GetWindowInfo(hwnd HWND, info *WINDOWINFO) int {
	ret, _, _ := procGetWindowInfo.Call(
		hwnd,
		uintptr(unsafe.Pointer(info)),
	)
	return int(ret)
}


// ff:
// hwnd:
func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1

	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}


// ff:
// hwnd:
func GetWindowRect(hwnd HWND) *RECT {
	var rect RECT
	procGetWindowRect.Call(
		hwnd,
		uintptr(unsafe.Pointer(&rect)))

	return &rect
}


// ff:
// repaint:
// height:
// width:
// y:
// x:
// hwnd:
func MoveWindow(hwnd HWND, x, y, width, height int, repaint bool) bool {
	ret, _, _ := procMoveWindow.Call(
		uintptr(hwnd),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(BoolToBOOL(repaint)))

	return ret != 0

}


// ff:
// ok:
// Y:
// X:
// y:
// x:
// hwnd:
func ScreenToClient(hwnd HWND, x, y int) (X, Y int, ok bool) {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procScreenToClient.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y), ret != 0
}


// ff:
// lParam:
// wParam:
// msg:
// hwnd:
// preWndProc:
func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := procCallWindowProc.Call(
		preWndProc,
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam)

	return ret
}


// ff:
// value:
// index:
// hwnd:
func SetWindowLong(hwnd HWND, index int, value uint32) uint32 {
	ret, _, _ := procSetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index),
		uintptr(value))

	return uint32(ret)
}


// ff:
// value:
// index:
// hwnd:
func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
	ret, _, _ := procSetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index),
		value)

	return ret
}


// ff:
// index:
// hwnd:
func GetWindowLong(hwnd HWND, index int) int32 {
	ret, _, _ := procGetWindowLong.Call(
		uintptr(hwnd),
		uintptr(index))

	return int32(ret)
}


// ff:
// index:
// hwnd:
func GetWindowLongPtr(hwnd HWND, index int) uintptr {
	ret, _, _ := procGetWindowLongPtr.Call(
		uintptr(hwnd),
		uintptr(index))

	return ret
}


// ff:
// b:
// hwnd:
func EnableWindow(hwnd HWND, b bool) bool {
	ret, _, _ := procEnableWindow.Call(
		uintptr(hwnd),
		uintptr(BoolToBOOL(b)))
	return ret != 0
}


// ff:
// hwnd:
func IsWindowEnabled(hwnd HWND) bool {
	ret, _, _ := procIsWindowEnabled.Call(
		uintptr(hwnd))

	return ret != 0
}


// ff:
// hwnd:
func IsWindowVisible(hwnd HWND) bool {
	ret, _, _ := procIsWindowVisible.Call(
		uintptr(hwnd))

	return ret != 0
}


// ff:
// hwnd:
func SetFocus(hwnd HWND) HWND {
	ret, _, _ := procSetFocus.Call(
		uintptr(hwnd))

	return HWND(ret)
}


// ff:
// hwnd:
func SetActiveWindow(hwnd HWND) HWND {
	ret, _, _ := procSetActiveWindow.Call(
		uintptr(hwnd))

	return HWND(ret)
}


// ff:
// hwnd:
func BringWindowToTop(hwnd HWND) bool {
	ret, _, _ := procBringWindowToTop.Call(uintptr(hwnd))
	return ret != 0
}


// ff:
// hwnd:
func SetForegroundWindow(hwnd HWND) HWND {
	ret, _, _ := procSetForegroundWindow.Call(
		uintptr(hwnd))

	return HWND(ret)
}


// ff:
func GetFocus() HWND {
	ret, _, _ := procGetFocus.Call()
	return HWND(ret)
}


// ff:
// erase:
// rect:
// hwnd:
func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
	ret, _, _ := procInvalidateRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(rect)),
		uintptr(BoolToBOOL(erase)))

	return ret != 0
}


// ff:
// hwnd:
func GetClientRect(hwnd HWND) *RECT {
	var rect RECT
	ret, _, _ := procGetClientRect.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)))

	if ret == 0 {
		panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
	}

	return &rect
}


// ff:
// hwnd:
func GetDC(hwnd HWND) HDC {
	ret, _, _ := procGetDC.Call(
		uintptr(hwnd))

	return HDC(ret)
}


// ff:
// hDC:
// hwnd:
func ReleaseDC(hwnd HWND, hDC HDC) bool {
	ret, _, _ := procReleaseDC.Call(
		uintptr(hwnd),
		uintptr(hDC))

	return ret != 0
}


// ff:
// hwnd:
func SetCapture(hwnd HWND) HWND {
	ret, _, _ := procSetCapture.Call(
		uintptr(hwnd))

	return HWND(ret)
}


// ff:
func ReleaseCapture() bool {
	ret, _, _ := procReleaseCapture.Call()

	return ret != 0
}


// ff:
// HANDLE:
// hwnd:
func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
	var processId int
	ret, _, _ := procGetWindowThreadProcessId.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)))

	return HANDLE(ret), processId
}


// ff:
// flags:
// caption:
// title:
// hwnd:
func MessageBox(hwnd HWND, title, caption string, flags uint) int {
	ret, _, _ := procMessageBox.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(flags))

	return int(ret)
}


// ff:
// index:
func GetSystemMetrics(index int) int {
	ret, _, _ := procGetSystemMetrics.Call(
		uintptr(index))

	return int(ret)
}


// ff:
// nIndex:
func GetSysColorBrush(nIndex int) HBRUSH {
	/*
		ret, _, _ := procSysColorBrush.Call(1,
			uintptr(nIndex),
			0,
			0)

		return HBRUSH(ret)
	*/
	ret, _, _ := syscall.Syscall(getSysColorBrush, 1,
		uintptr(nIndex),
		0,
		0)

	return HBRUSH(ret)
}


// ff:
// src:
// dst:
func CopyRect(dst, src *RECT) bool {
	ret, _, _ := procCopyRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)))

	return ret != 0
}


// ff:
// rect2:
// rect1:
func EqualRect(rect1, rect2 *RECT) bool {
	ret, _, _ := procEqualRect.Call(
		uintptr(unsafe.Pointer(rect1)),
		uintptr(unsafe.Pointer(rect2)))

	return ret != 0
}


// ff:
// dy:
// dx:
// rect:
func InflateRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procInflateRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}


// ff:
// src2:
// src1:
// dst:
func IntersectRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procIntersectRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}


// ff:
// rect:
func IsRectEmpty(rect *RECT) bool {
	ret, _, _ := procIsRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}


// ff:
// dy:
// dx:
// rect:
func OffsetRect(rect *RECT, dx, dy int) bool {
	ret, _, _ := procOffsetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(dx),
		uintptr(dy))

	return ret != 0
}


// ff:
// y:
// x:
// rect:
func PtInRect(rect *RECT, x, y int) bool {
	pt := POINT{X: int32(x), Y: int32(y)}
	ret, _, _ := procPtInRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(&pt)))

	return ret != 0
}


// ff:
// bottom:
// right:
// top:
// left:
// rect:
func SetRect(rect *RECT, left, top, right, bottom int) bool {
	ret, _, _ := procSetRect.Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(left),
		uintptr(top),
		uintptr(right),
		uintptr(bottom))

	return ret != 0
}


// ff:
// rect:
func SetRectEmpty(rect *RECT) bool {
	ret, _, _ := procSetRectEmpty.Call(
		uintptr(unsafe.Pointer(rect)))

	return ret != 0
}


// ff:
// src2:
// src1:
// dst:
func SubtractRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procSubtractRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}


// ff:
// src2:
// src1:
// dst:
func UnionRect(dst, src1, src2 *RECT) bool {
	ret, _, _ := procUnionRect.Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src1)),
		uintptr(unsafe.Pointer(src2)))

	return ret != 0
}


// ff:
// lpDialogProc:
// hWndParent:
// lpTemplate:
// hInstance:
func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
	ret, _, _ := procCreateDialogParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplate)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return HWND(ret)
}


// ff:
// lpDialogProc:
// hWndParent:
// lpTemplateName:
// hInstance:
func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
	ret, _, _ := procDialogBoxParam.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpTemplateName)),
		uintptr(hWndParent),
		lpDialogProc,
		0)

	return int(ret)
}


// ff:
// nIDDlgItem:
// hDlg:
func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
	ret, _, _ := procGetDlgItem.Call(
		uintptr(unsafe.Pointer(hDlg)),
		uintptr(nIDDlgItem))

	return HWND(ret)
}


// ff:
// hIcon:
// y:
// x:
// hDC:
func DrawIcon(hDC HDC, x, y int, hIcon HICON) bool {
	ret, _, _ := procDrawIcon.Call(
		uintptr(unsafe.Pointer(hDC)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(hIcon)))

	return ret != 0
}


// ff:
func CreateMenu() HMENU {
	ret, _, _ := procCreateMenu.Call(0,
		0,
		0,
		0)

	return HMENU(ret)
}


// ff:
// hMenu:
// hWnd:
func SetMenu(hWnd HWND, hMenu HMENU) bool {
	ret, _, _ := syscall.Syscall(setMenu, 2,
		uintptr(hWnd),
		uintptr(hMenu),
		0)

	return ret != 0
}

// 这是微软官方文档链接，指向Windows Win32 API中关于CheckMenuRadioItem函数的说明：
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-checkmenuradioitem
// （该链接描述的是Windows API中的一个函数，用于在菜单栏中切换单选按钮式的菜单项。）

// ff:
// hwnd:
// endID:
// startID:
// menuID:
func SelectRadioMenuItem(menuID uint16, startID uint16, endID uint16, hwnd HWND) bool {
	ret, _, _ := procCheckMenuRadioItem.Call(
		hwnd,
		uintptr(startID),
		uintptr(endID),
		uintptr(menuID),
		MF_BYCOMMAND)
	return ret != 0

}


// ff:
func CreatePopupMenu() HMENU {
	ret, _, _ := procCreatePopupMenu.Call(0,
		0,
		0,
		0)

	return HMENU(ret)
}


// ff:
// lptpm:
// hWnd:
// y:
// x:
// fuFlags:
// hMenu:
func TrackPopupMenuEx(hMenu HMENU, fuFlags uint32, x, y int32, hWnd HWND, lptpm *TPMPARAMS) BOOL {
	ret, _, _ := syscall.Syscall6(trackPopupMenuEx, 6,
		uintptr(hMenu),
		uintptr(fuFlags),
		uintptr(x),
		uintptr(y),
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lptpm)))

	return BOOL(ret)
}


// ff:
// hWnd:
func DrawMenuBar(hWnd HWND) bool {
	ret, _, _ := syscall.Syscall(drawMenuBar, 1,
		uintptr(hWnd),
		0,
		0)

	return ret != 0
}


// ff:
// lpmii:
// fByPosition:
// uItem:
// hMenu:
func InsertMenuItem(hMenu HMENU, uItem uint32, fByPosition bool, lpmii *MENUITEMINFO) bool {
	ret, _, _ := syscall.Syscall6(insertMenuItem, 4,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(BoolToBOOL(fByPosition)),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)

	return ret != 0
}


// ff:
// lpmii:
// fByPosition:
// uItem:
// hMenu:
func SetMenuItemInfo(hMenu HMENU, uItem uint32, fByPosition bool, lpmii *MENUITEMINFO) bool {
	ret, _, _ := syscall.Syscall6(setMenuItemInfo, 4,
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(BoolToBOOL(fByPosition)),
		uintptr(unsafe.Pointer(lpmii)),
		0,
		0)

	return ret != 0
}


// ff:
// y:
// x:
// hwnd:
func ClientToScreen(hwnd HWND, x, y int) (int, int) {
	pt := POINT{X: int32(x), Y: int32(y)}

	procClientToScreen.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pt)))

	return int(pt.X), int(pt.Y)
}


// ff:
// msg:
// hwnd:
func IsDialogMessage(hwnd HWND, msg *MSG) bool {
	ret, _, _ := procIsDialogMessage.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(msg)))

	return ret != 0
}


// ff:
// hwnd:
func IsWindow(hwnd HWND) bool {
	ret, _, _ := procIsWindow.Call(
		uintptr(hwnd))

	return ret != 0
}


// ff:
// nResult:
// hwnd:
func EndDialog(hwnd HWND, nResult uintptr) bool {
	ret, _, _ := procEndDialog.Call(
		uintptr(hwnd),
		nResult)

	return ret != 0
}


// ff:
// wRemoveMsg:
// wMsgFilterMax:
// wMsgFilterMin:
// hwnd:
// lpMsg:
func PeekMessage(lpMsg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := procPeekMessage.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg))

	return ret != 0
}


// ff:
// lpMsg:
// hAccTable:
// hwnd:
func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
	ret, _, _ := procTranslateMessage.Call(
		uintptr(hwnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)))

	return ret != 0
}


// ff:
// uFlags:
// cy:
// cx:
// y:
// x:
// hWndInsertAfter:
// hwnd:
func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint) bool {
	ret, _, _ := procSetWindowPos.Call(
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags))

	return ret != 0
}


// ff:
// hbr:
// lprc:
// hDC:
func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
	ret, _, _ := procFillRect.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr))

	return ret != 0
}


// ff:
// uFormat:
// lpRect:
// uCount:
// text:
// hDC:
func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
	ret, _, _ := procDrawText.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
		uintptr(uCount),
		uintptr(unsafe.Pointer(lpRect)),
		uintptr(uFormat))

	return int(ret)
}


// ff:
// hwnd:
func AddClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procAddClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}


// ff:
// hwnd:
func RemoveClipboardFormatListener(hwnd HWND) bool {
	ret, _, _ := procRemoveClipboardFormatListener.Call(
		uintptr(hwnd))
	return ret != 0
}


// ff:
// hWndNewOwner:
func OpenClipboard(hWndNewOwner HWND) bool {
	ret, _, _ := procOpenClipboard.Call(
		uintptr(hWndNewOwner))
	return ret != 0
}


// ff:
func CloseClipboard() bool {
	ret, _, _ := procCloseClipboard.Call()
	return ret != 0
}


// ff:
// format:
func EnumClipboardFormats(format uint) uint {
	ret, _, _ := procEnumClipboardFormats.Call(
		uintptr(format))
	return uint(ret)
}


// ff:
// uFormat:
func GetClipboardData(uFormat uint) HANDLE {
	ret, _, _ := procGetClipboardData.Call(
		uintptr(uFormat))
	return HANDLE(ret)
}


// ff:
// hMem:
// uFormat:
func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
	ret, _, _ := procSetClipboardData.Call(
		uintptr(uFormat),
		uintptr(hMem))
	return HANDLE(ret)
}


// ff:
func EmptyClipboard() bool {
	ret, _, _ := procEmptyClipboard.Call()
	return ret != 0
}


// ff:
// format:
func GetClipboardFormatName(format uint) (string, bool) {
	cchMaxCount := 255
	buf := make([]uint16, cchMaxCount)
	ret, _, _ := procGetClipboardFormatName.Call(
		uintptr(format),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(cchMaxCount))

	if ret > 0 {
		return syscall.UTF16ToString(buf), true
	}

	return "Requested format does not exist or is predefined", false
}


// ff:
// format:
func IsClipboardFormatAvailable(format uint) bool {
	ret, _, _ := procIsClipboardFormatAvailable.Call(uintptr(format))
	return ret != 0
}


// ff:
// paint:
// hwnd:
func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) HDC {
	ret, _, _ := procBeginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
	return HDC(ret)
}


// ff:
// paint:
// hwnd:
func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
	procEndPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)))
}


// ff:
// keyState:
func GetKeyboardState(keyState []byte) bool {
	if len(keyState) != 256 {
		panic("keyState slice must have a size of 256 bytes")
	}
	ret, _, _ := procGetKeyboardState.Call(uintptr(unsafe.Pointer(&keyState[0])))
	return ret != 0
}


// ff:
// dwhkl:
// uMapType:
// uCode:
func MapVirtualKeyEx(uCode, uMapType uint, dwhkl HKL) uint {
	ret, _, _ := procMapVirtualKey.Call(
		uintptr(uCode),
		uintptr(uMapType),
		uintptr(dwhkl))
	return uint(ret)
}


// ff:
// vKey:
func GetAsyncKeyState(vKey int) uint16 {
	ret, _, _ := procGetAsyncKeyState.Call(uintptr(vKey))
	return uint16(ret)
}


// ff:
// uFlags:
// lpChar:
// lpKeyState:
// uScanCode:
// uVirtKey:
func ToAscii(uVirtKey, uScanCode uint, lpKeyState *byte, lpChar *uint16, uFlags uint) int {
	ret, _, _ := procToAscii.Call(
		uintptr(uVirtKey),
		uintptr(uScanCode),
		uintptr(unsafe.Pointer(lpKeyState)),
		uintptr(unsafe.Pointer(lpChar)),
		uintptr(uFlags))
	return int(ret)
}


// ff:
// fSwap:
func SwapMouseButton(fSwap bool) bool {
	ret, _, _ := procSwapMouseButton.Call(
		uintptr(BoolToBOOL(fSwap)))
	return ret != 0
}


// ff:
// ok:
// y:
// x:
func GetCursorPos() (x, y int, ok bool) {
	pt := POINT{}
	ret, _, _ := procGetCursorPos.Call(uintptr(unsafe.Pointer(&pt)))
	return int(pt.X), int(pt.Y), ret != 0
}


// ff:
// y:
// x:
func SetCursorPos(x, y int) bool {
	ret, _, _ := procSetCursorPos.Call(
		uintptr(x),
		uintptr(y),
	)
	return ret != 0
}


// ff:
// cursor:
func SetCursor(cursor HCURSOR) HCURSOR {
	ret, _, _ := procSetCursor.Call(
		uintptr(cursor),
	)
	return HCURSOR(ret)
}


// ff:
// XORbits:
// ANDbits:
// cBitsPerPixel:
// cPlanes:
// nHeight:
// nWidth:
// instance:
func CreateIcon(instance HINSTANCE, nWidth, nHeight int, cPlanes, cBitsPerPixel byte, ANDbits, XORbits *byte) HICON {
	ret, _, _ := procCreateIcon.Call(
		uintptr(instance),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(cPlanes),
		uintptr(cBitsPerPixel),
		uintptr(unsafe.Pointer(ANDbits)),
		uintptr(unsafe.Pointer(XORbits)),
	)
	return HICON(ret)
}


// ff:
// icon:
func DestroyIcon(icon HICON) bool {
	ret, _, _ := procDestroyIcon.Call(
		uintptr(icon),
	)
	return ret != 0
}


// ff:
// dwFlags:
// y:
// x:
func MonitorFromPoint(x, y int, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromPoint.Call(
		uintptr(x),
		uintptr(y),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}


// ff:
// dwFlags:
// rc:
func MonitorFromRect(rc *RECT, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromRect.Call(
		uintptr(unsafe.Pointer(rc)),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}


// ff:
// dwFlags:
// hwnd:
func MonitorFromWindow(hwnd HWND, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		uintptr(hwnd),
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}


// ff:
// lmpi:
// hMonitor:
func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}


// ff:
// dwData:
// fnEnum:
// clip:
// hdc:
func EnumDisplayMonitors(hdc HDC, clip *RECT, fnEnum uintptr, dwData unsafe.Pointer) bool {
	ret, _, _ := procEnumDisplayMonitors.Call(
		hdc,
		uintptr(unsafe.Pointer(clip)),
		fnEnum,
		uintptr(dwData),
	)
	return ret != 0
}


// ff:
// dwFlags:
// devMode:
// iModeNum:
// szDeviceName:
func EnumDisplaySettingsEx(szDeviceName *uint16, iModeNum uint32, devMode *DEVMODE, dwFlags uint32) bool {
	ret, _, _ := procEnumDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(iModeNum),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(dwFlags),
	)
	return ret != 0
}


// ff:
// lParam:
// dwFlags:
// hwnd:
// devMode:
// szDeviceName:
func ChangeDisplaySettingsEx(szDeviceName *uint16, devMode *DEVMODE, hwnd HWND, dwFlags uint32, lParam uintptr) int32 {
	ret, _, _ := procChangeDisplaySettingsEx.Call(
		uintptr(unsafe.Pointer(szDeviceName)),
		uintptr(unsafe.Pointer(devMode)),
		uintptr(hwnd),
		uintptr(dwFlags),
		lParam,
	)
	return int32(ret)
}

/*

// ff:
// inputs:
func SendInput(inputs []INPUT) uint32 {
	var validInputs []C.INPUT

	for _, oneInput := range inputs {
		input := C.INPUT{_type: C.DWORD(oneInput.Type)}

		switch oneInput.Type {
		case INPUT_MOUSE:
			(*MouseInput)(unsafe.Pointer(&input)).mi = oneInput.Mi
		case INPUT_KEYBOARD:
			(*KbdInput)(unsafe.Pointer(&input)).ki = oneInput.Ki
		case INPUT_HARDWARE:
			(*HardwareInput)(unsafe.Pointer(&input)).hi = oneInput.Hi
		default:
			panic("unkown type")
		}

		validInputs = append(validInputs, input)
	}

	ret, _, _ := procSendInput.Call(
		uintptr(len(validInputs)),
		uintptr(unsafe.Pointer(&validInputs[0])),
		uintptr(unsafe.Sizeof(C.INPUT{})),
	)
	return uint32(ret)
}*/


// ff:
// dwThreadId:
// hMod:
// lpfn:
// idHook:
func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := procSetWindowsHookEx.Call(
		uintptr(idHook),
		uintptr(syscall.NewCallback(lpfn)),
		uintptr(hMod),
		uintptr(dwThreadId),
	)
	return HHOOK(ret)
}


// ff:
// hhk:
func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := procUnhookWindowsHookEx.Call(
		uintptr(hhk),
	)
	return ret != 0
}


// ff:
// lParam:
// wParam:
// nCode:
// hhk:
func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := procCallNextHookEx.Call(
		uintptr(hhk),
		uintptr(nCode),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
}


// ff:
// nVirtKey:
func GetKeyState(nVirtKey int32) int16 {
	ret, _, _ := syscall.Syscall(getKeyState, 1,
		uintptr(nVirtKey),
		0,
		0)

	return int16(ret)
}


// ff:
// hMenu:
func DestroyMenu(hMenu HMENU) bool {
	ret, _, _ := procDestroyMenu.Call(1,
		uintptr(hMenu),
		0,
		0)

	return ret != 0
}


// ff:
// lpwndpl:
// hWnd:
func GetWindowPlacement(hWnd HWND, lpwndpl *WINDOWPLACEMENT) bool {
	ret, _, _ := syscall.Syscall(getWindowPlacement, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)

	return ret != 0
}


// ff:
// lpwndpl:
// hWnd:
func SetWindowPlacement(hWnd HWND, lpwndpl *WINDOWPLACEMENT) bool {
	ret, _, _ := syscall.Syscall(setWindowPlacement, 2,
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpwndpl)),
		0)

	return ret != 0
}


// ff:
// fRedraw:
// lpsi:
// fnBar:
// hwnd:
func SetScrollInfo(hwnd HWND, fnBar int32, lpsi *SCROLLINFO, fRedraw bool) int32 {
	ret, _, _ := syscall.Syscall6(setScrollInfo, 4,
		hwnd,
		uintptr(fnBar),
		uintptr(unsafe.Pointer(lpsi)),
		uintptr(BoolToBOOL(fRedraw)),
		0,
		0)

	return int32(ret)
}


// ff:
// lpsi:
// fnBar:
// hwnd:
func GetScrollInfo(hwnd HWND, fnBar int32, lpsi *SCROLLINFO) bool {
	ret, _, _ := syscall.Syscall(getScrollInfo, 3,
		hwnd,
		uintptr(fnBar),
		uintptr(unsafe.Pointer(lpsi)))

	return ret != 0
}
