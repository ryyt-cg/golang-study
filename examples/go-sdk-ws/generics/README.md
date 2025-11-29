# Generics

Prior to go1.18, empty interface `interface{}` was using as generics.
```go
    package main

    import "fmt"

    func printAnything(data interface{}) {
        fmt.Println(data)
    }

    type person struct {
        Name string
        Age  int
    }

    func main() {
		printAnything("hello")
		printAnything(123)
		printAnything(true)
		printAnything(person{
			Name: "Jay",
			Age: 23,
        })
    }
```
* data type as `interface{}`
* printAnything will accept parameter data as any types.

output
```bash
hello
123
true
{Jay 23}
```

## Using Empty Interface (interface{} or any)
* **Runtime Polymorphism:** The empty interface represents a value of any type. When you use interface{}, you are essentially saying that the variable can hold a value of any concrete type.
* **Type Assertions:** To work with the underlying concrete type, you need to perform runtime type assertions or type switches, which dynamically check the type and extract the value.
* **Loss of Compile-Time Type Safety:** The compiler cannot verify the type of data stored in an interface{} variable at compile time, leading to potential runtime errors if type assertions fail.
* **Performance Overhead:** Assigning values to and retrieving values from interface{} involves boxing and unboxing operations, which can introduce some performance overhead.
* **Use Cases:** Suitable for situations where the exact type is unknown until runtime, such as deserializing JSON or handling diverse data from external sources.


## Generics

```go
    func print[T any](data T) }
        fmt.Println(data)
	}
```
or

```go
    func print[T int | string | bool | person](data T) {
        fmt.Println(data)
    }
```


* **Compile-Time Polymorphism:** Generics allow you to write functions, types, and methods that operate on a set of types defined by type parameters and constraints (interfaces). The specific types are determined at compile time.
* **Type Safety:** Generics maintain Go's strong type safety by enforcing type constraints at compile time, catching type-related errors before runtime.
* **No Runtime Type Assertions:** You typically don't need explicit type assertions or switches when working with generic code, as the compiler handles the type specifics.
* **Performance:** Generics often result in better performance than using interface{} with type assertions because the compiler can generate optimized code for specific types at compile time, avoiding runtime boxing/unboxing.
* **Use Cases:** Ideal for creating reusable and type-safe data structures (like lists, maps, queues), algorithms (like sorting), and functions that operate on various types while maintaining type guarantees.