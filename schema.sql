-- clean db
DROP TABLE IF EXISTS cars;

-- make table
CREATE TABLE cars (
  id TEXT PRIMARY KEY,
  name TEXT NOT NULL,
  class TEXT NOT NULL,
  star INTEGER NOT NULL
);

-- insert data
-- D class
INSERT INTO cars (id, name, class, star) VALUES
('lancer', 'MITSIBISHI LANCER EVOLUTION X', 'D', 3),
('bmwz4', 'BMW Z4 LCI E89', 'D', 3),
('camaro_lt', 'CHEVROLET CAMARO LT', 'D', 3);

-- C class
INSERT INTO cars (id, name, class, star) VALUES
('srt8', 'DODGE CHALLENGER SRT8', 'C', 3),
('hommage', 'BMW 3.0 CSL HOMMAGE', 'C', 3),
('boxster', 'PORSCHE BOXSTER 25TH', 'C', 3);

-- B class
INSERT INTO cars (id, name, class, star) VALUES
('gts_coupe', 'PORSCHE 911 GTS COUPE', 'B', 3),
('db11', 'ASTON MARTIN DB11', 'B', 3),
('ftype_svr', 'JAGUAR F-TYPE SVR', 'B', 4);

-- A class
INSERT INTO cars (id, name, class, star) VALUES
('vulcan', 'ASTON MARTIN VULCAN', 'A', 4),
('gtr_nismo', 'NISSAN GT-R NISMO', 'A', 4),
('nio', 'NIO EP9', 'A', 4);

-- S class
INSERT INTO cars (id, name, class, star) VALUES
('centenario', 'LAMBORGHINI CENTENARIO', 'S', 5),
('fxxk', 'FERRARI FXX K', 'S', 5),
('autentica', 'LAMBORGHINI AUTENTICA', 'S', 5);

