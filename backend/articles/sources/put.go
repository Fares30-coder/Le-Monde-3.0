package sources

// import (
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// 	"net/http"
// )

// func AddIpfs(c *gin.Context, db *gorm.DB) {
// 	sh := shell.NewShell("localhost:5001")

//     file, header, err := c.Request.FormFile("file")
//     if err != nil {
//         c.JSON(500, gin.H{"error": err.Error()})
//         return
//     }
//     defer file.Close()

//     hash, err := sh.Add(file)
//     if err != nil {
//         c.JSON(500, gin.H{"error": err.Error()})
//         return
//     }

//     c.JSON(200, gin.H{"filename": header.Filename, "hash": hash})
// }

// func AddDB(c *gin.Context, db *gorm.DB) {
// 	article := new(database.Article)

// 	if err := c.Bind(&article); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	result := db.Create(&article)
// 	if result.Error != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{"Created" : "Article created successfully"})
// }