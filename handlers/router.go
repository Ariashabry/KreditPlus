package handlers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Context struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

type Result struct {
	Success bool
	Data    interface{}
	Code    int
	Message string
}

func (c *Context) Api(group string) {
	public := c.Echo.Group(group)
	{
		public.POST("/konsumen", c.InsertKonsumen)
		public.PUT("/konsumen", c.UpdateKonsumen)
		public.DELETE("/konsumen/:id", c.DeleteKonsumen)
		public.GET("/konsumen", c.GetAllMember)
		public.POST("/pinjam", c.Pinjam)
		public.PUT("/approve/:id", c.UpdatePinjam)
		public.GET("/konsumen/:id", c.SeeStatus)
		public.POST("/payment/:id", c.Pinjam)
	}

}
