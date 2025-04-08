## Functions, Conditionals, and Loops

### Feedback

1. math.Mod is more suitable for taking modulus instead of converting floats to strings, taking modulus, and converting it back to float as that truncates the decimal part and can have side-effects and produce unexpected results.

2. strings.Builder is more suitable for building strings in Go as using concatenation with + is creating a new copy of the string in memory and consumes more resources for very large strings,so strings.Builder is more efficient and performance based.
