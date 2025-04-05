package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()

	url := "http://localhost:8080/cotacao"
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao fazer requisição:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Resposta com status inesperado", errors.New(resp.Status))
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Erro ao ler corpo da resposta:", err)
		return
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Erro ao decodificar JSON:", err)
		return
	}

	bid, ok := result["bid"]
	if !ok {
		log.Println("Campo 'bid' não encontrado na resposta")
		return
	}

	content := fmt.Sprintf("Dólar: %s\n", bid)
	if err := os.WriteFile("cotacao.txt", []byte(content), 0644); err != nil {
		log.Println("Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Println("Cotação salva em cotacao.txt", content)

}
