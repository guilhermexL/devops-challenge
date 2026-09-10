# Projeto Looper - DevOps Portfolio

## Descrição

Projeto demonstrativo que implementa um serviço HTTP simples escrito em Go e focado em práticas de DevOps: containerização, orquestração local, e observabilidade com Prometheus/Grafana. Serve como referência prática para construir, empacotar e monitorar um pequeno serviço em ambiente conteinerizado.

## O que a aplicação faz

- Expõe o endpoint HTTP GET `/projeto-looper` que retorna um JSON com o nome do projeto e o horário atual (UTC).
- Expõe métricas Prometheus em `/metrics`, incluindo:
  - `http_requisicoes_total` (counter): total de requisições HTTP recebidas
  - `http_servico_disponivel` (gauge): disponibilidade do serviço (1 = up, 0 = down)

## Principais arquivos

- `main.go` — código-fonte do serviço HTTP e métricas Prometheus.
- `main_test.go` — testes automatizados unitários e de integração para validar o endpoint HTTP.
- `Dockerfile` — build multi-stage para gerar imagem mínima em Alpine.
- `docker-compose.yml` — orquestração local com: app, nginx (proxy), Prometheus e Grafana (dashboard provisionado).
- `prometheus/prometheus.yml` — configuração do Prometheus para raspar o serviço.
- `grafana/provisioning` — datasource e dashboard pré-configurados.
- `playbook.yml` — playbook Ansible para instalar dependências locais, subir o compose e validar o endpoint.

## Requisitos

- Git
- Docker e Docker Compose
- Go (para build/local) — versão >= 1.25 (opcional se usar Docker)
- (Opcional) Ansible para executar `playbook.yml`

## Executando localmente (modo recomendado: Docker Compose)

1. Clone o repositório e entre na pasta:

   ```bash
   git clone https://github.com/guilhermexL/devops-challenge.git
   cd devops-challenge
   ```

2. Suba os serviços (build da imagem da aplicação + Prometheus/Grafana/nginx):

   ```bash
   docker compose up --build
   ```

3. Acesse:

- **API:** http://localhost/projeto-looper (via Nginx, porta 80)
- **Prometheus:** http://localhost:9090 (se exposto localmente pelo compose)
- **Grafana:** http://localhost:3000

## Exemplo de chamada ao endpoint

```bash
curl -s http://localhost/projeto-looper | jq
```

Resposta esperada:

```json
{
  "nome": "Projeto Looper",
  "horario": "2026-09-09T12:00:00Z"
}
```

## Build e execução sem Docker

1. Fazer build local (necessário Go instalado):

   ```bash
   go build -o http-server-projeto-looper .
   ./http-server-projeto-looper
   ```

2. O serviço ficará disponível na porta 8080, e as métricas em `/metrics`.

## Observabilidade

- **Prometheus** é configurado (`prometheus/prometheus.yml`) para raspar o endpoint do serviço.
- **Grafana** tem um dashboard provisionado em `grafana/provisioning/dashboards/` que mostra disponibilidade e taxa de requisições.
- Métricas relevantes:
  - `http_requisicoes_total`
  - `http_servico_disponivel`

## Playbook Ansible

O playbook `playbook.yml` automatiza:
- Instalação de dependências do sistema (apt)
- Instalação do Docker e plugin docker-compose
- Execução de `docker compose up --build`
- Teste HTTP para validar que `/projeto-looper` responde 200

Uso:

```bash
ansible-playbook playbook.yml --connection=local
```

## Testes e lint

- Há testes automatizados incluídos neste repositório em `main_test.go`. Para executá-los:
  ```bash
  go test -v ./...
  ```
- Recomenda-se aplicar `gofmt` antes de enviar alterações.

## Boas práticas e notas de produção

- A imagem usa build multi-stage para reduzir tamanho final (alpine runtime).
- Produção: adicionar healthchecks, logs estruturados, e configuração de variáveis via ENV.
- Segurança: evite expor serviços de administração sem autenticação; proteger Grafana/Prometheus em ambientes públicos.

---
*Este repositório serve como portfólio de engenharia e simulação do desafio técnico para o cargo de DevOps na Looper.*
