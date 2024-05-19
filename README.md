
# playground for graphql gql

Requiresites
golang version 1.20.6

---

Steps

run 

    <project root> go run server.go

go to 

    localhost:8080

You should see a client with submission to the running graphql

these are supported, they dont do anything interesting thouh

    # subscription subs{
    #   currentTodos{
    #     done
    #   }
    # }

    # query q {
    #   todos{
    #     id
    #   }
    # }

    # mutation m {
    #   createTodo(input:{
    #     userId:"123"
    #     text:"file"
    #   }){
    #     id
    #   }
    # }