/* 
  aula focada em tipos de variaveis
*/
package main

import (
   "fmt"
   "reflect"
   "math"
)

func main() {
   //inteiros
   fmt.Println(1, 2, 1000)
   fmt.Println("Literal inteiro e:", reflect.TypeOf(320000))

   numero := 2.12121
   fmt.Println("Literal inteiro e:", reflect.TypeOf(numero))

   // somenteo positivos
   var b byte = 255
   fmt.Println("O byte e:   ", reflect.TypeOf(b))

   i1 := math.MaxInt64
   fmt.Println("O valor maximo do int e:", i1)
   fmt.Println("O tipo de i1 e: ", reflect.TypeOf(i1))

   // mapeamento unicode
   var i2 rune = 'a'
   fmt.Println("O rune e: ", reflect.TypeOf(i2))
   fmt.Println(i2)

   // numeros reais
   var x float32 = 49.99
   fmt.Println("O tipo de x e: " , reflect.TypeOf(x))
   fmt.Println("O tipo do literal 49.99 e: ",  reflect.TypeOf(49.99))

   // boolean
   bo := true
   fmt.Println("O tipo de bo e: ", reflect.TypeOf(bo))
   fmt.Println(bo)
   fmt.Println(!bo)

  // Sting - aspas duplas
  s1 := "Ola, meu nome e Bruno"
  fmt.Println(s1 + "!!!")
  fmt.Println("O tamanho da string e: ", len(s1))



}
