# Gerador de números para jogos da megasena 

## Como rodar o programa

```
go run main.go
```

Por padrão é gerado apenas um jogo, porém pode fazer a geração de vários jogos, para isso rodar o seguinte comando: 

```
go run main.go -j 5
```

Também é possível fazer a geração de mais dezenas para isso você pode usar o seguinte comando: 

```
go run main.go -d 8
```

Observação: 
- Por padrão é gerado 6 dezenas 
- Por padrão é gerado 1 jogo

## Informações sobre o projeto

Verificação de entrada

```
	if n < 0 || n > maiorNumero {
		return nil, errors.New(fmt.Sprintf("impossivel gerar %v numeros", n))
	}

```
- Verifica se "n" é menor que o 0 ou maior que a constante "maiorNumero"
- Se for, retorna "nil" um erro é informado que é impossível gerar os números
<hr>

Criando um mapa vazio

```
	mapa := make(map[int]bool)
```
Criação de mapa vazio para garantir que nenhum número seja repetido. 
<hr>

Gerando número únicos
```
	for len(mapa) < n {
		dezena := rand.Intn(maiorNumero) + 1
		mapa[dezena] = true
	}
```
- Usa um "for" que continua até que o mapa tenha "n" números únicos
- Dentro do laço
  - Gera um número aleatório entre 1 a constante "maiorNumero"
  - Adiciona esse número ao mapa. A chave do mapa é o número, e o vlaor é "true"
<hr>

Convertendo o mapa para um slice

```
	numeros := make([]int, 0, n)
	for k := range mapa {
		numeros = append(numeros, k)
	}
```
- Criar um slice vazio chamado "numeros" com capacidade para "n" elementos
- Itera sobre as chaves do mapa (os números únicos) e as adiciona ao slice "numero"
<hr>

Ordenação do slice 

```
	sort.Ints(numeros)
```

- Ordena o slice "numeros" para ordem crescente. 
