package neander

import (
	"fmt"
)

type Result struct {
	Ac, Pc int8
}

const (
	NOP = 0   // Nenhuma operação
	STA = 16  // Armazena Acumulador no endereço “end” da memória
	LDA = 32  // Carrega o Acumulador com o conteúdo do endereço “end” da memória
	ADD = 48  // Soma o conteúdo do endereço “end” da memória ao Acumulador
	OR  = 64  // Efetua operação lógica “OU” do conteúdo do endereço “end” da memória ao Acumulador
	AND = 80  // Efetua operação lógica “E” do conteúdo do endereço “end” da memória ao Acumulador
	NOT = 96  // Inverte todos os bits do Acumulador
	JMP = 128 // Desvio incondicional para o endereço “end” da memória
	JN  = 144 // Desvio condicional, se “Ac!=0”, para o endereço “end” da memória
	JZ  = 160 // Desvio condicional, se "Ac==0", para o endereço “end” da memória
	HLT = 240 // Encerra o ciclo de busca-decodificação-execução
)

func RunProgram(program []byte) Result {
	padding := 4
	var result Result
	for i := padding; i < len(program); i += padding {
		mnemonic := program[i]
		if len(program) <= i+2 {
			break
		}
		addr := program[i+2]
		addrValueIndex := int(addr)*2 + padding
		switch mnemonic {
		case NOP:
			result.Pc += 1
			continue
		case STA:
			result.Pc += 2
			program[addrValueIndex] = byte(result.Ac)
			break
		case LDA:
			result.Pc += 2
			result.Ac = int8(program[addrValueIndex])
			break
		case ADD:
			result.Pc += 2
			result.Ac += int8(program[addrValueIndex])
			break
		case OR:
			result.Pc += 2
			result.Ac |= int8(program[addrValueIndex])
			break
		case AND:
			result.Pc += 2
			result.Ac &= int8(program[addrValueIndex])
			break
		case NOT:
			i -= 2
			result.Pc += 1
			result.Ac = ^result.Ac
			break
		case JMP:
			result.Pc = int8(addr)
			i = addrValueIndex - padding
			break
		case JN:
			result.Pc = int8(addr)
			if result.Ac < 0 {
				i = addrValueIndex - padding
			}
			continue
		case JZ:
			result.Pc = int8(addr)
			if result.Ac == 0 {
				i = addrValueIndex - padding
			}
			continue
		case HLT:
			result.Pc += 1
			i = len(program)
			break
		default:
			i -= 2
		}
	}

	return result
}

func PrintProgram(program []byte, hexa, continueAfterHlt bool) {
	padding := 4
	fmt.Print("\nProgram:\n")
	for i := padding; i < len(program); i += padding {
		mnemonic := program[i]
		if len(program) <= i+2 {
			break
		}
		addr := int(program[i+2])
		addrV := program[addr*2+padding]
		if hexa {
			fmt.Printf("\t[0x%.2x]", i/2-2)
		} else {
			fmt.Printf("\t[%.3d]", i/2-2)
		}
		switch mnemonic {
		case NOP:
			fmt.Printf(" NOP\n")
			i -= 2
			break
		case STA:
			str := fmt.Sprintf(" STA %.3d\n", addr)
			if hexa {
				str = fmt.Sprintf(" STA 0x%.2x\n", addr)
			}
			fmt.Print(str)
			break
		case LDA:
			str := fmt.Sprintf(" LDA %.3d(value = %d)\n", addr, int8(addrV))
			if hexa {
				str = fmt.Sprintf(" LDA 0x%.2x(value = 0x%.2x)\n", addr, addrV)
			}
			fmt.Print(str)
			break
		case ADD:
			str := fmt.Sprintf(" ADD %.3d(value = %d)\n", addr, int8(addrV))
			if hexa {
				str = fmt.Sprintf(" ADD 0x%.2x(value = 0x%.2x)\n", addr, addrV)
			}
			fmt.Print(str)
			break
		case OR:
			str := fmt.Sprintf(" OR %.3d(value = %d)\n", addr, int8(addrV))
			if hexa {
				str = fmt.Sprintf(" OR 0x%.2x(value = 0x%.2x)\n", addr, addrV)
			}
			fmt.Print(str)
			break
		case AND:
			str := fmt.Sprintf(" AND %.3d(value = %d)\n", addr, int8(addrV))
			if hexa {
				str = fmt.Sprintf(" AND 0x%.2x(value = 0x%.2x)\n", addr, addrV)
			}
			fmt.Print(str)
			break
		case NOT:
			i -= 2
			fmt.Printf(" NOT\n")
			break
		case JMP:
			str := fmt.Sprintf(" JMP %.3d\n", addr)
			if hexa {
				str = fmt.Sprintf(" JMP 0x%.2x\n", addr)
			}
			fmt.Print(str)
			break
		case JN:
			str := fmt.Sprintf(" JN  %.3d\n", addr)
			if hexa {
				str = fmt.Sprintf(" JN  0x%.2x\n", addr)
			}
			fmt.Print(str)
			break
		case JZ:
			str := fmt.Sprintf(" JZ  %.3d\n", addr)
			if hexa {
				str = fmt.Sprintf(" JZ  0x%.2x\n", addr)
			}
			fmt.Print(str)
			break
		case HLT:
			fmt.Printf(" HLT\n")
			if !continueAfterHlt {
				i = len(program)
			}
			break
		default:
			str := fmt.Sprintf(" %.3d\n", mnemonic)
			if hexa {
				str = fmt.Sprintf(" 0x%.2x\n", mnemonic)
			}
			i -= 2
			fmt.Print(str)
		}
	}

	fmt.Println()
}
