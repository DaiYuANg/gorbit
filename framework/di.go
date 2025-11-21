package framework

import (
	"fmt"
	"reflect"

	"github.com/samber/do/v2"
)

func ProvideFunc[T any](i do.Injector, fn interface{}) {
	do.Provide(i, func(inj do.Injector) (T, error) {
		return callConstructor[T](inj, fn)
	})
}

func callConstructor[T any](inj do.Injector, fn interface{}) (T, error) {
	var zero T

	vfn := reflect.ValueOf(fn)
	tfn := vfn.Type()

	if tfn.Kind() != reflect.Func {
		return zero, fmt.Errorf("constructor must be a function")
	}

	// 构建参数
	args := make([]reflect.Value, tfn.NumIn())
	for idx := 0; idx < tfn.NumIn(); idx++ {
		paramType := tfn.In(idx)

		// --- 核心：让 do 根据 reflect.Type 解析依赖 ---
		dep := reflect.ValueOf(do.MustInvoke[T](inj))
		if !dep.IsValid() {
			return zero, fmt.Errorf("cannot resolve dependency: %v", paramType)
		}
		args[idx] = dep
	}

	// 调用构造函数
	outs := vfn.Call(args)

	if len(outs) == 1 {
		return outs[0].Interface().(T), nil
	}

	if len(outs) == 2 {
		if err, ok := outs[1].Interface().(error); ok && err != nil {
			return zero, err
		}
		return outs[0].Interface().(T), nil
	}

	return zero, fmt.Errorf("constructor must return T or (T, error)")
}
