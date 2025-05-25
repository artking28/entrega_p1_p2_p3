# P1 - Neander - Arthur Andrade


## Estrutura

- `srcParserLPN/`:  Có que faz o parse do LNP para o ASM próprio.
- `srcAssemblr/`: Código que faz o parse do ASM próprio para o BIN de Neander. 
- `srcNeanderVM/`: Código fonte da VM do Neander puro.
- `misc`: Arquivos variados.

## Comandos `make`
- `make all`: Compila o projeto inteiro. Gera os binários `bfc` e `bfe`.
- `make test`: Testa o 
- `make clean`: Remove os binários gerados e limpa a tela.


### Testar na raiz do projeto
```
$ exp programa.lpn output.asm
$ asmp1 output.asm output.bin
$ neander output.bin
```


## Programa LPN

Esta e a sintaxe basica para funcionar.
Os operadores permitidos `+`, `-`, `&` e `|`, representando soma, subtração, and e or respetivamente.

A unica funcao permitida é a main, que é criada como demonstrado abaixo, o print coloca o valor resultado na posicao de memoria 200 do Neander

### Funcoes permitidas 
- `print`: coloca o valor na memoria 200 do neander, pode receber um ID um valor ou uma expressao
- `exit`: Força um halt no neander, n pode receber argumentos

### Exemplo

```text
func main() {

    def foo = 1
    foo += 1
    foo -= 1
    foo |= 1
    foo &= 1

    def bar = 1 - 1 + 3 & 1 | 1 
    bar = 3

    def chaos = foo + bar

    print(foo)
}
```

### BNF

```bnf
<letter> ::= [a-z] | [A-Z]
<digit> ::= [0-9]
<opers> ::= "+" | "-" | "&" | "|"
<space> ::= " "+

<ID> ::= <letter> (<letter> | <digit>)*
<number> ::= <digit>+
<value> ::= <number> | <ID>
<exp> ::= <value> (<opers> <exp>)* | "(" <exp> ")"
<arg_assign> ::= "=" | ( <opers> "=" )

<var_decl> ::= "def" <space> <ID> <space> "=" <space> <exp>
<assign> ::= <ID> <space> <arg_assign> <space> <exp>
<func> ::= "print(" <exp> ")" | "exit()"

<stmt> ::= <space> (<var_decl> | <func> | <assign> ) <space>
<stmts> ::= (<stmt> ("\n" | ";") )+

<program> ::= "fun main()" <space> "{" <stmts> "}"

```
