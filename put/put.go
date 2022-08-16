package main

type schema struct {
	url  string `dynamo:"url"`
	hash string `dynamo:"hash"`
}

func main() {
	sess := session.Must(session.NewSession())
	db := dynamo.New(sess)
	table := db.Table("test")

	// put item
	w := widget{UserID: 613, Time: time.Now(), Msg: "hello"}
	err := table.Put(w).Run()
}
