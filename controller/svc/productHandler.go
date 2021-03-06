package svc

import (
	"encoding/json"
	"net/http"

	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/util/log"
)

type ProductInfo struct {
	ProductKey   string `json:"ProductKey"`
	ProductName  string `json:"ProductName"`
	Description  string `json:"Description"`
	DeviceCount  int32  `json:"DeviceCount"`
	AccessPoints string `json:"AccessPoints"`
	CreateAt     string `json:"CreateAt"`
	UpdateAt     string `json:"UpdateAt"`

	ProductSecret string `json:"ProductSecret,omitempty"`
	DeleteFlag    int8
}

type ProductResp struct {
	Code        string
	Message     string
	ProductInfo ProductInfo
}

func handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("ProductName")
	productDesc := r.URL.Query().Get("Description")
	// create product
	productInfo := &ProductInfo{
		ProductKey:    uuid.New().String(),
		ProductSecret: uuid.New().String(),
		ProductName:   productName,
		Description:   productDesc,
		DeviceCount:   0,
		AccessPoints:  "",
		CreateAt:      time.Now().Format(time.RFC3339),
		UpdateAt:      time.Now().Format(time.RFC3339),
		DeleteFlag:    0,
	}
	log.Infof("create product name : %s, description : %s", productInfo.ProductName, productInfo.Description)
	w.Header().Set("Content-Type", "application/json")
	_, err := Ctx.Dbsvc.CreateProduct(*productInfo)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		var productResp ProductResp
		productResp.Code = "200"
		productResp.ProductInfo = *productInfo
		b, _ := json.Marshal(productResp)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleQueryProduct(w http.ResponseWriter, r *http.Request) {
	productKey := r.URL.Query().Get("ProductKey")
	log.Infof("query product info, product key : %s", productKey)
	w.Header().Set("Content-Type", "application/json")
	productInfo, err := Ctx.Dbsvc.QueryProductInfo(productKey)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		productInfo.ProductKey = productKey
		var productResp ProductResp
		productResp.Code = "200"
		productResp.ProductInfo = *productInfo
		b, _ := json.Marshal(productResp)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleQueryProductList(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("Offset"))
	keyword := r.URL.Query().Get("Keyword")
	log.Infof("query product list, limit : %d, offset : %d, keyword : %s", limit, offset, keyword)
	w.Header().Set("Content-Type", "application/json")
	productList, err := Ctx.Dbsvc.QueryProductList(int32(offset), int32(limit), keyword)
	if err != nil {
		errInfo := &ErrInfo{
			Code:    "500",
			Message: err.Error(),
		}
		b, _ := json.Marshal(errInfo)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	} else {
		type ProductListResp struct {
			Code            string
			Message         string
			TotalCount      int
			ProductInfoList []ProductInfo
		}
		var productListResp ProductListResp
		productListResp.Code = "200"
		productListResp.TotalCount = len(productList)
		productListResp.ProductInfoList = productList
		b, _ := json.Marshal(productListResp)
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	}
}

func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {
}

func handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
}

func handleQueryProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleModifyProductQuota(w http.ResponseWriter, r *http.Request) {
}

func handleCheckProductName(w http.ResponseWriter, r *http.Request) {
}
