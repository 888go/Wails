//go:build linux
// +build linux

package webview

/*
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0 gio-unix-2.0

#include "gtk/gtk.h"
#include "webkit2/webkit2.h"
#include "gio/gunixinputstream.h"

*/
import "C"
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

type responseWriter struct {
	req *C.WebKitURISchemeRequest

	header      http.Header
	wroteHeader bool
	finished    bool

	w    io.WriteCloser
	wErr error
}


// ff:请求头
func (rw *responseWriter) X请求头() http.Header {
	if rw.header == nil {
		rw.header = http.Header{}
	}
	return rw.header
}


// ff:
// buf:
func (rw *responseWriter) Write(buf []byte) (int, error) {
	if rw.finished {
		return 0, errResponseFinished
	}

	rw.WriteHeader(http.StatusOK)
	if rw.wErr != nil {
		return 0, rw.wErr
	}
	return rw.w.Write(buf)
}


// ff:
// code:
func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader || rw.finished {
		return
	}
	rw.wroteHeader = true

	contentLength := int64(-1)
	if sLen := rw.X请求头().Get(HeaderContentLength); sLen != "" {
		if pLen, _ := strconv.ParseInt(sLen, 10, 64); pLen > 0 {
			contentLength = pLen
		}
	}

// 在这里我们不能使用os.Pipe，因为它返回的文件带有用于关闭文件描述符的终结器。但是，读取文件描述符的控制权交给了InputStream，并将在那里被关闭。
// 此外，我们特别不希望拥有FD_CLOEXEC
	rFD, w, err := pipe()
	if err != nil {
		rw.finishWithError(http.StatusInternalServerError, fmt.Errorf("unable to open pipe: %s", err))
		return
	}
	rw.w = w

	stream := C.g_unix_input_stream_new(C.int(rFD), C.gboolean(1))
	defer C.g_object_unref(C.gpointer(stream))

	if err := webkit_uri_scheme_request_finish(rw.req, code, rw.X请求头(), stream, contentLength); err != nil {
		rw.finishWithError(http.StatusInternalServerError, fmt.Errorf("unable to finish request: %s", err))
		return
	}
}


// ff:
func (rw *responseWriter) Finish() error {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusNotImplemented)
	}

	if rw.finished {
		return nil
	}
	rw.finished = true
	if rw.w != nil {
		rw.w.Close()
	}
	return nil
}

func (rw *responseWriter) finishWithError(code int, err error) {
	if rw.w != nil {
		rw.w.Close()
		rw.w = &nopCloser{io.Discard}
	}
	rw.wErr = err

	msg := C.CString(err.Error())
	gerr := C.g_error_new_literal(C.g_quark_from_string(msg), C.int(code), msg)
	C.webkit_uri_scheme_request_finish_error(rw.req, gerr)
	C.g_error_free(gerr)
	C.free(unsafe.Pointer(msg))
}

type nopCloser struct {
	io.Writer
}


// ff:关闭
func (nopCloser) X关闭() error { return nil }

func pipe() (r int, w *os.File, err error) {
	var p [2]int
	e := syscall.Pipe2(p[0:], 0)
	if e != nil {
		return 0, nil, fmt.Errorf("pipe2: %s", e)
	}

	return p[0], os.NewFile(uintptr(p[1]), "|1"), nil
}
