package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// CreateTablee()
	r.HandleFunc("/", CreateItem).Methods("GET")
	// r.HandleFunc("/movie/1", function.ReadingItemid).Methods("GET")
	// r.HandleFunc("/movie", function.CreateItem).Methods("POST")
	// r.HandleFunc("/movie/2", function.UpdateItems).Methods("PUT")
	// r.HandleFunc("/movie/2", function.Softdelete).Methods("DELETE")
	log.Fatal(http.ListenAndServe("Localhost:5000", r))

}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	// w.Write([]byte("Item Created"))
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	item := map[string]interface{}{
		"Movieid":  2022,
		"Title":    "kgf2",
		"Hero":     "yash",
		"isactive": true,
	}

	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		log.Fatalf("Got error marshalling new movie item: %s", err)
	}

	tableName := "Movies"

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		log.Fatalf("Got error calling PutItem: %s", err)
	}

	// year := strconv.Itoa(item["Movieid"].(int))

	// fmt.Println("Successfully added '" + item["Title"].(string) + "' (" + year + ") to table " + tableName)
}

type Item struct {
	Movieid int
	Title   string
	Hero    string
}

func ReadingItem(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf"
	movieid := "2010"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
	})
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	// id, err := strconv.Atoi(movieid)

	// if id != 0 {
	// 	log.Fatalf("No item: %s", err)
	// }

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Id:  ", item.Movieid)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Hero:  ", item.Hero)

}

func ReadingItemid(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf3"
	movieid := "2010"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
	})
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	//id, err := strconv.Atoi(movieid)

	if Title == "" {
		log.Fatalf("No item: %s", err)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Id:  ", item.Movieid)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Hero:  ", item.Hero)

}

func Softdelete(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf2"
	Movieid := "2022"
	//Hero := "yash"
	Isactive := false

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				BOOL: aws.Bool(Isactive),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(Movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Isactive = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

}

func UpdateItems(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf2"
	Movieid := "2010"
	Hero := "yash"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(Hero),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(Movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Hero = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	fmt.Println("Successfully updated '" + Title + "' (" + Movieid + ") rating to " + Title)

}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	//delete an item in database

	tableName := "Movies"
	movieName := "xyz"
	movieid := "2001"

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(movieName),
			},
		},
		TableName: aws.String(tableName),
	}
	_, err := svc.DeleteItem(input)
	if err != nil {
		log.Fatalf("Got error calling DeleteItem:%s", err)
	}
	fmt.Println("Deleted'" + movieName + "'(" + movieid + ")from table" + tableName)
}