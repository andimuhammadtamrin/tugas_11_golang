package main

import "fmt"
import "strconv"

func main(){
  var input1 string
  fmt.Print("Masukkan angka 1 :", input1)
  fmt.Scanln(&input1)
  var input2 string
  fmt.Print("Masukkan angka 2 :", input2)
  fmt.Scanln(&input2)

  var num1,err = strconv.Atoi(input1)
  var num2,err1 = strconv.Atoi(input2)

  if err  != nil{
    fmt.Println(err.Error())
  }else if err1 != nil {
    fmt.Println(err.Error())
  }
  var jumlah = num1 + num2
  var kurang = num1 - num2
  var kali = num1 * num2
  var bagi = num1 / num2
  fmt.Println(input1," + ",input2," = ",jumlah)
  fmt.Println(input1," - ",input2," = ",kurang)
  fmt.Println(input1," * ",input2," = ",kali)
  fmt.Println(input1," / ",input2," = ",bagi)
}
