# Endpoints

### Get All Snippets
```
GET /s
```

returns all snippets in the database<br>
Example Response: 
```
[
    {
        "id": "String",
        "title": "String",
        "content": "String",
        "cookie": "String"
    },
    {
        "id": "SSawJG",
        "title": "Title",
        "content": "Content",
        "cookie": ""
    }
]
```

### Get Single Snippet by ID
```
GET /s/[id]
```
Parameters: id -> string<br>
returns a single snippet from the database with the matching id<br>
Example Response:
```
[
    {
        "id": "String",
        "title": "String",
        "content": "String",
        "cookie": "String"
    }
]
```

### Create Snippet
```
POST /create
```
Parameters: title -> string, content -> string, cookie -> string<br>
create a snippet with the provided info<br>

### Edit Snippet
```
POST /edit
```
Parameters: id -> string, title -> string, content -> string, cookie -> string<br>
edit an existing snippet which matches<br>