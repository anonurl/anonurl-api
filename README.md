# AnonURL API

[FrontEnd repository](https://github.com/anonurl/anonurl)<br>  

## 📖 Summary
[Changelog](https://github.com/anonurl/anonurl-api#changelog)<br>
[Used technologies](https://github.com/anonurl/anonurl-api#usedtech)<br>
[Endpoints](https://github.com/anonurl/anonurl-api#endpoints)<br>
[Setup](https://github.com/anonurl/anonurl-api#setup)<br>

<br><a name="changelog"></a>
## ♻️ Changelog (v1.0.0c)
- Some fixes

<br><a name="usedtech"></a>
## ⚙️ Used technologies:
- Golang
- MongoDB

<br><a name="endpoints"></a>
## 🔌 Endpoints
<strong>/api/create (POST):</strong> To create shorted URLs<br>
<strong>/api/track/ID (GET):</strong> To get basic information of a shorten URL<br>
<strong>/api/redirect/ID (GET):</strong> To get information of redirect to shorted URL<br>
<strong>/api/report (POST):</strong> To send a report<br>

<br><a name="setup"></a>
## 🔧 Setup
### Clone this repository:
`git clone https://github.com/anonurl/anonurl-api`

### Move to repository:
`cd anonurl-api`

### Configurate:
- Add your Mongo URI to "MONGO_URI" on enviroments variables
- Add a port to "PORT" enviroment variable

### Install all dependencies:
`go get`

### Run:
`go run main.go`
