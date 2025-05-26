**# P3 - Brainfuck - Arthur Andrade**

Este projeto é parte da P3 da disciplina de Compiladores. O objetivo é implementar um compilador (`bfc`) e um interpretador (`bfe`) para a linguagem Brainfuck.

## Estrutura

- `cmd-bfc/`: contém o código-fonte do compilador de Brainfuck.
- `cmd-bfe/`: contém o código-fonte do interpretador de Brainfuck.
- `bfc`: binário do compilador.
- `bfe`: binário do interpretador.

## Comandos `make`

- `make all`: compila o projeto inteiro. Gera os binários `bfc` e `bfe`.
- `make test`: compila e executa um programa simples de exemplo: `CREDITO=2*5+10`.
- `make clean`: remove os binários gerados e limpa a tela.

## Bug Reports

- O operador de **divisão (`/`) não é suportado**.
- A **precedência entre multiplicação e adição/subtração pode falhar** em alguns casos.

Caso encontre outros problemas, reporte diretamente ao autor.

---
