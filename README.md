# Student Registry
For the Programming club secretary task.

The backend can perform various functions such as adding an entry to the databse, deleting an entry, editing an entry and finding an entry.

The database that is used in the current repository is MongoDB Atlas an online free databasing platform. I created a client of the online MongoDB Atlas server. A major point to pay attention to is that the context variable that is used to update each entry in the database has a fixed life time. 

`ctx, _ = context.WithTimeout(context.Background(), 10*time.Minute)`

Right now I have set it to 10 minutes but it can be extended as per need so that context does not timeout while we are still accessing our database. The client is then used to establish the connection which is disconnected after all functions are executed.

`e := echo.New()`

e is a Multiplexer used to create an HTTP server that would listen, match and handle `GET` and `POST` requests in go. `GET` refers to the request to get data from a resources while `POST` requests are uesd to send data to a particular server.

While using the database I created a collection that contains all the details as documents which are nothing but key and value stored in binary format `bson` in MongoDB. The various functions of `mongo.Collection` like `InsertOne`,`DeleteOne`,`UpdateOne` and `Find` that help to perform the required tasks on the database.

In order to get input from the user without the use of front end I used the `POSTMAN` app that hepls us to send `POST` requests in the form of form values which are then interpreted by `c.FormValue()`.

The server used is `localhost:4000` and there are various URL used by the user to do various tasks `localhost:4000\add` to add the student entry, `localhost:4000\delete` to delete an entry, `localhost:4000\edit` to edit the entry and `localhost:4000/find` to find the entry while any other URL works as an instruction to the user.
