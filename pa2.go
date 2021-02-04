/*
"I William Askew (wi357066) afﬁrm that this program is entirely my own work and
 that I have neither developed my code together with any another person, nor copied
 any code from any other person, nor permitted my code to be copied or otherwise
 used by any other person, nor have I copied, modiﬁed,or otherwise used programs
 created by others. I acknowledge that any violation of the above terms will be
 treated as academic dishonesty.”
*/

package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
  "log"
  "sort"
)

func main() {

  if len(os.Args) < 2 {
    fmt.Println("Missing parameter. Please provide input file name or save file name.\n")
    return
  }
  file, err := os.Open(os.Args[1])

  if err != nil {
    fmt.Println("Can't read file: ", os.Args[1])
    return
  }
  defer file.Close();

  readInputFile(file);
}

func readInputFile(input *os.File) {

  var cylreq []int
  alg := ""
  upperCYL := 0
  lowerCYL := 0
  initCYL := 0
  abort := false

  scanner := bufio.NewScanner(input)
  var s [][]string
  for scanner.Scan() {
    var x = strings.Fields(scanner.Text())
    s = append(s, x)
   }
   if err := scanner.Err(); err != nil {
     log.Fatal(err)
   }

   for i := 0; i < len(s); i++ {
     if s[i][0] == "end" {
       break
     } else if s[i][0] == "use" {
       alg = s[i][1]
     } else if s[i][0] == "upperCYL" {
       u, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       upperCYL = u
     } else if s[i][0] == "lowerCYL" {
       l, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       lowerCYL = l
     } else if s[i][0] == "initCYL" {
       i, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       initCYL = i
     } else if s[i][0] == "cylreq" {
       c, err := strconv.Atoi(s[i][1])

       if err != nil {
         fmt.Println("error on string to int conversion")
       }
       cylreq = append(cylreq, c)
     }
   }

  if (upperCYL < lowerCYL) || (initCYL > upperCYL) || (initCYL < lowerCYL) {
    if (upperCYL < lowerCYL) {
      fmt.Printf("ABORT(13):upper (%d) < lower (%d)\n", upperCYL, lowerCYL)
      abort = true
    } else if (initCYL > upperCYL) {
      fmt.Printf("ABORT(11):initial (%d) > upper (%d)\n", initCYL, upperCYL)
      abort = true
    } else {
      fmt.Printf("ABORT(12):initial (%d) < lower (%d)\n", initCYL, lowerCYL)
      abort = true
    }
  }

  if !abort {
    for i := 0; i < len(cylreq); i++ {
      if cylreq[i] > upperCYL || cylreq[i] < lowerCYL {
        fmt.Printf("ERROR(15):Request out of bounds: req (%d) > upper (%d) or  < lower (%d)\n", cylreq[i], upperCYL, lowerCYL)
        cylreq = remove(cylreq, i)
      }
    }

    helper(cylreq, alg, upperCYL, lowerCYL, initCYL)
  }

  return
}

func helper(cylreq []int, alg string, upperCYL int, lowerCYL int, initCYL int) {

  if alg == "fcfs" {
    fcfs(cylreq, upperCYL, lowerCYL, initCYL)
  } else if alg == "sstf" {
    sstf(cylreq, upperCYL, lowerCYL, initCYL)
  } else if alg == "scan" {
    scan(cylreq, upperCYL, lowerCYL, initCYL)
  } else if alg == "c-scan" {
    c_scan(cylreq, upperCYL, lowerCYL, initCYL)
  } else if alg == "look" {
    look(cylreq, upperCYL, lowerCYL, initCYL)
  } else if alg == "c-look" {
    c_look(cylreq, upperCYL, lowerCYL, initCYL)
  }

  return
}

