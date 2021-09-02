CREATE DATABASE flickr;

CREATE TABLE flickr.photos (
	id INT NOT NULL AUTO_INCREMENT,
	filename VARCHAR(255) NOT NULL,
	flickr_id VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
);

CREATE TABLE flickr.albums (
	id INT NOT NULL AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL,
	PRIMARY KEY (id)
)

CREATE TABLE flickr.album_photos (
	id INT NOT NULL AUTO_INCREMENT,
	photo_id INT NOT NULL,
	album_id INT NOT NULL,
	INDEX ph_id (photo_id),
	INDEX al_id (album_id),
	PRIMARY KEY (id)
	-- FOREIGN KEY (photo_id)
	-- 	REFERENCES flickr.photos(id)
	-- 	ON DELETE CASCADE,
	-- FOREIGN KEY (album_id)
	-- 	REFERENCES flickr.albums(id)
	-- 	ON DELETE CASCADE
)