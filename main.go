// englishvocabulary project main.go
package main  
import (
"database/sql"
"fmt"
"time"
_ "github.com/lib/pq"
)

const (
host      ="localhost"
port      = 5432
user 	  = "postgres"
password  = "v19951162020"
dbname 	  = "englishwords"
)

type Words struct {
	engver string
	rusver string
	number int
}

func addWord(){
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
fmt.Println("Enter first english and then russian varaint.")
engWord := ""
rusWord := ""
fmt.Scanf("%s\n", &engWord)
fmt.Scanf("%s\n", &rusWord)
sqlStatement := `
INSERT INTO words (engversion, rusversion, number)
VALUES ($1, $2, 0)`
_, err = db.Exec(sqlStatement, engWord, rusWord)
if err != nil {
  panic(err)
}
return
}

func delWord() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  fmt.Println("Enter english or russian word.")
Word := ""
fmt.Scanf("%s\n", &Word)
sqlStatement := `
DELETE FROM words WHERE engversion = $1 OR rusversion = $1`
_, err = db.Exec(sqlStatement, Word)
if err != nil {
  panic(err)
}
return
}

func selWord() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  fmt.Println("Enter english or russian word.\n")
Word := ""
fmt.Scanf("%s\n", &Word)
sqlStatement := `
SELECT engversion, rusversion FROM words WHERE engversion = $1 OR rusversion = $1`
eng := ""
rus := ""
row := db.QueryRow(sqlStatement, Word)
switch err := row.Scan(&eng, &rus); err {
case sql.ErrNoRows:
  fmt.Println("There is no such word in the dictionary.\n")
case nil:
  if Word != eng {
  	fmt.Printf("%s\n", eng)
  	return
  }
  fmt.Printf("%s\n", rus)
default:
  panic(err)
}

return
}

func selAllWords(){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

    rows, err := db.Query("SELECT * FROM words")
    if err != nil {
        panic(err)
    }
    id := 0
   
    wos := make([]*Words, 0)
    for rows.Next() {
        wk := new(Words)
        err := rows.Scan(&id, &wk.engver, &wk.rusver, &wk.number)
        if err != nil {
            panic(err)
        }
        wos = append(wos, wk)
    }
    if err = rows.Err(); err != nil {
        panic(err)
    }

    for _, wk := range wos {
        fmt.Printf("%s, %s, %v\n", wk.engver, wk.rusver, wk.number)
    }
}

func exersizeeng() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  words := make(map[string]string)
  var key string
  var value string
  var number int
  rows, err := db.Query("SELECT * FROM words")
    if err != nil {
        panic(err)
    }
    id := 0
    for rows.Next() {
        err := rows.Scan(&id, &key, &value, &number)
        if err != nil {
            panic(err)
        }
        words[key]=value
    }
    if err = rows.Err(); err != nil {
        panic(err)
    }
    count := 0
    win := 0
	var startTime = time.Now()
	num := 0
    for eng, rus := range words {
    	count++
        fmt.Printf("%s ", eng)
        fmt.Scanf("%s\n", &key)
        if key == rus {
        	
        	win++
			sqlStatement := `
        SELECT number FROM words WHERE engversion = $1 `
        row := db.QueryRow(sqlStatement, eng)
        err := row.Scan(&num)
        if err != nil {
            panic(err)
        }		
        num++
	    sqlStatement = `
        UPDATE words SET number = $1 WHERE engversion = $2 `
	    _, err = db.Exec(sqlStatement, num, eng)
        if err != nil {
        panic(err)
        }
        continue
        }
		sqlStatement := `
        SELECT number FROM words WHERE engversion = $1 `
        row := db.QueryRow(sqlStatement, eng)
        err := row.Scan(&num)
        if err != nil {
            panic(err)
        }		
        num--
	    sqlStatement = `
        UPDATE words SET number = $1 WHERE engversion = $2 `
	    _, err = db.Exec(sqlStatement, num, eng)
        if err != nil {
        panic(err)
        }
	fmt.Printf("Incorrect, correct version %s.\n", rus)
	}
	var duration = time.Since(startTime)
    fmt.Printf("%v correct answers from %v, elapsed time %v.\n", win, count, duration)
    
}

func exersizerus() {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
  words := make(map[string]string)
  var key string
  var value string
  var number int
  rows, err := db.Query("SELECT * FROM words")
    if err != nil {
        panic(err)
    }
    id := 0
    for rows.Next() {
        err := rows.Scan(&id, &key, &value, &number)
        if err != nil {
            panic(err)
        }
        words[key]=value
    }
    if err = rows.Err(); err != nil {
        panic(err)
    }
    count := 0
    win := 0
	num := 0
	var startTime = time.Now()
    for eng, rus := range words {
    	count++
        fmt.Printf("%s ", rus)
        fmt.Scanf("%s\n", &key)
        if key == eng {
        	win++
			sqlStatement := `
        SELECT number FROM words WHERE rusversion = $1 `
        row := db.QueryRow(sqlStatement, rus)
        err := row.Scan(&num)
        if err != nil {
            panic(err)
        }		
        num++
	    sqlStatement = `
        UPDATE words SET number = $1 WHERE rusversion = $2 `
	    _, err = db.Exec(sqlStatement, num, rus)
        if err != nil {
        panic(err)
        }
        continue
        }
		sqlStatement := `
        SELECT number FROM words WHERE rusversion = $1 `
        row := db.QueryRow(sqlStatement, rus)
        err := row.Scan(&num)
        if err != nil {
            panic(err)
        }		
        num--
	    sqlStatement = `
        UPDATE words SET number = $1 WHERE rusversion = $2 `
	    _, err = db.Exec(sqlStatement, num, rus)
        if err != nil {
        panic(err)
        }
        fmt.Printf("Incorrect, correct version %s.\n", eng)
    }
	var duration = time.Since(startTime)
	
    fmt.Printf("%v correct answers from %v, elapsed time %v.\n", win, count, duration)
    
}

func deleteLearnedWords () {
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable",
host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()
_, err = db.Query("DELETE FROM words WHERE number > 3")
    if err != nil {
        panic(err)
    }
}


func menu() {
for {
	fmt.Println("\nWelcome to programm \"English with Vladosik\".\n To add a new word, press 1.\n To find a word in the database, press 2.\n To display the entire dictionary, press 3.\n To begin the exercise of translating from English to Russian, press 4.\n To begin the exercise of translating from Russian to English, press 5.\n To delete a word, press 6.\n To delete all learned words, press 7.\n To exit the program, press 8.")
answer := ""
fmt.Scanf("%s\n", &answer)
switch answer {
case "1":
	addWord()
case "2":
	selWord()
case "3":
    selAllWords()
case "4":
    exersizeeng()
case "5":
    exersizerus()
case "6":
    delWord()
case "7":
    deleteLearnedWords()
case "8":
    return
default:
    fmt.Println("Enter the correct number.\n")	
}
}
}

func main() {
menu()

}
