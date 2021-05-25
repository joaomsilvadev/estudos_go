package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  
  var isHeistOn bool
  
  rand.Seed(time.Now().UnixNano())
  isHeistOn = true
  
  if eludeGuards := rand.Intn(100); eludeGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  if openedVault := rand.Intn(100); isHeistOn && openedVault >= 70 {
    fmt.Println("Grab an go")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Println("The vault can't be opened")
  }

  if leftSafely := rand.Intn(5) ; isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("The cops arrived too early. You got busted!")
      case 1:
        isHeistOn = false
        fmt.Println("You got locked inside the vault. You got busted!")
      case 2:
        isHeistOn = false
        fmt.Println("The getway car was robbed. You got busted!")
      default:
        successCounter++
        fmt.Println("Start the getaway car!")
    }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Print("You are rich! You just robbed $", amtStolen, " dollars.\n")
  }

  fmt.Println("Heist status:", isHeistOn)
}
