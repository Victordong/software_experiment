package underscore

import (
	"reflect"
)

// Keys is 获取map的所有key
func Keys(source interface{}) interface{} {
	sourceRV := reflect.ValueOf(source)
	if sourceRV.Kind() != reflect.Map {
		return nil
	}

	return Map(source, func(_, key interface{}) Facade {
		return Facade{reflect.ValueOf(key)}
	})
}

// Keys is Queryer's method
func (q *Query) Keys() Queryer {
	q.source = Keys(q.source)
	return q
}
