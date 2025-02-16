func myFunc(x int) (int, error) {
  if x == 0 {
    return 0, fmt.Errorf("cannot divide by zero")
  }
  return 10 / x, nil
}

func main() {
  result, err := myFunc(0)
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }

  result, err = myFunc(2) 
  if err != nil {
    fmt.Println("Error:", err)
  } else {
    fmt.Println("Result:", result)
  }
} 