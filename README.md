# Information Security course at Saratov State University, 7th semester
## Tasks
1. Implement console programs `put-message` and `get-message`. The first one should embed the message into the container and create a stegocontainer. The second one restores the message from the stegocontainer exactly bit by bit.
2. OpenSSL (WIP)
3. Web application attacks (WIP)
4. Implement a microservice REST app with following features:
    * `/registrate?usr=username&passwd=password`: registers an user by his username and password and stores user data (username, salt and hashed password) in database (I use MongoDB as DMS). The response contains the code 201 if user is successfully registered, and 400 otherwise. 
    * `/auth?usr=username&passwd=password`: authentification and authorization. Sets access and refresh tokens using cookies if authentification is completed. Responses: 200 if user is succesfully logged in, and 403 otherwise.
    * `/refresh`: updates access and refresh tokens. Only a refresh token is required. The responses have the same codes as in `/auth`.