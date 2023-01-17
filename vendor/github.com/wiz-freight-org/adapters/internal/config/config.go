package config

import (
	"os"
	"strconv"

	"github.com/wiz-freight-org/adapters/internal/config/globals"
	"github.com/wiz-freight-org/adapters/utils"
)

var (
	AppName                 string
	Port                    string
	ServiceToken            string
	StaticFileCacheDuration string

	LogLevel        string
	LogFilePathSize string

	NotificationURL string
	IDURL           string
	BookingURL      string
	DocsURL         string
	QuotesURL       string
	AssetURL        string
	ScheduleURL     string
	PaymentURL      string
	RateManagerURL  string
	WebsocketURL    string

	DatabaseURL       string
	DatabaseURLTest   string
	MigrationFilePath string
	MaxDBConn         int

	LocationURL string
	RateURL     string
	APIBaseURL  string

	OpenExchangeURL string
	OpenExchangeKey string

	DocumentURL string

	URLShortnerURL   string
	URLShortnerToken string

	BaseURL string
	PDFURL  string

	MaerskSpotURL string
	MaerskSpotKey string

	InvoiceURL string

	ThirdPartyFreightMargin float64

	GooglePlacesURL      string
	GooglePlaceDetailURL string
	GoogleKey            string

	RedisURL        string
	RedisDB         int
	RedisMaxRetries int
	RedisPoolSize   int

	ZohoBooksClientID     string
	ZohoBooksClientSecret string
	ZohoBooksRefreshToken string
	ZohoBooksBaseURL      string

	ZohoBooksCustomerContactSQSURL         string
	ZohoBooksVendorContactSQSURL           string
	ZohoBooksCustomerProFormaInvoiceSQSURL string
	ZohoBooksCustomerTaxInvoiceSQSURL      string
	ZohoBooksCustomerPaymentSQSURL         string
	ZohoBooksPartnerPaymentSQSURL          string
	ZohoBooksOrgID                         string

	LocalisationURL string

	TaskURL      string
	MilestoneURL string

	OnGridURL         string
	OnGridToken       string
	EnableZohoSandBox bool

	IfactURL  string
	RedashURL string
)

func init() {
	AppName = os.Getenv("AppName")
	Port = os.Getenv("Port")
	ServiceToken = os.Getenv("ServiceToken")
	StaticFileCacheDuration = os.Getenv("StaticFileCacheDuration")

	// Log related configuration
	LogLevel = os.Getenv("LogLevel")
	LogFilePathSize = os.Getenv("LogFilePathSize")

	// Redis config
	NotificationURL = os.Getenv("NotificationURL")
	IDURL = os.Getenv("IDURL")
	BookingURL = os.Getenv("BookingURL")
	DocsURL = os.Getenv("DocsURL")
	QuotesURL = os.Getenv("QuotesURL")
	ScheduleURL = os.Getenv("ScheduleURL")

	AssetURL = os.Getenv("AssetURL")
	PaymentURL = os.Getenv("PaymentURL")
	RateManagerURL = os.Getenv("RateManagerURL")

	LocationURL = os.Getenv("LocationURL")
	RateURL = os.Getenv("RateURL")
	APIBaseURL = os.Getenv("APIBaseURL")

	OpenExchangeURL = os.Getenv("OpenExchangeURL")
	OpenExchangeKey = os.Getenv("OpenExchangeKey")

	DocumentURL = os.Getenv("DocumentURL")

	URLShortnerURL = os.Getenv("URLShortnerURL")
	URLShortnerToken = os.Getenv("URLShortnerToken")

	BaseURL = os.Getenv("BaseURL")
	PDFURL = os.Getenv("PDFURL")

	MaerskSpotURL = os.Getenv("MaerskSpotURL")
	MaerskSpotKey = os.Getenv("MaerskSpotKey")
	ThirdPartyFreightMargin, _ = strconv.ParseFloat(os.Getenv("ThirdPartyFreightMargin"), 64)

	GooglePlacesURL = os.Getenv("GooglePlacesURL")
	GoogleKey = os.Getenv("GoogleKey")
	GooglePlaceDetailURL = os.Getenv("GooglePlaceDetailURL")

	InvoiceURL = os.Getenv("InvoiceURL")

	RedisURL = os.Getenv("RedisURL")
	RedisMaxRetries, _ = strconv.Atoi(os.Getenv("RedisMaxRetries"))
	RedisPoolSize, _ = strconv.Atoi(os.Getenv("RedisPoolSize"))
	RedisDB, _ = strconv.Atoi(os.Getenv("RedisDB"))

	DatabaseURL = os.Getenv("DatabaseURL")
	DatabaseURLTest = os.Getenv("DatabaseURLTest")
	MigrationFilePath = os.Getenv("MigrationFilePath")
	MaxDBConn, _ = strconv.Atoi(os.Getenv("MaxDBConn"))

	// Setup global logger
	config := utils.NewLoggerConfig(AppName)
	config.SetLevel(LogLevel)
	globals.Logger = utils.NewLogger(config)

	ZohoBooksClientID = os.Getenv("ZohoBooksClientID")
	ZohoBooksClientSecret = os.Getenv("ZohoBooksClientSecret")
	ZohoBooksRefreshToken = os.Getenv("ZohoBooksRefreshToken")
	ZohoBooksBaseURL = os.Getenv("ZohoBooksBaseURL")

	WebsocketURL = os.Getenv("WebsocketURL")

	ZohoBooksCustomerContactSQSURL = os.Getenv("ZohoBooksCustomerContactSQSURL")
	ZohoBooksVendorContactSQSURL = os.Getenv("ZohoBooksVendorContactSQSURL")
	ZohoBooksCustomerProFormaInvoiceSQSURL = os.Getenv("ZohoBooksCustomerProFormaInvoiceSQSURL")
	ZohoBooksCustomerTaxInvoiceSQSURL = os.Getenv("ZohoBooksCustomerTaxInvoiceSQSURL")
	ZohoBooksCustomerPaymentSQSURL = os.Getenv("ZohoBooksCustomerPaymentSQSURL")
	ZohoBooksPartnerPaymentSQSURL = os.Getenv("ZohoBooksPartnerPaymentSQSURL")
	ZohoBooksOrgID = os.Getenv("ZohoBooksOrgID")

	LocalisationURL = os.Getenv("LocalisationURL")
	TaskURL = os.Getenv("TaskURL")
	MilestoneURL = os.Getenv("MilestoneURL")

	OnGridURL = os.Getenv("OnGridURL")
	OnGridToken = os.Getenv("OnGridToken")
	EnableZohoSandBox, _ = strconv.ParseBool(os.Getenv("EnableZohoSandBox"))

	IfactURL = os.Getenv("IfactURL")
	RedashURL = os.Getenv("RedashURL")
}
