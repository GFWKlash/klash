package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/hub/executor"
	"github.com/Dreamacro/clash/config"
	P "github.com/Dreamacro/clash/listener"
	"github.com/Dreamacro/clash/tunnel"
	"github.com/Dreamacro/clash/component/resolver"
)

var (
	// Cached general config, update on modified
	cachedConfig *config.General = nil
)

//export forceUpdateGeneralConfig
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

func notZeroOrDefault(p int, def int) int {
	if p != 0 {
		return p
	}

	return def
}

//export setConfigHTTPPort
func setConfigHTTPPort(port *C.uint16_t) bool {
	lazyLoadGeneralConfig()
	portInt := int(*port)

	ports := P.GetPorts()
	tcpIn := tunnel.TCPIn()
	if cachedConfig != nil {
		if cachedConfig.Inbound.Port == portInt {
			P.ReCreateHTTP(notZeroOrDefault(portInt, ports.Port), tcpIn)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.ReCreateHTTP(notZeroOrDefault(portInt, ports.Port), tcpIn)
	}
	return false
}

//export setConfigSocksPort
func setConfigSocksPort(port *C.uint16_t) bool {
	lazyLoadGeneralConfig()
	portInt := int(*port)

	ports := P.GetPorts()
	tcpIn := tunnel.TCPIn()
	udpIn := tunnel.UDPIn()
	if cachedConfig != nil {
		if cachedConfig.Inbound.SocksPort == portInt {
			P.ReCreateSocks(notZeroOrDefault(portInt, ports.SocksPort), tcpIn, udpIn)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.ReCreateSocks(notZeroOrDefault(portInt, ports.SocksPort), tcpIn, udpIn)
	}
	return false
}


//export setConfigRedirPort
func setConfigRedirPort(port *C.uint16_t) bool {
	lazyLoadGeneralConfig()
	portInt := int(*port)

	ports := P.GetPorts()
	tcpIn := tunnel.TCPIn()
	udpIn := tunnel.UDPIn()
	if cachedConfig != nil {
		if cachedConfig.Inbound.RedirPort == portInt {
			P.ReCreateRedir(notZeroOrDefault(portInt, ports.RedirPort), tcpIn, udpIn)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.ReCreateRedir(notZeroOrDefault(portInt, ports.RedirPort), tcpIn, udpIn)
	}
	return false
}

//export setConfigTProxyPort
func setConfigTProxyPort(port *C.uint16_t) bool {
	lazyLoadGeneralConfig()
	portInt := int(*port)

	ports := P.GetPorts()
	tcpIn := tunnel.TCPIn()
	udpIn := tunnel.UDPIn()
	if cachedConfig != nil {
		if cachedConfig.Inbound.TProxyPort == portInt {
			P.ReCreateTProxy(notZeroOrDefault(portInt, ports.TProxyPort), tcpIn, udpIn)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.ReCreateTProxy(notZeroOrDefault(portInt, ports.TProxyPort), tcpIn, udpIn)
	}
	return false
}

//export setConfigMixedPort
func setConfigMixedPort(port *C.uint16_t) bool {
	lazyLoadGeneralConfig()
	portInt := int(*port)

	ports := P.GetPorts()
	tcpIn := tunnel.TCPIn()
	udpIn := tunnel.UDPIn()
	if cachedConfig != nil {
		if cachedConfig.Inbound.MixedPort == portInt {
			P.ReCreateMixed(notZeroOrDefault(portInt, ports.MixedPort), tcpIn, udpIn)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.ReCreateMixed(notZeroOrDefault(portInt, ports.MixedPort), tcpIn, udpIn)
	}
	return false
}

//export setConfigTunnelMode
func setConfigTunnelMode(mode *C.int32_t) bool {
	lazyLoadGeneralConfig()
	modeEnum := tunnel.TunnelMode(int(*mode))
	if cachedConfig != nil {
		if cachedConfig.Mode != modeEnum {
			tunnel.SetMode(modeEnum)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		tunnel.SetMode(modeEnum)
	}
	return false
}

/*
TODO: Set log level
func setConfigLogLevelImpl(level *C.int32_t) bool {
	lazyLoadGeneralConfig()
}
*/

//export setConfigEnableIPv6
func setConfigEnableIPv6(enabled *C.int32_t) bool {
	lazyLoadGeneralConfig()
	enabledBool := (int(*enabled) != 0)
	if cachedConfig != nil {
		if cachedConfig.IPv6 != enabledBool {
			resolver.DisableIPv6 = !enabledBool
			cachedConfig.IPv6 = enabledBool
			return true
		}
	} else {
		// Force set and ignore the result
		resolver.DisableIPv6 = !enabledBool
	}
	return false
}

//export setConfigAllowLan
func setConfigAllowLan(allowed *C.int32_t) bool {
	lazyLoadGeneralConfig()
	allowedBool := (int(*allowed) != 0)
	if cachedConfig != nil {
		if cachedConfig.AllowLan != allowedBool {
			P.SetAllowLan(allowedBool)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.SetAllowLan(allowedBool)
	}
	return false
}

//export setConfigBoundAddress
func setConfigBoundAddress(addr *C.char) bool {
	lazyLoadGeneralConfig()
	addrString := C.GoString(addr)
	if cachedConfig != nil {
		if cachedConfig.BindAddress != addrString {
			P.SetBindAddress(addrString)
			forceUpdateGeneralConfig()
			return true
		}
	} else {
		// Force set and ignore the result
		P.SetBindAddress(C.GoString(addr))
	}
	return false
}
