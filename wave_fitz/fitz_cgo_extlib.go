//go:build extlib && !pkgconfig
// +build extlib,!pkgconfig

package wave_fitz

/*
#cgo !static LDFLAGS: -lmupdf -lm
#cgo static LDFLAGS: -lmupdf -lm -lmupdf-third
#cgo android LDFLAGS: -llog
#cgo windows LDFLAGS: -lcomdlg32 -lgdi32
*/
import "C"
