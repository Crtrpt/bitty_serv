package api



import (
	"fmt"
	"log"

	"os"


	"bitty/model"

	"github.com/bwmarrin/snowflake"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)


var engine *xorm.Engine
var err error
var node *snowflake.Node

func Init(){
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	node, err = snowflake.NewNode(1)

	engine, err = xorm.NewEngine("mysql", os.Getenv("db"))
	engine.ShowSQL(true)

	engine.SetTableMapper(names.SnakeMapper{})
	engine.SetColumnMapper(names.SnakeMapper{})

	engine.Sync2(new(model.User))
	engine.Sync2(new(model.Endpoint))
	engine.Sync2(new(model.UserToken))
	engine.Sync2(new(model.Msg))

	var rows, _ = engine.Query("select version() `version`")
	fmt.Printf("\n=========================================================\n\n")

	fmt.Printf("DBVersion: %s", rows[0]["version"])

	fmt.Printf("\n=========================================================\n\n")
}