func fcfs(cylreq []int, upperCYL int, lowerCYL int, initCYL int)  {
  fmt.Printf("Seek algorithm: FCFS\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  traversed := 0
  currentCYL := initCYL

  for len(queue) != 0 {
      fmt.Printf("Servicing %5d\n", queue[0])
      temp := (queue[0] - currentCYL)

      if temp < 0 {
        temp = temp * -1
      }

      traversed += temp
      currentCYL = queue[0]
      queue = remove(queue, 0)
  }

  fmt.Printf("FCFS traversal count = %5d\n", traversed)
}

func sstf(cylreq []int, upperCYL int, lowerCYL int, initCYL int) {
  fmt.Printf("Seek algorithm: SSTF\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  traversed := 0
  n := 0
  currentCYL := initCYL

  for len(queue) != 0 {
    tempTraversed := upperCYL + 1
    for i := 0; i < len(queue); i++ {
        temp := queue[i] - currentCYL

        if temp < 0 {
          temp = temp * -1
        }

        if temp < tempTraversed {
          tempTraversed = temp
          n = i
        }
    }
      fmt.Printf("Servicing %5d\n", queue[n])
      traversed += tempTraversed
      currentCYL = queue[n]
      queue = remove(queue, n)
  }

  fmt.Printf("SSTF traversal count = %5d\n", traversed)
}

func scan(cylreq []int, upperCYL int, lowerCYL int, initCYL int) {
  fmt.Printf("Seek algorithm: SCAN\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  sort.Ints(queue)
  traversed := 0
  i := 0
  currentCYL := initCYL
  largestCYLRemoved := false
  largestCYL := queue[len(queue)-1]
  smallestNearInit := -1

  for n := 0; n < len(queue); n++{
    if queue[n] < initCYL {
      smallestNearInit = n
    }
  }


  for len(queue) != 0 {
    if queue[i] > initCYL {
      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)

      if currentCYL == largestCYL{
        largestCYLRemoved = true
        i = 0
      }
    } else if largestCYLRemoved && currentCYL == largestCYL {
      temp := upperCYL - currentCYL

      traversed += temp

      currentCYL = upperCYL
    } else if largestCYLRemoved {

      //check for cylreqs lower and service using smallestNearInit and go down until there is none
      if smallestNearInit >= 0 {
        fmt.Printf("Servicing %5d\n", queue[smallestNearInit])

        temp := queue[smallestNearInit] - currentCYL

        if temp < 0 {
          temp = temp * -1
        }

        traversed += temp

        currentCYL = queue[smallestNearInit]

        queue = remove(queue, smallestNearInit)
        smallestNearInit--

      }
    } else {
      i++
    }
  }

  fmt.Printf("SCAN traversal count = %5d\n", traversed)
}

func c_scan(cylreq []int, upperCYL int, lowerCYL int, initCYL int) {
  fmt.Printf("Seek algorithm: C-SCAN\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  sort.Ints(queue)
  traversed := 0
  i := 0
  currentCYL := initCYL
  largestCYLRemoved := false
  largestCYL := queue[len(queue)-1]

  for len(queue) != 0 {
    if queue[i] > initCYL {
      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)

      if currentCYL == largestCYL{
        largestCYLRemoved = true
        i = 0
      }
    } else if largestCYLRemoved && currentCYL == largestCYL {
      temp := upperCYL - currentCYL

      traversed += temp

      currentCYL = upperCYL
    } else if largestCYLRemoved {
      if currentCYL == upperCYL {
        temp := lowerCYL - upperCYL

        if temp < 0 {
          temp = temp * -1
        }

        traversed += temp

        currentCYL = lowerCYL
      }

      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      if temp < 0 {
        temp = temp * -1
      }

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)
    } else {
      i++
    }
  }
  fmt.Printf("C-SCAN traversal count = %5d\n", traversed)
}

func look(cylreq []int, upperCYL int, lowerCYL int, initCYL int) {
  fmt.Printf("Seek algorithm: LOOK\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  sort.Ints(queue)
  traversed := 0
  i := 0
  currentCYL := initCYL
  largestCYLRemoved := false
  largestCYL := queue[len(queue)-1]
  smallestNearInit := -1

  for n := 0; n < len(queue); n++{
    if queue[n] < initCYL {
      smallestNearInit = n
    }
  }


  for len(queue) != 0 {
    if queue[i] > initCYL {
      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)

      if currentCYL == largestCYL{
        largestCYLRemoved = true
        i = 0
      }
    } else if largestCYLRemoved {
      //check for cylreqs lower and service using smallestNearInit and go down until there is none
      if smallestNearInit >= 0 {
        fmt.Printf("Servicing %5d\n", queue[smallestNearInit])

        temp := queue[smallestNearInit] - currentCYL

        if temp < 0 {
          temp = temp * -1
        }

        traversed += temp

        currentCYL = queue[smallestNearInit]

        queue = remove(queue, smallestNearInit)
        smallestNearInit--
      }
    } else {
      i++
    }
  }

  fmt.Printf("LOOK traversal count = %5d\n", traversed)
}

func c_look(cylreq []int, upperCYL int, lowerCYL int, initCYL int) {
  fmt.Printf("Seek algorithm: C-LOOK\n")
  fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
  fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
  fmt.Printf("\tInit cylinder: %6d\n", initCYL)
  fmt.Printf("\tCylinder requests:\n")
  for i := 0; i < len(cylreq); i++ {
    fmt.Printf("\t\tCylinder %5d\n", cylreq[i])
  }

  queue := cylreq
  sort.Ints(queue)
  traversed := 0
  i := 0
  currentCYL := initCYL
  largestCYLRemoved := false
  largestCYL := queue[len(queue)-1]

  for len(queue) != 0 {
    if queue[i] > initCYL {
      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)

      if currentCYL == largestCYL{
        largestCYLRemoved = true
        i = 0
      }
    } else if largestCYLRemoved {

      fmt.Printf("Servicing %5d\n", queue[i])

      temp := queue[i] - currentCYL

      if temp < 0 {
        temp = temp * -1
      }

      traversed += temp

      currentCYL = queue[i]

      queue = remove(queue, i)
    } else {
      i++
    }
  }
  fmt.Printf("C-LOOK traversal count = %5d\n", traversed)
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}
