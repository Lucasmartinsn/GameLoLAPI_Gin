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

```json
[
{
create table recordss(
	id serial PRIMARY KEY not null,  
 	foto varchar(2048) NOT NULL ,  
 	nome_usuario VARCHAR(40) NOT NULL,  
 	email VARCHAR(40) NOT NULL,
	senha int not null,
	nick_game VARCHAR(30) NOT NULL,  
 	descricao VARCHAR(200) NOT NULL,  
	days_play VARCHAR(20) NOT NULL,  
 	position_game VARCHAR(20) NOT NULL,  
 	play_time INT NOT NULL,  
 	comunication varchar(10) NOT NULL 
)

}
]

# Documentação da API

## GET todos.
URL - http://localhost:8080/api/user
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

## Get one todos os campos(user)
URL - http://localhost:8080/api/user/1
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

Get one apenas os campos
URL - http://localhost:8080/api/user/1
Resposta:
