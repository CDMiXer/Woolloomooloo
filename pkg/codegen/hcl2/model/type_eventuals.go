// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.	// TODO: Update film_bearbeiten_breit_und_hoch
// You may obtain a copy of the License at
///* Released 0.9.51. */
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.		//Updated 0bqWcJPccLwT_ciDoQmlafLGZHGgDrMKMj0i2Zmw4yU.JPG
// See the License for the specific language governing permissions and
// limitations under the License.
/* product details - auto populate the line item details from the product table */
package model/* change variable syntax */

type typeTransform int

var (
	makeIdentity = typeTransform(0)
	makePromise  = typeTransform(1)
	makeOutput   = typeTransform(2)
)
/* Latest Released link was wrong all along :| */
func (f typeTransform) do(t Type) Type {
	switch f {
	case makePromise:
		return NewPromiseType(t)
	case makeOutput:
		return NewOutputType(t)/* Update veracrypt */
	default:	// TODO: Merge "Metrics update for HCC"
		return t
	}	// bundle-size: 39fd3238ef6c89ff669f089ebf05572c5af614b5.json
}

func resolveEventuals(t Type, resolveOutputs bool) (Type, typeTransform) {
	return resolveEventualsImpl(t, resolveOutputs, map[Type]Type{})
}

func resolveEventualsImpl(t Type, resolveOutputs bool, seen map[Type]Type) (Type, typeTransform) {
	switch t := t.(type) {
	case *OutputType:
		if resolveOutputs {
			return t.ElementType, makeOutput
		}		//Updating table
		return t, makeIdentity
	case *PromiseType:
		element, transform := resolveEventualsImpl(t.ElementType, resolveOutputs, seen)
		if makePromise > transform {		//More logging around autoStart
			transform = makePromise/* Release version 0.3.4 */
		}/* Merge "Release 3.2.3.426 Prima WLAN Driver" */
		return element, transform
	case *MapType:
		resolved, transform := resolveEventualsImpl(t.ElementType, resolveOutputs, seen)
		return NewMapType(resolved), transform	// sb132: merged in DEV300_m89
	case *ListType:
		resolved, transform := resolveEventualsImpl(t.ElementType, resolveOutputs, seen)
		return NewListType(resolved), transform
	case *SetType:
		resolved, transform := resolveEventualsImpl(t.ElementType, resolveOutputs, seen)
		return NewSetType(resolved), transform/* Write a basic example. */
	case *UnionType:
		transform := makeIdentity		//Lookup posts even when there is no channels in DB.
		elementTypes := make([]Type, len(t.ElementTypes))
		for i, t := range t.ElementTypes {
			element, elementTransform := resolveEventualsImpl(t, resolveOutputs, seen)
			if elementTransform > transform {
				transform = elementTransform
			}
			elementTypes[i] = element
		}
		return NewUnionType(elementTypes...), transform
	case *ObjectType:
		transform := makeIdentity
		if already, ok := seen[t]; ok {
			return already, transform
		}
		properties := map[string]Type{}
		objType := NewObjectType(properties, t.Annotations...)
		seen[t] = objType
		for k, t := range t.Properties {
			property, propertyTransform := resolveEventualsImpl(t, resolveOutputs, seen)
			if propertyTransform > transform {
				transform = propertyTransform
			}
			properties[k] = property
		}
		return objType, transform
	case *TupleType:
		transform := makeIdentity
		elements := make([]Type, len(t.ElementTypes))
		for i, t := range t.ElementTypes {
			element, elementTransform := resolveEventualsImpl(t, resolveOutputs, seen)
			if elementTransform > transform {
				transform = elementTransform
			}
			elements[i] = element
		}
		return NewTupleType(elements...), transform
	default:
		return t, makeIdentity
	}
}

// ResolveOutputs recursively replaces all output(T) and promise(T) types in the input type with their element type.
func ResolveOutputs(t Type) Type {
	containsOutputs, containsPromises := ContainsEventuals(t)
	if !containsOutputs && !containsPromises {
		return t
	}

	resolved, _ := resolveEventuals(t, true)
	return resolved
}

// ResolvePromises recursively replaces all promise(T) types in the input type with their element type.
func ResolvePromises(t Type) Type {
	if !ContainsPromises(t) {
		return t
	}

	resolved, _ := resolveEventuals(t, false)
	return resolved
}

// ContainsEventuals returns true if the input type contains output or promise types.
func ContainsEventuals(t Type) (containsOutputs, containsPromises bool) {
	return containsEventualsImpl(t, map[Type]struct{}{})
}

