-- Table: user

-- DROP TABLE user;
DROP SEQUENCE IF EXISTS entusers_id_seq;
CREATE SEQUENCE entusers_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;
  
CREATE TABLE IF NOT EXISTS entusers
(
    id integer NOT NULL DEFAULT nextval('entusers_id_seq'::regclass),
    username text COLLATE pg_catalog."default",
    password text COLLATE pg_catalog."default",
    firstname text COLLATE pg_catalog."default",
    lastname text COLLATE pg_catalog."default",
    rolename text COLLATE pg_catalog."default",
    jwttoken text COLLATE pg_catalog."default",
    CONSTRAINT entusers_pkey PRIMARY KEY (id),
    CONSTRAINT entusers_username_key UNIQUE (username)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE entusers
    OWNER to postgres;
	
	
	
-- Create Sensor TABLE
-- Table: entsensors

-- DROP TABLE entsensors;
DROP SEQUENCE IF EXISTS entsensors_id_seq;
CREATE SEQUENCE entsensors_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;
  
CREATE TABLE IF NOT EXISTS entsensors
(
    id bigint NOT NULL DEFAULT nextval('entsensors_id_seq'::regclass),
    userid integer,
    sensorname text COLLATE pg_catalog."default",
    sensortype text COLLATE pg_catalog."default",
    sensormode text COLLATE pg_catalog."default",
    location text COLLATE pg_catalog."default",
    sensorowner text COLLATE pg_catalog."default",
    sensorauth text COLLATE pg_catalog."default",
    sensorproto text COLLATE pg_catalog."default",
    localurl text COLLATE pg_catalog."default",
    securelevel text COLLATE pg_catalog."default",
    sensorstate integer,
    sensorpid integer,
--    user jsonb,
    CONSTRAINT entsensors_pkey PRIMARY KEY (id)
--    CONSTRAINT sensor_userid_s_fkey FOREIGN KEY (userid)
--        REFERENCES user (id) MATCH SIMPLE
--        ON UPDATE NO ACTION
--        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE entsensors
    OWNER to postgres;
	
	
	
-- Create Video TABLE
-- Table: videoinfo

-- DROP TABLE videoinfo;
DROP SEQUENCE IF EXISTS entvideos_id_seq;
CREATE SEQUENCE entvideos_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;
  
CREATE TABLE entvideos
(
    id bigint NOT NULL DEFAULT nextval('entvideos_id_seq'::regclass),
    videoname text COLLATE pg_catalog."default",
    videotype text COLLATE pg_catalog."default",
    sensorid integer,
    userid integer,
    starttime text COLLATE pg_catalog."default",
    endtime text COLLATE pg_catalog."default",
    inputurl text COLLATE pg_catalog."default",
    filepath text COLLATE pg_catalog."default",
    CONSTRAINT entvideos_pkey PRIMARY KEY (id)
   -- CONSTRAINT entvideos_sensorsid_fkey FOREIGN KEY (sensorid)
   --     REFERENCES entsensors (id) MATCH SIMPLE
   --     ON UPDATE NO ACTION
   --     ON DELETE NO ACTION,
   -- CONSTRAINT entvideos_userid_fkey FOREIGN KEY (userid)
   --     REFERENCES entusers (id) MATCH SIMPLE
   --     ON UPDATE NO ACTION
   --     ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE entvideos
    OWNER to postgres;	

-- Create Video TABLE
-- Table: entchunks

-- DROP TABLE entchunks;
DROP SEQUENCE IF EXISTS entchunks_id_seq;
CREATE SEQUENCE entchunks_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;

CREATE TABLE entchunks
(
    id bigint NOT NULL DEFAULT nextval('entchunks_id_seq'::regclass),
    chunkname text COLLATE pg_catalog."default",
    chunktype text COLLATE pg_catalog."default",
    sensorid integer,
    userid integer,
    ownername text COLLATE pg_catalog."default",
    filepath text COLLATE pg_catalog."default",
    chunkstate integer,
    chunkpid integer,
    -- users jsonb,
    CONSTRAINT entchunks_pkey PRIMARY KEY (id)
   -- CONSTRAINT entchunks_sensorsid_fkey FOREIGN KEY (sensorid)
   --     REFERENCES entsensors (id) MATCH SIMPLE
   --     ON UPDATE NO ACTION
   --     ON DELETE NO ACTION,
   -- CONSTRAINT entchunks_userid_fkey FOREIGN KEY (userid)
   --     REFERENCES entusers (id) MATCH SIMPLE
   --     ON UPDATE NO ACTION
   --     ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE entchunks
    OWNER to postgres;	



-- Create Video TABLE
-- Table: entchunks

-- DROP TABLE entchunks;
DROP SEQUENCE IF EXISTS entproducts_id_seq;
CREATE SEQUENCE entproducts_id_seq
  INCREMENT 1
  MINVALUE 1
  MAXVALUE 9223372036854775807
  START 1
  CACHE 1;

CREATE TABLE entproducts
(
    id bigint NOT NULL DEFAULT nextval('entproducts_id_seq'::regclass),
    productname text COLLATE pg_catalog."default",
    producttype text COLLATE pg_catalog."default",
    sensorid integer,
    userid integer,
    ownername text COLLATE pg_catalog."default",
    filepath text COLLATE pg_catalog."default",
    productstate integer,
    productpid integer,
    CONSTRAINT entproducts_pkey PRIMARY KEY (id)
  --  CONSTRAINT entproducts_sensorsid_fkey FOREIGN KEY (sensorid)
  --      REFERENCES entsensors (id) MATCH SIMPLE
  --      ON UPDATE NO ACTION
  --      ON DELETE NO ACTION,
  --  CONSTRAINT entproducts_userid_fkey FOREIGN KEY (userid)
  --      REFERENCES entusers (id) MATCH SIMPLE
  --      ON UPDATE NO ACTION
  --      ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE entproducts
    OWNER to postgres;	

-- Table: sensorgrantinfo

-- DROP TABLE sensorgrantinfo;
--CREATE SEQUENCE IF NOT EXISTS sen_user_item_userid_seq
--  INCREMENT 1
--  MINVALUE 1
--  MAXVALUE 9223372036854775807
--  START 1
--  CACHE 1;
  
CREATE TABLE sensor_to_users
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    sensor_id integer,
    user_id integer
--    CONSTRAINT sensoruser_userid_fkey FOREIGN KEY (userid)
--        REFERENCES entusers (id) MATCH SIMPLE
--        ON UPDATE NO ACTION
--        ON DELETE NO ACTION,
--    CONSTRAINT sensoruser_sensorid_fkey FOREIGN KEY (sensorid)
--        REFERENCES entsensors (id) MATCH SIMPLE
--        ON UPDATE NO ACTION
--        ON DELETE NO ACTION
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE sensor_to_users
    OWNER to postgres;	

-- CREATE TABLE sensor_to_users
CREATE TABLE user_to_sensors
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    user_id integer,
    sensor_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE user_to_sensors
    OWNER to postgres;	

-- CREATE TABLE sensor_to_users
CREATE TABLE user_to_videos
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    user_id integer,
    video_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE user_to_videos
    OWNER to postgres;	



-- CREATE TABLE sensor_to_users
CREATE TABLE video_to_users
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    video_id integer,
    user_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE video_to_users
    OWNER to postgres;	



-- CREATE TABLE sensor_to_users
CREATE TABLE product_to_users
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    product_id integer,
    user_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE product_to_users
    OWNER to postgres;	

-- user_to_products
-- CREATE TABLE sensor_to_users
CREATE TABLE user_to_products
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('sen_user_item_userid_seq'::regclass),
    user_id integer,
    product_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE user_to_products
    OWNER to postgres;	

-- chunk_to_videos
CREATE TABLE chunk_to_videos
(
    chunk_id integer,
    video_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE chunk_to_videos
    OWNER to postgres;	


-- chunk_to_products
-- CREATE TABLE chunk_to_products
CREATE TABLE chunk_to_products
(
    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('chunk_to_products_userid_seq'::regclass),
    chunk_id integer,
    product_id integer
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE chunk_to_videos
    OWNER to postgres;	

-- user_to_chunks
-- CREATE TABLE user_to_chunks
--CREATE TABLE user_to_chunks
--(
--    -- No Primary Key for relation -- id bigint NOT NULL DEFAULT nextval('chunk_to_products_userid_seq'::regclass),
--    user_id integer,
--    chunk_id integer
--)
--WITH (
--    OIDS = FALSE
--)
--TABLESPACE pg_default;
--
--ALTER TABLE user_to_chunks
--    OWNER to postgres;	
--
    
INSERT INTO entusers (username, password, firstname, lastname, rolename, jwttoken) VALUES
    ('Admin1', '123456', 'Admin', 'One', 'Admin', 'Token'),
    ('AdminUser', '123456', 'John', 'Doe', 'Admin', 'Token'),
    ('Operator1', '123456', 'Operator', 'One', 'Operator', 'Token'),
    ('OperatorUser', '123456', 'Jean', 'Smith', 'Operator', 'Token'),
    ('Reviewer1', '123456', 'Miken', 'One', 'Reviewer', 'Token'),
    ('ReviewUser', '123456', 'Mike', 'Johnson', 'Reviewer', 'Token'),
    ('Creator1', '123456', 'Kim', 'Creator', 'One', 'Token'),
    ('ProdUser', '123456', 'Kan', 'Hann', 'Creator', 'Token');

	
INSERT INTO entsensors (userid, sensorname, sensortype, sensormode, location, sensorowner, sensorproto, sensorauth, localurl, securelevel, sensorstate, sensorpid) VALUES  -- , sensorstate, sensorpid
	
    (2, 'HP TrueVision HD','Webcam',    'Live',  'Bi7',  'Albert', 'rtsp', 'admin:admin', 'n/a', 'high',1,1),   --  , 1, 99999999
    (2, 'HP TrueVision HD', 'Webcam',   'File',  'Bi7',  'Mike',   'rtsp', 'admin:admin', 'n/a', 'medium',1,1),   --  , 1, 99999999
    (3, 'HP TrueVision HD',  'Webcam',  'Chunk', 'Bi7',  'Mike',   'rtsp', 'admin:admin', 'n/a', 'low',1,1),   --  , 1, 99999999
	
    (2, 'Integrated Webcam','Webcam',   'Live',  'VIT',  'Albert', 'rtsp', 'admin:admin', 'n/a', 'high',1,1),   --  , 1, 99999999
    (2, 'Integrated Webcam', 'Webcam',  'File',  'VIT',  'Mike',   'rtsp', 'admin:admin', 'n/a', 'medium',1,1),   --  , 1, 99999999
    (3, 'Integrated Webcam',  'Webcam', 'Chunk', 'VIT',  'Mike',   'rtsp', 'admin:admin', 'n/a', 'low',1,1),   --  , 1, 99999999
	
    (1, '3EA1', 'Avigilon', 'Live',  'yard2',    'Peter',  'rtsp', 'admin:admin', '192.168.0,2', 'high',1,1),  -- , 1, 99999999
    (2, '3EA1', 'Avigilon', 'File',  'living',   'Mike',   'rtsp', 'admin:admin', '192.168.0.2', 'medium',1,1),  -- , 1, 99999999
    (2, '3EA1', 'Avigilon', 'Chunk', 'Tampa',    'Tim',    'rtsp', 'admin:admin', '192.168.0.2', 'low',1,1),  --, 1, 99999999

    (2, 'A5B3','SV3C',     'Live',  'Clearwater', 'Albert', 'rtsp', 'admin:admin', '192.168.0.21', 'high',1,1),   --  , 1, 99999999
    (2, 'A5B3','SV3C',     'File',  'Tarpon',     'Mike',   'rtsp', 'admin:admin', '192.168.0.21', 'medium',1,1),   --  , 1, 99999999
    (3, 'A5B3','SV3C',     'Chunk', 'Largo',      'Mike',   'rtsp', 'admin:admin', '192.168.0.21', 'low',1,1),   --  , 1, 99999999
    
    (2, 'Integrated Webcam','Webcam',    'Live',  'JCR', 'Albert', 'rtsp', 'admin:admin', 'n/a', 'high',1,1),   --  , 1, 99999999
    (2, 'Integrated Webcam', 'Webcam',   'File',  'JCR', 'Mike',   'rtsp', 'admin:admin', 'n/a', 'medium',1,1),   --  , 1, 99999999
    (3, 'Integrated Webcam',  'Webcam',  'Chunk', 'JCR', 'Mike',   'rtsp', 'admin:admin', 'n/a', 'low',1,1),   --  , 1, 99999999
	
    (2, 'GreenSensor','SV3C',  'Live',  'Clearwater', 'Albert', 'rtsp', 'admin:admin', '192.168.1.11', 'high',1,1),   --  , 1, 99999999
    (2, 'BlueSensor', 'SV3C',  'File',  'Tarpon',     'Mike',   'rtsp', 'admin:admin', '192.168.1.11', 'medium',1,1),   --  , 1, 99999999
    (3, 'RedSensor',  'SV3C',  'Chunk', 'Largo',      'Mike',   'rtsp', 'admin:admin', '192.168.1.11', 'low',1,1);   --  , 1, 99999999

    --(1, 'AnSensor',   'Avigilon', 'Live',  'yard2',      'Peter',  'rtsp', 'admin:admin', '192.168.0.2', 'high',23, 2),  -- , 1, 99999999
    --(2, 'TonySensor', 'Avigilon', 'File',  'living',     'Mike',   'rtsp', 'admin:admin', '192.168.0.2', 'medium',23, 2),  -- , 1, 99999999
    --(2, 'TimSensor',  'Avigilon', 'Chunk', 'Tampa',      'Tim',    'rtsp', 'admin:admin', '192.168.0.2', 'low',23, 2),  --, 1, 99999999
    --(2, 'Adm Green','SV3C',     'Live',  'Clearwater', 'Albert', 'rtsp', 'admin:admin', '192.168.0.14', 'high',23, 2),   --  , 1, 99999999
    --(2, 'Adm Blue', 'SV3C',     'File',  'Tarpon',     'Mike',   'rtsp', 'admin:admin', '192.168.0.14', 'medium',23, 2),   --  , 1, 99999999
    --(3, 'Adm Red',  'SV3C',     'Chunk', 'Largo',      'Mike',   'rtsp', 'admin:admin', '192.168.0.14', 'low',23, 2);   --  , 1, 99999999
INSERT INTO entvideos ( videoname, videotype, sensorid, userid, starttime, endtime, inputurl, filepath) VALUES
    ('Melonga Dancers', 'Body Cam1', 1,1, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.10', 'Melonga_160x120.mp4'),
    ('Tango Dancers', 'Body Cam2', 2, 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.11', 'Tango_160x120.mp4'),
    ('Peter', 'Body Cam3', 2, 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.18', 'Melonga_160x120.mp4'),
    ('Paul', 'Body Cam4', 2, 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.19', 'Tango_160x120.mp4');
    --('Mike', 'Body Cam5', '3', 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.20', 'Melonga_160x120.mp4'),
    --('Sue', 'Body Cam6', '4', 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.21', 'street_160x120.mp4'),
    --('Ann', 'Body Cam7', '4', 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.22', 'corner_160x120.mp4'),
    --('Jennifer', 'Body Cam8', '5', 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.23', 'beach_160x120.mp4'),
    --('Alice', 'Body Cam9', '5', 2, '2018/07/24 11:18:49', '2018/07/24 12:18:49 ', '198.162.0.24', 'house_160x120.mp4');
--

-- entchunks
-- id 
-- chunkname
-- chunktype
-- sensorid 
-- userid 
-- ownername
-- filepath 
-- chunkstate
-- chunkpid
-- entchunks
INSERT INTO entchunks (chunkname, chunktype, sensorid, userid, ownername, filepath, chunkstate, chunkpid) VALUES
    ('Chunkone','SV3C', 1, 2, 'SomeOwner', 'Melonga_160x120.mp4', 1, 1234),
--    ('chunk two','Avigilon', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','CudaEye', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','SV3C', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','Avigilon', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three', 'CudaEye', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','SV3C', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','AVigilon', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
--    ('chunk three','CUDAEYE', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234);


-- id 
-- productnam
-- producttyp
-- sensorid
-- userid
-- ownername 
-- filepath
-- productsta
-- productpid
-- entchunks
INSERT INTO entproducts (productname, producttype, sensorid, userid, ownername, filepath, productstate, productpid) VALUES
    ('product one','SV3C', 1, 2,'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product two','SV3C', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product four','CudaEye', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product five','SV3C', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product six','Avigilon', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product seven', 'CudaEye', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product eight','SV3C', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product nine','AVigilon', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234),
    ('product ten','CUDAEYE', 1, 2, 'SomeOwner', 'assets/Products/dir1/', 1, 1234);

INSERT INTO sensor_to_users (sensor_id, user_id) VALUES
    (1,1),
    (2,2),
    (3,2),
    (4,2),
    (5,2),
    (6,3),
    (1,4),
    (5,5),
    (1,6),
    (6,6),
    (1,7),
    (5,7),
    (1,8),
    (2,8),
    (5,8),
    (6,8);

--INSERT INTO userowner_to_sensors (user_id, sensor_id) VALUES
--    (1, 1),
--    (2, 2),
--    (2, 3),
--    (2, 4),
--    (2, 5),
--    (8, 6);

INSERT INTO user_to_sensors (user_id, sensor_id) VALUES
    (1, 1),
    (2, 2),
    (2, 3),
    (2, 4),
 --   (2, 5),
    (8, 6);

INSERT INTO user_to_videos (user_id, video_id) VALUES
    (1, 1),
    (2, 2),
    (2, 3),
    (2, 4),
    (2, 5),
    (3, 6),
    (4, 6),
    (5, 7),
    (6, 7),
    (7, 7),
    (8, 8),
    (8, 9),
    (8, 1),
    (8, 2);

INSERT INTO video_to_users (video_id, user_id) VALUES
    (1, 1),
    (2, 2),
    (2, 3),
    (2, 4),
    (2, 5),
    (3, 6),
    (4, 6),
    (5, 7),
    (6, 7),
    (7, 7),
    (8, 8),
    (9, 8),
    (9, 1);

--INSERT INTO uservideo (user_id, sensor_id, video_id) VALUES
--    (1, 1, 1),
--    (2, 2, 2),
--    (2, 2, 3),
--    (2, 2, 4),
--    (2, 3, 5),
--    (2, 4, 6),
--    (2, 4, 7),
--    (2, 5, 8),
--    (2, 5, 9);

INSERT INTO product_to_users (product_id, user_id) VALUES
    (1, 1),
    (2, 2),
    (2, 3),
    (2, 4),
    (2, 5),
    (3, 6),
    (4, 6),
    (5, 7),
    (6, 7),
    (1, 7),
    (2, 8),
    (3, 8),
    (6, 1);

INSERT INTO user_to_products (user_id, product_id) VALUES
    (1,1),
    (2,2),
    (3,2),
    (4,2),
    (5,2),
    (6,3),
    (2,4),
    (7,5),
    (7,6),
    (7,1),
    (8,2),
    (8,3),
    (1,6);


INSERT INTO chunk_to_videos (chunk_id, video_id) VALUES
    (1,1),
    (2,2),
    (3,3),
    (4,4),
    (1,5),
    (2,6),
    (3,7),
    (4,8),
    (1,9),
    (2,1),
    (3,2),
    (4,3);


INSERT INTO chunk_to_products (chunk_id, product_id) VALUES
    (1,1),
    (2,2),
    (3,3),
    (3,4),
    (3,5),
    (3,6),
    (3,7),
    (8,8),
    (9,9);

--INSERT INTO user_to_chunks (user_id, chunk_id) VALUES
--    (1,1),
--    (2,2),
--    (2,3),
--    (2,4),
--    (3,5),
--    (3,6),
--    (3,7),
--    (4,8),
--    (4,9);
--