# devops-challenge

Visão geral

Este repositório contém a solução para um desafio DevOps cujo objetivo é demonstrar competências em automação, containerização, infraestrutura como código, integração contínua e observabilidade. A aplicação aqui presente é uma base de referência para demonstrar práticas recomendadas de desenvolvimento, implantação e operação em ambientes modernos.

Principais componentes

- Código da aplicação: microserviço/serviço monolítico (dependendo da implementação incluída neste repositório).
- Containerização: Dockerfile(s) para criação de imagens reprodutíveis.
- Orquestração local: docker-compose para facilitar execução e testes locais.
- Infraestrutura (opcional): exemplos de IaC (Terraform/CloudFormation) e manifests Kubernetes quando aplicáveis.
- CI/CD: pipeline de integração contínua (GitHub Actions) para build, testes e publicação de artefatos.
- Observabilidade: recomendações para logs e métricas; integração com Prometheus/Grafana/ELK quando disponível.

Requisitos

- Git
- Docker & Docker Compose (para execução local)
- (Opcional) Terraform, kubectl, Helm — caso utilize infra/cluster Kubernetes
- (Opcional) Conta em Docker Hub / registro de container para publicação de imagens

Instalação e execução local

1. Clonar o repositório:

   git clone https://github.com/guilhermexL/devops-challenge.git
   cd devops-challenge

2. Inspecionar arquivos de apoio (Dockerfile, docker-compose.yml, Makefile, scripts/)

3. Executar com Docker Compose:

   docker-compose up --build

4. Acessar a aplicação em http://localhost:PORT (substituir PORT conforme docker-compose.yml)

Testes

- Executar a suíte de testes automatizados (unit / integração):

  make test

  ou

  ./scripts/test.sh

- Validar pipelines de lint/format conforme as ferramentas configuradas (eslint, flake8, gofmt, etc.).

CI/CD

- Pipeline GitHub Actions (/.github/workflows) realiza:
  - Build da imagem
  - Execução de testes e lint
  - Publicação de imagem em registry (quando marcado)
  - Deploy automático em ambiente de staging/production conforme configuração dos secrets

Deployment

- Para deploy em Kubernetes, verificar manifests em k8s/ ou helm/ e aplicar:

  kubectl apply -f k8s/

- Para deploy em VM ou serviço PaaS, adaptar Dockerfile e usar o registry configurado.

Arquitetura e boas práticas

- Imagens otimizadas e multi-stage builds para reduzir tamanho final.
- Variáveis de ambiente para configuração (12-factor app).
- Healthchecks e readiness probes (quando aplicado em Kubernetes).
- Logs estruturados (JSON) para facilitar ingestão por sistemas de observabilidade.

Estrutura do repositório (exemplo)

- /app                      -> código-fonte da aplicação
- /Dockerfile               -> imagem da aplicação
- /docker-compose.yml       -> orquestração local
- /k8s                     -> manifests Kubernetes (se aplicável)
- /.github/workflows        -> definições de CI/CD
- /scripts                  -> scripts de suporte (build, test, deploy)
- /docs                     -> documentação adicional

Como contribuir

Contribuições são bem-vindas. Para mudanças:

1. Abrir uma issue descrevendo o objetivo
2. Criar um branch com nome descritivo
3. Abrir um Pull Request com descrição, testes e evidências

Licença

Este projeto está licenciado sob a licença MIT — ver arquivo LICENSE para detalhes.

Contato

Para dúvidas ou informações adicionais, abra uma issue no repositório ou contate o mantenedor no perfil do GitHub.

Notas finais

O README oferece um panorama completo e pode ser adaptado para refletir detalhes técnicos específicos da stack usada neste repositório (linguagem, portas, comandos exatos de build/test). Atualize as seções de execução e testes conforme os scripts e arquivos presentes na árvore do projeto.