insert into Brands (name) values ("Brand1"), ("Brand2");
insert into Suppliers (name) values ("Supplier1"), ("Supplier2");
insert into PurchaseTypes (name) values ("PurchaseType1"), ("PurchaseType2");

insert into Records (description, brandID, supplierID, purchaseTypeID, price, currency, recordDate, createdAt) 
    values ("Desc1", 1, 1, 1, 12.12, "USD", CURRENT_DATE(),  CURRENT_DATE()),
    ("Desc2", 2, 2, 2, 22.22, "USD", CURRENT_DATE(),  CURRENT_DATE());
