//go:build tools

package main

import (
	"flag"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	var outputPath string
	flag.StringVar(&outputPath, "output", "./internal/infra/postgres/model", "output path")
	flag.Parse()
	fmt.Printf("outputPath: %s\n", outputPath)

	g := gen.NewGenerator(gen.Config{
		OutPath: outputPath,
		Mode:    gen.WithDefaultQuery, // generate mode
	})

	// 以下の環境変数を直接読み込み
	// dbHost=localhost
	// DB_PORT=5432
	// DB_NAME=plusweb
	// DB_USER=***
	// DB_PASS=***
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	dbConnInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", dbHost, dbUser, dbPass, dbName, dbPort)
	fmt.Printf("dbConnInfo: %s\n", dbConnInfo)
	db, _ := gorm.Open(postgres.Open(dbConnInfo))
	g.UseDB(db) // reuse your gorm db
	g.GenerateAllTable()
	g.Execute()
}
