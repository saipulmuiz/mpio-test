generate-mocks:
	# repositories
	@mockgen -destination=./service/repository/mocks/mock_user_repository.go -package=mocks github.com/saipulmuiz/mpio-test/api UserRepository
	@mockgen -destination=./service/repository/mocks/mock_product_repository.go -package=mocks github.com/saipulmuiz/mpio-test/api ProductRepository
	@mockgen -destination=./service/repository/mocks/mock_category_repository.go -package=mocks github.com/saipulmuiz/mpio-test/api CategoryRepository
	@mockgen -destination=./service/repository/mocks/mock_product_category_repository.go -package=mocks github.com/saipulmuiz/mpio-test/api ProductCategoryRepository
	@mockgen -destination=./service/repository/mocks/mock_cart_repository.go -package=mocks github.com/saipulmuiz/mpio-test/api CartRepository