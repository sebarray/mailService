package met

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"servicemail/mail"
)

type reclut struct {
	Msg   string `json:Msg`
	Reclu string `json:Reclu`
}

func SendMail(w http.ResponseWriter, r *http.Request) {
	var newReclut reclut
	reqbody, err := ioutil.ReadAll(r.Body) //recibo datos que envia el cliente
	if err != nil {
		fmt.Fprintln(w, "error al recibir datos del cliente")
	}

	json.Unmarshal(reqbody, &newReclut) // asigno los valores  recibido al struct
	fmt.Print(newReclut)
	mail.Mail(newReclut.Msg, newReclut.Reclu)

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newReclut)

}
