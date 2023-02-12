<h1 align="center">
    Bot Discord WIth Go 🐹
</h1>

<p align="center">
  <a href="#-projeto">Projeto</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-rodando">Rodando</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-como-contribuir">Como contribuir</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

<br>

<a id="-projeto"></a>

## 💻 Projeto

Um bot de música para discord feito com golang, utilizando libs como FFMPEG e youtube-DL

<a id="-rodando"></a>

## Requerimentos:

- [Golang](https://go.dev/)

## 📂 Instalando as dependências:

```bash
go mod download
```

Para limpar dependências inúteis

```bash
go mod tidy
```

## Executando o Go-Wallpaper 🌇

Primeiro clone o projeto

```bash
git clone https://github.com/Drack112/Go-Music-Discord.git
```

Pegue o TOKEN do seu bot na plataforma de desenvolvedores do discord, crie um arquivo .env de acordo com o .env.example e então execute o bot.

```
TOKEN=
```

```bash
go build -o main . && ./main
```

<a id="-como-contribuir"></a>

## 🤔 Como contribuir

- Faça um fork desse repositório;
- Cria uma branch com a sua feature: `git checkout -b minha-feature`;
- Faça commit das suas alterações: `git commit -m 'feat: Minha nova feature'`;
- Faça push para a sua branch: `git push origin minha-feature`.

Depois que o merge da sua pull request for feito, você pode deletar a sua branch.
