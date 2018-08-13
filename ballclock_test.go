package main

import (
  "testing"
  "reflect"
)

func TestGetNewClock(t *testing.T){
  expected := new (BallClock)
  expected.Main = []int{1 , 2, 3, 4}

  var actual = GetNewClock(4)

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("Item not initalized correctly \n actual: %d\n expected: %d", actual, expected)
  }
}

func TestEmptyQueue(t *testing.T){
  testFrom := []int{1, 2, 3, 4}
  testTo := []int{8}

  expectedFrom := []int{}
  expectedTo := []int{8, 4, 3, 2, 1}

  actualFrom, actualTo := EmptyQueue(testFrom, testTo)

  if !reflect.DeepEqual(actualFrom, expectedFrom) {
    t.Errorf("From not correct when emptying queue.\n actual: %d\n expected: %d", actualFrom, expectedFrom)
  }

  if !reflect.DeepEqual(actualTo, expectedTo) {
    t.Errorf("To not correct when emptying queue.\n actual: %d\n expected: %d", actualTo, expectedTo)
  }
}


func TestFindCycle1(t *testing.T){
  var result = FindCycle(30)
  if result != 15 {
    t.Errorf("TesDays1 failed, got: %d, want: 30.", result)
  }
}

func TestFindCycle2(t *testing.T){
  var result = FindCycle(45)
  if result != 378 {
    t.Errorf("TesDays1 failed, got: %d, want: 378.", result)
  }
}


func TestRunSimulation(t *testing.T){
  expected := new (BallClock)
  expected.Min = []int{}
  expected.FiveMin = []int{22,13,25,3,7}
  expected.Hour = []int{6,12,17,4,15}
  expected.Main = []int{11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9}

  actual := RunSimulation(30, 325)

  if !reflect.DeepEqual(actual, expected) {
    t.Errorf("RunSimulation Failed.\n actual: %d\n expected: %d", actual, expected)
  }
}


