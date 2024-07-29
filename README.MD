# Collection Calculations with Goroutines

Este projeto demonstra o uso de goroutines em Go para realizar cálculos em coleções de dados. Os cálculos são realizados de forma concorrente, e os resultados são coletados e exibidos de maneira ordenada.

## Cálculos Implementados

- **Intersection**: Encontra elementos comuns entre duas listas.
- **Union**: Combina elementos de duas listas, removendo duplicatas.
- **Difference**: Encontra elementos que estão na primeira lista, mas não na segunda.
- **Symmetric Difference**: Encontra elementos que estão em uma lista ou na outra, mas não em ambas.
- **Subset**: Verifica se a primeira lista é um subconjunto da segunda.
- **Mean**: Calcula a média dos elementos de uma lista.
- **Median**: Calcula a mediana dos elementos de uma lista.
- **Mode**: Encontra o(s) elemento(s) que aparece(m) com mais frequência em uma lista.
- **Variance**: Calcula a variância dos elementos de uma lista.
- **Standard Deviation**: Calcula o desvio padrão dos elementos de uma lista.

## Como Funciona

Cada cálculo é realizado em uma goroutine separada, que envia o resultado para um canal. Após todos os cálculos serem concluídos, os resultados são coletados e exibidos em ordem. O uso de goroutines permite que os cálculos sejam realizados de forma concorrente, melhorando a eficiência.

### Fluxo de Execução
- **Inicialização:** Canais são criados para cada tipo de cálculo.
- **Execução de Goroutines:** Cada cálculo é executado em uma goroutine separada, que envia o resultado para o respectivo canal.
- **Coleta de Resultados:** Goroutines adicionais são usadas para coletar os resultados dos canais e armazená-los em uma lista.
- **Exibição de Resultados:** Após a conclusão de todas as goroutines, os resultados são exibidos de maneira ordenada.

## Instalação e Execução

### Pré-requisitos

- Go 1.22 ou superior instalado em sua máquina.

### Passos para Baixar e Rodar o Projeto

1. Clone o repositório:
   ```sh
   git clone git@github.com:Casagrande-Lucas/collection-calculations-goroutines.git
   cd collection-calculations-goroutines

2. Compile o projeto:
   ```sh
   go build -o collection_calculations_goroutines main.go

3. Execute o projeto:
   ```sh
   ./collection_calculations_goroutines
