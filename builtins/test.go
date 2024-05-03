package builtins

import (
    "fmt"
    "io"
    "strconv"
)

// Test evaluates conditional expressions, including numerical comparisons.
func Test(w io.Writer, args ...string) error {
    if len(args) != 3 {
        fmt.Fprintln(w, "test: expected three arguments for comparisons")
        return nil
    }

    // Extract arguments
    leftOperand := args[0]
    operator := args[1]
    rightOperand := args[2]

    // Parse operands as integers
    left, err1 := strconv.Atoi(leftOperand)
    right, err2 := strconv.Atoi(rightOperand)

    if err1 != nil || err2 != nil {
        fmt.Fprintln(w, "test: both operands must be integers for comparison")
        return nil
    }

    // Perform comparison based on the operator
    switch operator {
    case "-eq":
        fmt.Fprintln(w, left == right)
    case "-ne":
        fmt.Fprintln(w, left != right)
    case "-gt":
        fmt.Fprintln(w, left > right)
    case "-ge":
        fmt.Fprintln(w, left >= right)
    case "-lt":
        fmt.Fprintln(w, left < right)
    case "-le":
        fmt.Fprintln(w, left <= right)
    default:
        fmt.Fprintf(w, "test: unknown operator %s\n", operator)
    }

    return nil
}
