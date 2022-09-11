CREATE DATABASE IF NOT EXISTS paging CHARACTER SET 'utf8';

CREATE TABLE IF NOT EXISTS paging.computer (
	id int NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
	PRIMARY KEY (id)
) CHARACTER SET utf8
  COLLATE utf8mb3_general_ci;

-- copied from Wikipedia section on computer models
INSERT INTO paging.computer (name) VALUES
    ('Commodore 64'),
    ('Altair 8800'),
    ('Apple I and also Replica 1'),
    ('Applix 1616'),
    ('Compukit UK101'),
    ('Dick Smith Super-80 Computer'),
    ('Educ-8 non-microprocessor kit computer'),
    ('Elektor Junior Computer'),
    ('Elektor TV Games Computer'),
    ('Ferguson Big Board'),
    ('Galaksija'),
    ('Heathkit H8 and relations'),
    ('Heathkit H11'),
    ('Heath ET-100 8088 trainer'),
    ('Kenbak-1'),
    ('KIM-1'),
    ('LNW-80'),
    ('MK14'),
    ('Mark-8'),
    ('Micro-Professor MPF-I'),
    ('Nascom 1 and Nascom 2'),
    ('Newbear 77-68'),
    ('Processor Technology SOL 20'),
    ('PSI Comp 80 (computer)'),
    ('SCELBI'),
    ('Sinclair ZX80 kit'),
    ('Tangerine MICROTAN 65'),
    ('TEC-1'),
    ('Wave Mate Bullet');