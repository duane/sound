package main

import (
  "fmt"
  "math"
  "os"

  "ao"
)

func main() {
  fmt.Println("libao example program.")

  ao.Initialize()
  default_driver := ao.DefaultDriverId()
  fmt.Println("default_driver: ", default_driver)
  format := ao.MakeFormat(16, 2, 44100, ao.FMT_LITTLE, "L,R");

  dev := ao.OpenLive(default_driver, &format, nil)
  if dev == nil {
    fmt.Fprintln(os.Stderr, "Error opening device: ", ao.Error())
    os.Exit(1)
  }

  var buf [2 * 2 * 44100]byte
  freq := 440.0
  for i := 0; i < 44100; i++ {
    sample := int(0.75 * 32768.0 * math.Sin(2 * math.Pi * freq * float64(i)/44100))

    buf[i * 4] = byte(sample & 0xff)
    buf[i * 4 + 2] = buf[i * 4]
    buf[i * 4 + 1] = byte((sample >> 8) & 0xff)
    buf[i * 4 + 3] = buf[i * 4 + 1]
  }
  ao.Play(dev, buf[:])
  ao.Close(dev)
  ao.Shutdown()
}
