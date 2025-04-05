package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type ApiResponse struct {
	USDBRL Cotacao `json:"USDBRL"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "cotacao.db")
	if err != nil {
		log.Fatal("Erro ao abrir conexão com o banco:", err)
	}
	defer db.Close()

	if err := criarTabela(); err != nil {
		log.Fatal("Erro ao criar tabela:", err)
	}

	http.HandleFunc("/cotacao", cotacaoHandler)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func criarTabela() error {
	query := `
	CREATE TABLE IF NOT EXISTS cotacoes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`
	_, err := db.Exec(query)
	return err
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctxAPI, cancelAPI := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Erro ao criar requisição para API", http.StatusInternalServerError)
		log.Println("Erro ao criar requisição:", err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao buscar cotação", http.StatusInternalServerError)
		log.Println("Erro ao buscar cotação:", err)
		return
	}
	defer resp.Body.Close()

	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		http.Error(w, "Erro ao decodificar resposta da API", http.StatusInternalServerError)
		log.Println("Erro ao decodificar resposta:", err)
		return
	}

	bid := apiResp.USDBRL.Bid

	ctxDB, cancelDB := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancelDB()

	if err := salvarCotacao(ctxDB, bid); err != nil {
		log.Println("Erro ao salvar cotação no banco:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"bid": bid}); err != nil {
		log.Println("Erro ao enviar resposta para o cliente:", err)
	}
}

func salvarCotacao(ctx context.Context, bid string) error {
	_, err := db.ExecContext(ctx, `INSERT INTO cotacoes (bid) VALUES (?)`, bid)
	return err
}
