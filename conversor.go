package main

import (
	"fmt"
)

func calcTemperatura(escolhatemp int, temperatura float64) (float64, float64, float64) {
	var Celsius, Fahrenheit, Kelvin float64
	switch escolhatemp {
	case 1:
		Celsius = temperatura
		Fahrenheit = (temperatura * 1.8) + 32
		Kelvin = temperatura + 273.15
		return Celsius, Fahrenheit, Kelvin
	case 2:
		Fahrenheit = temperatura
		Celsius = (temperatura - 32) / 1.8
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
func calcComprimento(escolhacomp int, comprimento float64) (float64, float64, float64) {
	var m, km, cm float64
	switch escolhacomp {
	case 1:
		m = comprimento
		km = comprimento / 1000
		cm = comprimento * 100
		return m, km, cm
	case 2:
		km = comprimento
		m = km * 1000
		cm = m * 100
		return m, km, cm
	case 3:
		cm = comprimento
		m = cm / 100
		km = m / 1000
		return m, km, cm
	default:
		m = 0
		cm = 0
		km = 0
		return m, km, cm
	}
}
func calcVelocidade(escolhavel int, velocidade float64) (float64, float64) {
	var ms, kmh float64
	switch escolhavel {
	case 1:
		ms = velocidade
		kmh = ms * 3.6
		return ms, kmh
	case 2:
		kmh = velocidade
		ms = kmh / 3.6
		return ms, kmh
	default:
		kmh = 0
		ms = 0
		return ms, kmh
	}
}
func calcTempo(escolhatempo int, tempo float64) (float64, float64, float64) {
	var segundos, minutos, horas float64
	switch escolhatempo {
	case 1:
		segundos = tempo
		minutos = segundos / 60
		horas = minutos / 60
		return segundos, minutos, horas
	case 2:
		minutos = tempo
		segundos = minutos * 60
		horas = minutos / 60
		return segundos, minutos, horas
	case 3:
		horas = tempo
		minutos = horas * 60
		segundos = minutos * 60
		return segundos, minutos, horas
	default:
		segundos = 0
		minutos = 0
		horas = 0
		return segundos, minutos, horas
	}
}
func calcVolume(escolhavol int, volume float64) (float64, float64, float64) {
	var ml, l, m3 float64
	switch escolhavol {
	case 1:
		m3 = volume
		l = m3 * 1000
		ml = l * 1000
		return ml, l, m3
	case 2:
		l = volume
		m3 = l / 1000
		ml = l * 1000
		return ml, l, m3
	case 3:
		ml = volume
		l = ml / 1000
		m3 = l / 1000
		return ml, l, m3
	default:
		ml = 0
		l = 0
		m3 = 0
		return ml, l, m3
	}
}
func main() {
	var escolha, escolhatemp, escolhacomp, escolhavel, escolhatempo, escolhavol int
	var temperatura, comprimento, velocidade, tempo, volume float64
	fmt.Println(`
		Escolha o tipo de conversão de unidade:
		1. Temperatura (Celsius, Fahrenheit, Kelvin)
		2. Comprimentos (m, km, cm)
		3. Velocidade (m/s, km/h)
		4. Tempo (segundos, minutos, horas)
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
				1. m
				2. km
				3. cm
				`)
		fmt.Scan(&escolhacomp)
		fmt.Println("Qual o valor do comprimento?")
		fmt.Scan(&comprimento)
		var m, cm, km = calcComprimento(escolhacomp, comprimento)
		fmt.Printf(
			"\n================= CONVERSÃO DE COMPRIMENTO =================\n"+
				"m    : %.3f m\n"+
				"km   : %.4f km\n"+
				"cm   : %.2f cm\n"+
				"============================================================\n",
			m, cm, km)
	case 3:
		fmt.Println(`
				Qual sua unidade de medida de Velocidade?
				1. m/s
				2. km/h
		`)
		fmt.Scan(&escolhavel)
		fmt.Println("Qual a velocidade?")
		fmt.Scan(&velocidade)
		var ms, kmh = calcVelocidade(escolhavel, velocidade)
		fmt.Printf(
			"\n================= CONVERSÃO DE VELOCIDADE =================\n"+
				"m/s   : %.2f m/s\n"+
				"km/h  : %.2f km/h\n"+
				"============================================================\n",
			ms, kmh)
	case 4:
		fmt.Println(`
				Qual sua unidade de medida de Tempo?
					1. Segundos (s)
					2. Minutos (min)
					3. Horas (h)
		`)
		fmt.Scan(&escolhatempo)
		fmt.Println("Qual o valor do tempo?")
		fmt.Scan(&tempo)
		var segundos, minutos, horas = calcTempo(escolhatempo, tempo)
		fmt.Printf(
			"\n================= CONVERSÃO DE TEMPO =================\n"+
				"Segundos     : %.2f s\n"+
				"Minutos      : %.2f min\n"+
				"Horas        : %.2f h\n"+
				"======================================================\n",
			segundos, minutos, horas)
	case 5:
		fmt.Println(`
				Qual sua unidade de medida de Volume?
				1. m³
				2. L
				3. mL`)
		fmt.Scan(&escolhavol)
		fmt.Println("Qual o valor do Volume?")
		fmt.Scan(&volume)
		var ml, l, m3 = calcVolume(escolhavol, volume)
		fmt.Printf(
			"\n================= CONVERSÃO DE VOLUME =================\n"+
				"mL     : %.2f mL\n"+
				"L      : %.3f L\n"+
				"m³     : %.6f m³\n"+
				"============================================================\n",
			ml, l, m3)
	default:
		fmt.Println(`Opção Inválida`)
	}
}
