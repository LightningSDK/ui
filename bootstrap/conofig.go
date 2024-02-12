package bootstrap

import (
	"reflect"
)

var Components = map[string]reflect.Type{
	"columns": reflect.TypeOf(Columns{}),
}
