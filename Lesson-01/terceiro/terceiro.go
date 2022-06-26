/* 
  aula focada em print
*/
package main

import (
   "fmt"
)

func main() {
   // Exibe na tela
   fmt.Print("Conteudo")
   fmt.Print(" na mesma linha")
   // Exibe quebrando linha
   fmt.Println(" Quebra linha ao fim do texto")
   fmt.Println("Nova linha")

   // Exibindo variavel na tela
   x := 3.141516

   xs := fmt.Sprint(x)
   fmt.Println("O valor de X e: " + xs)
   fmt.Println("O valor de x e: ", x)
   fmt.Printf("O valor de X e: %.2f", x)

   // Exibindo varias variaveis

   a := 1
   b := 1.9999
   c := false
   d := "texto"

   fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d)

   // o mesmo, porem mais generio
   fmt.Printf("\n %v %v %v %v", a, b, c, d)
}
