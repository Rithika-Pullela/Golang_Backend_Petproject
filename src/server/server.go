package main

import (
	s "beprj/src/Structs"
	dataBase "beprj/src/db"
	pb "beprj/src/proto"
	"beprj/src/serverfuncs"
	
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	
)


func main(){

	err := godotenv.Load("/home/rithika/Desktop/backend_petproject/src/.env")
	if err != nil {
		log.Fatal("Error loading .env file ", err)
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 7005))
	if err != nil {
		panic("failed to listen")
	}
	host := flag.String("host", os.Getenv("HOST_VALUE"), "PostgreSQL host")
	port := flag.String("port", os.Getenv("PORT_VALUE"), "PostgreSQL port")
	user := flag.String("user", os.Getenv("USER_VALUE"), "PostgreSQL user")
	password := flag.String("password", os.Getenv("PASSWORD_VALUE"), "PostgreSQL password")
	database := flag.String("database", os.Getenv("DATABASE_VALUE"), "PostgreSQL database name")
	flag.Parse()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s   sslmode=disable", *host, *port, *user, *password, *database)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Connection Successful")
	fmt.Println("server application here!")
	db.DropTable(&s.Assignment{},&s.Course{},&s.Faculty{},&s.Student{})
	db.AutoMigrate(&s.Student{},&s.Course{},&s.Faculty{},&s.Assignment{})
	if err := serverfuncs.InsertMockData(db); err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	 pb.RegisterUniversityFuncServer(grpcServer,&serverfuncs.UniversityServer{
		Db: &dataBase.DBClient{Db:db},
		DB:db,
	 })
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())
}

func AddMockData() {
	panic("unimplemented")
}