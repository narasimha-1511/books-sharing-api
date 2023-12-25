package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/narasimha-1511/zolo-backend/controller"
)

// now lets create a routes function that returns a http.Handler
func Routes(router *gin.Engine){
	//create another route for way /api/v1/
	api := router.Group("/api/v1/")

	api.GET("/booky", controller.GetBook);	
	// api.PUT("/booky/:name/:title/:author", controller.CreateBookParams);
	api.PUT("/booky", controller.CreateBookPostForm);
	api.PUT("/booky/:book_id/borrow", controller.BorrowBook);
	api.GET("/booky/borrowed", controller.GetBorrowedBooks);
	api.POST("/booky/:book_id/borrow/:borrow_id", controller.ReturnBook);
	
}