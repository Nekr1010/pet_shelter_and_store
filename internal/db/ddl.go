package db

import "pet_shelter_and_store/logger"

func InitMigrations() error {
	// 1. Таблица пользователей
	usersTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            full_name VARCHAR(255) NOT NULL,
            username VARCHAR(255) NOT NULL UNIQUE,
            password VARCHAR(255) NOT NULL,
            role VARCHAR(20) NOT NULL CHECK (role IN ('superadmin', 'shop_owner', 'admin', 'volunteer', 'user')),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(usersTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating users table: %v", err)
		return err
	}

	// 2. Таблица магазинов/приютов
	shopsTableQuery := `
        CREATE TABLE IF NOT EXISTS shops (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            address TEXT NOT NULL,
            owner_id INT REFERENCES users(id) ON DELETE CASCADE,
            is_shelter BOOLEAN DEFAULT FALSE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(shopsTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating shops table: %v", err)
		return err
	}

	// 3. Таблица товаров
	productsTableQuery := `
        CREATE TABLE IF NOT EXISTS products (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            description TEXT,
            price FLOAT NOT NULL,
            category VARCHAR(100),
            status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'reserved', 'sold')),
            shop_id INT REFERENCES shops(id) ON DELETE CASCADE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(productsTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating products table: %v", err)
		return err
	}

	// 4. Таблица животных
	animalsTableQuery := `
        CREATE TABLE IF NOT EXISTS animals (
            id SERIAL PRIMARY KEY,
            name VARCHAR(255) NOT NULL,
            animal_type VARCHAR(100) NOT NULL,
            status VARCHAR(20) DEFAULT 'available' CHECK (status IN ('available', 'reserved', 'adopted')),
            shop_id INT REFERENCES shops(id) ON DELETE CASCADE,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(animalsTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating animals table: %v", err)
		return err
	}

	// 5. Таблица заказов (товаров)
	ordersTableQuery := `
        CREATE TABLE IF NOT EXISTS orders (
            id SERIAL PRIMARY KEY,
            product_id INT REFERENCES products(id) ON DELETE SET NULL,
            user_id INT REFERENCES users(id) ON DELETE CASCADE,
            status VARCHAR(20) DEFAULT 'created' CHECK (status IN ('created', 'completed', 'canceled')),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(ordersTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating orders table: %v", err)
		return err
	}

	// 6. Таблица заявок (на животных)
	requestsTableQuery := `
        CREATE TABLE IF NOT EXISTS requests (
            id SERIAL PRIMARY KEY,
            type VARCHAR(20) NOT NULL CHECK (type IN ('adoption', 'surrender')),
            animal_id INT REFERENCES animals(id) ON DELETE SET NULL,
            user_id INT REFERENCES users(id) ON DELETE CASCADE,
            status VARCHAR(20) DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
            processed_by INT REFERENCES users(id) ON DELETE SET NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP,
            deleted_at TIMESTAMP
        );`

	if _, err := db.Exec(requestsTableQuery); err != nil {
		logger.Error.Printf("[db] InitMigrations(): error creating requests table: %v", err)
		return err
	}

	// 7. Индексы для ускорения запросов
	indexQueries := []string{
		`CREATE INDEX IF NOT EXISTS idx_products_shop_id ON products(shop_id) WHERE deleted_at IS NULL;`,
		`CREATE INDEX IF NOT EXISTS idx_animals_shop_id ON animals(shop_id) WHERE deleted_at IS NULL;`,
		`CREATE INDEX IF NOT EXISTS idx_orders_user_id ON orders(user_id) WHERE deleted_at IS NULL;`,
		`CREATE INDEX IF NOT EXISTS idx_requests_animal_id ON requests(animal_id) WHERE deleted_at IS NULL;`,
	}

	for _, query := range indexQueries {
		if _, err := db.Exec(query); err != nil {
			logger.Error.Printf("[db] InitMigrations(): error creating index: %v", err)
			return err
		}
	}

	logger.Info.Println("[db] All migrations executed successfully")
	return nil
}
