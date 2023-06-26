package handlers

import (
	"github.com/ariashabry/KreditPlus/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

// request a loan
func (c *Context) Pinjam(ctx echo.Context) error {
	// Ambil data pinjaman dari body request
	p := models.Pinjaman{}

	err := ctx.Bind(&p)

	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed sending data",
		}
		log.Println("Error request a loan")
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	// Validasi nilai pinjaman dan tenor
	if p.Amount <= 0 || p.Tenor <= 0 {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Invalid loan amount or tenor",
		}
		log.Println("Invalid loan amount or tenor")
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	// Cek batas limit pinjaman
	// Simulasi batas limit dengan jumlah pinjaman maksimal 1.000.000
	if p.Amount > 1000000 {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Loan amount exceeds the limit",
		}
		log.Println("Loan amount exceeds the limit")
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	// Cek tenor pinjaman
	// Simulasi batas tenor dengan tenor maksimal 12 bulan
	if p.Tenor > 12 {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Loan tenor exceeds the limit",
		}
		log.Println("Loan tenor exceeds the limit")
		return ctx.JSON(http.StatusBadRequest, Results)
	}

	p.Status = "Pending"
	p.TotalDebt = 0
	p.InterestRate = 0.02 // anggap saja simulasi bunga bersifat flat, diangka 2%.
	p.MonthlyPayment = 0

	//lakukan penyimpanan ke db
	err = p.Create(c.DB)

	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed request a loan",
		}
		log.Println("Failed request a loan", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	Results := &Result{
		Success: true,
		Data:    p,
		Code:    http.StatusCreated,
		Message: "Success request a loan",
	}
	return ctx.JSON(http.StatusCreated, Results)
}

//approve a loan
func (c *Context) UpdatePinjam(ctx echo.Context) error {
	//get loan id
	loanID, _ := strconv.Atoi(ctx.Param("id"))

	// Cari pinjaman berdasarkan ID
	pdb := models.Pinjaman{}
	err := pdb.GetById(c.DB, loanID)
	//cek apakah user tersebut ada atau tidak
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: "Loan Not Found",
			}
			log.Println("Loan Not Found")
			return ctx.JSON(http.StatusNotFound, Results)
		} else {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusInternalServerError,
				Message: "Loan Not Found",
			}
			log.Println("Failed Approve loan", err.Error())
			return ctx.JSON(http.StatusInternalServerError, Results)
		}
	}

	// Setujui pinjaman
	pdb.Status = "Approved"

	// hitung jumlah pembayaran bulanan
	pdb.MonthlyPayment = int(float64(pdb.Amount) * (1 + pdb.InterestRate) / float64(pdb.Tenor))

	//lakukan update
	err = pdb.Update(c.DB)

	//cek success or error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Failed Approve a loan",
		}
		log.Println("Failed updating data", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}
	Results := &Result{
		Success: true,
		Data:    pdb,
		Code:    http.StatusCreated,
		Message: "Success UApproving a loan",
	}
	return ctx.JSON(http.StatusCreated, Results)
}

//see a loan
func (c *Context) SeeStatus(ctx echo.Context) error {
	UserID, _ := strconv.Atoi(ctx.Param("id"))

	// Cari pinjaman berdasarkan ID
	u := models.Konsumen{}
	err := u.GetById(c.DB, UserID)
	//cek apakah user tersebut ada atau tidak
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: "Loan Not Found",
			}
			log.Println("Loan Not Found")
			return ctx.JSON(http.StatusNotFound, Results)
		} else {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusInternalServerError,
				Message: "Loan Not Found",
			}
			log.Println("Failed Approve loan", err.Error())
			return ctx.JSON(http.StatusInternalServerError, Results)
		}
	}

	Results := &Result{
		Success: true,
		Data:    u,
		Code:    http.StatusOK,
		Message: "Success Showing Konsumen Data",
	}
	return ctx.JSON(http.StatusOK, Results)

}

func (c *Context) PayMent(ctx echo.Context) error {
	//get loan id
	loanID, _ := strconv.Atoi(ctx.Param("id"))

	// Cari pinjaman berdasarkan ID
	pdb := models.Pinjaman{}
	err := pdb.GetById(c.DB, loanID)
	//cek apakah user tersebut ada atau tidak
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusNotFound,
				Message: "Loan Not Found",
			}
			log.Println("Loan Not Found")
			return ctx.JSON(http.StatusNotFound, Results)
		} else {
			Results := &Result{
				Success: false,
				Data:    nil,
				Code:    http.StatusInternalServerError,
				Message: "Loan Not Found",
			}
			log.Println("Payment Failed", err.Error())
			return ctx.JSON(http.StatusInternalServerError, Results)
		}
	}
	// Ambil data pembayaran dari body request
	pu := models.Pinjaman{}
	err = ctx.Bind(&pu)

	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusBadRequest,
			Message: "Failed sending data",
		}
		log.Println("Error doing payment")
		return ctx.JSON(http.StatusBadRequest, Results)
	}
	// Validasi nilai pembayaran
	if pu.Amount <= 0 {
		return ctx.JSON(http.StatusBadRequest, "Invalid payment amount")
	}

	// Hitung jumlah total hutang yang belum terbayar
	remainingDebt := pdb.Amount - pdb.TotalDebt

	// Cek apakah pembayaran lebih dari atau sama dengan total hutang yang belum terbayar
	if pu.Amount >= remainingDebt {
		// Bayar hutang secara penuh
		pdb.TotalDebt = pdb.Amount
		pdb.Status = "Paid"
	} else {
		// Bayar sebagian hutang
		pdb.TotalDebt += pu.Amount
	}

	// Simpan data pembayaran ke database
	transaction := &models.Transaction{
		LoanID:      uint(loanID),
		UserID:      pdb.UserID,
		Amount:      pu.Amount,
		PaymentDate: time.Now(),
	}

	err = transaction.Create(c.DB)
	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Payment Failed",
		}
		log.Println("Payment Failed", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	err = pdb.Update(c.DB)

	//cek error
	if err != nil {
		Results := &Result{
			Success: false,
			Data:    nil,
			Code:    http.StatusInternalServerError,
			Message: "Payment Failed",
		}
		log.Println("Payment Failed", err.Error())
		return ctx.JSON(http.StatusInternalServerError, Results)
	}

	Results := &Result{
		Success: true,
		Data:    pdb,
		Code:    http.StatusCreated,
		Message: "Success request a loan",
	}
	return ctx.JSON(http.StatusCreated, Results)

}
