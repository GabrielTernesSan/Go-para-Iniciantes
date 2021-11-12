# JSON Marshal (ordenação)

- A função que cria um json é json.Marshal. Podemos traduzir Marshal como ordenar. Quando você cria um Json você transforma um objeto em um conjunto de bytes enfileirados. Muitas vezes representamos esse conjunto de bytes ordenados por uma string.
- A função json.Marshal retorna uma slide de bytes e erro se tiver algum, então, para vermos o que foi gerado é necessário converter para string antes de executar o Print.

- ```go
  package main
  
  import (
  	"encoding/json"
  	"fmt"
  )
  
  type Pessoa struct {
  	Nome          string
  	Sobrenome     string
  	Idade         int
  	Profissao     string
  	Contabancaria float64
  }
  
  func main() {
  	jamesbond := pessoa{
  		Nome:          "James",
  		Sobrenome:     "Bond",
  		Idade:         40,
  		Profissao:     "Agente Secreto",
  		Contabancaria: 1000000.50,
  	}
  
  	darthvader := pessoa{"Anakin", "Skywalker", 50, "Vilão Intergaláctico", 50000000000000.83}
  
  	j, err := json.Marshal(jamesbond)
  	if err != nil {
  		fmt.Println(err)
  	}
  	d, err := json.Marshal(darthvader)
  	if err != nil {
  		fmt.Println(err)
  	}
  
  	fmt.Println(string(j))
  	fmt.Println(string(d))
  }
  OUTPUT: 
  {"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}
  
  {"Nome":"Anakin","Sobrenome":"Skywalker","Idade":50,"Profissao":"Vilão Intergaláctico","Contabancaria":50000000000000.83}
  
  ```

- **Structs exportados para JSON precisam começar com letra maiúscula**.