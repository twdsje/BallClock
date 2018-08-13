package main

import (
  "fmt"
  "reflect"
)

func main(){
  PrintFindCycle(30)
  PrintFindCycle(45)
  PrintRunSimulation(30, 325)
  PrintRunSimulation(30, 21599)
}

type BallClock struct{
  Min []int
  FiveMin []int
  Hour []int
  Main []int
}

func GetNewClock(balls int) *BallClock {
  returnValue := new (BallClock)

  for i := 1; i <= balls; i++ {
    returnValue.Main = append(returnValue.Main, i)
  }

  return returnValue
}

func Iterate(myClock *BallClock) {
  //Get the first ball from the main queue.
  var next int = myClock.Main[0]
  myClock.Main = myClock.Main[1:]

  //Minutes
  if len(myClock.Min) < 4 {
    myClock.Min = append(myClock.Min, next)
  } else {
    myClock.Min, myClock.Main = EmptyQueue(myClock.Min, myClock.Main)

    //Five Minutes
    if len(myClock.FiveMin) < 11 {
      myClock.FiveMin = append(myClock.FiveMin, next)
    } else {
      myClock.FiveMin, myClock.Main = EmptyQueue(myClock.FiveMin, myClock.Main)

      //Hours
      if len(myClock.Hour) < 11 {
        myClock.Hour = append(myClock.Hour, next)
      } else {
        myClock.Hour, myClock.Main = EmptyQueue(myClock.Hour, myClock.Main)

        //Add the final ball.
        myClock.Main = append(myClock.Main, next)
      }
    }
  }
}

func EmptyQueue(from []int, to []int) ([]int, []int) {
  for len(from) > 0 {
    n := len(from) - 1
    to = append(to, from[n])
    from = from[:n]
  }
  return from, to
}

func RunSimulation(balls int, minutes int) *BallClock{
  myClock := GetNewClock(balls)
  for i := 0; i < minutes; i++ {
    Iterate(myClock)
  }

  return myClock
}

func FindCycle(balls int) int {
  myClock := GetNewClock(balls)
  initial := GetNewClock(balls)
  Iterate(myClock)

  cycles := 1
  for !reflect.DeepEqual(initial.Main, myClock.Main) {
    Iterate(myClock)
    cycles++
  }
  return cycles / 1440
}

func PrintFindCycle(balls int){
  result := FindCycle(balls)
  fmt.Printf("%d balls cycle after %d days.\n", balls, result)
}

func PrintRunSimulation(balls int, minutes int){
  result := RunSimulation(balls, minutes)
  fmt.Println(result)
}
