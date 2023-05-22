package util

import (
	"fmt"
	"time"
)

func GetHeader() string {
	txt := fmt.Sprintf("%-*s\t", 6, "QTD")
	txt += fmt.Sprintf("%-*s\t", 40, "DATABASE")
	txt += fmt.Sprintf("%-*s\t", 10, "DATA")
	txt += fmt.Sprintf("%-*s\t", 8, "INICIO")
	txt += fmt.Sprintf("%-*s\t", 8, "FIM")
	txt += fmt.Sprintf("%-*s", 8, "TEMPO")

	txt += fmt.Sprintf("\n%-*s\t", 6, "------")
	txt += fmt.Sprintf("%-*s\t", 40, "------------------------------")
	txt += fmt.Sprintf("%-*s\t", 10, "----------")
	txt += fmt.Sprintf("%-*s\t", 8, "--------")
	txt += fmt.Sprintf("%-*s\t", 8, "--------")
	txt += fmt.Sprintf("%-*s", 8, "--------")

	return txt
}

func PrintEnd(timeStart time.Time) string {
	timeEnd := time.Now()
	tempo := time.Time{}.Add(timeEnd.Sub(timeStart))

	txt := fmt.Sprintf("%-*s\t", 8, timeEnd.Format("15:04:05"))
	txt += fmt.Sprintf("%-*s", 8, tempo.Format("15:04:05"))

	fmt.Print(txt)

	return txt
}

func PrintStart(index int, qtd int, db string) (time.Time, string) {
	timeStart := time.Now()

	txt := fmt.Sprintf("\n%-*s\t", 6, fmt.Sprintf("%d/%d", index, qtd))
	txt += fmt.Sprintf("%-*s\t", 40, db)
	txt += fmt.Sprintf("%-*s\t", 10, timeStart.Format("02/01/2006"))
	txt += fmt.Sprintf("%-*s\t", 8, timeStart.Format("15:04:05"))

	fmt.Print(txt)

	return timeStart, txt
}

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
