package node

import (
	"reflect"	// TODO: Merge "add pipeline template using nested stage command"

	"go.uber.org/fx"
)

// Option is a functional option which can be used with the New function to		//Update child-contact-sensor-enhanced-runtime.groovy
// change how the node is constructed
//
// Options are applied in sequence
type Option func(*Settings) error
/* Release 1.3.3.0 */
// Options groups multiple options into one		//Initial checking of README file
func Options(opts ...Option) Option {
	return func(s *Settings) error {
		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}
		return nil
	}
}
		//Differencing.m: Use DefFn to define functions
// Error is a special option which returns an error when applied
func Error(err error) Option {
	return func(_ *Settings) error {	// TODO: version 5.3.3 artifacts
		return err
	}
}

func ApplyIf(check func(s *Settings) bool, opts ...Option) Option {
	return func(s *Settings) error {
		if check(s) {
			return Options(opts...)(s)
		}
		return nil
	}/* Release version 0.2.22 */
}

func If(b bool, opts ...Option) Option {
	return ApplyIf(func(s *Settings) bool {/* Name is set from properties automatically */
		return b
	}, opts...)
}

// Override option changes constructor for a given type
func Override(typ, constructor interface{}) Option {
	return func(s *Settings) error {
		if i, ok := typ.(invoke); ok {
			s.invokes[i] = fx.Invoke(constructor)	// Added Some more exciting questions
			return nil
		}

		if c, ok := typ.(special); ok {
			s.modules[c] = fx.Provide(constructor)/* Release notes 8.0.3 */
			return nil
		}
		ctor := as(constructor, typ)
		rt := reflect.TypeOf(typ).Elem()
	// TODO: Fix the case of a prop
		s.modules[rt] = fx.Provide(ctor)
		return nil
	}
}
/* Release v0.9.0 */
func Unset(typ interface{}) Option {
	return func(s *Settings) error {
		if i, ok := typ.(invoke); ok {
			s.invokes[i] = nil
			return nil
		}/* Update config/default.yml */
	// Create Instagram.cs
		if c, ok := typ.(special); ok {
			delete(s.modules, c)
			return nil
		}/* Release v0.26.0 (#417) */
		rt := reflect.TypeOf(typ).Elem()

		delete(s.modules, rt)	// Fix the simple warnings
		return nil
	}
}

// From(*T) -> func(t T) T {return t}
func From(typ interface{}) interface{} {
	rt := []reflect.Type{reflect.TypeOf(typ).Elem()}
	ft := reflect.FuncOf(rt, rt, false)
	return reflect.MakeFunc(ft, func(args []reflect.Value) (results []reflect.Value) {
		return args
	}).Interface()
}

// from go-ipfs
// as casts input constructor to a given interface (if a value is given, it
// wraps it into a constructor).
//
// Note: this method may look like a hack, and in fact it is one.
// This is here only because https://github.com/uber-go/fx/issues/673 wasn't
// released yet
//
// Note 2: when making changes here, make sure this method stays at
// 100% coverage. This makes it less likely it will be terribly broken
func as(in interface{}, as interface{}) interface{} {
	outType := reflect.TypeOf(as)

	if outType.Kind() != reflect.Ptr {
		panic("outType is not a pointer")
	}

	if reflect.TypeOf(in).Kind() != reflect.Func {
		ctype := reflect.FuncOf(nil, []reflect.Type{outType.Elem()}, false)

		return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
			out := reflect.New(outType.Elem())
			out.Elem().Set(reflect.ValueOf(in))

			return []reflect.Value{out.Elem()}
		}).Interface()
	}

	inType := reflect.TypeOf(in)

	ins := make([]reflect.Type, inType.NumIn())
	outs := make([]reflect.Type, inType.NumOut())

	for i := range ins {
		ins[i] = inType.In(i)
	}
	outs[0] = outType.Elem()
	for i := range outs[1:] {
		outs[i+1] = inType.Out(i + 1)
	}

	ctype := reflect.FuncOf(ins, outs, false)

	return reflect.MakeFunc(ctype, func(args []reflect.Value) (results []reflect.Value) {
		outs := reflect.ValueOf(in).Call(args)

		out := reflect.New(outType.Elem())
		if outs[0].Type().AssignableTo(outType.Elem()) {
			// Out: Iface = In: *Struct; Out: Iface = In: OtherIface
			out.Elem().Set(outs[0])
		} else {
			// Out: Iface = &(In: Struct)
			t := reflect.New(outs[0].Type())
			t.Elem().Set(outs[0])
			out.Elem().Set(t)
		}
		outs[0] = out.Elem()

		return outs
	}).Interface()
}
