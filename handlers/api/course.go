package api

import (
	"fmt"
	"net/http"

	"github.com/herdgolf/api/db"
	"github.com/herdgolf/api/models"
	"github.com/labstack/echo/v4"
)

func CreateCourseHandler(c echo.Context) error {
	cs := new(models.Course)
	db := db.DB()

	if err := c.Bind(cs); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	course := &models.Course{
		Name:  cs.Name,
		Par:   cs.Par,
		Holes: cs.Holes,
	}

	if err := db.Create(&course).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &course,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateCourseHandler(c echo.Context) error {
	id := c.Param("id")
	p := new(models.Course)
	db := db.DB()

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	existingCourse := new(models.Course)

	if err := db.First(&existingCourse, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	if p.Name != "" {
		existingCourse.Name = p.Name
	}

	if p.Par != 0 {
		existingCourse.Par = p.Par
	}

	if err := db.Save(&existingCourse).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &existingCourse,
	}

	return c.JSON(http.StatusOK, response)
}

func GetCourseHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	var courses []*models.Course

	if res := db.Model(&models.Course{}).Preload("Holes").Find(&courses, id); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": courses,
	}

	return c.JSON(http.StatusOK, response)
}

func DeleteCourseHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	Course := new(models.Course)

	err := db.Delete(&Course, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"message": "Course deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}

func GetHoleByNumberHandler(c echo.Context) error {
	courseID := c.Param("id")
	holeNumber := c.Param("number")
	db := db.DB()

	var hole *models.Hole

	if res := db.Find(&hole, "course_id = ? AND number = ?", courseID, holeNumber); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": hole,
	}

	return c.JSON(http.StatusOK, response)
}

func GetHolesByCourseIDHandler(c echo.Context) error {
	id := c.Param("id")
	db := db.DB()

	var holes []*models.Hole

	if res := db.Find(&holes, "course_id = ?", id); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	response := map[string]interface{}{
		"data": holes,
	}

	return c.JSON(http.StatusOK, response)
}

func GetCoursesHandler(c echo.Context) error {
	database := db.DB()
	var courses []*models.Course
	database.Scopes(db.Paginate(c)).Find(&courses)

	return c.JSON(http.StatusOK, courses)
}

func UpdateHoleHandler(c echo.Context) error {
	courseID := c.Param("id")
	holeNumber := c.Param("number")
	p := new(models.Hole)
	db := db.DB()

	if err := c.Bind(p); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	existingHole := new(models.Hole)

	if res := db.Find(&existingHole, "course_id = ? AND number = ?", courseID, holeNumber); res.Error != nil {
		return c.JSON(http.StatusInternalServerError, res.Error.Error())
	}

	if p.StrokeIndex != 0 {
		existingHole.StrokeIndex = p.StrokeIndex
	}

	if p.Par != 0 {
		existingHole.Par = p.Par
	}

	if p.Number != 0 {
		existingHole.Number = p.Number
	}
	fmt.Println(existingHole)

	if err := db.Save(&existingHole).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := map[string]interface{}{
		"data": &existingHole,
	}

	return c.JSON(http.StatusOK, response)
}
