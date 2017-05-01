# lsjd

## Data Structures

### Review (object)

+ uuid: (string, required)
+ author: (string, required)
+ signature: (string, required)
+ body: (string, required)

### Article (object)

+ uuid: (string, required)
+ author: (string, required)
+ signature: (string, required)
+ version: (int, required)
+ tags: (array[string])
+ body: (string, required)
+ reviews: (array[reviews])

## Articles

# GET /articles

Returns a list of 10 articles uploaded to the server, the query string should
be used to select order and filters

# GET /articles/<uuid:string>

Return the article identified by its uuid string, if available

# POST /articles

+ Request:
	+Body: Article

Create a new article, assign it a uuid and return that uuid.

# PUT /articles/<uuid:string>

Update the article

# POST /article/<uuid:string>/comment

+ Request:
	+ Body: Comment

Add a comment for the article.

# POST /article/<uuid:string>/reviewed

+ Request
	+ header: Reviewed <string signature of the article>

Marks the last version of the article as reviewed by the user issuing
the request.

# GET /tags

Return a list of tags, the tags should be a tree, e.g.: "math",
"math.real_analysis", "math.discrete"

