package helper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func BcryptPassword(passwordSalt string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordSalt), 8)
	return string(newPassword)
}

func ShowMessage(httpStatus int, message string, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, message)
}

func ShowDataWithTypeJson(data interface{}, w http.ResponseWriter, code int) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ShowDataWithTypeXml(data interface{}, w http.ResponseWriter, code int) {
	w.Header().Add("Content-type", "application/xml")
	xml.NewEncoder(w).Encode(data)
}

func ClearDoubleCode(str string) string {
	regex, _ := regexp.Compile(`[\"]+`)
	str = regex.ReplaceAllStringFunc(str, func(s string) string {
		if s == "\"" {
			return ""
		}
		return s
	})
	return str
}

func TruncateTable(db *gorm.DB, table []string) {
	for _, t := range table {
		db.Exec(fmt.Sprintf("TRUNCATE TABLE %s restart identity cascade", t))
	}
	fmt.Println("====== TRUNCATE SUCCESS =======")
}
