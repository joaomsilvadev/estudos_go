package main

import "fmt"

func fuelGauge(fuel int){
  fmt.Println("You have", fuel, "spatial galons of fuel left!")
}

func calculateFuel (planet string) int{
  var fuel int
  
  switch planet {
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  return fuel
}

func greetPlanet (planet string){
  fmt.Print("Dear passengers, we just arrived on ", planet, "!\n")
  fmt.Println("Zygma Zin baud na - Welcome")
}

func cantFly(){
  fmt.Println("We do not have the enough fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  var fuelRemaining, fuelCost int
  fuelRemaining = fuel
  fuelCost = calculateFuel(planet)
  if fuelRemaining >= fuelCost {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }
  return fuelRemaining
}

func main() {
  var fuel int 
  var planetChoice string
  fuel = 1000000
  planetChoice = "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  
}
