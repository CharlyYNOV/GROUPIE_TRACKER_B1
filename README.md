# Groupie Tracker  - Projet B1

<img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go" alt="Go Version" />
<img src="https://img.shields.io/badge/Platform-Linux%20Windows-lightgrey" alt="Platform" />

## Description 
Groupie Tracker est une application web permettant de suivre vos groupes de musique préférés. Découvrez les membres des groupes, leurs dates de création, le premier album, ainsi que les dates et lieux de leurs prochains concerts.

---
## Technologies utilisées 
- Backend: Go ([Golang](https://go.dev/dl/))
- Frontend: HTML, CSS
- APIs: [Groupietrackers's API](https://groupietrackers.herokuapp.com/api)

---
## Installation
```bash
git clone https://github.com/CharlyYNOV/GROUPIE_TRACKER_B1.git
go build -o Groupie-Tracker #.exe si vous êtes sous windows
```

---
## Architecture
```bash
.
├── go.mod
├── internals
│   └── api.go
├── main.go
├── static
│   └── css
│       └── accueil.css
└── templates
    └── accueil.html
```
