package sqlop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	defaultResourceTable = "resource"
)

type UserResource struct {
	Name    string
	CpusUse int
	MemUse  int
}

func (ur *UserResource) Insert(db *sql.DB) error {
	qstr := "INSERT INTO " + defaultResourceTable + " (user_id, cpus_use, mem_use) (SELECT user_id, ?, ? from userpasswd WHERE name=?)"
	_, err := db.Exec(qstr, ur.CpusUse, ur.MemUse, ur.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ur *UserResource) Delete(db *sql.DB) error {
	qstr := "DELETE FROM " + defaultResourceTable + " WHERE user_id=(SELECT user_id from userpasswd WHERE name=?"
	_, err := db.Exec(qstr, ur.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ur *UserResource) Update(db *sql.DB) error {
	qstr := "UPDATE " + defaultResourceTable + " set cpus_use=cpus_use+? , mem_use=mem_use+?" +
		" WHERE user_id=(SELECT user_id FROM userpasswd WHERE name=?)"

	_, err := db.Exec(qstr, ur.CpusUse, ur.MemUse, ur.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (ur *UserResource) GetRS(db *sql.DB) error {
	qstr := "SELECT cpus_use, mem_use FROM " + defaultResourceTable + " WHERE user_id=(SELECT user_id FROM userpasswd WHERE name=?)"
	row := db.QueryRow(qstr, ur.Name)
	err := row.Scan(&ur.CpusUse, &ur.MemUse)
	if err != nil {
		log.Println("resource GetRS error :", err)
		return err
	}
	return nil
}

/**
 * real connect mysql
 * @param {[type]} mydb *MysqlCon) (*sql.DB, error [description]
 */
func (ur UserResource) Connect(mydb *MysqlCon) (*sql.DB, error) {
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
