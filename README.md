![chat](https://user-images.githubusercontent.com/36485235/185006091-4fb542c2-a61a-4854-85c2-cf883d15c52b.png)

# Chat Room
Real time multi user chatting app with stored chat history.

## Technologies
- Languages: Go, JavaScript, HTML, CSS
- Backend: 
  - Websocket and REST API built with Socket.IO, Gorilla, and MongoDB
  - Deployed with Docker on Azure Cloud Kubernetes Cluster
- Frontend: 
  - Built with Vue.js and Bootstrap
  
  
## Functionality
- The Chat API provides two endpoints to receive and delete the chat history from the database in MongoDB.
- API CONNECTION:
 1. GET: BASE_URL/API/MESSAGES
    - RESPONSE: ID, USER, TEXT, DATE
 2. DELETE: BASE_URL/API/DELETE
 
 - Users can connect to the chat room and send or receive messages in real time
 - New users are greeted with a welcoming message. Messages are displayed with the user, text, and date 
 - The user can also close the chat for all connected user by sending a disconnecting message that disables further messages
 - Users can see messages from different users in real time
 - The chat history is fetched or deleted via an own REST API connected to MongoDB
