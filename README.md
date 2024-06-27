
# CastroCoin Blockchain CLI

Este é um projeto de blockchain em Go chamado CastroCoin, que inclui uma CLI (interface de linha de comando) para adicionar blocos e imprimir a cadeia de blocos.

## Requisitos

- Go 1.20 ou superior (última versão estável no momento da escrita)
- Biblioteca `github.com/cleicastro/castrocoin/blockchain`

## Instalação

1. **Clone o repositório:**

    ```sh
    git clone https://github.com/seu-usuario/castrocoin.git
    cd castrocoin
    ```

2. **Instale as dependências:**

    ```sh
    go get ./...
    ```

3. **Compile o projeto:**

    ```sh
    go build -o castrocoin
    ```

## Uso

### Comandos Disponíveis

- `add -block BLOCK_DATA` : Adiciona um bloco à cadeia com os dados fornecidos.
- `print` : Imprime todos os blocos na cadeia.

### Exemplo de Uso

#### Adicionar um Bloco

Para adicionar um bloco com dados, use o comando `add` seguido dos dados do bloco:

```sh
./castrocoin add -block "Dados do Bloco"
```

#### Imprimir a Cadeia de Blocos

Para imprimir todos os blocos na cadeia, use o comando `print`:

```sh
./castrocoin print
```

## Estrutura do Projeto

- `main.go`: O ponto de entrada da aplicação, que contém a definição da CLI e os comandos.
- `blockchain`: O pacote que contém a lógica da blockchain.

## Exemplo de Saída

Ao adicionar um bloco:

```sh
./castrocoin add -block "Primeiro Bloco"
Added block!
```

Ao imprimir a cadeia de blocos:

```sh
./castrocoin print
Previous Hash: 0000000000000000
Data in Block: Genesis Block
Hash: 1b2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p
PoW: true

Previous Hash: 1b2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p
Data in Block: Primeiro Bloco
Hash: 2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p1a2b
PoW: true
```

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para mais detalhes.
