package main

import(
  "math/rand"
  "strings"
  "time"
)

func getRandomColor(i int) string{

  var letters = []string{"0", "1", "2", "3", "4", "5", "6", "7",
                         "8", "9", "A", "B", "C", "D", "E", "F"}
  color := "#"

  rand.Seed(int64(i) + time.Now().UnixNano())

  for x := 0; x < 6; x++ {
    color += letters[rand.Intn(16)];
  }
  return color
}


func lower(s string) string{
  return strings.ToLower(s)
}
