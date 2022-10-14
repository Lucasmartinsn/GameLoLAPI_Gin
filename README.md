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
