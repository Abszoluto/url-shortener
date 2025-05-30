
# Encurtador de URLs em Go

Um projeto simples e funcional de encurtador de URLs com interface web, desenvolvido em Go.
O projeto foi desenvolvido com o objetivo de demonstrar a linguagem Go, em um trabalho para a disciplina de Linguagens de Programação.

![Go Version](https://img.shields.io/badge/Go-1.16%2B-blue)
![Status](https://img.shields.io/badge/status-em%20desenvolvimento-yellow)
![Windows Support](https://img.shields.io/badge/platform-Windows%20only-blueviolet)

---

## Funcionalidades

- Interface moderna e responsiva (HTML + CSS)
- Geração de URLs curtas aleatórias
- Redirecionamento automático para a URL original
- Armazenamento em memória (sem banco de dados)

---

## Requisitos

- [Go instalado](https://go.dev/dl/) (versão 1.16 ou superior)
- Git (opcional)

---

## Instalação e Execução (Windows)

### 1. Clone o projeto

```bash
git clone https://github.com/seuusuario/url-shortener-go.git
cd url-shortener-go
```

Ou baixe o `.zip` pelo botão verde **"Code > Download ZIP"**, extraia e abra no terminal.

### 2. Inicialize o módulo Go

```bash
go mod init encurtador
go get github.com/gorilla/mux
```

### 3. Execute o projeto

```bash
go run main.go router.go handlers.go storage.go
```

---

## Acesse no navegador

```
http://localhost:8080
```

---

## Estrutura do Projeto

```
url-shortener-go/
│
├── main.go
├── router.go
├── handlers.go
├── storage.go
│
├── templates/
│   └── index.html
│
└── public/
    └── style.css
```

---

## Exemplo de uso

1. Acesse `http://localhost:8080`
2. Cole uma URL no campo
3. Clique em **"Encurtar"**
4. Use a URL gerada para redirecionar

---

## Compilar para `.exe` (opcional)

```bash
go build -o encurtador.exe
.\encurtador.exe
```

---

## Observações

- Este projeto **não utiliza banco de dados** — os dados são perdidos ao reiniciar.
- Código simples e didático, ideal para aprender os fundamentos de Go Web.

---

## 📝 Licença

Este projeto é livre para uso educacional e pessoal.
