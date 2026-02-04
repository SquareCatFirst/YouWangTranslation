package util

import (
	"strconv"
	"strings"
)

func AsMap(v interface{}) (map[string]interface{}, bool) {
	m, ok := v.(map[string]interface{})
	return m, ok
}

func AsSlice(v interface{}) ([]interface{}, bool) {
	s, ok := v.([]interface{})
	return s, ok
}

func GetString(m map[string]interface{}, key string) string {
	if m == nil {
		return ""
	}
	v, ok := m[key]
	if !ok || v == nil {
		return ""
	}
	s, ok := v.(string)
	if !ok {
		return ""
	}
	return strings.TrimSpace(s)
}

func GetPtrString(m map[string]interface{}, key string) *string {
	s := GetString(m, key)
	if s == "" {
		return nil
	}
	return &s
}

func GetBool(m map[string]interface{}, key string) bool {
	if m == nil {
		return false
	}
	v, ok := m[key]
	if !ok || v == nil {
		return false
	}
	b, ok := v.(bool)
	if !ok {
		return false
	}
	return b
}

func AsInt(v interface{}, def int) int {
	switch t := v.(type) {
	case float64:
		return int(t)
	case int:
		return t
	case int64:
		return int(t)
	case string:
		n, err := strconv.Atoi(strings.TrimSpace(t))
		if err == nil {
			return n
		}
	}
	return def
}

func AsInt64(v interface{}, def int64) int64 {
	switch t := v.(type) {
	case float64:
		return int64(t)
	case int64:
		return t
	case int:
		return int64(t)
	case string:
		n, err := strconv.ParseInt(strings.TrimSpace(t), 10, 64)
		if err == nil {
			return n
		}
	}
	return def
}

func GetInt(m map[string]interface{}, key string) int {
	if m == nil {
		return 0
	}
	v, ok := m[key]
	if !ok || v == nil {
		return 0
	}
	return AsInt(v, 0)
}

func GetInt64(m map[string]interface{}, key string) int64 {
	if m == nil {
		return 0
	}
	v, ok := m[key]
	if !ok || v == nil {
		return 0
	}
	return AsInt64(v, 0)
}

func GetPtrInt64(m map[string]interface{}, key string) *int64 {
	id := GetInt64(m, key)
	if id == 0 {
		return nil
	}
	x := id
	return &x
}
