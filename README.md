# Backend assignment MobileApi spec
Steps to run the application
1. Install golang and set $GOPATH .
2. git clone this project in the $GOPATH repository
3. navigate to the folder mobileService
4. run the following commands-

   $ go build ./cmd/mobile-server/
   
   $ ./mobile-server --port=8000
   
   this will start the application on localhost 8000
   
# Dependencies
Swagger is used to generate code from the mobileapi.yaml.</br>
Mastermind/glide is used to manage dependencies.</br>
Swagger code is generated from the given yaml file using 'swagger generate server -f mobileapi.yaml'. </br>


## Gets article categories

EndPoint url : http://localhost:8000/v1/categories

Request type: HTTP GET

output- 
{
    "categories": [
        {
            "id": "Cat Id 1",
            "imageUrl": "Cat 1 Image url",
            "title": "Cat 1 Title"
        },
        {
            "id": "Cat Id 2",
            "imageUrl": "Cat 2 Image url",
            "title": "Cat 2 Title"
        }
    ]
}

## Gets all the articles for an category

Request type: HTTP GET

EndPoint url : http://localhost:8000/v1/categories/{categoryId}

output- 
{
    "articles": [
        {
            "id": "Art Id 1",
            "source": "Art 1 Source 1",
            "starred": false,
            "timestamp": "Art ts 1",
            "title": "Art title 1"
        },
        {
            "id": "Art Id 2",
            "source": "Art 2 Source 2",
            "starred": false,
            "timestamp": "Art ts 2",
            "title": "Art title 2"
        }
    ]
}

## A list of Article overview objects

Request type: HTTP GET

EndPoint url : http://localhost:8000/v1/articles/{articleId}

output- 
{
    "body": "Art 2 Body 2",
    "imageUrl": "Art 1 Image 1",
    "overview": {
        "id": "Art Id 1",
        "source": "Art 1 Source 1",
        "starred": false,
        "timestamp": "Art ts 1",
        "title": "Art title 1"
    }
}


## Updates article's starred property

Request type: HTTP PATCH

Endpoint url: http://localhost:8000/v1/articles/{articleId}

{
	"starred":true
}

Content-type: application/json

Input Params: article Id, starred value (true or false)

request url: http://localhost:8080/api/games/getRankHistory?name=Gardenscapes&market=fi

Output json: 

{
    "message": "Value Updated OK",
    "status": "Success"
}
