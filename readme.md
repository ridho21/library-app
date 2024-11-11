# Manajemen Perpustakaan

API untuk manajemen aplikasi perpustakaan, terdiri dari fitur
- Register
- Login
- Insert Book
- Get All Book
- Update Book
- Delete Book
- Borrow Book (not implemented yet)
- Return Book (not implemented yet)
- Calculate Penalty (not implemented yet)

## API Documentation

### Register

Method : POST <br />
Endpoint : /auth/register

Body Requests :
```json
{
    "full_name" : "Ridho Mahendra",
    "email" : "ridho123@gmail.com",
    "password" : "ridho123",
    "phone_number": "081372919506",
    "address": "Riau",
    "role" : "ADMIN"
}
```
Response : 
```json
{
  "status": {
    "code": 201,
    "description": "Success Register new User"
  },
  "data": {
    "id": "815b8717-bce0-45f4-8854-57b9002f0b95",
    "full_name": "Ridho Mahendra",
    "email": "ridho123@gmail.com",
    "phone_number": "081372919506",
    "address": "Riau",
    "role": "ADMIN",
    "created_at": "2024-11-05T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z"
  }
}
```

### Login

Method : POST <br />
Endpoint : /auth/login

Body Requests :
```json
{
  "email" : "ridho123@gmail.com",
  "password": "ridho123"
}
```
Response : 
```json
{
  "status": {
    "code": 201,
    "description": "Success Login as User"
  },
  "data": {
    "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzA4MDA0MzksImlzcyI6InJtIiwidXNlcklkIjoiODE1Yjg3MTctYmNlMC00NWY0LTg4NTQtNTdiOTAwMmYwYjk1IiwiZW1haWwiOiJyaWRobzEyM0BnbWFpbC5jb20iLCJyb2xlIjoiQURNSU4ifQ.cVr-tBieGZnD8fqgZAk4JYHQQSQ4983a6H8sPIFwb-4",
    "userId": "815b8717-bce0-45f4-8854-57b9002f0b95"
  }
}
```

### Insert Book
Put Token from Login to Auth Bearer <br />

Method : POST <br />
Endpoint : /book

Body Requests :
```json
{
  "title": "Kambing Jantan",
  "publication_year": "2024-11-04T00:00:00Z",
  "stock": 2,
  "total_pages": 100,
  "publisher": {
    "id": "81e02bc1-0095-4cb8-8963-90be589e2fd6"
  },
  "author": {
    "id": "af694ba4-22b6-4dfa-a269-7c5d88d816c1"
  },
  "category": {
    "id": "87ba173f-b9d8-4f54-b19f-1a70ceefc646"
  }
}
```
Response : 
```json
{
  "data": {
    "id": "23c43f4d-8470-4f81-a5c0-96af3545d5fb",
    "title": "Kambing Jantan",
    "publication_year": "2024-11-04T00:00:00Z",
    "stock": 2,
    "total_pages": 100,
    "created_at": "2024-11-05T00:00:00Z",
    "updated_at": "0001-01-01T00:00:00Z",
    "publisher": {
      "id": "81e02bc1-0095-4cb8-8963-90be589e2fd6",
      "publisher_name": "",
      "email": "",
      "phone_number": ""
    },
    "author": {
      "id": "af694ba4-22b6-4dfa-a269-7c5d88d816c1",
      "author_name": "",
      "email": "",
      "phone_number": ""
    },
    "category": {
      "id": "87ba173f-b9d8-4f54-b19f-1a70ceefc646",
      "category_name": ""
    }
  },
  "message": "Success Insert New Book"
}
```

### Get All Book
Put Token from Login to Auth Bearer <br />

Method : GET <br />
Endpoint : /book

Body Requests :
```json

```
Response : 
```json
{
  "status": {
    "code": 200,
    "description": "Success"
  },
  "data": [
    [
      {
        "id": "23c43f4d-8470-4f81-a5c0-96af3545d5fb",
        "title": "Kambing Jantan",
        "publication_year": "2024-11-04T00:00:00Z",
        "stock": 2,
        "total_pages": 100,
        "created_at": "2024-11-05T00:00:00Z",
        "updated_at": "0001-01-01T00:00:00Z",
        "publisher": {
          "id": "81e02bc1-0095-4cb8-8963-90be589e2fd6",
          "publisher_name": "Jendela Dunia",
          "email": "jd@gmail.com",
          "phone_number": "0899999999"
        },
        "author": {
          "id": "af694ba4-22b6-4dfa-a269-7c5d88d816c1",
          "author_name": "Raditya Dika",
          "email": "rd@gmail.com",
          "phone_number": "088888888"
        },
        "category": {
          "id": "87ba173f-b9d8-4f54-b19f-1a70ceefc646",
          "category_name": "Novel"
        }
      }
    ]
  ]
}
```

### Update Book
Put Token from Login to Auth Bearer <br />

Method : PUT <br />
Endpoint : /book

Body Requests :
```json
{
  "id": "23c43f4d-8470-4f81-a5c0-96af3545d5fb",
  "title": "Kambing Betina",
  "publication_year": "2024-11-04T00:00:00Z",
  "stock": 10,
  "total_pages": 152
}
```
Response : 
```json
{
  "data": "",
  "message": "Success Update Book"
}
```

### Delete Book
Put Token from Login to Auth Bearer <br />

Method : DELETE <br />
Endpoint : /book/:id

Path Requests :
```json
/23c43f4d-8470-4f81-a5c0-96af3545d5fb
```
Response : 
```json
{
  "data": "",
  "message": "Success Delete Book"
}
```