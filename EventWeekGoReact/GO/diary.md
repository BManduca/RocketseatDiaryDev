# Projeto Ask me anything (back-end/server)

## Primeiros passos

- Com o repositório já criado, é preciso efetuar o clone do mesmo
- Após clonar o repositório, basta acessar e efetuar o seguinte comando: go mod init 'url do repositorio', para inicializar o módulo de go
- Criar um arquivo docker-compose e configurar da seguinte forma(configs usada na minha aplicação):
    >
        services:
            db:
                image: postgres:latest
                restart: unless-stopped
                ports:
                # {porta publica}:{porta acessada dentro do container}
                # porta publica -> acessamos por fora do container
                # porta acessada do container -> para networking interna que o docker cria
                - ${WSBM_DATABASE_PORT:-5432}:5432
                environment:
                POSTGRES_USER: ${WSBM_DATABASE_USER}
                POSTGRES_PASSWORD: ${WSBM_DATABASE_PASSWORD}
                POSTGRES_DB: ${WSBM_DATABASE_NAME}
                volumes:
                - db:/var/lib/postgresql/data
            
            pgadmin:
                image: dpage/pgadmin4:latest
                restart: unless-stopped
                # o serviço depende do DB, ou seja, o serviço so vai subir, assim que o DB tenha subido.
                depends_on:
                - db
                ports:
                - ${WSBM_PGADMIN_PORT}:80
                environment:
                PGADMIN_DEFAULT_EMAIL: ${WSBM_PGADMIN_EMAIL}
                PGADMIN_DEFAULT_PASSWORD: ${WSBM_PGADMIN_PASSWORD}
                volumes:
                - pgadmin:/var/lib/pgadmin
                
                
            volumes:
            db:
                # Armazenamento persistente dos dados do banco
                # clonar os arquivos file system para o files system da maquína host
                driver: local

            pgadmin:
                driver: local

- Para as varíaveis de ambiente, iremos definir as mesmas em uma arquivo .env
    >
        WSBM_DATABASE_PORT=5432
        WSBM_DATABASE_NAME="your_database_name"
        WSBM_DATABASE_USER="your_user_name"
        WSBM_DATABASE_PASSWORD="your_password"
        WSBM_PGADMIN_PORT=8081
        WSBM_PGADMIN_EMAIL="your_email"
        WSBM_PGADMIN_PASSWORD="your_password"

