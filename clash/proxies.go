package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/tunnel"
	ClashConstant "github.com/Dreamacro/clash/constant"
)

var (
	// Cached general config, update on modified
	cachedProxiesName []string = nil
)


//export forceUpdateProxies
func forceUpdateProxies() {
	proxies := tunnel.Proxies()
	updateProxies(proxies)
}

func updateProxies(proxies map[string]ClashConstant.Proxy) {
	cachedProxiesName = make([]string, 0, len(proxies))
	for k, _ := range proxies {
		cachedProxiesName = append(cachedProxiesName, k)
	}
}

func lazyLoadProxies() {
	if cachedProxiesName == nil {
		forceUpdateProxies()
	}
}

//export getProxiesCount
func getProxiesCount() C.uint64_t {
	lazyLoadProxies()
	if cachedProxiesName != nil {
		return C.uint64_t(len(cachedProxiesName))
	}
	return C.uint64_t(0)
}

//export getProxyKeyAt
func getProxyKeyAt(cIndex *C.uint64_t) *C.char {
	lazyLoadProxies()
	if cachedProxiesName != nil {
		index := int(*cIndex)
		if index < len(cachedProxiesName) {
			return C.CString(cachedProxiesName[index])
		}
		return C.CString("")
	}
	return C.CString("")
}

//export getProxyName
func getProxyName(k *C.char) *C.char {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return C.CString(proxy.Name())
	} else {
		// Update list
		updateProxies(proxies)
	}

	return C.CString("")
}

//export getProxyType
func getProxyType(k *C.char) C.int32_t {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return C.int32_t(proxy.Type())
	} else {
		// Update list
		updateProxies(proxies)
	}

	return C.int32_t(-1)
}

// TODO: Convert type to type string

//export getProxyLastDelay
func getProxyLastDelay(k *C.char) C.uint16_t {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return C.uint16_t(proxy.LastDelay())
	} else {
		// Update list
		updateProxies(proxies)
	}

	return C.uint16_t(0xFFFF)
}

//export getProxyAlive
func getProxyAlive(k *C.char) bool {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return proxy.Alive()
	} else {
		// Update list
		updateProxies(proxies)
	}

	return false
}

//export getProxySupportUDP
func getProxySupportUDP(k *C.char) bool {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return proxy.SupportUDP()
	} else {
		// Update list
		updateProxies(proxies)
	}

	return false
}

//export getProxyAddr
func getProxyAddr(k *C.char) *C.char {
	key := C.GoString(k)

	proxies := tunnel.Proxies()
	if proxy, found := proxies[key]; found {
		return C.CString(proxy.Addr())
	} else {
		// Update list
		updateProxies(proxies)
	}

	return C.CString("")
}
