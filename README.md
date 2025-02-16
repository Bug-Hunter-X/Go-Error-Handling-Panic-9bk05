# Go Error Handling Panic

This example demonstrates a common error in Go's error handling.  When a function returns an error, the returned value might still be used if not properly handled, potentially causing a panic.  The solution addresses this by checking for the error before using the return value. This subtle bug can be difficult to trace without proper debugging.