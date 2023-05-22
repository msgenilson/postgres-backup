package main

import (
	"backup/postgres/config"
	"backup/postgres/util"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() error {
	var err error
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.PG_IMPORT_HOST,
			config.PG_IMPORT_PORT,
			config.PG_IMPORT_USER,
			config.PG_IMPORT_PASSWORD,
			config.PG_IMPORT_DBNAME,
		),
	)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func CloseConnect() error {
	return db.Close()
}

func dumpExecute(path string, name string) {
	cmd := exec.Command(
		config.PATH_PG_DUMP,
		"-h",
		config.PG_IMPORT_HOST,
		"-p",
		config.PG_IMPORT_PORT,
		"-U",
		config.PG_IMPORT_USER,
		"-F",
		"c",
		"-b",
		"-O",
		"-v",
		"-f",
		path+"/"+name+".backup",
		name,
	)

	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start cmd: %v", err)
		return
	}
	if err := cmd.Wait(); err != nil {
		log.Printf("\nCmd returned error: %v\n", err.Error())
		panic(err.Error())
	}
}

func execTask() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	theTime := time.Now()

	dataHora := theTime.Format("2006-01-02-150405")

	var path string
	// if config.MODO_RESTORE == "1" {
	// 	path = config.PATH_DB_RESTORE + "/"
	// } else {
	path = config.PATH_DB_BK + "/" + dataHora
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		util.CheckErr(err)
	}
	// }

	fLog, err := os.Create(path + "/data.log")
	if err != nil {
		log.Fatal(err)
	}
	defer fLog.Close()

	_, err = fLog.WriteString(util.GetHeader())
	util.CheckErr(err)

	rows := getDatabaseList()

	index := 0

	for rows.Next() {
		var qtd int
		var datName string
		index++
		err := rows.Scan(&datName, &qtd)
		util.CheckErr(err)

		timeStart, textStart := util.PrintStart(index, qtd, datName)

		dumpExecute(path, datName)

		textEnd := util.PrintEnd(timeStart)

		_, err = fLog.WriteString(textStart + textEnd)
		util.CheckErr(err)
	}
}

func getDatabaseList() *sql.Rows {
	rows, err := db.Query(
		`SELECT
			datname,
			count(1) over() qtd
		FROM
			pg_database
		WHERE
			datname NOT IN ('template1', 'template0', 'postgres')
			AND datname NOT ILIKE '%_old'
		ORDER BY
			datname
		limit $1`,
		config.LIMIT_DATABASE,
	)

	util.CheckErr(err)

	return rows
}

func main() {
	fmt.Println("Init backup")
	fmt.Println()

	d := util.GetHeader()
	fmt.Print(d)
	os.Setenv("PGPASSWORD", config.PG_IMPORT_PASSWORD)
	execTask()
	fmt.Println()
}
