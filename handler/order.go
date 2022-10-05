package handler

import (
	"assignment2/helper"
	"assignment2/models"
	"assignment2/param"
	"assignment2/repo"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	orderRepo repo.OrderRepo
	itemRepo  repo.ItemRepo
	db        *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{
		orderRepo: repo.NewOrderRepo(db),
		itemRepo:  repo.NewItemRepo(db),
		db:        db,
	}
}

// Order godoc
// @Summary Create Order
// @Schemes
// @Description Create Order
// @Accept json
// @Produce json
// @Param data body param.OrderReq true "body data"
// @Success 200 {object} map[string]interface{}
// @Router /order [POST]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req = param.OrderReq{}
	var err = c.Bind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), helper.Map{})
		return
	}

	if len(req.Items) == 0 { // CEK ITEM
		helper.JSON(c, http.StatusBadRequest, "items not found", helper.Map{})
		return
	}

	var order = models.Order{
		CustomerName: req.CustomerName,
		OrderAt:      req.OrderAt,
	}

	var items = []models.OrderItem{}
	var trx = h.db.Begin()
	defer trx.Commit()

	err = h.orderRepo.Create(&order, trx) //CREATE ORDER
	if err != nil {
		trx.Rollback()
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	for _, v := range req.Items { //CREATE MODEL ITEM
		items = append(items, models.OrderItem{
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
			OrderId:     order.Id,
		})
	}

	err = h.itemRepo.Creates(&items, trx) //CREATE ITEM
	if err != nil {
		trx.Rollback()
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	//RESPONSE
	var itemResp []param.ItemResp
	for _, v := range items {
		itemResp = append(itemResp, param.ItemResp{
			Id:          v.Id,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
		})
	}

	helper.JSON(c, http.StatusOK, http.StatusText(http.StatusOK), param.OrderResp{
		OrderId:      order.Id,
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        itemResp,
	})
	return
}

// Order godoc
// @Summary Get Order
// @Schemes
// @Description Get Order
// @Accept json
// @Produce json
// @Param id  path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Router /order/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	var tmp = c.Param("order_id")
	var orderId, err = strconv.Atoi(tmp)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), helper.Map{})
		return
	}
	order, err := h.orderRepo.Get(orderId) //GET ORDER

	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	if order.Id == 0 { //CEK ORDER
		helper.JSON(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), helper.Map{})
		return
	}

	//RESP
	var itemResp []param.ItemResp
	for _, v := range order.Item {
		itemResp = append(itemResp, param.ItemResp{
			Id:          v.Id,
			ItemCode:    v.ItemCode,
			Description: v.Description,
			Quantity:    v.Quantity,
		})
	}

	helper.JSON(c, http.StatusOK, http.StatusText(http.StatusOK), param.OrderResp{
		OrderId:      order.Id,
		OrderAt:      order.OrderAt,
		CustomerName: order.CustomerName,
		Items:        itemResp,
	})
	return
}

// Order godoc
// @Summary Get Order
// @Schemes
// @Description Get Order
// @Accept json
// @Produce json
// @Param order_id path int true "Order Id"
// @Param data body param.OrderReq true "body data"
// @Success 200 {object} map[string]interface{}
// @Router /order/{order_id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	tmp := c.Param("order_id")
	var orderId, err = strconv.Atoi(tmp)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), helper.Map{})
		return
	}
	var req = param.OrderReq{}

	err = c.Bind(&req)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), helper.Map{})
		return
	}

	order, err := h.orderRepo.Get(orderId) //GET ORDER
	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	if order.Id == 0 { // CEK ORDER
		helper.JSON(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), helper.Map{})
		return
	}

	//SET DATA ORDER
	order.CustomerName = req.CustomerName
	order.OrderAt = req.OrderAt

	err = h.orderRepo.Update(&order) //UPDATE ORDER
	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	if len(req.Items) > 0 {
		var updatesItem []models.OrderItem
		for _, v := range req.Items {
			updatesItem = append(updatesItem, models.OrderItem{
				Id:          v.Id,
				ItemCode:    v.ItemCode,
				Description: v.Description,
				Quantity:    v.Quantity,
				OrderId:     order.Id,
			})
		}

		err = h.itemRepo.Updates(&updatesItem) // UPDATE ITEM
		if err != nil {
			helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
			return
		}
	}

	helper.JSON(c, http.StatusOK, http.StatusText(http.StatusOK), helper.Map{})
	return

}

// Order godoc
// @Summary Delete Order
// @Schemes
// @Description Delete Order
// @Accept json
// @Produce json
// @Param order_id path int true "Order Id"
// @Success 200 {object} map[string]interface{}
// @Router /order/{order_id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	var tmp = c.Param("order_id")
	var orderId, err = strconv.Atoi(tmp)
	if err != nil {
		helper.JSON(c, http.StatusBadRequest, err.Error(), helper.Map{})
		return
	}

	order, err := h.orderRepo.Get(orderId)
	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	if order.Id == 0 {
		helper.JSON(c, http.StatusNotFound, http.StatusText(http.StatusNotFound), helper.Map{})
		return
	}

	var trx = h.db.Begin()
	defer trx.Commit()

	err = h.itemRepo.DeleteByOrder(order.Id, trx)
	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	err = h.orderRepo.Delete(order.Id, trx)
	if err != nil {
		helper.JSON(c, http.StatusInternalServerError, err.Error(), helper.Map{})
		return
	}

	helper.JSON(c, http.StatusOK, http.StatusText(http.StatusOK), helper.Map{})
	return

}
