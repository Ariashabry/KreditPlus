package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/ariashabry/KreditPlus/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDB adalah struct untuk mock DB
type MockDB struct {
	mock.Mock
}

// GetById adalah method mock untuk mencari pinjaman berdasarkan ID
func (m *MockDB) GetById(id int) (*models.Pinjaman, error) {
	args := m.Called(id)
	result := args.Get(0)
	if result != nil {
		return result.(*models.Pinjaman), args.Error(1)
	}
	return nil, args.Error(1)
}

// Update adalah method mock untuk mengupdate data pinjaman
func (m *MockDB) Update(pinjaman *models.Pinjaman) error {
	args := m.Called(pinjaman)
	return args.Error(0)
}

// Create adalah method mock untuk membuat transaksi pembayaran
func (m *MockDB) Create(transaction *models.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (c *Context) TestPayMent(t *testing.T) {
	// Membuat instance mock DB
	mockDB := new(MockDB)

	// Membuat instance Echo
	e := echo.New()

	// Menyiapkan data dummy pinjaman
	pinjaman := &models.Pinjaman{
		IDPinjaman: 123,
		Amount:     1000,
		TotalDebt:  500,
		UserID:     456,
		Status:     "Pending",
	}
	mockDB.On("GetById", pinjaman.IDPinjaman).Return(pinjaman, nil)

	// Menyiapkan data dummy pembayaran
	pembayaran := &models.Pinjaman{
		Amount: 200,
	}
	body, _ := json.Marshal(pembayaran)
	req := httptest.NewRequest(http.MethodPost, "/payment/123", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues(strconv.Itoa(int(pinjaman.IDPinjaman)))

	// Membuat instance pinjaman dan transaksi
	pdb := &models.Pinjaman{
		IDPinjaman: pinjaman.IDPinjaman,
		Amount:     pinjaman.Amount,
		TotalDebt:  pinjaman.TotalDebt,
		UserID:     pinjaman.UserID,
		Status:     pinjaman.Status,
	}
	transaction := &models.Transaction{
		LoanID:      uint(pinjaman.IDPinjaman),
		UserID:      pinjaman.UserID,
		Amount:      pembayaran.Amount,
		PaymentDate: time.Now(),
	}

	// Mock pemanggilan metode
	mockDB.On("Update", pdb).Return(nil)
	mockDB.On("Create", transaction).Return(nil)

	// Eksekusi fungsi PayMent
	err := c.PayMent(ctx)
	assert.NoError(t, err)

	// Memeriksa response code
	assert.Equal(t, http.StatusCreated, rec.Code)

	// Memeriksa data pinjaman setelah pembayaran
	var response struct {
		Data *models.Pinjaman `json:"data"`
	}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, pinjaman.Amount, response.Data.TotalDebt)
	assert.Equal(t, "Paid", response.Data.Status)

	// Memeriksa eksekusi method mock
	mockDB.AssertExpectations(t)
}
