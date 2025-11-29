# Reflection
Go reflection is a powerful feature that allows you to inspect objects at runtime. It located inside the `reflect` package. Reflection can be used to dynamically access fields, kinds/types, and values, which is useful for tasks like serialization, deserialization, and generic programming.

**Kinds vs Types:**
* Kind: It represents a broader category of types, like "car". In Go, kinds are represented by the `reflect.Kind` type. Examples of kinds include `reflect.Int`, `reflect.String`, `reflect.Struct`, and `reflect.Slice`. Each kind represents a specific category of data types.
* Type: It represents a specific implementation of a kind, like "Toyota" or "Honda".  In Go, types are represented by the `reflect.Type` type such as int, string, struct, slice, etc. Each type has a specific implementation and behavior.


