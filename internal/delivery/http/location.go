package http

import (
	"directory-service/internal/domain"
	"directory-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LocationHandler struct {
	usc *usecase.LocationUseCase
}

func NewLocationHandler(u *usecase.LocationUseCase) *LocationHandler {
	return &LocationHandler{
		usc: u,
	}
}

// CreateLocation создает новое местоположение офиса
// @Summary Создать офис
// @Tags locations
// @Accept json
// @Produce json
// @Param input body domain.Location true "Данные офисаа"
// @Success 201 {object} domain.Location
// @Router /locations [post]
func (h *LocationHandler) Create(c *gin.Context) {
	var input domain.Location

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(400, "Ошибка в JSON файле")
		return
	}

	err = h.usc.Create(c, &input)

	if err != nil {
		c.JSON(500, "Ошибка при вызове функции")
		return
	}

	c.JSON(201, input)
}

// GetLocationsByID ищет офис по UUID
// @Summary Получить фоис
// @Tags locations
// @Produce json
// @Param id path string true "UUID офиса"
// @Success 200 {object} domain.Location
// @Router /locations/{id} [get]
func (h *LocationHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")

	parsedID, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(400, "Ошибка при получении данных")
		return
	}

	result, err := h.usc.GetByID(c, parsedID)

	if err != nil {
		c.JSON(404, "Информация не найдена")
		return
	}

	c.JSON(200, result)
}

// UpdateLocations обновляет список локаций
// @Summary Обновить локацию
// @Tags locations
// @Produce json
// @Param id path string true "UUID офиса"
// @Param input body domain.Location true "Новые данные офиса"
// @Success 200 {object} domain.Location
// @Router /locations/{id} [put]
func (h *LocationHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	parsedID, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(400, "Ошибка 400 при извлечении ID")
		return

	}
	var input domain.Location

	input.ID = parsedID

	err = h.usc.Update(c, &input)

	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось обновить"})
		return
	}

	c.Status(200)

}

// GetAll возвращает список локаций
// @Summary Получить все локации
// @Tags locations
// @Produce json
// @Success 200 {array} domain.Location
// @Router /locations [get]
func (h *LocationHandler) GetAll(c *gin.Context) {

	locations, err := h.usc.GetAll(c, 100, 0)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if locations == nil {
		locations = []domain.Location{}
	}

	c.JSON(200, locations)
}

// DeleteLocations удаляет офис по UUID
// @Summary Удаляет локацию
// @Tags locations
// @Produce json
// @Param id path string true "UUID офиса"
// @Success 200 {object} domain.Location
// @Router /locations/{id} [delete]
func (h *LocationHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")

	parsedID, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(400, "Ошибка JSON")
		return
	}

	err = h.usc.Delete(c, parsedID)

	if err != nil {
		c.JSON(500, "Нету данных")
		return
	}

	c.JSON(200, gin.H{"status": "deleted"})

}
