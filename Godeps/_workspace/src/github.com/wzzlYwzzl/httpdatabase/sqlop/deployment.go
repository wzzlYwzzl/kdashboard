package sqlop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	defaultDeployTable = "deployment"
)

type Deploy struct {
	Name      string
	AppName   string
	CpusUse   int
	MemoryUse int
}

func (d *Deploy) Insert(db *sql.DB) error {
	qstr := "INSERT INTO " + defaultDeployTable + " (user_id, app_name, cpus_use, mem_use) (SELECT user_id, ?, ?, ? from userpasswd WHERE name=?)"
	_, err := db.Exec(qstr, d.AppName, d.CpusUse, d.MemoryUse, d.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (d *Deploy) Delete(db *sql.DB) error {
	qstr := "DELETE FROM " + defaultDeployTable + " WHERE app_name=?"
	_, err := db.Exec(qstr, d.AppName)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

/**
 * real connect mysql
 * @param {[type]} mydb *MysqlCon) (*sql.DB, error [description]
 */
func (d *Deploy) Connect(mydb *MysqlCon) (*sql.DB, error) {
	db, err := sql.Open("mysql", mydb.Name+":"+mydb.Password+"@tcp("+mydb.Host+")/"+defaultDB+"?charset=utf8")
	if err != nil {
		log.Println("open mysql with error: ", err)
		return db, err
	}
	err = db.Ping()
	if err != nil {
		log.Println("ping mysql with error:", err)
		return db, err
	}
	return db, nil
}
