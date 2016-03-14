package main

type broken struct {
  code int
}

func main() {
  var brok broken
  // broken - can't assign to field with short variable declaration.
  // computer says non-name brok.code on left side of :=.
  brok.code := 100
}
