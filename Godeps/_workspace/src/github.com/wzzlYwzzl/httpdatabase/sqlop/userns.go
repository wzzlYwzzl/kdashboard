package sqlop

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

/**
 * This file is for the table (id int, name string, namespace string)
 */
const (
	defaultDB    = "httpdb"
	defaultTable = "userns"
)

type User struct {
	Name      string
	Namespace string
}

type MysqlCon struct {
	Host     string
	Name     string
	Password string
}

func (user *User) Insert(db *sql.DB) error {
	qstr := "INSERT INTO " + defaultTable + " (user_id, namespace) (SELECT user_id , ? from userpasswd WHERE name=?)"
	_, err := db.Exec(qstr, user.Namespace, user.Name)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (user *User) Delete(db *sql.DB) error {
	qstr := "DELETE FROM " + defaultTable + " WHERE namespace=?"
	stmt, err := db.Prepare(qstr)
	if err != nil {
		log.Println("delete with error: ", err)
		return err
	}
	_, err = stmt.Exec(user.Namespace)
	if err != nil {
		log.Println("Exec error: ", err)
		return err
	}

	return nil
}

/**
 * real connect mysql
 * @param {[type]} mydb *MysqlCon) (*sql.DB, error [description]
 */
func (user User) Connect(mydb *MysqlCon) (*sql.DB, error) {
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

/**
 * query operation for namespaces
 * @param {[type]} db   *sql.DB [description]
 * @param {[type]} user *User)  ([]string,    error [description]
 */
func (user *User) Query(db *sql.DB) ([]string, error) {
	result := make([]string, 0, 10)
	var tmpstr string

	qstr := "SELECT namespace FROM " + defaultTable + "," + defaultUserTable + " WHERE " +
		defaultTable + ".user_id=" + defaultUserTable + ".user_id and " + defaultUserTable +
		".name=" + "'" + user.Name + "'"

	rows, err := db.Query(qstr)
	if err != nil {
		log.Println("Query error with :", err)
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&tmpstr)
		result = append(result, tmpstr)
	}

	log.Printf("length of find namespace result is %d", len(result))
	return result, nil
}

func (user *User) QueryAll(db *sql.DB) ([]string, error) {
	result := make([]string, 0, 10)
	var tmpstr string

	qstr := "SELECT namespace FROM " + defaultTable
	log.Println(qstr)
	rows, err := db.Query(qstr)
	if err != nil {
		log.Println("Query error with :", err)
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&tmpstr)
		result = append(result, tmpstr)
		log.Println("find namespace", tmpstr)
	}

	log.Printf("length of find namespace result is %d", len(result))
	return result, nil
}
