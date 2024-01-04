package router

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	_ "github.com/go-sql-driver/mysql"
)

func SetUp() {

     var dbName string = "database-1"
     var dbUser string = "admin"
     var dbHost string = "database-1.cluster-c7iiq4siobuj.ap-southeast-2.rds.amazonaws.com"
     var dbPort int = 3306
     var dbEndpoint string = fmt.Sprintf("%s:%d", dbHost, dbPort)
     var region string = "ap-southeast-2"

    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
    	panic("configuration error: " + err.Error())
    }

    authenticationToken, err := auth.BuildAuthToken(
    	context.TODO(), dbEndpoint, region, dbUser, cfg.Credentials)
    if err != nil {
	    panic("failed to create authentication token: " + err.Error())
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
        dbUser, authenticationToken, dbEndpoint, dbName,
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }
}