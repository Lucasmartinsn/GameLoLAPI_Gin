# new-api-gin

comando uteis:

install pgadmin

# Instalando chaves:

sudo curl https://www.pgadmin.org/static/packages_pgadmin_org.pub | sudo apt-key add

#Configurar repositorios:
sudo sh -c 'echo "deb https://ftp.postgresql.org/pub/pgadmin/pgadmin4/apt/$(lsb_release -cs) pgadmin4 main" > /etc/apt/sources.list.d/pgadmin4.list && apt update'

#Instalar PgAdmin4 Web e Desktop
sudo apt install pgadmin4

#erro do pgadmin
//nao conectar ao server postgres
usar o endereço IP 127.0.0.1 em ADD NEW SERVER

# Documentação da API

## Metodos GETs

## GET todos.

URL - http://localhost:8080/user
Resposta:

```json
[
  {
    "id": 1,
    "foto": "50",
    "nome_usuario": "sakura uchiha",
    "Email": "temarelove@",
    "nick_game": "wpk",
    "descricao": "",
    "days_play": "domingo",
    "position_game": "camper",
    "play_time": 2,
    "comunication": "1"
  }
]
```

## Get one

URL - http://localhost:8080/user/1
Resposta:

```json
{
  "id": 1,
  "foto": "50",
  "nome_usuario": "sakura uchiha",
  "Email": "temarelove@",
  "nick_game": "wpk",
  "descricao": "",
  "days_play": "domingo",
  "position_game": "camper",
  "play_time": 2,
  "comunication": "1"
}
```

## Get one apenas os campos(nick_game,descricao,days_play,position_game,Play_time,comunication)
URL - http://localhost:8080/user/info/1
Resposta:

```json
{
  "id": 1,
  "nick_game": "wpk",
  "descricao": "",
  "days_play": "domingo",
  "position_game": "camper",
  "play_time": 2,
  "comunication": "1"
}
```

## Get runas principais

URL - http://localhost:8080/user/runes/principais
Resposta:
```json
{
  "id": 8000,
  "img": "https://encrypted-tbn0.gstatic.com/images?        q=tbn:ANd9GcTEGqAj7vuYaV7Fvpz3HA9TW1dQmDvsf7X2xw&usqp=CAU",
  "nome": "Precision"
}
```
## Get runas secundarias

URL - http://localhost:8080/user/runes/secundarias
Resposta:
```json
  [
    {
		"id": 8100,
		"id_rune": 8128,
		"img": "nao achei",
		"nome": "DarkHarvest"
	},
	{
		"id": 8100,
		"id_rune": 9923,
		"img": "nao achei",
		"nome": "Hail of Blades"
	},
	{
		"id": 8300,
		"id_rune": 8351,
		"img": "nao achei",
		"nome": "Glacial Augment"
	}
  ]
```
## Metodo POST 

URL - http://localhost:8080/user
Forma de envio:

```json
  {
    "foto": "https://www.kindpng.com/picc/m/378-3787490_overlord-anime-poster-png-download-overlord-season-       2.png",
    "nome_usuario": "kakashi hatake",
    "email": "rinandobito@",
    "senha": 987854,
    "nick_game":"logt",
    "descricao":"pegador de casada",
    "days_play":"segunda",
    "position_game":"camper",
    "Play_time": 4,
    "comunication":"chat"
  }
```
## Metodos PUTs

## Put todos os campos

URL - http://localhost:8080/user/3
Forma de envio:
```json
  {
    "foto": "https://www.kindpng.com/picc/m/378-3787490_overlord-anime-poster-png-download-overlord-season-       2.png",
    "nome_usuario": "kakashi hatake",
    "email": "rinandobito@",
    "senha": 987854,
    "nick_game":"logt",
    "descricao":"pegador de casada",
    "days_play":"segunda",
    "position_game":"camper",
    "Play_time": 4,
    "comunication":"chat"
  }
```

## Put campo Foto

URL - http://localhost:8080/user/1/foto
Forma de envio:
```json
  {
    "foto": "https://www.kindpng.com/picc/m/378-3787490_overlord-anime-poster-png-download-overlord-season-       2.png",
  }
  ```
  ## Put campo de Nome
  
  URL - http://localhost:8080/user/1/name
  Forma de envio:
  ```json
  {
    "nome_usuario": "sakura uchiha"
  }
  ```
  ## Put informações sobre o Jogo
  
  URL - http://localhost:8080/user/info/1
  Forma de envio:
  ```json
    {
      "nick_game": "logt",
      "descricao": "",
      "days_play": "segunda",
      "position_game": "camper",
      "play_time": 4,
      "comunication": "chat"
     }
    ```
    ## Metodo DELETE
    URL - http://localhost:8080/user/1
