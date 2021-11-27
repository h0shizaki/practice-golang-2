# Practice GO lang with MySQL

How to setup
1. Import database name test.sql  
2. Setting databse connection in main.go  
```go
    sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>")
```  
3. Run main.go by  
``` Powershell
    go run main.go
```  

