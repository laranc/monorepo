package ecs

import "reflect"

func ComponentType[T Component]() reflect.Type {
	t := *new(T)
	return t.Type()
}

func ComponentCast[T any](component Component) (T, bool) {
	v, ok := component.(T)
	return v, ok
}

func ComponentCastN[T any](components []Component) []T {
	result := make([]T, len(components))
	var ok bool
	for i, c := range components {
		result[i], ok = c.(T)
		if !ok {
			return nil
		}
	}
	return result
}

func ResourceType[T Resource]() reflect.Type {
	t := *new(T)
	return t.Type()
}

func ResourceCast[T any](resource Resource) (T, bool) {
	v, ok := resource.(T)
	return v, ok
}
