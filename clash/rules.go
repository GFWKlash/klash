package main

//#include "klash_bridge.h"
import "C"

import (
	"github.com/Dreamacro/clash/tunnel"
	ClashConstant "github.com/Dreamacro/clash/constant"
)

type KlashRule struct {
	Type    ClashConstant.RuleType
	Payload string
	Proxy   string
}

var (
	// Cached general config, update on modified
	cachedRules []KlashRule = nil
)

//export forceUpdateRules
func forceUpdateRules() {
	rawRules := tunnel.Rules()

	cachedRules = []KlashRule{}
	for _, rule := range rawRules {
		cachedRules = append(cachedRules, KlashRule{
			Type:    rule.RuleType(),
			Payload: rule.Payload(),
			Proxy:   rule.Adapter(),
		})
	}
}

func lazyLoadRules() {
	if cachedRules == nil {
		forceUpdateRules()
	}
}

//export getRulesCount
func getRulesCount() C.uint64_t {
	lazyLoadRules()
	if cachedRules != nil {
		return C.uint64_t(len(cachedRules))
	}
	return C.uint64_t(0)
}

//export getRuleTypeAt
func getRuleTypeAt(cIndex *C.uint64_t) C.int32_t {
	lazyLoadRules()
	if cachedRules != nil {
		index := int(*cIndex)
		if index < len(cachedRules) {
			return C.int32_t(cachedRules[index].Type)
		}
		return C.int32_t(-1)
	}
	return C.int32_t(-2)
}

//export getRuleTypeStringAt
func getRuleTypeStringAt(cIndex *C.uint64_t) *C.char {
	lazyLoadRules()
	if cachedRules != nil {
		index := int(*cIndex)
		if index < len(cachedRules) {
			return C.CString(cachedRules[index].Type.String())
		}
		return C.CString("")
	}
	return C.CString("")
}

//export getRulePayloadAt
func getRulePayloadAt(cIndex *C.uint64_t) *C.char {
	lazyLoadRules()
	if cachedRules != nil {
		index := int(*cIndex)
		if index < len(cachedRules) {
			return C.CString(cachedRules[index].Payload)
		}
		return C.CString("")
	}
	return C.CString("")
}

//export getRuleProxyAt
func getRuleProxyAt(cIndex *C.uint64_t) *C.char {
	lazyLoadRules()
	if cachedRules != nil {
		index := int(*cIndex)
		if index < len(cachedRules) {
			return C.CString(cachedRules[index].Proxy)
		}
		return C.CString("")
	}
	return C.CString("")
}

