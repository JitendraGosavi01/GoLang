In Go, Maps is an unordered collection of key and its associated value. They are very good for looking up values fast. The key type must have the operations == and != defined, like string, int, float.

Hence arrays, slices and structs cannot be used as key type, but pointers and interface types can.

structs can be used as a key when we provide Key() or Hash() method, so that a unique numeric or string key can be calculated from the struct's fields.
