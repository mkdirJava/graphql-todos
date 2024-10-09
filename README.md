
# playground for graphql gql

Requiresites
golang version 1.20.6

---

Steps

run 

    <project root> go run server.go

go to 

    localhost:8080

Supported Queries

query todosWithUser{
  todos{
    id
    text
    done
    user{
      id
      name
    }
  }
}

query todoWithoutUser{
  todos{
    id
    text
    done
  }
}

mutation assosaiteUserWithTodo{
  assignTodo(todoId:"0711e7c1-d3dd-4872-b808-2fbcbb0a585d",userId:"668cae22-0a58-4ade-b18a-4e24685ae6ef") {
    id
    text
    user{
      id
      name
    }
  }
}

mutation createUser {
  createUser(userName:"John"){
    id
    name
  }
}


mutation createTodo {
createTodo(input:{
  text: "Go Shopping"
}){
id
  text
  done
  user{
id
  name}

}
}

mutation createTodoWithUser {
createTodo(input:{
  text: "Go Cook Somthing"
  userId:"dad00df0-7adb-4aa4-a9b7-0b618821befc"
}){
id
  text
  done
  user{
id
  name}

}
}

subscription subscribeToData{
  currentTodos{
    id
    text
    done
    user{
      id
      name
    }
  }
}


