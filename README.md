# Todo_APP_GO
This is a to do app  with Go 


# Use docker for the database
1.Run the docker compose 
```shell
docker-compose up -d 
```
2.Connect to the docker container

```shell
docker exec -it dbTodo bash 
```

3.Connect to Postgresql in the terminal 

```shell
psql -U caleb -d test 
```

#database 

1 . create the table 
```shell
CREATE TABLE tasks(
 id SERIAL PRIMARY KEY ,
 tittle  VARCHAR(50) NOT NULL UNIQUE,
 text    VARCHAR(255) NOT NULL,
 date   DATE
);

```

2. insert data example to the data 
```shell 
INSERT INTO tasks (tittle, text, date)
VALUES ('homework','This is easy homework', '2023-05-22');
```
3. OK now you can run 

```shell 
go run .
```
Or compile the app
```shell
go build .
```
now run 
```shell 
./todoApp 
```


