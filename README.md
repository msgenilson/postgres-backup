# Backup PostgreSQL

Programa escrito em Go para realizar backups de vários ou todos os bancos de dados

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
# Linux
PATH_DB_BK="/home/username/postgres-backup/database"
# Windows
PATH_DB_BK="C:\\Users\\username\\postgres-backup\\database"
```
Caminho do pg_dumb da instalação local do PostgreSQL
```env
# Linux
PATH_PG_DUMP="/usr/bin/pg_dump"
# Windows
PATH_PG_DUMP="C:\\Program Files\\PostgreSQL\\15\\bin\\pg_dump.exe"
```
Limite de bancos que deverá buscar para backup (apenas para testes)
```env
LIMIT_DATABASE=6
```

## Execução
Script bash no Windows para execução de tarefas agendadas. Isso após realizar o comando build no seu programa Go.
```bash
# start.bat
@echo on
title Backup
cd C:\Users\username\postgres-backup\
start /min backup.exe
EXIT
```
