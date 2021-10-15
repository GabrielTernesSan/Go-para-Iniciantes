# Println

- O Println() é uma função do Package format ("fmt"), ela exibe na tela o valor passado dentro dos "()".

  - Exemplo:

    ```GO
    func main() {
    	const name, age = "Gabriel", 21
    	fmt.Println(name, "tem", age, "anos de idade.")
    }
    Output:
    	Gabriel tem 21 anos de idade.
    ```

  - Além de exibir na tela, ela retorna o número de bytes gravados e qualquer erro de gravação encontrado.