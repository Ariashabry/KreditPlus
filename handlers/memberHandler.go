package handlers

import (
	"github.com/ariashabry/KreditPlus/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func (c *Context) GetAllMember(ctx echo.Context) error {
	m := models.Konsumens{}

	err := m.GetAll(c.DB)

	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed getting data",
		}
		log.Println("Error getting data konsumens:", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	if len(m) > 0 {
		Results := &Result{
			Success: true,
			Data:    &m,
			Code:    http.StatusOK,
			Message: "Success getting data",
		}
		return ctx.JSON(http.StatusOK, Results)
	} else {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusNotFound,
			Message: "Data Not Found",
		}
		return ctx.JSON(http.StatusNotFound, Results)
	}
}

func (c *Context) InsertKonsumen(ctx echo.Context) error {
	m := models.Konsumen{}
	err := ctx.Bind(&m)
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed send data",
		}
		log.Println("Error creating new konsumen:", err.Error())
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	err = m.Create(c.DB)

	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed Adding data",
		}
		log.Println("Failed Adding data", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	Results := &Result{
		Success: true,
		Data:    m,
		Code:    http.StatusCreated,
		Message: "Success adding new konsumen",
	}
	return ctx.JSON(http.StatusCreated, Results)
}

func (c *Context) UpdateKonsumen(ctx echo.Context) error {
	//get from request
	m := models.Konsumen{}
	err := ctx.Bind(&m)
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed send data",
		}
		log.Println("Error creating member:", err.Error())
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	//check data to db using id from request
	mdb := models.Konsumen{}
	err = mdb.GetById(c.DB, int(m.IdKonsumen))

	//check error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User Not Found")
			return ctx.JSON(http.StatusNotFound, nil)
		} else {
			log.Println("Error", err.Error())
			return ctx.JSON(http.StatusInternalServerError, nil)
		}
	}

	//ganti data yang mau diupdate
	mdb.FullName = m.FullName
	mdb.LegalName = m.LegalName
	mdb.TempatLahir = m.TempatLahir
	mdb.TanggalLahir = m.TanggalLahir
	mdb.Gaji = m.Gaji
	mdb.FotoKTP = m.FotoKTP
	mdb.FotoSelfie = m.FotoSelfie
	mdb.Gender = m.Gender

	//lakukan update
	err = mdb.Update(c.DB)

	//cek success or error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed Updating Data",
		}
		log.Println("Failed updating data", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}
	Results := &Result{
		Success: true,
		Data:    mdb,
		Code:    http.StatusCreated,
		Message: "Success Updating Data",
	}
	return ctx.JSON(http.StatusCreated, Results)
}

func (c *Context) DeleteKonsumen(ctx echo.Context) error {
	//get parameter
	id, _ := strconv.Atoi(ctx.Param("id"))

	//get data from db using id
	m := models.Konsumen{}
	err := m.GetById(c.DB, id)

	//cek apakah user tersebut ada atau tidak
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("User Not Found")
			return ctx.JSON(http.StatusNotFound, nil)
		} else {
			log.Println("Failed deleting member", err.Error())
			return ctx.JSON(http.StatusInternalServerError, nil)
		}
	}

	//delete
	err = m.Delete(c.DB)

	//cek success or error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed Deleting Data",
		}
		log.Println("Failed Deleting data", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}
	Results := &Result{
		Success: true,
		Data:    m,
		Code:    http.StatusOK,
		Message: "Success Deleting Data",
	}
	return ctx.JSON(http.StatusOK, Results)
}
