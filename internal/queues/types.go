package queues

const (
	TypeScraperProductQueue = "scraper_product"
	TypeImageProcessorQueue = "image_processor"
	TypeImageUploaderQueue  = "image_uploader"
)

const (
	Concurrency               = 100
	ConcurrencyScraperProduct = 30
	ConcurrencyImageProcessor = 60
	ConcurrencyImageUploader  = 10
)