- Criando as migrations do nosso DB
  - Para esse passo será usado a ferramenta [tern](https://github.com/jackc/tern)
  - Basicamente é uma ferramenta utilizada para fazer as migrações do nosso banco
  - instalação: go install github.com/jackc/tern/v2@latest
  - Criando um diretório internal
    - Todo o código dentro deste diretório é interno para este pacote, desta forma, não é possível utilizar nenhuma funcionalidade dentro deste internal fora do meu módulo.
    - Normalmente quando criamos pacotes executaveis, ou seja, binários, nos os colocamos dentro de internal, pois, nào iremos importar um binário como dependência.
  - inicializando o tern: tern init ./internal/store/pgstore/migrations
  - após a inicialização, vamos configurar o arquivo tern.conf da seguinte forma:
    >
        [database]
        port = {{ env WSBM_DATABASE_PORT }}
        database = {{ env WSBM_DATABASE_NAME }}
        user = {{ env WSBM_DATABASE_USER }}
        password = {{ env WSBM_DATABASE_PASSWORD }}
        host = {{ env WSBM_DATABASE_HOST }}
  - Após a configuração, iremos criar nossas migrations da seguinte forma
    - tern new --migrations ./internal/store/pgstore/migrations/ create_rooms_table
    - tern new --migrations ./internal/store/pgstore/migrations/ create_messages_table
  - Depois de criado, iremos configurar cada migration
  - O tern ele não tem a opção de ler as varíaveis de ambiente através de um arquivo .env
  - E para setar as varíaveis de ambiente no meu ambiente, será criada uma ferramenta, que vai ser um wrap do nosso tern, ou seja, vai literalmente chamar o tern por debaixo dos panos e vai declarar essas varíaveis de ambiente lidas desse arquivo .env, para que o tern seja capaz entender tudo isso.
  - Para fazer esse 'carregamento' das varíaveis, estarei utilizando a biblioteca [godotenv](https://github.com/joho/godotenv)
    - Que basicamente ela permite que seja carregado as varíaveis de ambiente através de um arquivo .env
  - Pode ser que aqui gere um erro, pois, o godotenv, ainda não se encontrar em nossas dependências e para solucionar isso, basta executar: go mod tidy
  - E no final nosso wrap, ficará da seguinte forma:
    >
        package main

        import "github.com/joho/godotenv"

        func main() {
                if err := godotenv.Load(); err != nil {
                    panic(err)
                }

                cmd := exec.Command(
                    "tern", 
                    "migrate", 
                    "--migrations", 
                    "./internal/store/pgstore/migrations", 
                    "--config", 
                    "./internal/store/pgstore/migrations/tern.conf"
                )

                if err := cmd.Run(); err != nil {
                    panic(err)
                }
        }
- Para executar, basta rodar o comando => go run cmd/tools/terndotenv/main.go
- teste para validação das configurações do tenf:
  - tern migrate --migrations ./internal/store/pgstore/migrations --config ./internal/store/pgstore/migrations/tern.conf
- Verificar processos que estão rodando em determinada porta
  - lsof -i :{numero_porta}

## Representando entidades do banco de dados
- Para isso será utilizado o pacote [sqlc](https://docs.sqlc.dev/en/stable/index.html)
  - Basicamente é um gerador de código
  - Focado principalmente em GO, mas funciona para outras linguagens
  - Principal funcionalidade: Gera código GO, para interagir com nosso banco de dados
  - Ele não é um ORM (em go não é aconselhado utilizar ORM)
  - Instalação:
    - Há duas formas hoje em que eu posso estar instalando o sqlc, via homebrew ou através do go mesmo:
      - Como para o tern, eu utilizei via go, manterei esse padrão
        - go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
  
## Diretivas
- Diretivas -> diz para nosso codigo go, quais comandos ele tem que rodar, quando for rodado o comando go:generate

- basicamente estaremos utilizando desta diretiva para simplificar a execução dos comandos do tern para fazer a migração do nosso banco de dados e do sqlc para gerar todas as nossas entidades e o meu código para interagir com meu banco de dados

- para rodar os dois comandos: go generate ./... -> vai olhar no diretorio que estamos atualmente e em todos os sub-diretorios, para ver se tem alguma diretiva para rodar o go:generate

## Criando routers http em GO
- Precisamos da pacote [chi](github.com/go-chi/chi/v5)
- Pacote para criar routers http em GO
- Através do chi, importaremos o Mux, onde o Mux é um multiplexador de rota HTTP simples que analisa um caminho de solicitação, registra quaisquer parâmetros de URL e executa um manipulador final. Ele implementa a interface http.Handler e é amigável com a biblioteca padrão.


## [Configurando o VSCode para GO](https://efficient-sloth-d85.notion.site/Configurando-o-VSCode-para-Go-ed177054908d48f58095ff54d303f44c)

## Evitando problema com CORS

- A api que estamos construindo vai ser consumida por um front-end, que será feito em React
- O front-end sendo feito em React, pode ocasionar os famosos CORS (Cross-Origin Resource Sharing) que é um sistema que consiste na transmissão de HTTP headers, que determina se navegadores vão bloquear código JavaScript de acessarem respostas provindas de requisições entre origens.
- Será preciso fazer o handle desse cors, pois, se caso a nossa API não responder nada para o pre-flight request, simplesmente o CORS não vai funcionar e o browser por consequência vai bloquear os requests para aquela API
  - Pre-flight request => Determine se a request atual para o servidor do domínio solicitante é segura
- Para tratar esse situação o próprio CHI, tem um pacote chamado [CORS](https://github.com/go-chi/cors)
  - basicamente é um middleware

## Criando websockets em go
- Precisamos utilizar o pacote [Gorilla/WebSocket](https://github.com/gorilla/websocket)
- O pacote gorilla/websocket é uma implementação WebSocket rápida, bem testada e amplamente utilizada para Go.
- Instalação: go get github.com/gorilla/websocket

## Extensões para auxiliar o projeto
- [GO](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Error lens](https://marketplace.visualstudio.com/items?itemName=usernamehw.errorlens)

--------------------------------------------------------------------------------

## Create alias to Generate uuid 

- alias uuid="uuidgen | tr -d '\n'"

## Criando uma "Ordered List"

- No Projeto AMA (Ask Me Anything), depois que criarmos a sala de pergutntas, as perguntas serão apresentadas dentro de um lista e para isso estaremos utilizando o ordered list, da seguinte forma:
    >
      <ol className='list-decimal list-inside'>

      </ol>

  - Aonde teremos um list type como decimal, que assim cada item da lista começará com um numeral, por exemplo 1,2,3 e assim por diante...
  - teremos também o list-inside, que irá fazer com que o list-position seja inside e fará com que o número da listagem, ou seja, de cada item esteja dentro do texto e não fora.

## Trabalhando com requisições no react

- Uma forma de trabalhar com requisições e api dentro do Vite é utilizando o react-query e podemos instalar da seguinte forma:
  - npm i @tanstack/react-query -f
    - como aqui ainda depende da versão 18 do react, teremos que utilizar o -f