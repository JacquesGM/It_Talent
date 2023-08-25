package main

import (
	"fmt"
	"math"
)

func main() {
	var tamanhoArquivo float64
	var velocidadeInternet float64

	// Solicitar o tamanho do arquivo para download (em MB)
	fmt.Print("Digite o tamanho do arquivo para download (em MB): ")
	fmt.Scan(&tamanhoArquivo)

	// Solicitar a velocidade da Internet (em Mbps)
	fmt.Print("Digite a velocidade da Internet (em Mbps): ")
	fmt.Scan(&velocidadeInternet)

	// Calcular o tempo aproximado de download (em minutos)
	tempoDownload := (tamanhoArquivo * 8) / velocidadeInternet // Converter de MB para bits
	tempoDownloadMinutos := math.Ceil(tempoDownload / 60)      // Converter de segundos para minutos

	// Exibir o tempo aproximado de download
	if tempoDownloadMinutos <= 1 {
		fmt.Printf("O tempo aproximado de download é de %.2f minuto.\n", tempoDownloadMinutos)
	} else {
		fmt.Printf("O tempo aproximado de download é de %.2f minutos.\n", tempoDownloadMinutos)
	}

}
