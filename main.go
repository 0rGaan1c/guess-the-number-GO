package main

import (
    "fmt"
    "math/rand"
    "time"
)



func main() {
  isPlaying := true

  for isPlaying {
    fmt.Println("\nGuess the number between(1 and 100)?")
    fmt.Println("Choose Difficulty:")
    fmt.Println("1. Easy\n2. Medium\n3. Hard")
    fmt.Println("Input 1, 2 or 3 to chosse difficulty level.")
    fmt.Print("> ")

    var difficulty int
    fmt.Scan(&difficulty)
    
    var chances int

    if difficulty == 1 {
      chances = 6
      fmt.Println("Easy Mode: You have", chances, "chances to guess the number.")
      isPlaying = playGame(chances)
    } else if difficulty == 2 {
      chances = 4
      fmt.Println("Medium Mode: You have", chances, "chances to guess the number.")
      isPlaying = playGame(chances)
    } else if difficulty == 3 {
      chances = 3
      fmt.Println("Hard Mode: You have", chances, "chances to guess the number.")
      isPlaying = playGame(chances)
    } else {
      fmt.Println("Invalid Input")
    }

  }
}

func playGame(chances int) bool {
  rand.Seed(time.Now().UnixNano())
  randomNumber := rand.Intn(100 - 1) + 1

  for chances > 0 {
    var guess int
    fmt.Print("Guess: ")
    fmt.Scan(&guess)

    if guess == randomNumber {
      fmt.Println("You Win!")
      chances = -1
    } else if guess > randomNumber {
      fmt.Println("Your guess is greater than the secret number. Chances remain:", chances - 1, "\n")
      chances -= 1
    } else if guess < randomNumber {
      fmt.Println("Your guess is smaller than the secret number. Chances remain:",chances - 1, "\n")
      chances -= 1
    }
    if chances == 0 {
      fmt.Println("You Lose! The Secret Number was", randomNumber, ".\n")
    }
  }

  fmt.Print("Want to try again(y/n)? ")
  var again string
  fmt.Scan(&again)

  if "y" == again {
    return true
  } else if "n" == again {
    return false
  }
  return true
}
