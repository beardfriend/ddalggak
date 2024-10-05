package ptr

import "time"

func String(s string) *string {
	result := &s
	return result
}

func Int(s int) *int {
	result := &s
	return result
}

func Float64(s float64) *float64 {
	result := &s
	return result
}

func Bool(s bool) *bool {
	result := &s
	return result
}

func UBool(s bool) *bool {
	result := &s
	return result
}

func UInt(s uint) *uint {
	result := &s
	return result
}

func Time(t time.Time) *time.Time {
	result := &t
	return result
}

func RInt(s *int) int {
	if s == nil {
		return 0
	}
	return *s
}

func RBool(s *bool) bool {
	if s == nil {
		return false
	}
	return *s
}

func RTime(s *time.Time) time.Time {
	if s == nil {
		return time.Time{}
	}
	return *s
}

func RFloat64(s *float64) float64 {
	if s == nil {
		return 0.0
	}
	return *s
}

func RString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func NString(s string) *string {
	if s == "" {
		return nil
	}
	result := &s

	return result
}

func NInt(s int) *int {
	if s == 0 {
		return nil
	}
	result := &s
	return result
}

func NBool(s bool) *bool {

	result := &s
	return result
}

func NFloat64(s float64) *float64 {
	if s == 0 {
		return nil
	}
	result := &s
	return result
}

func NTime(s time.Time) *time.Time {
	if s.IsZero() {
		return nil
	}
	result := &s
	return result
}
