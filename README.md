
<h1 align="center"> SMTP Hub </h1>

<p align="center"> <strong>Open-Source Send Email Assistant</strong> </p>

[![link to Go version](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-smtp-hub)](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-smtp-hub)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/fonteeBoa/go-smtp-hub)
[![go report card](https://goreportcard.com/badge/github.com/fonteeBoa/go-smtp-hub "go report card")](https://goreportcard.com/report/github.com/fonteeBoa/go-smtp-hub)

<span style="color:red;">**English README**</span>: Dive into the English version of the project's documentation [here](https://github.com/fonteeboa/go-smtp-hub/blob/master/README_en_us.md)

O SMTP Hub é uma biblioteca desenvolvida em Go focada em simplificar o processo de envio de e-mails por SMTP. Projetada para oferecer uma experiência simplificada, esta biblioteca pode ser configurada e integrada com quatro tipos de bancos de dados diferentes.

<h2 align="center"> <strong>Funcionalidades</strong> </h2>

🔹 Envio Simplificado de E-mails: Fornece uma API simplificada para configurar e enviar e-mails através de servidores SMTP.

🔹 Suporte a Diferentes Servidores SMTP: Compatível com uma variedade de provedores de serviços de e-mail.

🔹 Conexão com Banco de Dados: O SMTP Hub pode se integrar a diferentes tipos de banco de dados, como PostgreSQL, MySQL, SQLite e MongoDB. A configuração é simples, utilizando variáveis de ambiente para especificar os detalhes de conexão.

🔹 Personalização de Mensagens: Permite personalizar o conteúdo e os detalhes dos e-mails a serem enviados.

<h2 align="center"> <strong>Configuração</strong> </h2>

<h4 align="center"> <strong>Banco de Dados Relacional</strong> </h4>

#### PostgreSQL
```
POSTGRES_HOST: Define o endereço do host para o PostgreSQL.
POSTGRES_EXTERNAL_PORT: Especifica a porta externa para o PostgreSQL.
POSTGRES_USER: Nome de usuário para autenticação no PostgreSQL.
POSTGRES_PASSWORD: Senha para autenticação no PostgreSQL.
POSTGRES_DB: Nome do banco de dados PostgreSQL a ser utilizado.
```
#### MySQL
```
MYSQL_HOST: Define o endereço do host para o MySQL.
MYSQL_PORT: Especifica a porta para o MySQL.
MYSQL_USER: Nome de usuário para autenticação no MySQL.
MYSQL_PASSWORD: Senha para autenticação no MySQL.
MYSQL_DBNAME: Nome do banco de dados MySQL a ser utilizado.
```
#### SQLite
```
SQLITE_PATH: Caminho do arquivo SQLite, se for o banco de dados escolhido.
```

<h4 align="center"> <strong>Banco de Dados NoSQL</strong> </h4>

#### MongoDB
```
MONGODB_URI: Define o URI de conexão para o MongoDB.
MONGODB_DBNAME: Nome do banco de dados MongoDB a ser utilizado.
```

<h4 align="center"> <strong>Configuração Geral</strong> </h4>

```
DATABASE_TYPE: Especifica o tipo de banco de dados a ser utilizado (Valores: sqlite, postgres, mysql, mongodb).
```

<h4 align="center"> <strong>Observações</strong> </h4>

Para utilizar as funções vinculadas ao banco do SMTP Hub, é obrigatório o uso da variável DATABASE_TYPE, pois algumas validações são executadas com base nesta variável antes de chamar as rotinas.

Certifique-se de fornecer valores válidos e corretos para cada uma dessas variáveis de ambiente. Isso garante uma conexão adequada e o funcionamento correto da biblioteca com o banco de dados desejado.
