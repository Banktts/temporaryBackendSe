CREATE TABLE restaurant (
	R_ID int  NOT NULL PRIMARY KEY ,
	R_name varchar(30) NOT NULL,
	R_rating  FLOAT(2) NOT NULL,
	R_votes int NOT NULL,
	R_latitude FLOAT(6) NOT NULL ,
	R_longitude FLOAT(6 ) NOT NULL,
	R_isRecomend BOOLEAN NOT NULL,
	R_image_url varchar(255) NOT NULL 

);

CREATE TABLE menu (
	M_ID int NOT NULL PRIMARY KEY,
	R_ID int NOT NULL,
	M_name varchar(25)NOT NULL,
	M_price int NOT NULL,
	M_type varchar(6) NOT NULL,
	M_image_url varchar(255) NOT NULL

);
CREATE TABLE customer (
	C_ID int   NOT NULL PRIMARY KEY,
	C_latitude  FLOAT(6) NOT NULL,
	C_longitude  FLOAT(6) NOT NULL
);

CREATE TABLE ordert (
	O_ID int NOT NULL AUTO_INCREMENT,
	C_ID int NOT NULL,
	R_ID int NOT NULL,
	D_ID int NOT NULL,
	created_at TIMESTAMP,
	primary key(O_ID)
);
CREATE TABLE orderline (
	O_ID int NOT NULL,
	M_ID int NOT NULL,
	E_ID int NOT NULL,
	amount int NOT NULL,
	special_inst varchar(255) NOT NULL,
	primary key(O_ID, M_ID, E_ID)
);

CREATE TABLE delivery_man (
	D_ID int   NOT NULL PRIMARY KEY AUTO_INCREMENT,
	D_name varchar(20)NOT NULL,
	D_phone varchar(15) NOT NULL,
	D_rating float NOT NULL,
	D_latitude  FLOAT(6) NOT NULL,
	D_longitude  FLOAT(6) NOT NULL
);

ALTER TABLE menu
ADD FOREIGN KEY (R_ID) REFERENCES restaurant(R_ID);
ALTER TABLE orderline ADD FOREIGN KEY (O_ID) REFERENCES ordert(O_ID);
ALTER TABLE ordert ADD FOREIGN KEY (D_ID) REFERENCES delivery_man(D_ID);