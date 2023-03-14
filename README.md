# MoneyTrans
Simple GoLang project for transferring money between accounts through RESTful API.

This project provides a simple RESTful API for transferring money between accounts. It allows you to transfer money between accounts, view a list of all accounts, and view the details of a specific account.

## Installation

1. Clone the repository:
```
git clone git@github.com:obadmatar/moneytrans.git
```

2. Install dependencies:
```
go mod download
```

3. Run the server:
```
go run main.go
```

4. Access the API at
```
http://localhost:8080/api
```


## Available Endpoints

### 1. `GET` Accounts
Returns a list of accounts
```
https://localhost:8080/api/accounts
```

- Response
```json
[
  {
      "id": "17f904c1-806f-4252-9103-74e7a5d3e340",
      "name": "Fivespan",
      "balance": "946.15"
  }
]
```



### 2. `GET` Account By ID
Returns the details of a specific account with the given ID.
```
https://localhost:8080/api/accounts/{id}
```

- Response
```json
{
    "id": "17f904c1-806f-4252-9103-74e7a5d3e340",
    "name": "Fivespan",
    "balance": "946.15"
}
```



### 3. `POST` Transfere
Transfers an amount of money from one account to another.
```
https://localhost:8080/api/transfer
```

- Request
```json
{
	"senderId": "3d253e29-8785-464f-8fa0-9e4b57699db9",
	"receiverId": "17f904c1-806f-4252-9103-74e7a5d3e340",
	"amount": "80.00"
}
```



- Response
```json
{
  "message": "Transfer Done"
}
```

## Dependencies

- This project uses the following dependencies:
`github.com/shopspring/decimal v1.3.1`
`github.com/julienschmidt/httprouter v1.3.0`


## Repository
- This project uses an in-memory repository to store the account information. When the server is started, a few sample accounts are added to the repository from the json file `data/accounts.json`.








