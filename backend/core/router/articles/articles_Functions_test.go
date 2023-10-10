package articles_test

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/golang/mock/gomock"
	"github.com/gin-gonic/gin"
	"github.com/your-package-path/articles"
	//mock_articles "github.com/your-package-path/articles/mock_articles" // Import du package de mock généré
)

func TestAddArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContext := mock_articles.NewMockContext(ctrl)

	//   instance de votre routeur Gin et  créer une demande HTTP de test.
	router := gin.Default()
	router.POST("/articles", articles.AddArticle)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/articles", nil)

	// Définissez le comportement attendu du mock de contexte
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)
	mockContext.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil)

	// Utilisez le routeur Gin pour servir la demande HTTP de test avec le mock de contexte
	router.ServeHTTP(w, req)

	// Vérifiez le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}


func TestDeleteArticle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContext := mock_articles.NewMockContext(ctrl)

	// Créez une instance de votre routeur Gin et utilisez-la pour créer une demande HTTP de test.
	router := gin.Default()
	router.DELETE("/articles/:id", articles.DeleteArticle)

	// Simulez une demande HTTP DELETE avec un paramètre d'ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/articles/123", nil)
	req = req.WithContext(gin.WrapH(mockContext))

	// Définissez le comportement attendu du mock de contexte
	mockContext.EXPECT().Param("id").Return("123")
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)

	// Utilisez le routeur Gin pour servir la demande HTTP de test avec le mock de contexte
	router.ServeHTTP(w, req)

	// Vérifiez le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAddLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContext := mock_articles.NewMockContext(ctrl)

	// Créez une instance de votre routeur Gin et utilisez-la pour créer une demande HTTP de test.
	router := gin.Default()
	router.POST("/articles/:id/likes", articles.AddLike)

	// Simulez une demande HTTP POST avec un paramètre d'ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/articles/123/likes", nil)
	req = req.WithContext(gin.WrapH(mockContext))

	// Définissez le comportement attendu du mock de contexte
	mockContext.EXPECT().Param("id").Return("123")
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)

	// Utilisez le routeur Gin pour servir la demande HTTP de test avec le mock de contexte
	router.ServeHTTP(w, req)

	// Vérifiez le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestRemoveLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockContext := mock_articles.NewMockContext(ctrl)

	// Créez une instance de votre routeur Gin et utilisez-la pour créer une demande HTTP de test.
	router := gin.Default()
	router.DELETE("/articles/:id/likes", articles.RemoveLike)

	// Simulez une demande HTTP DELETE avec un paramètre d'ID
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/articles/123/likes", nil)
	req = req.WithContext(gin.WrapH(mockContext))

	// Définissez le comportement attendu du mock de contexte
	mockContext.EXPECT().Param("id").Return("123")
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)
	mockContext.EXPECT().ShouldBindJSON(gomock.Any()).Return(nil)

	// Utilisez le routeur Gin pour servir la demande HTTP de test avec le mock de contexte
	router.ServeHTTP(w, req)

	// Vérifiez le code de statut de la réponse
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}