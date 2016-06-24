package sqlop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
 * This file is for the table (name string, password string, cpus int, memory int)
 */

const (
	defaultUserTable = "userpasswd"
)

type UserInfo struct {
	Name     string
	Password string
	Cpus     int
	Mem      int
}

func (u UserInfo) Insert(db *sql.DB) error {
	qstr := "INSERT INTO " + defaultUserTable + " VALUE(NULL,?,?,?,?)"
	_, err := db.Exec(qstr, u.Name, u.Password, u.Cpus, u.Mem)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (u UserInfo) Delete(db *sql.DB) error {
	qstr := "DELETE FROM " + defaultUserTable + " WHERE name=?"
	log.Println(qstr)
	log.Println(u.Name)
	_, err := db.Exec(qstr, u.Name)
	if err != nil {
		log.Println("Exec error :", err)
		return err
	}

	log.Println("userpasswd.go delete operation is call success")

	return nil
}

/**
 * real connect mysql
 * @param {[type]} mydb *MysqlCon) (*sql.DB, error [description]
 */
func (u UserInfo) Connect(mydb *MysqlCon) (*sql.DB, error) {
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

func (u *UserInfo) Query(db *sql.DB) error {
	qstr := "SELECT name,password,cpus,memory FROM " + defaultUserTable + " WHERE name=?"
	row := db.QueryRow(qstr, u.Name)

	err := row.Scan(&u.Name, &u.Password, &u.Cpus, &u.Mem)
	if err != nil {
		log.Println("row.Scan error :", err)
		return err
	}

	return nil
}

func (u *UserInfo) QueryUsers(db *sql.DB) ([]string, error) {
	result := make([]string, 0)
	var tmpstr string

	qstr := "SELECT name FROM " + defaultUserTable
	rows, err := db.Query(qstr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&tmpstr)
		result = append(result, tmpstr)
	}

	return result, nil
}
