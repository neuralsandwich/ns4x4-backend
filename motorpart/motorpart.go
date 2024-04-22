package motorpart

import (
    "errors"
)

var (
	ErrBadId = errors.New("ID is invalid")
	ErrBadField = errors.New("Field is invalid")
	ErrNoId = errors.New("ID is missing")
	ErrPartExists = errors.New("Part already exists")
)

type Part struct {
    Id string `json:"id"`
    
    Tagging string `json:"tagging"`
    PartsUnique string `json:"parts_unique"`
    PartsUniqueAdmin string `json:"parts_unique_admin"`
    Type string `json:"type"`
    PartsComboSKU string `json:"parts_combo_sku"`
    IsCombo string `json:"is_combo"`
    GenuinePartNumber string `json:"genuine_part_number"`
    ConsumerName string `json:"consumer_name"`
    PartName string `json:"part_name"`
    PurchasePrice string `json:"purchase_price"`
    PurchageCurrency string `json:"purchase_currency"`
    PurchasePriceEurp string `json:"purchase_price_euro"`
    Tariff int `json:"tariff"`
    ProductCategory string `json:"product_category"`
    ProductSubcategoryOne string `json:"product_subcategory_one"`
    ProductSubcategoryTwo string `json:"product_subcategory_two"`
    ProductSubcategoryThree string `json:"product_subcategory_three"`
    ProductSubcategoryFour string `json:"product_subcategory_four"`
    Description string `json:"description"`
    Specification string `json:"specification"`
    Notes string `json:"notes"`
    UKDescription string `json:"uk_description"`
    UKSpecification string `json:"uk_specification"`
    UKnotes string `json:"uk_notes"`
    Shipping string `json:"shipping"`
    WholesaleMargin string `json:"wholesale_margin"`
    MotorFactorMargin string `json:"motor_factor_margin"`
    TradeMargin string `json:"trade_margin"`
    OnlineMargin string `json:"online_margin"`
    OnlinePlusMargin string `json:"online_plus_margin"`
    TradePlusMargin string `json:"trade_plus_margin"`
    WholesalePriceEuro string `json:"wholesale_price_euro"`
    MotorFactorPriceEuro string `json:"motor_factor_price_euro"`
    TradePriceEuro string `json:"trade_price_euro"`
    TradePlusPriceEurp string `json:"trade_plus_price_euro"`
    OnlinePriceEuro string `json:"online_price_euro"`
    OnlinePlusPriceEuro string `json:"online_plus_price_euro"`
    CurrentStock string `json:"current_stock"`
    StockAlert string `json:"stock_alert"`
    LotSize string `json:"lot_size"`
    LocationDescription string `json:"location_description"`
    SupplierName string `json:"supplier_name"`
    ProductStatus string `json:"product_status"`
    ProductPublishStatus string `json:"product_publish_status"`
    UKProductPublishStatus string `json:"uk_product_publish_status"`
    Years string `json:"years"`
    StartYear string `json:"start_year"`
    EndYear string `json:"end_year"`
    VATRate string `json:"vat_rate"`
    YenEuroConverionRate string `json:"yen_euro_conversion_rate"`
    DollarEuroConversionRate string `json:"dollar_euro_conversion_rate"`
    EuroSterlingConverionRate string `json:"euro_sterling_conversion_rate"`
    RelatedProductId string `json:"related_product_id"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
    DeleteAt string `json:"deleted_at"`
    CurrentStockMessage string `json:"current_stock_message"`
    TarrifCode string `json:"tariff_code"`
    MetaDescription string `json:"meta_description"`
    MetaKeywords string `json:"meta_keywords"`
    MetaTitle string `json:"meta_title"`
    MoveToUK string `json:"move_to_uk"`
    UKPriceExVAT string `json:"uk_price_ex_vat"`
    ReadyUK int `json:"ready_uk"`
    ComboUpdate string `json:"combo_update"`
    SubCatOneName string `json:"sub_cat_one_name"`
    SubCatTwoName string `json:"sub_cat_two_name"`
    SubCatThreeName string `json:"sub_cat_three_name"`
    CategoryName string `json:"category_name"`
    PriceEuro string `json:"price_euro"`
    TelephonePriceEuro string `json:"telephone_price_euro"`
    PriceEuroVAT string `json:"price_euro_vat"`
    PriceSterling string `json:"price_sterling"`
    TelephonePriceSterling string `json:"telephone_price_sterling"`
    PriceSterlingVAT string `json:"price_sterling_vat"`
    MakeName string `json:"make_name"`
    ModelName string `json:"model_name"`
    ImageName string `json:"image_name"`
    Checkbox string `json:"checkbox"`
    IrishStatus string `json:"irish_status"`
    UKStatus string `json:"uk_status"`
    Combo string `json:"combo"`
    Stock string `json:"stock"`
    Sub4 string `json:"sub4"`
}
