package main
import ("fmt")
func calcTemperatura (escolhatemp int, temperatura float64) (float64, float64, float64) {
 	var Celsius, Fahrenheit, Kelvin float64
	switch escolhatemp {
	case 1:
		Celsius = temperatura
		Fahrenheit = (temperatura*1.8) + 32
		Kelvin = temperatura + 273.15
		return Celsius, Fahrenheit, Kelvin
	case 2:
		Fahrenheit = temperatura
		Celsius = (temperatura - 32)/1.8
		Kelvin = Celsius + 273.15
		return Celsius, Fahrenheit, Kelvin
	case 3:
		Kelvin = temperatura
		Celsius = temperatura - 273.15
		Fahrenheit = (Celsius * 1.8) + 32
		return Celsius, Fahrenheit, Kelvin
	default:
		Celsius = 0
		Kelvin = 0
		Fahrenheit = 0
		return Celsius, Fahrenheit, Kelvin
	}
}
func calcComprimento (escolhacomp int, comprimento float64) (float64, float64, float64) {
	var m2, km2, cm2 float64
	switch escolhacomp {
		case 1:
			m2 = comprimento
			km2 = comprimento/1000000
			cm2 = comprimento * 10000
			return m2, km2, cm2
		case 2:
			km2 = comprimento
			m2 = comprimento * 1000000
			cm2 = comprimento * 10000000000
			return m2, km2, cm2
		case 3:
			cm2 = comprimento
			m2 = comprimento/10000
			km2 = comprimento/10000000000
			return m2, km2, cm2
		default:
			m2 = 0
			cm2 = 0
			km2 = 0
			return m2, km2, cm2
	}
}
func main() {
	var escolha, escolhatemp, escolhacomp, escolhavel, escolhatempo, escolhavol int
	var temperatura, comprimento float64
	fmt.Println(`
		Escolha o tipo de conversão de unidade:
		1. Temperatura (Celsius, Fahrenheit, Kelvin)
		2. Comprimentos (m², km², cm²)
		3. Velocidade (m/s, km/h)
		4. Tempo (segundos, minutos, horas, dias)
		5. Volume (m³, L, mL)
	`)
	fmt.Scan(&escolha)
	switch escolha {
	case 1:
		fmt.Println(`
				Qual sua unidade de medida de Temperatura?
				1. Celsius
				2. Fahrenheit
				3. Kelvin
				`)
		fmt.Scan(&escolhatemp)
		fmt.Println("Qual a temperatura?")
		fmt.Scan(&temperatura)
		var Celsius, Fahrenheit, Kelvin = calcTemperatura(escolhatemp, temperatura)
		fmt.Printf(
		"\n================= CONVERSÃO DE TEMPERATURA =================\n"+
		"Celsius     : %.2f °C\n"+
		"Fahrenheit  : %.2f °F\n"+
		"Kelvin      : %.2f K\n"+
		"============================================================\n",
		Celsius, Fahrenheit, Kelvin)
	case 2:
		fmt.Println(`
				Qual sua unidade de medida de Comprimento?
				1. m²
				2. km²
				3. cm²
				`)
		fmt.Scan(&escolhacomp)
		fmt.Println("Qual o valor do comprimento?")
		fmt.Scan(&comprimento)
		var m2, cm2, km2 = calcComprimento (escolhacomp, comprimento)
		fmt.Printf(
		"\n================= CONVERSÃO DE COMPRIMENTO =================\n"+
		"m²     : %.2f m²\n"+
		"km²    : %.10f km²\n"+
		"cm²    : %.2f cm²\n"+
		"============================================================\n",
		m2, cm2, km2,
		)

	case 3:
			fmt.Println(`
				Qual sua unidade de medida de Velocidade?
				1. m/s
				2. km/h
		`)
			fmt.Scan(&escolhavel)

			var velocidade float64

			switch escolhavel {
			case 1:
				fmt.Println("Digite a velocidade em m/s:")
				fmt.Scan(&velocidade)
				fmt.Printf("Velocidade em m/s: %.2f m/s", velocidade)
				fmt.Printf("Velocidade em km/h: %.2f km/h", velocidade * 3.6)
			case 2:
				fmt.Println("Digite a velocidade em km/h:")
				fmt.Scan(&velocidade)
				fmt.Printf("Velocidade em km/h: %.2f km/h", velocidade)
				fmt.Printf("Velocidade em m/s: %.2f m/s", velocidade / 3.6)
			default:
				fmt.Println("Escolha Inválida")
			}
	case 4:
		fmt.Println(`
				Qual sua unidade de medida de Tempo?
					1. Segundos (s)
					2. Minutos (min)
					3. Horas (h)
					4. Dias (D)
		`)
		fmt.Scan(&escolhatempo)

		var tempo float64

		switch escolhatempo {
			case 1:
				fmt.Println("Insira o tempo em Segundos:")
				fmt.Scan(&tempo)
				fmt.Printf("Tempo em Segundos: %.2f s | ", tempo)
				fmt.Printf("Tempo em Minutos: %.2f min | ", tempo / 60)
				fmt.Printf("Tempo em Horas: %.3f h | ", tempo / 3600)
				fmt.Printf("Tempo em Dias: %.4f Dias", tempo / 86400)
			case 2:
				fmt.Println("Insira o tempo em Minutos:")
				fmt.Scan(&tempo)
				fmt.Printf("Tempo em Segundos: %.1f s | ", tempo * 60)
				fmt.Printf("Tempo em Minutos: %.0f min | ", tempo)
				fmt.Printf("Tempo em Horas: %.2f h | ", tempo / 60)
				fmt.Printf("Tempo em Dias: %.3f", tempo / 1440)
			case 3:
				fmt.Println("Insirad o tempo em Horas:")
				fmt.Scan(&tempo)
				fmt.Printf("Tempo em Segundos: %.3f", tempo * 3600)
				fmt.Printf("Tempo em Minutos: %.2f", tempo * 60)
				fmt.Printf("Tempo em Horas: %.1f", tempo)
				fmt.Printf("Tempo em Dias: %.2f", tempo / 24)
			case 4:
				fmt.Println("Insira o tempo em Dias:")
				fmt.Scan(&tempo)
				fmt.Printf("Tempo em Segundos: %.1f | ", tempo * 86400)
				fmt.Printf("Tempo em Minutos: %.1f | ", tempo * 1440)
				fmt.Printf("Tempo em Horas: %.1f | ", tempo * 24)
				fmt.Printf("Tempo em Dias: %.1f", tempo)
			default:
				fmt.Println("Escolha Inválida")
			}
	case 5:
		fmt.Println(`
				Qual sua unidade de medida de Volume?
				1. m³
				2. L
				3. mL`)
		fmt.Scan(&escolhavol)

		var volume float64

		switch escolhavol {
			case 1:
				fmt.Println("Digite a medida em m³:")
				fmt.Scan(&volume)
				fmt.Printf("Volume em m³: %.2fm³ | ", volume)
				fmt.Printf("Volume em L: %.fL | ", volume * 1000)
				fmt.Printf("Volume em mL: %.fmL", volume * 1000000)
			case 2:
				fmt.Println("Digite a medida em L:")
				fmt.Scan(&volume)
				fmt.Printf("Volume em L: %.2f | ", volume)
				fmt.Printf("Volume em m³: %.4f | ", volume/1000)
				fmt.Printf("Volume em mL: %.f", volume * 1000)
			case 3:
				fmt.Println("Digite a medida em mL:")
				fmt.Scan(&volume)
				fmt.Printf("Volume em mL: %.2f | ", volume)
				fmt.Printf("Volume em L: %.3f | ", volume / 1000)
				fmt.Printf("Volume em m³: %.6f", volume / 1000000)
			default:
				fmt.Println("Escolha Inválida")
		}
	default:
		fmt.Println(`Opção Inválida`)

	}
}

