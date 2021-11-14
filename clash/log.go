package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/log"
)

var (
	enableRedirect = false
)

//export redirectLogToKlash
func redirectLogToKlash() {
	enableRedirect = true
	go func() {
		sub := log.Subscribe()
		defer log.UnSubscribe(sub)

		for item := range sub {
			msg := item.(*log.Event)

			payload := C.CString(msg.Payload)

			switch msg.LogLevel {
			case log.INFO:
				C.klash_log_info(payload)
			case log.ERROR:
				C.klash_log_error(payload)
			case log.WARNING:
				C.klash_log_warn(payload)
			case log.DEBUG:
				C.klash_log_debug(payload)
			case log.SILENT:
				C.klash_log_verbose(payload)
			}

			if !enableRedirect {
				break
			}
		}
	}()
}

//export stopRedirectLogToKlash
func stopRedirectLogToKlash() {
	enableRedirect = false
}
