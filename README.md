# go-book-manage


# Book Manage

This project is create for learning syntax of Golang Programming, Fiber Web framework and GORM. 


## Installation

Install this project by : git clone https://github.com/khomkritz/go-book-manage.git

In project have 2 folder (backned, frontend)

## Backend

```bash
  open folder : cd backned
  run command : go run .
```


    
## API Reference

#### Sign Up

```https
  POST /signup
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required** |
| `password` | `string` | **Required** |

Statuscode : 201

Response : 
```
{
    "ID": int(),
    "CreatedAt": str(),
    "UpdatedAt": str(),
    "DeletedAt": str(),
    "id": int(),
    "username": str(),
    "password": str()
}
```

#### Login

```https
  POST /login
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `username` | `string` | **Required** |
| `password` | `string` | **Required** |

Statuscode : 200

Response : 
```
{
    "jwtToken" : str()
}
```
#### Get all books 

```https
  GET /api/books
```
| Header key        | Description                              |
| ----------------- | ---------------------------------------- |
| `Authorization`   | **jwtToken** The authorized user’s token. This is used to gain access to protected endpoint. |

Statuscode : 200

Response : 
```
[
    {
        "ID": int(),
        "CreatedAt": str(),
        "UpdatedAt": str(),
        "DeletedAt": str(),
        "title": str(),
        "author": str(),
        "genre": str(),
        "year": int(),
        "user_id": int()
    }
]
```


#### Get item by id

```http
  GET /api/books/:id
```
| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** |

| Header key        | Description                              |
| ----------------- | ---------------------------------------- |
| `Authorization`   | **jwtToken** The authorized user’s token. This is used to gain access to protected endpoint. |

Statuscode : 200

Response : 
```
{
  "ID": int(),
  "CreatedAt": str(),
  "UpdatedAt": str(),
  "DeletedAt": str(),
  "title": str(),
  "author": str(),
  "genre": str(),
  "year": int(),
  "user_id": int()
}
```

#### Create Books

```https
  POST /api/books
```

| Body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `title` | `string` | **Required** |
| `author` | `string` | **Required** |
| `genre` | `string` | **Required** |
| `year` | `int` | **Required** |

| Header key        | Description                              |
| ----------------- | ---------------------------------------- |
| `Authorization`   | **jwtToken** The authorized user’s token. This is used to gain access to protected endpoint. |

Statuscode : 201

Response : 
```
{
    "ID": int(),
    "CreatedAt": str(),
    "UpdatedAt": str(),
    "DeletedAt": str(),
    "id": int(),
    "username": str(),
    "password": str()
}
```

#### Update Books

```https
  POST /api/books/:id
```

| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** |


| Header key        | Description                              |
| ----------------- | ---------------------------------------- |
| `Authorization`   | **jwtToken** The authorized user’s token. This is used to gain access to protected endpoint. |

Statuscode : 200

Response : 
```
{
    "ID": int(),
    "CreatedAt": str(),
    "UpdatedAt": str(),
    "DeletedAt": str(),
    "id": int(),
    "username": str(),
    "password": str()
}
```

#### Delete Books

```https
  POST /api/books/:id
```

| Parameters | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `id` | `string` | **Required** |


| Header key        | Description                              |
| ----------------- | ---------------------------------------- |
| `Authorization`   | **jwtToken** The authorized user’s token. This is used to gain access to protected endpoint. |

Statuscode : 204


