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

type RespostaLooper struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	totalRequisicoes = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requisicoes_total",
		Help: "O volume total de requisições HTTP recebidas",
	})

	disponibilidadeServico = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_servico_disponivel",
		Help: "Disponibilidade do serviço (1 = Up, 0 = Down)",
	})
)

func handlerProjetoLooper(w http.ResponseWriter, r *http.Request) {
	totalRequisicoes.Inc()

	resposta := RespostaLooper{
		Nome:    "Projeto Looper",
		Horario: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	disponibilidadeServico.Set(1)

	http.HandleFunc("/projeto-looper", handlerProjetoLooper)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Servidor http-server-projeto-looper iniciado na porta 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
