# postgres-backup
Backup PostgreSQL

### Variáveis
Configurações definidas para o banco de dados
```env
PGI_HOST=0.0.0.0
PGI_PORT=5432
PGI_USER=user
PGI_PASSWORD=***********
PGI_DBNAME=db
```
Caminho destino dos arquivos de backup
```env
PATH_DB_BK="/home/username/postgres-backup/database"
```
Caminho do pg_dumb da instalação local do PostgreSQL
```env
PATH_PG_DUMP="/usr/bin/pg_dump"
```
Limite de bancos que deverá buscar para backup (apenas para testes)
```env
LIMIT_DATABASE=6
```
