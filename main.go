package main

func main() {
	app := App{}
	app.Initialise(BDuser, DBpasswd, DBhost, DBname)
	app.Run("localhost:10000")

}
