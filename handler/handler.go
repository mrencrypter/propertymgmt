package handler

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/v12"
	"github.com/murugan-dev/propertymgmt/models"
	logger2 "github.com/sirupsen/logrus"
)

const RentType string = "RENT"
const SellType string = "SELL"

type HouseHandler struct {
	db  *sqlx.DB
	log *logger2.Logger
}

func NewHandler(d *sqlx.DB, l *logger2.Logger) *HouseHandler {
	return &HouseHandler{
		db:  d,
		log: l,
	}
}

func (h *HouseHandler) RentHouse(ctx iris.Context) {
	h.log.Info("received request for renthouse")
	var req models.HouseDetails
	if err := ctx.ReadJSON(&req); err != nil {
		h.log.WithError(err).Error("error occured while reading request")
		ctx.StatusCode(500)
		ctx.Text("Internal Server error")
		return
	}
	if len(req.Name) == 0 || len(req.Address) == 0 {
		ctx.StatusCode(401)
		ctx.Text("Invalid Request")
		return
	}
	if err := h.insertHouseDetail(req, RentType); err != nil {
		h.log.WithError(err).Error("error occured")
		ctx.StatusCode(500)
		ctx.Text("Error occured")
		return
	}

	ctx.StatusCode(200)
	ctx.Text("Added")
}

func (h *HouseHandler) SellHouse(ctx iris.Context) {
	h.log.Info("received request for sellhouse")
	var req models.HouseDetails
	if err := ctx.ReadJSON(&req); err != nil {
		h.log.WithError(err).Error("error reading request")
		ctx.StatusCode(500)
		ctx.Text("Error occured")
		return
	}
	if len(req.Name) == 0 || len(req.Address) == 0 {
		ctx.StatusCode(401)
		ctx.Text("Invalid Request")
		return
	}
	if err := h.insertHouseDetail(req, SellType); err != nil {
		h.log.WithError(err).Error("error occcured")
		ctx.StatusCode(500)
		ctx.Text("Error occured")
		return
	}

	ctx.StatusCode(200)
	ctx.Text("Added")
}

func (h *HouseHandler) FindHouse(ctx iris.Context) {
	h.log.Info("received request for findhouse")
	params := ctx.Params()
	country := params.Get("country")
	locality := params.Get("locality")
	htype := params.Get("type")

	if len(country) == 0 || len(locality) == 0 || len(htype) == 0 {
		ctx.StatusCode(401)
		ctx.Text("invalid request params")
		return
	}

	query := `select hd.name, hd.address, hd.locality, hd.pincode, cd.name country, hd.amount 
			from house_details hd join country_details cd on hd.country_id = cd.id
			where cd.name = $1 and hd.type = $2 and hd.locality = $3`

	var houses []models.HouseDetails
	if err := h.db.Select(&houses, query, country, htype, locality); err != nil {
		h.log.WithError(err).Error("error occured")
		ctx.StatusCode(500)
		ctx.Text("Internal Server Error")
		return
	}
	ctx.StatusCode(200)
	ctx.JSON(houses)
}

func (h *HouseHandler) insertHouseDetail(req models.HouseDetails, htype string) error {
	var cid string
	if err := h.db.Get(&cid, "select id from country_details where name = $1", req.Country); err != nil {
		return fmt.Errorf("error finding country details")
	}
	query := "insert into house_details(id, name, address, locality, pincode, country_id, amount, type)" +
		"values($1, $2, $3, $4, $5, $6, $7, $8)"
	if _, err := h.db.Exec(query, uuid.NewString(), req.Name, req.Address, req.Locality, req.PinCode, cid,
		req.Amount, htype); err != nil {
		return err
	}
	return nil
}
