# Exercícios de fixação

### Exercício 1.1

- Esses são seus primeiros exercícios, e seus primeiros passos.
- Completando os exercícios dessa seção, você será um ninja nível 1.
- É o seu primeiro passo pra se tornar um developer ninja.
- Esses exercícios servem pra reforçar seu aprendizado. Só se aprende a programar programando. Ninguem aprende a andar de bicicleta assistindo vídeos de pessoas andando de bicicleta. É necessário botar a mão na massa.
- Eu vou começar explicando qual é o exercício. Aí vou pedir pra você dar pausa. Esse é o momento de por os miolos pra trabalhar, encontrar sua solução, tec-tec-tec, e rodar pra ver se funciona. Depois é só dar play novamente, ver a minha abordagem para a mesma questão, e comparar nossas soluções.
- Vamos lá:
- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
  1. 42
  2. "James Bond"
  3. true
- Agora demonstre os valores nestas variáveis, com:
  1. Uma única declaração print
  2. Múltiplas declarações print

###  Exercício 1.2

- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
  1. Identificador "x" deverá ter tipo int
  2. Identificador "y" deverá ter tipo string
  3. Identificador "z" deverá ter tipo bool
- Na função main:
  1. Demonstre os valores de cada identificador
  2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?

### Exercício 1.3

- Utilizando a solução do exercício anterior:
  1. Em package-level scope, atribua os seguintes valores às variáveis:
     1. para "x" atribua 42
     2. para "y" atribua "James Bond"
     3. para "z" atribua true
  2. Na função main:
     1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
     2. Demonstre a variável "s".

### Exercício 1.4

- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
  1. Demonstre o valor da variável "x"
  2. Demonstre o tipo da variável "x"
  3. Atribua 42 à variável "x" utilizando o operador "="
  4. Demonstre o valor da variável "x"

### Exercício 1.5

- Utilizando a solução do exercício anterior:
  1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
  2. Na função main:
     1. Isto já deve estar feito:
        1. Demonstre o valor da variável "x"
        2. Demonstre o tipo da variável "x"
        3. Atribua 42 à variável "x" utilizando o operador "="
        4. Demonstre o valor da variável "x"
     2. Agora faça também:
        1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
        2. Demonstre o valor de "y"
        3. Demonstre o tipo de "y"

### Exercício 2.1

- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.

### Exercício 2.2

- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
  - "=="
  - "!="
  - "<="
  - "<"
  - ">="
  - "<"
- Demonstre estes valores.

### Exercício 2.3

- Crie constantes tipadas e não-tipadas.
- Demonstre seus valores.

### Exercício 2.4

- Crie um programa que:
  - Atribua um valor int a uma variável
  - Demonstre este valor em decimal, binário e hexadecimal
  - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
  - Demonstre esta outra variável em decimal, binário e hexadecimal

### Exercício 2.5

- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.

### Exercício 2.6

- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.

### Exercício 3.1

- Põe na tela: todos os números de 1 a 10000.

### Exercício 3.2

- Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
- Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.

### Exercício 3.3

- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.

### Exercício 3.4

- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.

### Exercício 3.5

- Demonstre o resto da divisão por 4 de todos os números entre 10 e 100

### Exercício 3.6

- Crie um programa que demonstre o funcionamento da declaração if.

### Exercício 3.7

- Utilizando a solução anterior, adicione as opções else if e else.

### Exercício 3.8

- Crie um programa que utilize a declaração switch, sem switch statement especificado.

### Exercício 3.9

- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

### Exercício 3.10

- Anote (à mão) o resultado das expressões:
  - fmt.Println(true && true)
  - fmt.Println(true && false)
  - fmt.Println(true || true)
  - fmt.Println(true || false)
  - fmt.Println(!true)