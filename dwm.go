// gradual translation of dwm to go via cgo
package main

/*
#cgo CFLAGS:  -std=c99 -Wall -Wno-deprecated-declarations -Os -I/usr/X11R6/include -I/usr/include/freetype2 -D_DEFAULT_SOURCE -D_BSD_SOURCE -D_POSIX_C_SOURCE=200809L -DVERSION="6.2" -DXINERAMA -Wno-unused-variable -Wno-unused-function 
#cgo LDFLAGS: -L/usr/X11R6/lib -lX11 -lXinerama -lfontconfig -lXft
#include <stddef.h>
#include <X11/cursorfont.h>
#include <X11/keysym.h>
#include <X11/Xatom.h>
#include <X11/Xlib.h>
#include <X11/Xproto.h>
#include <X11/Xutil.h>
#include <X11/Xft/Xft.h>
#include "util.h"
#include "drw.h"
#include "dwm.h"
void try_open_display(void);
void close_display(void);
*/
import "C"

import "os"

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-v" {
		die("dwm-custom")
	} else if len(os.Args) != 1 {
		die("usage: dwm [-v]")
	}
	C.try_open_display()

	C.checkotherwm()
	C.setup()
	C.scan()
	C.run()
	C.cleanup()
	C.close_display()
	os.Exit(0)
}

func die(s string) {
	println(s)
	os.Exit(1)
}
