package products

import (
	"go-clean-arc-sauna-shop-app/app/application/product"
	"go-clean-arc-sauna-shop-app/app/presentation/settings"

	validator "go-clean-arc-sauna-shop-app/pkg/validator"

	"github.com/gin-gonic/gin"
)

type handler struct {
	saveProductUseCase  *product.SaveProductUseCase
	fetchProductUseCase *product.FetchProductUseCase
}

func NewHandler(
	saveProductUseCase *product.SaveProductUseCase,
	fetchProductUseCase *product.FetchProductUseCase,
) handler {
	return handler{
		saveProductUseCase:  saveProductUseCase,
		fetchProductUseCase: fetchProductUseCase,
	}
}

type PostProductsParams struct {
	OwnerID     string `json:"owner_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}

func (h handler) PostProducts(ctx *gin.Context) {

	// リクエストパラメータの取得
	var params PostProductsParams
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		settings.ReturnBadRequest(ctx, err)
	}

	// バリデーションの実行
	validate := validator.GetValidator()
	err = validate.Struct(params)
	if err != nil {
		settings.ReturnStatusBadRequest(ctx, err)
	}

	// ユースケースで定義した InputDto へ データを詰め替え
	input := product.SaveProductUseCaseInputDto{
		OwnerID:     params.OwnerID,
		Name:        params.Name,
		Description: params.Description,
		Price:       params.Price,
		Stock:       params.Stock,
	}

	// ユースケースを呼び出す
	dto, err := h.saveProductUseCase.Run(ctx, input)
	if err != nil {
		settings.ReturnError(ctx, err)
	}

	// レスポンスを作成
	response := postProductResponse{
		productResponseModel{
			Id:          dto.ID,
			OwnerID:     dto.OwnerID,
			Name:        dto.Name,
			Description: dto.Description,
			Price:       dto.Price,
			Stock:       dto.Stock,
		},
	}
	settings.ReturnStatusCreated(ctx, response)

}
