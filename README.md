### Escrevendo em imagens com golang

Esse projeto tem como objetivo escrever textos em imagens, utilizando a linguagem de programação Go.

### Como usar

Para utilizar o projeto, basta clonar o repositório e executar o comando abaixo:

```bash
go run cmd/cli/main.go D:/TEMP/go/write-in-image/examples/nomes.txt D:/TEMP/saida D:/TEMP/go/write-in-image/examples/template1.jpg 72 650
```
> cmd/cli/main.go D:/TEMP/go/write-in-image/examples/nomes.txt = Caminho do arquivo que contém os nomes que serão escritos na imagem
> D:/TEMP/saida = Caminho do diretório onde as imagens com os nomes escritos serão salvas
> D:/TEMP/go/write-in-image/examples/template1.jpg = Caminho da imagem que será utilizada como base para escrever os nomes
> 72 = Tamanho da fonte que será utilizada para escrever os nomes
> 650 = Altura do texto que será escrito na imagem (parâmetro opcional)


Lembre de alterar os caminhos dos arquivos de acordo com o seu ambiente.
O caminho não poderá conter espaços, pois o programa não está tratando isso.

### Exemplo

O arquivo `nomes.txt` contém os nomes que serão escritos na imagem. O arquivo `template1.jpg` é a imagem que será utilizada como base para escrever os nomes. O diretório `saida` é onde as imagens com os nomes escritos serão salvas e o número `72` é o tamanho da fonte que será utilizada para escrever os nomes.