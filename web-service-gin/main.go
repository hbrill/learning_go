package main

/*
To keep things simple for the tutorial, you’ll store data in memory. A more typical API would interact with a database.

Note that storing data in memory means that the set of albums will be lost each time you stop the server, then recreated when you start it.
*/
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default() // Initialize a Gin router using Default.
	// Use the GET function to associate the GET HTTP method and /albums path with a handler function.
	//Note that you’re passing the name of the getAlbums function.
	//This is different from passing the result of the function, which you would do by passing getAlbums()
	router.GET("/albums", getAlbums)

	// Associate the /albums/:id path with the getAlbumByID function.
	// In Gin, the colon preceding an item in the path signifies that the item is a path parameter.
	router.GET("/albums/:id", getAlbumByID)

	//Associate the POST method at the /albums path with the postAlbums function.
	router.POST("/albums", postAlbums)

	// Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8000")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client.
	// Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)            // Append the album struct initialized from the JSON to the albums slice
	c.IndentedJSON(http.StatusCreated, newAlbum) // Add a 201 status code to the response, along with JSON representing the album you added.
}

func getAlbumByID(c *gin.Context) {
	// Use Context.Param to retrieve the id path parameter from the URL.
	// When you map this handler to a path, you’ll include a placeholder for the parameter in the path.
	id := c.Param("id")

	// Loop over the album structs in the slice, looking for one whose ID field value matches the id parameter value.
	// If it’s found, you serialize that album struct to JSON and return it as a response with a 200 OK HTTP code.
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// Return an HTTP 404 error with http.StatusNotFound if the album isn’t found.
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
