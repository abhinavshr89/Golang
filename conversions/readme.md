
## Program Explaination
```go
input = strings.TrimSpace(input)
```
- `string.Trimspace` removes extra whitespaces from the input


```go
numRating, err := strconv.ParseFloat(input, 64)
```
- Here we are converting the number to float

## Explanation of strconv Package Functions

In the provided Go code, the `strconv` package is used for converting strings to numeric types and vice versa. Here's an explanation of the functions used:

1. **strconv.Atoi(s string) (int, error)**:
   - This function converts the string `s` to an integer `int`. It returns the converted integer and an error if the string couldn't be converted to an integer.

2. **strconv.ParseFloat(s string, bitSize int) (float64, error)**:
   - This function converts the string `s` to a floating-point number `float64`. The `bitSize` parameter specifies the bit size of the floating-point representation (32 for float32, 64 for float64). It returns the converted floating-point number and an error if the string couldn't be converted to a floating-point number.

3. **strconv.ParseInt(s string, base int, bitSize int) (int64, error)**:
   - This function converts the string `s` to a signed integer `int64`, interpreting the string as a number in the specified `base` (e.g., base 10 for decimal, base 16 for hexadecimal). The `bitSize` parameter specifies the bit size of the integer representation. It returns the converted integer and an error if the string couldn't be converted to an integer.

4. **strconv.ParseUint(s string, base int, bitSize int) (uint64, error)**:
   - This function converts the string `s` to an unsigned integer `uint64`, interpreting the string as a number in the specified `base`. The `bitSize` parameter specifies the bit size of the integer representation. It returns the converted integer and an error if the string couldn't be converted to an unsigned integer.

5. **strconv.FormatInt(i int64, base int) string**:
   - This function returns the string representation of the integer `i` in the specified `base`. It converts the integer to a string using the digits of the given base.

6. **strconv.FormatUint(i uint64, base int) string**:
   - Similar to `FormatInt`, but it formats unsigned integers.

7. **strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string**:
   - This function formats the floating-point number `f` into a string representation. The `fmt` parameter specifies the format ('b' for decimal, 'e' for scientific notation, 'f' for decimal no exponent, 'g' for the most compact representation). The `prec` parameter specifies the precision for 'e', 'f', and 'g' formats, and the `bitSize` parameter specifies the bit size of the floating-point number.

These functions provide robust functionality for converting numeric data between strings and various numeric types in Go.
