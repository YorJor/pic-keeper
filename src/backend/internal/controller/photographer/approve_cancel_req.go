package photographer

import (
	"net/http"

	"github.com/Roongkun/software-eng-ii/internal/controller/util"
	"github.com/Roongkun/software-eng-ii/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *Resolver) ApproveCancelReq(c *gin.Context) {
	user := c.MustGet("user")
	userObj, ok := user.(model.User)
	if !ok {
		util.Raise400Error(c, "Invalid user type in context")
		return
	}

	paramId := c.Param("id")
	bookingId := uuid.MustParse(paramId)

	booking, err := r.BookingUsecase.BookingRepo.FindOneById(c, bookingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	if booking.Status != model.BookingCustomerReqCancelStatus {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "failed",
			"error":  "this booking is not waiting for your confirmation",
		})
		c.Abort()
		return
	}

	gallery, err := r.GalleryUsecase.GalleryRepo.FindOneById(c, booking.GalleryId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	if gallery.PhotographerId != userObj.Id {
		c.JSON(http.StatusForbidden, gin.H{
			"status": "failed",
			"error":  "this booking is not yours",
		})
		c.Abort()
		return
	}

	booking.Status = model.BookingCancelledStatus
	if err := r.BookingUsecase.BookingRepo.UpdateOne(c, booking); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   booking,
	})
}
