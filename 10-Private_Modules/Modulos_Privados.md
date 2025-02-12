Módulos Privados no Go

No Go, é possível utilizar módulos privados em seus projetos. Para isso, é necessário configurar o acesso aos repositórios privados. Existem duas principais configurações que precisam ser feitas: a criação de um arquivo `.netrc` e a configuração do `gitconfig`.

1. Arquivo .netrc

O arquivo `.netrc` é utilizado para armazenar credenciais de acesso a repositórios privados. Ele deve ser criado no diretório home do usuário (`~/.netrc`). O conteúdo do arquivo deve seguir o seguinte formato:

```
machine <hostname>
login <username>
password <token>
```

Por exemplo, se você estiver utilizando o GitHub, o arquivo `.netrc` pode ser configurado da seguinte forma:

```
machine github.com
login seu-usuario
password seu-token-de-acesso
```

2. Configuração do gitconfig

Além do arquivo `.netrc`, é necessário configurar o `gitconfig` para que o Go possa acessar os repositórios privados. Adicione as seguintes linhas ao seu arquivo `~/.gitconfig`:

```
[url "https://<username>:<token>@github.com/"]
    insteadOf = https://github.com/
```

Substitua `<username>` pelo seu nome de usuário e `<token>` pelo seu token de acesso.

3. Chave SSH

Outra forma de acessar repositórios privados é utilizando uma chave SSH. Primeiro, gere uma chave SSH (caso ainda não tenha uma) com o comando:

```
ssh-keygen -t rsa -b 4096 -C "seu-email@example.com"
```

Adicione a chave pública gerada ao seu serviço de hospedagem de repositórios (GitHub, Bitbucket, etc.). Em seguida, configure o `ssh-agent` para usar a chave privada:

```
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa
```

4. Configuração para Bitbucket

Para acessar repositórios privados no Bitbucket, é necessário adicionar uma configuração específica no arquivo `.netrc`:

```
machine app.bitbucket.org
login seu-usuario
password seu-token-de-acesso
```

Além disso, adicione a seguinte configuração ao seu `~/.gitconfig`:

```
[url "https://<username>:<token>@bitbucket.org/"]
    insteadOf = https://bitbucket.org/
```

5. Proxy do Go

O Go utiliza um proxy para gerenciar dependências. Para configurar o proxy, defina a variável de ambiente `GOPROXY`:

```
export GOPROXY=https://proxy.golang.org,direct
```

6. Uso do Vendor

O diretório `vendor` é utilizado para armazenar todas as dependências do projeto localmente. Isso garante que todas as dependências estejam disponíveis, mesmo sem acesso à internet. Para criar o diretório `vendor`, execute:

```
go mod vendor
```

Vantagens de utilizar o `vendor`:
- Garantia de que todas as dependências estão disponíveis localmente.
- Maior controle sobre as versões das dependências.
- Facilita a reprodução do ambiente de desenvolvimento em diferentes máquinas.

Com essas configurações, o Go será capaz de acessar e baixar módulos privados dos repositórios configurados, além de gerenciar dependências de forma eficiente.

