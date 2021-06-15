package kata

import (
//   "fmt"
  "strconv"
)

func CreatePhoneNumber(numbers [10]uint) string {
  var phoneNumber string 
 for i, char := range numbers {
//    fmt.Println("i is ===> %v", i)
   str := strconv.FormatUint(uint64(char), 10)
//    switch  case version
   switch i {
    case  0:
     phoneNumber = phoneNumber + "(" + str
    case 2:
     phoneNumber = phoneNumber + str + ") "
    case 5:
     phoneNumber = phoneNumber + str + "-"
    default:
     phoneNumber = phoneNumber + str 
   }
//    if version 
//    if i == 0 {
//      phoneNumber = phoneNumber + "(" + str
//    } else if i == 2 {
//      phoneNumber = phoneNumber + str + ") "
//    } else if i == 5 {
//      phoneNumber = phoneNumber + str + "-"
//    } else {
//      phoneNumber = phoneNumber + str 
//    }
 }
  return phoneNumber
}
