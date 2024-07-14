### Escrevendo em imagens com golang

Esse projeto tem como objetivo escrever textos em imagens, utilizando a linguagem de programação Go.

### Como usar

Para utilizar o projeto, basta clonar o repositório e executar o comando abaixo:

Você pode executar o comando abaixo para visualizar os parâmetros a serem informados:

```bash
go run .\cmd\cli\main.go writeTextInImage -h
```

Após isto você irá visualizar no terminal os parâmetros que devem ser informados para executar o programa.

```bash
You will be able to write text in an image with this command and some parameters like font, size, color, position, etc.

Usage:
  write-in-image writeTextInImage [flags]

Flags:
  -f, --file string           File path With the text to write in the image (required)
  -c, --font-color string     Font color (optional) default is black (default "000000")
  -s, --font-size string      Font size (optional) default 72 (default "72")
  -h, --help                  help for writeTextInImage
  -o, --outoput-path string   Output path to save the new image (required)
  -t, --template string       Template image path (required)
  -a, --text-height string    Text height (optional) default 650 (default "650")
  ```


Você pode utilizar os shorthands para informar os parâmetros ou os nomes completos. Abaixo um exemplo de como utilizar o comando para escrever textos em imagens:

Exempolo com ShortHand:
```bash
go run .\cmd\cli\main.go writeTextInImage -f "D:\Projetos_Pessoais\golang\write-in-image\examples\nomes.txt" -o D:/TEMP/saida -t  "D:\Projetos_Pessoais\golang\write-in-image\examples\template1.jpg" -s 120 -c 911414 
```

Exemplo com nome completo:
```bash
go run .\cmd\cli\main.go writeTextInImage --file "D:\Projetos_Pessoais\golang\write-in-image\examples\nomes.txt" --outoput-path D:/TEMP/saida --template "D:\Projetos_Pessoais\golang\write-in-image\examples\template1.jpg" --font-size 120 --font-color 911414
```


