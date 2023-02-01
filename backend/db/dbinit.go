package db

import "log"

func DBInit() {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS Brands (
			brandID INT UNSIGNED AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY(brandID)
		);`,

		`CREATE TABLE IF NOT EXISTS Suppliers (
			supplierID INT UNSIGNED AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY(supplierID)
		);`,

		`CREATE TABLE IF NOT EXISTS PurchaseTypes (
			purchaseTypeID INT UNSIGNED AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			PRIMARY KEY(purchaseTypeID)
		);`,

		`CREATE TABLE IF NOT EXISTS Records (
			recordID INT UNSIGNED AUTO_INCREMENT,
			description VARCHAR(255) NOT NULL,
			brandID INT UNSIGNED,
			supplierID INT UNSIGNED,
			purchaseTypeID INT UNSIGNED,		 
			price DECIMAL(65, 2),
			currency VARCHAR(3),
			recordDate DATETIME,
			createdAt DATETIME,
			
			PRIMARY KEY(recordID),
			CONSTRAINT fk_brand
				FOREIGN KEY(brandID)
					REFERENCES Brands(brandID)
					ON DELETE CASCADE,
			CONSTRAINT fk_supplier
				FOREIGN KEY(supplierID) 
					REFERENCES Suppliers(supplierID)
					ON DELETE CASCADE,
			CONSTRAINT fk_purchaseTypeID
				FOREIGN KEY(purchaseTypeID)
					REFERENCES PurchaseTypes(purchaseTypeID)
					ON DELETE CASCADE
		);`,
	}

	for _, query := range queries {
		// Executes the SQL query in our database. Check err to ensure there was no error.
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Error creating table: %v", err)
		}
	}

}
