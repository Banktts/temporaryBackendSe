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
	created_at TIMESTAMP
);

ALTER TABLE menu
ADD FOREIGN KEY (R_ID) REFERENCES restaurant(R_ID);
