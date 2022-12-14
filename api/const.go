package api

import (
	"bitty/model"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis/v8"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

var engine *xorm.Engine
var err error
var node *snowflake.Node
var ctx = context.Background()
var rdb *redis.Client

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	node, err = snowflake.NewNode(1)

	engine, err = xorm.NewEngine("mysql", os.Getenv("db"))
	engine.ShowSQL(true)

	//初始化redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err = rdb.Set(ctx, "key", "redis=====================", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	//初始化数据库``
	engine.SetTableMapper(names.SnakeMapper{})
	engine.SetColumnMapper(names.SnakeMapper{})

	engine.Sync2(new(model.User))
	engine.Sync2(new(model.Contact))
	engine.Sync2(new(model.UserToken))
	engine.Sync2(new(model.Msg))
	engine.Sync2(new(model.Session))
	engine.Sync2(new(model.SessionMember))
	engine.Sync2(new(model.Chat))
	engine.Sync2(new(model.Group))
	engine.Sync2(new(model.GroupMember))

	var rows, _ = engine.Query("select version() `version`")
	fmt.Printf("\n=========================================================\n\n")

	fmt.Printf("DBVersion: %s", rows[0]["version"])

	fmt.Printf("\n=========================================================\n\n")
}
