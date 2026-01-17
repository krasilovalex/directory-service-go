package http

import (
	"directory-service/internal/domain"
	"directory-service/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PositionHandler struct {
	usc *usecase.PositionUseCase
}

func NewPositionHandler(u *usecase.PositionUseCase) *PositionHandler {
	return &PositionHandler{
		usc: u,
	}
}

// CreatePosition создает новую должность
// @Summary Создать должность
// @Tags positions
// @Accept json
// @Produce json
// @Param input body domain.Position true "Данные должности"
// @Success 201 {object} domain.Position
// @Router /positions [post]
func (h *PositionHandler) Create(c *gin.Context) {
	var input domain.Position

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(400, "Ошибка в JSON файле")
		return
	}

	err = h.usc.Create(c, &input)

	if err != nil {
		fmt.Println("ОШИБКА В БАЗЕ", err)
		c.JSON(500, "Ошибка при вызове функции")
		return
	}

	c.JSON(201, input)
}

// GetPositionByID ищет должность по UUID
// @Summary Получить должность
// @Tags positions
// @Produce json
// @Param id path string true "UUID должности"
// @Success 200 {object} domain.Position
// @Router /positions/{id} [get]
func (h *PositionHandler) GetByID(c *gin.Context) {
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

// UpdatePositions обновляет список должностей
// @Summary Обновить должность
// @Tags positions
// @Produce json
// @Param id path string true "UUID должности"
// @Param input body domain.Position true "Новые данные должности"
// @Success 200 {object} domain.Position
// @Router /positions/{id} [put]
func (h *PositionHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	parsedID, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(400, "Ошибка 400 при извлечении ID")
		return
	}

	var input domain.Position

	input.ID = parsedID

	err = h.usc.Update(c, &input)

	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось обновить"})
		return
	}

	c.Status(200)
}

// GetAll возвращает список отделов
// @Summary Получить все отделы
// @Tags positions
// @Produce json
// @Success 200 {array} domain.Position
// @Router /positions [get]
func (h *PositionHandler) GetAll(c *gin.Context) {

	positions, err := h.usc.GetAll(c, 100, 0)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if positions == nil {
		positions = []domain.Position{}
	}

	c.JSON(200, positions)
}

// DeletePosition удаляет долнжость по uuid
// @Summary Удалить должность
// @Tags positions
// @Produce json
// @Param id path string true "UUID офиса"
// @Success 200 {object} domain.Position
// @Router /positions/{id} [delete]
func (h *PositionHandler) Delete(c *gin.Context) {
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
