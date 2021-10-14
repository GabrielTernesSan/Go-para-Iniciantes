# Explorando Tipos

- Tipos em GO são estáticos, ou seja, se você declarar uma variável com tipo Int ela será Int até o final.

- Ao declarar uma variável para conter valores de um tipo essa variável só poderá conter valores desse tipo.

- O tipo pode ser deduzido pelo compilador:

  ```GO
  x := 10 // Int
  var y := "a tia do Batima" // String
  ```

  

- **Tipos de dados primitivos**: Disponíveis na linguagem nativamente como blocos blocos básicos.

  ```GO
  int
  string
  bool
  ```

- **Tipos de dados compostos**: São tipos compostos de tipos primitivos e criados pelo usuário.

  - O ato de definir, criar, estruturar tipos compostos chama-se composição.

  ```Go
  slice
  array
  struct 
  map
  ```