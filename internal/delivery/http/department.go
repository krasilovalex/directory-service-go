package http

import (
	"directory-service/internal/domain"
	"directory-service/internal/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DepartmentHandler struct {
	usc *usecase.DepartmentUseCase
}

func NewDepartmentHandler(u *usecase.DepartmentUseCase) *DepartmentHandler {
	return &DepartmentHandler{
		usc: u,
	}
}

// CreateDepartment создает новый отдел
// @Summary Создать отдел
// @Tags departments
// @Accept json
// @Produce json
// @Param input body domain.Department true "Данные отдела"
// @Success 201 {object} domain.Department
// @Router /departments [post]
func (h *DepartmentHandler) Create(c *gin.Context) {
	var input domain.Department

	err := c.BindJSON(&input)

	if err != nil {
		c.JSON(400, "Ошибка в JSON файле")
		return
	}
	err = h.usc.Create(c, &input)

	if err != nil {
		fmt.Println("ОШИБКА В БАЗЕ", err)
		c.JSON(500, "Ошибка при вызове фукнции")
		return
	}

	c.JSON(201, input)
}

// GetDepartmentByID ищет отдел по UUID
// @Summary Получить отдел
// @Tags departments
// @Produce json
// @Param id path string true "UUID отдела"
// @Success 200 {object} domain.Department
// @Router /departments/{id} [get]
func (h *DepartmentHandler) GetByID(c *gin.Context) {
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

// GetAll возвращает список отделов
// @Summary Получить все отделы
// @Tags departments
// @Produce json
// @Success 200 {array} domain.Department
// @Router /departments [get]
func (h *DepartmentHandler) GetAll(c *gin.Context) {

	departments, err := h.usc.GetAll(c, 100, 0)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if departments == nil {
		departments = []domain.Department{}
	}

	c.JSON(200, departments)
}

// DeleteLocations удаляет отдел по UUID
// @Summary Удаляет отдел
// @Tags departments
// @Produce json
// @Param id path string true "UUID департамента"
// @Success 200 {object} domain.Department
// @Router /departments/{id} [delete]
func (h *DepartmentHandler) Delete(c *gin.Context) {
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

// UpdateDepartments обновляет список отделов
// @Summary Обновить отдел
// @Tags departments
// @Produce json
// @Param id path string true "UUID отдела"
// @Param input body domain.Department true "Новые данные отдела"
// @Success 200 {object} domain.Department
// @Router /departments/{id} [put]
func (h *DepartmentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")

	parsedID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(400, "Ошибка 400 при извлечении ID")
		return
	}
	var input domain.Department

	err = c.ShouldBind(&input)

	if err != nil {
		c.JSON(400, "Ошибка 400 при извлечении JSON")
		return
	}
	input.ID = parsedID

	err = h.usc.Update(c, &input)

	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось обновить"})
		return
	}

	c.Status(200)
}
