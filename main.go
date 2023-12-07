package main

import (
    "encoding/json"
    "context"
    "log"
    "fmt"
    "io/ioutil"
    "net/http"
    "ilsan/ent"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    client, err := ent.Open("mysql", "root:root@@tcp(172.17.0.1:3306)/ilsan?parseTime=True")
    if err != nil {
        log.Fatalf("failed opening connection to mysql: %v", err)
    }
    defer client.Close()
    // Run the auto migration tool.
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
    userRepository := UserRepository{Client: client.User}

    http.HandleFunc("/login", func(res http.ResponseWriter, req *http.Request) {
        if req.Method == http.MethodPost {
            body, err := ioutil.ReadAll(req.Body)
            if err != nil {
                http.Error(res, "Failed to read request body", http.StatusInternalServerError)
                return
            }
            var data map[string]interface{}
            if err := json.Unmarshal(body, &data); err != nil {
                http.Error(res, "Failed to parse JSON", http.StatusBadRequest)
                return
            }
            user := userRepository.FindByName(data["name"].(string))
            if user.Password != data["password"].(string) {
                res.Header().Set("Content-Type", "application/json")
                res.WriteHeader(http.StatusOK)
                json.NewEncoder(res).Encode(map[string]interface{}{
                    "isSuccess":  false,
                    "code": 500,
                    "messages": fmt.Sprintf("이름과 비밀번호가 올바르지 않습니다."),
                })
                return 
            }
            res.Header().Set("Content-Type", "application/json")
            res.WriteHeader(http.StatusOK)
            json.NewEncoder(res).Encode(map[string]interface{}{
               "isSuccess":  true,
               "code": 200,
               "messages": fmt.Sprintf("로그인에 성공하였습니다."),
               "userId": user.ID,
            })
        } else {
            res.WriteHeader(http.StatusMethodNotAllowed)
            return
        }
    
    })
    // http.Handle("/store/{storeId}/menu-review", menuReview)

    http.ListenAndServe(":8000", nil)
}
