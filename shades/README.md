Working examples of common [Go gotchas](http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/)

In most cases, we include a program that will fail to compile and
a working example, and note where Go tooling or a robust development setup would help prevent or mitigate these issues.

###[BEGINNER](https://github.com/buddhamagnet/shades/tree/master/beginner)

* [brace woes](https://github.com/buddhamagnet/shades/tree/master/beginner/braces)
* [unused variables](https://github.com/buddhamagnet/shades/tree/master/beginner/unused_vars)
* [unused imports](https://github.com/buddhamagnet/shades/tree/master/beginner/unused_imports)
* [short variable declarations](https://github.com/buddhamagnet/shades/tree/master/beginner/short_dec)
* [short variable redeclaration](https://github.com/buddhamagnet/shades/tree/master/beginner/short_dec_redeclare)
* [short variable declaration and fields](https://github.com/buddhamagnet/shades/tree/master/beginner/short_dec_fields)
* [accidental shadowing](https://github.com/buddhamagnet/shades/tree/master/beginner/shadow)
* [nil without explicit type](https://github.com/buddhamagnet/shades/tree/master/beginner/nil_type)
* [nil slices and maps](https://github.com/buddhamagnet/shades/tree/master/beginner/nil_collections)
* [map capacity](https://github.com/buddhamagnet/shades/tree/master/beginner/map_cap)
* [nil strings](https://github.com/buddhamagnet/shades/tree/master/beginner/nil_strings)
* [array function arguments](https://github.com/buddhamagnet/shades/tree/master/beginner/array_values)
* [unexpected values in range clauses](https://github.com/buddhamagnet/shades/tree/master/beginner/range_values)
* slices and arrays are one-dimensional
* non-existent map keys
* immutable strings
* string to byte slice conversion
* strings and index operator
* strings are not always UTF-8
* string length
* missing command in multi-line literals
* log.Fatal and log.Panic
* built-in data structure synchronization
* iteration values for strings in range clauses
* iterating maps
* switch fallthrough
* increments and decrements
* bitwise NOT
* operator preference
* unexported fields are not encoded
* app exits with active goroutines
* sending to an unbuffered channel
* sending to a closed channel
* nil channels
* methods with value receivers

###[INTERMEDIATE](https://github.com/buddhamagnet/shades/tree/master/intermediate)

* closing HTTP response body
* closing HTTP connections
* unmarshalling JSON numbers into interface values
* comparing structs, arrays, slices and maps
* panic recovery
* updating and referencing item values in range clauses
* hidden data in slices
* slice data corruption
* stale slices
* type declarations and methods
* breaking out of for switch and select
* iteration values and closures in for statements
* deferred function call argument evaluation
* deferred function call execution
* failed type assertions
* blocked goroutines and resource leaks

###[ADVANCED](https://github.com/buddhamagnet/shades/tree/master/advanced)

* using pointer receiver methods on values
* updating map value fields
* nil interfaces and nil interfaces values
* stack and heap variables
* GOMAXPROCS, concurrency, and parallelism
* read and write operation reordering
* preemptive scheduling