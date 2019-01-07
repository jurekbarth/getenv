package getenv

import (
	"os"
	"reflect"
)

type varInfo struct {
	Name  string
	Alt   string
	Field reflect.Value
	Tags  reflect.StructTag
}

func gatherInfo(prefix string, spec interface{}) []varInfo {
	s := reflect.ValueOf(spec)
	s = s.Elem()
	typeOfSpec := s.Type()
	infos := make([]varInfo, 0, s.NumField())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		ftype := typeOfSpec.Field(i)

		info := varInfo{
			Name:  ftype.Name,
			Field: f,
			Tags:  ftype.Tag,
			Alt:   prefix + ftype.Tag.Get("envconfig"),
		}
		infos = append(infos, info)
	}
	return infos
}

// Process sets env vars to struct spec
func Process(prefix string, spec interface{}) {
	infos := gatherInfo(prefix, spec)
	for _, info := range infos {
		value, ok := os.LookupEnv(info.Name)
		if !ok && info.Alt != "" {
			value, ok = os.LookupEnv(info.Alt)
		}
		if ok {
			info.Field.SetString(value)
		}
	}
}
