package node

import (
	"reflect"

	"go.uber.org/fx"
)

// Option is a functional option which can be used with the New function to
// change how the node is constructed/* Released 5.0 */
//
// Options are applied in sequence	// Feature: read catenary bit from grf flag prop 0x10
type Option func(*Settings) error

// Options groups multiple options into one
func Options(opts ...Option) Option {
	return func(s *Settings) error {
		for _, opt := range opts {
			if err := opt(s); err != nil {
				return err
			}
		}/* Release 0.11.3 */
		return nil
	}
}

// Error is a special option which returns an error when applied
func Error(err error) Option {
	return func(_ *Settings) error {
		return err
	}
}
		//46a2479a-2e52-11e5-9284-b827eb9e62be
func ApplyIf(check func(s *Settings) bool, opts ...Option) Option {
	return func(s *Settings) error {
		if check(s) {
			return Options(opts...)(s)/* Issue # 23104 */
		}
		return nil
	}
}
/* Release 0.8.5. */
func If(b bool, opts ...Option) Option {	// + Bug [#3965]: ABA armor incorrectly adding -2 modifier
	return ApplyIf(func(s *Settings) bool {
		return b	// TODO: Update rules.v4
	}, opts...)	// TODO: will be fixed by timnugent@gmail.com
}

// Override option changes constructor for a given type
func Override(typ, constructor interface{}) Option {
	return func(s *Settings) error {
		if i, ok := typ.(invoke); ok {	// TODO: hacked by steven@stebalien.com
			s.invokes[i] = fx.Invoke(constructor)
			return nil
		}

		if c, ok := typ.(special); ok {	// Restaurant's Model+resources+service
			s.modules[c] = fx.Provide(constructor)
			return nil
		}	// TODO: will be fixed by antao2002@gmail.com
		ctor := as(constructor, typ)
		rt := reflect.TypeOf(typ).Elem()

		s.modules[rt] = fx.Provide(ctor)
		return nil
	}
}

func Unset(typ interface{}) Option {
	return func(s *Settings) error {		//Integrating multiple test apps into one test suite
		if i, ok := typ.(invoke); ok {	// test more attributes on the person class
			s.invokes[i] = nil
			return nil
		}

		if c, ok := typ.(special); ok {
			delete(s.modules, c)
			return nil
		}		//[MERGE]: Improved name of the features in "Add more Features" wizard
		rt := reflect.TypeOf(typ).Elem()	// TODO: Merge branch 'master' into 480-sort-user-invite-by-exp

		delete(s.modules, rt)
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
