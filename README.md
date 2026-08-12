# Teste Técnico - Projeto Korp

Descrição

Projeto demonstrativo que implementa um serviço HTTP simples escrito em Go e focado em práticas de DevOps: containerização, orquestração local, e observabilidade com Prometheus/Grafana. Serve como referência prática para construir, empacotar e monitorar um pequeno serviço em ambiente conteinerizado.

O que a aplicação faz

- Expõe o endpoint HTTP GET /projeto-korp que retorna um JSON com o nome do projeto e o horário atual (UTC).
- Expõe métricas Prometheus em /metrics, incluindo:
  - http_requisicoes_total (counter): total de requisições HTTP recebidas
  - http_servico_disponivel (gauge): disponibilidade do serviço (1 = up, 0 = down)

Principais arquivos

- main.go — código-fonte do serviço HTTP e métricas Prometheus.
- Dockerfile — build multi-stage para gerar imagem mínima em Alpine.
- docker-compose.yml — orquestração local com: app, nginx (proxy), Prometheus e Grafana (dashboard provisionado).
- prometheus/prometheus.yml — configuração do Prometheus para raspar o serviço.
- grafana/provisioning — datasource e dashboard pré-configurados.
- playbook.yml — playbook Ansible para instalar dependências locais, subir o compose e validar o endpoint.

Requisitos

- Git
- Docker e Docker Compose
- Go (para build/local) — versão >= 1.25 (opcional se usar Docker)
- (Opcional) Ansible para executar playbook.yml

Executando localmente (modo recomendado: Docker Compose)

1. Clone o repositório e entre na pasta:

   git clone https://github.com/guilhermexL/devops-challenge.git
   cd devops-challenge

2. Suba os serviços (build da imagem da aplicação + Prometheus/Grafana/nginx):

   docker compose up --build

3. Acesse:

- API: http://localhost/projeto-korp  (via Nginx, porta 80)
- Prometheus: http://localhost:9090    (se exposto localmente pelo compose)
- Grafana: http://localhost:3000

Exemplo de chamada ao endpoint

curl -s http://localhost/projeto-korp | jq

Resposta esperada:

{
  "nome": "Projeto Korp",
  "horario": "2026-08-12T...Z"
}

Build e execução sem Docker

1. Fazer build local (necessário Go instalado):

   go build -o http-server-projeto-korp .
   ./http-server-projeto-korp

2. Serviço ficará disponível na porta 8080, métricas em /metrics.

Observabilidade

- Prometheus é configurado (prometheus/prometheus.yml) para raspar o endpoint do serviço.
- Grafana tem um dashboard provisionado em grafana/provisioning/dashboards/ que mostra disponibilidade e taxa de requisições.
- Métricas relevantes:
  - http_requisicoes_total
  - http_servico_disponivel

Playbook Ansible

O playbook playbook.yml automatiza:
- Instalação de dependências do sistema (apt)
- Instalação do Docker e plugin docker-compose
- Execução de docker compose up --build
- Teste HTTP para validar que /projeto-korp responde 200

Uso:

   ansible-playbook playbook.yml --connection=local

Testes e lint

- Atualmente não há testes automatizados incluídos neste repositório. Para adicionar testes, crie pacotes _test.go e execute `go test ./...`.
- Recomenda-se aplicar gofmt/golangci-lint em pipelines CI antes do build.

Boas práticas e notas de produção

- A imagem usa build multi-stage para reduzir tamanho final (alpine runtime).
- Produção: adicionar healthchecks, logs estruturados, e configuração de variáveis via ENV.
- Segurança: evite expor serviços de administração sem autenticação; proteger Grafana/Prometheus em ambientes públicos.

*Esse repositório serve como realização do desafio técnico para o cargo de DevOps na Korp.*