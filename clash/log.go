package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/log"
	"github.com/Dreamacro/clash/tunnel/statistic"
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

//export getRealtimeTrafficStatistic
func getRealtimeTrafficStatistic(upload, download *C.uint64_t) {
	up, down := statistic.DefaultManager.Now()

	*upload = C.uint64_t(up)
	*download = C.uint64_t(down)
}

//export getTotalTrafficStatistic
func getTotalTrafficStatistic(upload, download *C.uint64_t) {
	snapshot := statistic.DefaultManager.Snapshot()

	*upload = C.uint64_t(snapshot.UploadTotal)
	*download = C.uint64_t(snapshot.DownloadTotal)
}

//export resetTrafficStatistic
func resetTrafficStatistic() {
	statistic.DefaultManager.ResetStatistic()
}
