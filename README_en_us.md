<h1 align="center"> SMTP Hub </h1>

<p align="center"> <strong>Open-Source Send Email Assistant</strong> </p>

[![link to Go version](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-smtp-hub)](https://img.shields.io/github/go-mod/go-version/fonteeboa/go-smtp-hub)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/fonteeboa/go-smtp-hub)
[![go report card](https://goreportcard.com/badge/github.com/fonteeboa/go-smtp-hub "go report card")](https://goreportcard.com/report/github.com/fonteeboa/go-smtp-hub)

<span style="color:red;">**English README**</span>: Dive into the English version of the project's documentation [here](https://github.com/fonteeboa/go-smtp-hub/blob/master/README_en_us.md)

SMTP Hub is a Go library focused on simplifying the process of sending emails via SMTP. Designed to offer a streamlined experience, this library can be configured and integrated with four different types of databases.

<h2 align="center"> <strong>Features</strong> </h2>

🔹 Simplified Email Sending: Provides a simplified API to configure and send emails through SMTP servers.

🔹 Support for Different SMTP Servers: Compatible with a variety of email service providers.

🔹 Database Connection: SMTP Hub can integrate with different types of databases such as PostgreSQL, MySQL, SQLite, and MongoDB. Configuration is simple, using environment variables to specify connection details.

🔹 Message Customization: Allows customization of email content and details.

<h2 align="center"> <strong>Configuration</strong> </h2>

<h4 align="center"> <strong>Relational Database</strong> </h4>

#### PostgreSQL
```
POSTGRES_HOST: Sets the host address for PostgreSQL.
POSTGRES_EXTERNAL_PORT: Specifies the external port for PostgreSQL.
POSTGRES_USER: Username for authentication in PostgreSQL.
POSTGRES_PASSWORD: Password for authentication in PostgreSQL.
POSTGRES_DB: Name of the PostgreSQL database to be used.
```
#### MySQL
```
MYSQL_HOST: Sets the host address for MySQL.
MYSQL_PORT: Specifies the port for MySQL.
MYSQL_USER: Username for authentication in MySQL.
MYSQL_PASSWORD: Password for authentication in MySQL.
MYSQL_DBNAME: Name of the MySQL database to be used.
```
#### SQLite
```
SQLITE_PATH: Path to the SQLite file, if this is the chosen database.
```
<h4 align="center"> <strong>NoSQL Database</strong> </h4>

#### MongoDB
```
MONGODB_URI: Sets the connection URI for MongoDB.
MONGODB_DBNAME: Name of the MongoDB database to be used.
```
<h4 align="center"> <strong>General Configuration</strong> </h4>

```
DATABASE_TYPE: Specifies the type of database to be used (Values: sqlite, postgres, mysql, mongodb).
```

<h4 align="center"> <strong>Notes</strong> </h4>

To use functions related to the SMTP Hub's database, using the DATABASE_TYPE variable is mandatory, as some validations are performed based on this variable before calling the routines.

Make sure to provide valid and correct values for each of these environment variables. This ensures proper connection and correct functioning of the library with the desired database.
