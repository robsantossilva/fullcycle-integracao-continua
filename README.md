## Integração continua

É o processo de integrar modificações do codebase de forma contínua e automatizada, evitando assim erros humanos de verificação,
garantindo mais agilidade e segurança no processo de desenvolvimento de um software.

### Principais processos
- Execução de testes
- Linter
- Verificação de qualidade de código
- Verificação de segurança
- Geração de artefatos prontos para o processo de deploy
- Identificação da próxima versão a ser gerada no software
- Geração de tags e releases

### Status check Github
É a garantia de que uma Pull Request não poderá ser mergeada ao repositório sem antes ter passado pelo processo de CI ou mesmo no processo de Code Review

### Ferramentas populares
- Jenkins
- Github Actions
- Circle CI
- AWS Code Build
- Azure DevOps
- Google Cloud Build
- GitLab Pipelines / CI

### Github Actions
Automate your worlflow from idea to production

O Github Actions é um automatizador de workflow de desenvolvimento de software.
Ele utiliza os principais eventos gerados pelo GitHub para que possamos executar tarefas dos mais variadostipos, incluindo processos de CI.

### Workflow
- São conjunto de processos definidos por você. Ex: Fazer o build + rodar os testes da aplicação
- É possível ter mais do que um workflow por repositório
- Definidos em arquivos ".yml" em .github/workflows
- Possui um ou mais "Jobs"
- É iniciado baseado em eventos do GitHub ou através de agendamento

| Evento   | Filtros              | Ambiente        | Ações                                                         |
| -------- | -------------------- | --------------- | ------------------------------------------------------------- |
| on: push | branchs:<br>- master | runs-on: ubuntu | steps:<br>- uses: action/runs-composer<br>- run: npm run prod |

### Actions
É a ação que de fato será executada em um dos Steps de um Job em um Workflow.
Ela pode ser criada do zero ou ser reutilizada de actions pre-existentes.
**Action pode ser desenvolvida:**
- Javascript
- Docker Image

https://docs.github.com/pt/actions/automating-builds-and-tests/about-continuous-integration

### Strategy Matrix

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go: [ '1.14', '1.13' ]
    name: Go ${{ matrix.go }} sample
    steps:
      - uses: actions/checkout@v2
      - name: Setup go
        uses: actions/setup-go@v2
        with:
          go-version: ${{ matrix.go }}
      - run: go run hello.go
```

### Build and push Docker images
https://github.com/marketplace/actions/build-and-push-docker-images

