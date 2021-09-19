package rrule

import (
	"time"

	"github.com/teambition/rrule-go"
)

// RFC 5545规则检测

func CheckUseful(rule string) bool {
	rpc, err := rrule.StrToROptionInLocation(rule, time.Local)
	if err != nil {
		return false
	}
	_, err = rrule.NewRRule(*rpc)
	if err != nil {
		return false
	}
	return true
}

func GetRuleMaxTime(rule string) uint {
	rpc, err := rrule.StrToROptionInLocation(rule, time.Local)
	if err != nil {
		return 0
	}
	r, err := rrule.NewRRule(*rpc)
	if err != nil {
		return 0
	}
	allTimes := r.All()
	if len(allTimes) <= 0 {
		return 0
	}
	return uint(allTimes[len(allTimes)-1].Unix())
}

//inc 是判断开闭区间
func GetBetweenTimes(rule string, start, end time.Time, inc bool) []time.Time {
	rpc, err := rrule.StrToROptionInLocation(rule, time.Local)
	if err != nil {
		return nil
	}
	r, err := rrule.NewRRule(*rpc)
	if err != nil {
		return nil
	}
	return r.Between(start, end, inc)
}
