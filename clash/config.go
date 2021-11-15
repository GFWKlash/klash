package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/config"
)

var (
	// Cached general config, update on modified
	cachedConfig *config.General = nil
)

func forceUpdateGeneralConfig() {
	cachedConfig = executor.GetGeneral()
}

func lazyLoadGeneralConfig() {
	if cachedConfig == nil {
		cachedConfig = executor.GetGeneral()
	}
}

//export getConfigHTTPPort
func getConfigHTTPPort(port *C.uint16_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*port = C.uint16_t(cachedConfig.Inbound.Port)
	} else {
		*port = C.uint16_t(0)
	}
}

//export getConfigSocksPort
func getConfigSocksPort(port *C.uint16_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*port = C.uint16_t(cachedConfig.Inbound.SocksPort)
	} else {
		*port = C.uint16_t(0)
	}
}


//export getConfigRedirPort
func getConfigRedirPort(port *C.uint16_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*port = C.uint16_t(cachedConfig.Inbound.RedirPort)
	} else {
		*port = C.uint16_t(0)
	}
}

//export getConfigTProxyPort
func getConfigTProxyPort(port *C.uint16_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*port = C.uint16_t(cachedConfig.Inbound.TProxyPort)
	} else {
		*port = C.uint16_t(0)
	}
}

//export getConfigMixedPort
func getConfigMixedPort(port *C.uint16_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*port = C.uint16_t(cachedConfig.Inbound.MixedPort)
	} else {
		*port = C.uint16_t(0)
	}
}

//export getConfigTunnelMode
func getConfigTunnelMode(mode *C.int32_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*mode = C.int32_t(cachedConfig.Mode)
	} else {
		*mode = C.int32_t(-1)
	}
}

func getConfigLogLevelImpl(level *C.int32_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		*level = C.int32_t(cachedConfig.LogLevel)
	} else {
		*level = C.int32_t(-1)
	}
}

//export getConfigEnableIPv6
func getConfigEnableIPv6(enabled *C.int32_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		if cachedConfig.IPv6 {
			*enabled = C.int32_t(1)
		} else {
			*enabled = C.int32_t(0)
		}
	} else {
		*enabled = C.int32_t(0)
	}
}

//export getConfigAllowLan
func getConfigAllowLan(allowed *C.int32_t) {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		if cachedConfig.AllowLan {
			*allowed = C.int32_t(1)
		} else {
			*allowed = C.int32_t(0)
		}
	} else {
		*allowed = C.int32_t(0)
	}
}

//export getConfigBoundAddress
func getConfigBoundAddress() *C.char {
	lazyLoadGeneralConfig()
	if cachedConfig != nil {
		return C.CString(cachedConfig.BindAddress)
	}
	return C.CString("")
}

/*
TODO:
Inbound: config.Inbound{
	Authentication: authenticator,
},
*/
