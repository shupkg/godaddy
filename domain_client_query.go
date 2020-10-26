package godaddy

import (
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Q(key string, values ...interface{}) Query {
	return Query{}.Set(key, values...)
}

func ParseStruct(i interface{}) Query {
	var query = Query{}
	v := reflect.Indirect(reflect.ValueOf(i))
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		cacheKey := t.PkgPath() + "." + t.Name() + "[" + strconv.Itoa(i) + "]"
		qf, find := fieldCache[cacheKey]
		if !find {
			fieldT := t.Field(i)
			qf.Key = fieldT.Tag.Get("query")
			if qf.Key == "" {
				qf.Key = fieldT.Tag.Get("json")
			}
			if qf.Key != "" {
				qf.Omitempty = strings.Contains(qf.Key, "omitempty")
				if idx := strings.IndexAny(qf.Key, ",; "); idx != -1 {
					qf.Key = strings.TrimSpace(qf.Key[:idx])
				}
			}
			if qf.Key == "" {
				qf.Key = fieldT.Name
			}
			fieldCache[cacheKey] = qf
		}

		vi := v.Field(i).Interface()
		if qf.Omitempty && vi == nil {
			continue
		}
		query[qf.Key] = query.format(vi, qf.Omitempty)
	}
	return query
}

type Query map[string][]string

func (v Query) Get(key string) string {
	key = strings.TrimSpace(key)
	return strings.Join(v[key], ",")
}

// Set sets the key to value. It replaces any existing values.
func (v Query) Set(key string, values ...interface{}) Query {
	key = strings.TrimSpace(key)
	v[key] = v.format(values, true)
	return v
}

// Del deletes the values associated with key.
func (v Query) Del(key string) Query {
	key = strings.TrimSpace(key)
	delete(v, key)
	return v
}

// Add adds the value to key. It appends to any existing values associated with key.
func (v Query) Add(key string, values ...interface{}) Query {
	key = strings.TrimSpace(key)
	v[key] = append(v[key], v.format(values, true)...)
	return v
}

// Encode encodes the values into ``URL encoded'' form ("bar=baz&foo=quux") sorted by key.
func (v Query) Encode() string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		if len(vs) > 0 {
			keyEscaped := url.QueryEscape(k)
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(strings.Join(vs, ",")))
		}
	}
	return buf.String()
}

func (v Query) FormEncode() string {
	return url.Values(v).Encode()
}

func (v Query) format(i interface{}, omitempty bool) []string {
	if i == nil {
		if omitempty {
			return []string{""}
		}
		return nil
	}

	switch o := i.(type) {
	case []string:
		if len(o) == 0 && !omitempty {
			return []string{""}
		}
		return o
	case string:
		return v._format(o == "", omitempty, func() string { return o }, "")
	case time.Time:
		return v._format(o.IsZero(), omitempty, func() string { return o.Format(time.RFC3339) }, "")
	case []byte:
		return v.format(string(o), omitempty)
	case int:
		return v._format(o == 0, omitempty, func() string { return strconv.Itoa(o) }, "0")
	case int8:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatInt(int64(o), 10) }, "0")
	case int16:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatInt(int64(o), 10) }, "0")
	case int32:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatInt(int64(o), 10) }, "0")
	case int64:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatInt(o, 10) }, "0")
	case uint:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatUint(uint64(o), 10) }, "0")
	case uint8:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatUint(uint64(o), 10) }, "0")
	case uint16:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatUint(uint64(o), 10) }, "0")
	case uint32:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatUint(uint64(o), 10) }, "0")
	case uint64:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatUint(o, 10) }, "0")
	case float32:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatFloat(float64(o), 'f', -1, 64) }, "0")
	case float64:
		return v._format(o == 0, omitempty, func() string { return strconv.FormatFloat(o, 'f', -1, 64) }, "0")
	case bool:
		return v._format(!o, omitempty, func() string { return strconv.FormatBool(o) }, "false")
	case fmt.Stringer:
		s := o.String()
		return v._format(s == "", omitempty, func() string { return s })
	case fmt.GoStringer:
		s := o.GoString()
		return v._format(s == "", omitempty, func() string { return s })
	}

	refV := reflect.ValueOf(i)
	switch refV.Kind() {
	case reflect.Ptr:
		return v.format(refV.Elem().Interface(), omitempty)
	case reflect.Slice:
		var result []string
		for i := 0; i < refV.Len(); i++ {
			result = append(result, v.format(refV.Index(i).Interface(), omitempty)...)
		}
		return result
	}

	return []string{fmt.Sprintf("%v", i)}
}

func (v Query) _format(zero, omitempty bool, stringer func() string, zeros ...string) []string {
	if !zero {
		return []string{stringer()}
	}

	if !omitempty {
		if len(zeros) == 0 {
			return []string{""}
		}
		return zeros
	}
	return nil
}

var fieldCache = map[string]queryField{}

type queryField struct {
	Omitempty bool
	Key       string
}
