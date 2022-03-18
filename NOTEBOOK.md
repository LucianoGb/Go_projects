# GO

**Criada pelo Google.**

**Por que estudar Go?**<br>
Alta performace <br/>
Multiplataforma <br/>
Garbage collection <br/>

**Usada em:<br/>**

APIs, CLIs, micro serviços, processamento  de dados e Cloud.

**Fontes para estudo**

exercícios em: <https://github.com/crgimenes/Go-Hands-On>

gravações em: <https://www.youtube.com/cesargimenes>

**Playground**
<https://go.dev/play/>

**Olá mundo em Go:**

```go
  package main
  import "fmt"

  func main(){
    fmt.Println("Olá Mundo!")
  }
```

**Package main** Go trabalha com package, o main é o ponto de partida do go.

**func main()** Principal função do go, assim como java e c, go utiliza a função main para porta de entrada e saida.

**Notação da função Println()** fmt(pacote). Println(identificador da função). 

**Aprendendo ler a documentação **

Na documentação do Println temos: func Println(a ...any) (n int, err error)
os '...' -> quer dizer que é uma função variádica, que é uma função que consegue trabalhar com qualquer número de argumentos. (n int, err error) essa parte é do seu retorno, que é um inteiro que me retorna a quantidade de bytes lidos e erro caso possua alguma.

**OBS**

Sempre que uma função tiver um retorno e quisermos que esse retorno seja armazenado em uma variável, temos que adicionar uma váriavel para cada retorno. vamos usar o da função print.

```go
  b,erros := fmt.Println("Olá Mundo!")
  // retorno 12 <nil> -> nil = nulo
```

Caso a gente só queira imprimir os bytes da função temos por obrição passar duas variáveis, em go temos um jeito de declarar uma variável que não nos obriga usar o valor '_' -> é chamando de identificado branco ou BLANC IDENTIFY

```go
  b,_ := fmt.Println("Olá Mundo!")
  // retorno 12 <nil> -> nil = nulo
```

## Variavel

### Shorthand 
No go podemos declarar variável de duas formas, a mais usada é a shorthand.
Onde passamos o nome da variável, o operador gopher e o valor.

```go
  nome := "Luciano"
```

Dessa forma o operador gopher **":="** já atribui um tipo a variável de acordo com seu valor.

**OBS** o operador gopher é um operador de declaração. Se tentarmos declarar a mesma váriavel com um outro valor o go retornará um erro.

Erro retornado:
fundamentos\variavel\variavel.go:15:10: no new variables on left side of :=

Mas podemos usar o operador "=" que é um operador de atribuição para mudar o valor da variável sem problemas.

O Gopher ( Marmota como é chamado em português ) só funciona dentro de codeblocks ou scope. Ou seja dentro de uma função ou método.

