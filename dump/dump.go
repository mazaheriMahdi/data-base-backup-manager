package dump

import (
	"fmt"
	pg "github.com/habx/pg-commands"
	"log"
	"os"
)

func GetDump(dbConfig DBConfiguration) (*os.File, error) {
	log.Println("Running pg dump")
	dump, _ := pg.NewDump(&pg.Postgres{
		Host:     dbConfig.Host,
		Port:     dbConfig.Port,
		DB:       dbConfig.DB,
		Username: dbConfig.Username,
		Password: dbConfig.Password,
	})
	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: false})
	if dumpExec.Error != nil {
		log.Println("Dump Error:", dumpExec.Error)
		return nil, dumpExec.Error.Err
	} else {
		fmt.Println("Dump success")
		fmt.Println(dumpExec.Output)
	}
	open, err := os.Open(dumpExec.File)
	if err != nil {
		return nil, err
	}
	return open, nil
}
