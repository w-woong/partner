package route

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/w-woong/common"
	"github.com/w-woong/common/middlewares"
	commonport "github.com/w-woong/common/port"
	"github.com/w-woong/partner/delivery"
	"github.com/w-woong/partner/port"
)

func AddressRoute(router *mux.Router, conf common.ConfigHttp,
	tokenCookie commonport.TokenCookie, parser commonport.IDTokenParser, userSvc commonport.UserSvc,
	usc port.AddressFinder) *delivery.AddressHttpHandler {

	handler := delivery.NewAddressHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)
	router.HandleFunc("/v1/partner/address", handler.HandleFindAddress).Methods(http.MethodGet)
	router.HandleFunc("/v1/partner/address/_find",
		middlewares.AuthBearerHandler(handler.HandleFindAddress, conf.BearerToken)).Methods(http.MethodGet)
	return handler

	// handler := delivery.NewPayKcpHttpHandler(time.Duration(conf.Timeout)*time.Second, usc)

	// router.HandleFunc("/v1/payment", handler.HandleRegisterCard).Methods(http.MethodGet, http.MethodPost)
	// router.HandleFunc("/v1/payment/callback", handler.HandleOrderCardCallback).Methods(http.MethodGet, http.MethodPost)
	// router.HandleFunc("/v1/order/cart", middlewares.AuthIDTokenUserAccountHandler(
	// 	handler.HandleFindByUserID, validator, userSvc,
	// )).Methods(http.MethodGet)
	// router.HandleFunc("/v1/order/cart/_find-or-create", middlewares.AuthIDTokenUserAccountHandler(
	// 	handler.HandleFindOrCreateByUserID, validator, userSvc,
	// )).Methods(http.MethodGet)

	// router.HandleFunc("/test/order/cart", middlewares.AuthIDTokenUserAccountHandler(
	// 	handler.HandleTestRefreshError, validator, userSvc,
	// )).Methods(http.MethodGet)

	// router.HandleFunc("/v1/order/cart/product", middlewares.AuthIDTokenUserAccountHandler(
	// 	handler.HandleAddCartProduct, validator, userSvc,
	// )).Methods(http.MethodPost)

}