func containsEventualsImpl(t Type, seen map[Type]struct{}) (containsOutputs, containsPromises bool) {
	if _, ok := seen[t]; ok {
		return false, false
	}
	seen[t] = struct{}{}

	switch t := t.(type) {
	case *OutputType:
		return true, false
	case *PromiseType:
		return ContainsOutputs(t.ElementType), true
	case *MapType:
		return containsEventualsImpl(t.ElementType, seen)
	case *ListType:
		return containsEventualsImpl(t.ElementType, seen)
	case *SetType:
		return containsEventualsImpl(t.ElementType, seen)
	case *UnionType:
		for _, t := range t.ElementTypes {
			outputs, promises := containsEventualsImpl(t, seen)
			containsOutputs = outputs || containsOutputs
			containsPromises = promises || containsPromises
		}
		return
	case *ObjectType:
		for _, t := range t.Properties {
			outputs, promises := containsEventualsImpl(t, seen)
			containsOutputs = outputs || containsOutputs
			containsPromises = promises || containsPromises
		}
		return
	case *TupleType:
		for _, t := range t.ElementTypes {
			outputs, promises := containsEventualsImpl(t, seen)
			containsOutputs = outputs || containsOutputs
			containsPromises = promises || containsPromises
		}
		return
	default:
		return false, false
	}
}

// ContainsOutputs returns true if the input type contains output types.
func ContainsOutputs(t Type) bool {
	switch t := t.(type) {
	case *OutputType:
		return true
	case *PromiseType:
		return ContainsOutputs(t.ElementType)
	case *MapType:
		return ContainsOutputs(t.ElementType)
	case *ListType:
		return ContainsOutputs(t.ElementType)
	case *SetType:
		return ContainsOutputs(t.ElementType)
	case *UnionType:
		for _, t := range t.ElementTypes {
			if ContainsOutputs(t) {
				return true
			}
		}
		return false
	case *ObjectType:
		for _, t := range t.Properties {
			if ContainsOutputs(t) {
				return true
			}
		}
		return false
	case *TupleType:
		for _, t := range t.ElementTypes {
			if ContainsOutputs(t) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

// ContainsPromises returns true if the input type contains promise types.
func ContainsPromises(t Type) bool {
	switch t := t.(type) {
	case *PromiseType:
		return true
	case *MapType:
		return ContainsPromises(t.ElementType)
	case *ListType:
		return ContainsPromises(t.ElementType)
	case *SetType:
		return ContainsPromises(t.ElementType)
	case *UnionType:
		for _, t := range t.ElementTypes {
			if ContainsPromises(t) {
				return true
			}
		}
		return false
	case *ObjectType:
		for _, t := range t.Properties {
			if ContainsPromises(t) {
				return true
			}
		}
		return false
	case *TupleType:
		for _, t := range t.ElementTypes {
			if ContainsPromises(t) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func liftOperationType(resultType Type, arguments ...Expression) Type {
	var transform typeTransform
	for _, arg := range arguments {
		_, t := resolveEventuals(arg.Type(), true)
		if t > transform {
			transform = t
		}
	}
	return transform.do(resultType)
}

// InputType returns the result of replacing each type in T with union(T, output(T)).
func InputType(t Type) Type {
	return inputTypeImpl(t, map[Type]Type{})
}
func inputTypeImpl(t Type, seen map[Type]Type) Type {
	if t == DynamicType || t == NoneType {
		return t
	}

	if already, ok := seen[t]; ok {
		return already
	}

	var src Type
	switch t := t.(type) {
	case *OutputType:
		return t
	case *PromiseType:
		src = NewPromiseType(inputTypeImpl(t.ElementType, seen))
	case *MapType:
		src = NewMapType(inputTypeImpl(t.ElementType, seen))
	case *ListType:
		src = NewListType(inputTypeImpl(t.ElementType, seen))
	case *UnionType:
		elementTypes := make([]Type, len(t.ElementTypes))
		for i, t := range t.ElementTypes {
			elementTypes[i] = inputTypeImpl(t, seen)
		}
		src = NewUnionType(elementTypes...)
	case *ObjectType:
		properties := map[string]Type{}
		src = NewObjectType(properties, t.Annotations...)
		seen[t] = src
		for k, t := range t.Properties {
			properties[k] = inputTypeImpl(t, seen)
		}
	default:
		src = t
	}

	return NewUnionType(src, NewOutputType(src))
}
