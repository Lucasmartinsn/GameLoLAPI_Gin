# GameLoLAPI_Gin

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

# API documentation

###### User manipulation

## Metodos GETs

## Get all users

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

## Get one user

URL - http://localhost:8080/user/id

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

## Get one user, info the game.
URL - http://localhost:8080/user/info/id

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

## Put all fields

URL - http://localhost:8080/user/id

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

## Put fields Photograph

URL - http://localhost:8080/user/id/foto

Forma de envio:
```json
  {
    "foto": "https://www.kindpng.com/picc/m/378-3787490_overlord-anime-poster-png-download-overlord-season-       2.png",
  }
  ```
  ## Put fields name
  
  URL - http://localhost:8080/user/id/name
  
  Forma de envio:
  ```json
  {
    "nome_usuario": "sakura uchiha"
  }
  ```
  ## Put fields password
  
  URL - http://localhost:8080/user/id/password
  
  Forma de envio:
  ```json
  {
    "senha": "sakura uchiha"
  }
  ```

  ## Put info the game
  
  URL - http://localhost:8080/user/info/id
  
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
    URL - http://localhost:8080/user/id

    
### rune manipulation

## Get all rune mains

URL - http://localhost:8080/user/runes/principais

Resposta:
```json
{
  "id": 8000,
  "img": "https://encrypted-tbn0.gstatic.com/images?        q=tbn:ANd9GcTEGqAj7vuYaV7Fvpz3HA9TW1dQmDvsf7X2xw&usqp=CAU",
  "nome": "Precision"
}
```

## Get all runas secondary

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

## Get all rune secondary.

This method will return all child runes linked to a parent rune. The id parameter must be the id of the main rune.

URL - http://localhost:8080/user/runes/secundarias/id

Resposta:
```json
  [
    {
		"id": 8100,
		"id_rune": 8128,
		"img": "nao achei",
		"nome": "DarkHarvest"
	}
  ]
```
