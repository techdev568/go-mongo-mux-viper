# go-mongo-mux-viper


## How to run?
First, start a MongoDB instance using docker:

docker run --name mongodb -d -p 27017:27017 mongo

### Get User
This endpoint retrieves a user given the email.  
Send a `GET` request to `/users/:email`:
```sh
curl -X GET 'http://127.0.0.1:9080/users/bob@gmail.com'
```

### Create User
This endpoint inserts a document in the `users` collection of the `users` database.  
Send a `POST` request to `/users`:
```sh
curl -X POST 'http://127.0.0.1:9080/users' -H "Content-Type: application/json" -d '{"name": "Bob", "email": "bob@gmail.com", "password": "ilovealice"}'
```

### Update User
This endpoint updates the provided fields within the specified document filtered by email.  
Send a `PUT` request to `/users/:email`:
```sh
curl -X PUT 'http://127.0.0.1:9080/users/bob@gmail.com' -H "Content-Type: application/json" -d '{"password": "loveyoualice"}'
```

### Delete User
This endpoint deletes the user from database given the email.  
Send a `DELETE` request to `/users/:email`:
```sh
curl -X DELETE 'http://127.0.0.1:9080/users/bob@gmail.com'
```