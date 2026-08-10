package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type RespostaKorp struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	// Métrica 1: Volume de requisições (Counter)
	totalRequisicoes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requisicoes_total",
		Help: "O volume total de requisições HTTP recebidas",
	})

	// Métrica 2: Disponibilidade do serviço (Gauge)
	disponibilidadeServico = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_servico_disponivel",
		Help: "Disponibilidade do serviço (1 = Up, 0 = Down)",
	})
)

func handlerProjetoKorp(w http.ResponseWriter, r *http.Request) {
	totalRequisicoes.Inc()

	resposta := RespostaKorp{
		Nome:    "Projeto Korp",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	disponibilidadeServico.Set(1)

	http.HandleFunc("/projeto-korp", handlerProjetoKorp)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor http-server-projeto-korp iniciado na porta 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
