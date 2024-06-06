CREATE TABLE IF NOT EXISTS vehiculos(
  id SERIAL PRIMARY KEY,
  typ      TEXT,
  color     TEXT,
  model     TEXT,
  brand     TEXT,
  photo     TEXT,
  price     INT,
  disponibility BOOLEAN
);

INSERT INTO vehiculos (typ, color, model, brand, photo, price, disponibility) 
VALUES
('Sedan', 'Blanco', 'Corolla', 'Toyota', 'https://img.remediosdigitales.com/7f7da6/toyota-corolla_sedan-2023-1600-03/1366_2000.jpg', 20000, 1),
('SUV', 'Negro', 'CR-V', 'Honda', 'https://carnovo.com/wp-content/uploads/2021/02/honda-crv-ehev.jpg', 25000, 1),
('Camioneta', 'Azul', 'F-150', 'Ford','https://img.remediosdigitales.com/4f7ec0/ford-f-150-2018-1280-04/1366_2000.jpg', 30000, 1),
('Convertible', 'Rojo', 'Mustang', 'Ford', 'https://s1.1zoom.me/big0/230/Ford_Mustang_GT_490230.jpg', 35000, 1),
('Hatchback', 'Gris', 'Golf', 'Volkswagen', 'https://www.elcarrocolombiano.com/wp-content/uploads/2016/10/20161006-VOLKSWAGEN-GOLF-14-TSI-01A.jpg', 22000, 1),
('Coupé', 'Plateado', '911', 'Porsche', 'https://upload.wikimedia.org/wikipedia/commons/c/c1/Porsche_992_GT3_1X7A0323.jpg',100000, 1),
('SUV', 'Blanco', 'X5', 'BMW', 'https://www.elcarrocolombiano.com/wp-content/uploads/2020/03/20200327-BMW-X5-XDRIVE45E-HIBRIDA-ENCHUFABLE-COLOMBIA-PRECIO-Y-CARACTERISTICAS-01.jpg',60000, 1),
('Sedan', 'Negro', 'Accord', 'Honda', 'https://img.remediosdigitales.com/a0dc55/honda-accord-2018_/1366_2000.jpg',24000, 1),
('Minivan', 'Azul', 'Odyssey', 'Honda', 'https://acnews.blob.core.windows.net/imgnews/large/NAZ_a40e04f4c31540488cdd8b8bc5bc52f3.jpg',32000, 1),
('Camioneta', 'Rojo', 'Silverado', 'Chevrolet', 'https://www.elcarrocolombiano.com/wp-content/uploads/2023/04/20230413-CHEVROLET-SILVERADO-HD-ZR2-2024-PORTADA.jpg',28000, 1),
('Convertible', 'Amarillo', 'Camaro', 'Chevrolet', 'https://noticias.coches.com/wp-content/uploads/2013/08/Chevrolet-Camaro-Cabrio-2014-1.jpg',40000, 1),
('Sedan', 'Gris', 'Model S', 'Tesla', 'https://cdn.autobild.es/sites/navi.axelspringer.es/public/media/image/2015/08/450823-tesla-model-s-preparado-revozport.jpg',80000, 1),
('SUV', 'Negro', 'Range Rover', 'Land Rover', 'https://hips.hearstapps.com/hmg-prod/images/range-rover-velar-r-dynamic-black-edition-1-1575994075.jpg?crop=0.931xw:0.699xh;0.0335xw,0.179xh&resize=2048:*',90000, 1),
('Hatchback', 'Blanco', 'Civic', 'Honda', 'https://es.digitaltrends.com/wp-content/uploads/2019/12/y3239att.jpg?p=1',21000, 1),
('Coupé', 'Rojo', 'GT-R', 'Nissan', 'https://cdn.motor1.com/images/mgl/KqBrN/s1/nissan-gt-r-2020.webp',110000, 1),
('Sedan', 'Azul', 'Elantra', 'Hyundai', 'https://i.pinimg.com/originals/1b/22/8b/1b228bb87381550cf2b4c7f92aa61454.png',19000, 1),
('SUV', 'Plateado', 'Q5', 'Audi', 'https://www.elcarrocolombiano.com/wp-content/uploads/2020/11/Dise%C3%B1o-sin-t%C3%ADtulo-34-3.jpg',50000, 1),
('Camioneta', 'Negro', 'Tacoma', 'Toyota', 'https://www.elcarrocolombiano.com/wp-content/uploads/2023/11/20231128-TOYOTA-TACOMA-2024-PORTADA.jpg',27000, 1),
('Convertible', 'Blanco', 'Boxster', 'Porsche', 'https://www.diariomotor.com/imagenes/2016/04/porsche-718-boxster-2016-prueba-mdm-00.jpg',70000, 1),
('Sedan', 'Rojo', '3 Series', 'BMW', 'https://s1.1zoom.me/big0/98/BMW_F30_2015_3-Series_Red_Metallic_Sedan_525155_1280x853.jpg',45000, 1),
('SUV', 'Gris', 'Explorer', 'Ford', 'https://www.elcarrocolombiano.com/wp-content/uploads/2018/07/20180723-FORD-EXPLORER-ECOBOOST-2018-COLOMBIA-01.jpg',35000, 1),
('Hatchback', 'Azul', 'Focus', 'Ford', 'https://www.elcarrocolombiano.com/wp-content/uploads/2019/02/20190219-FORD-FOCUS-ST-2019-01.jpg',20000, 1),
('Coupé', 'Negro', 'Challenger', 'Dodge', 'https://noticias.coches.com/wp-content/uploads/2023/06/Dodge-Challenger-Black-Ghost-2023-1.jpeg',55000, 1),
('Sedan', 'Blanco', 'A4', 'Audi', 'https://i.pinimg.com/originals/5c/7e/9d/5c7e9d53460217543a0ca6c87feacc85.jpg',39000, 1),
('Minivan', 'Negro', 'Sienna', 'Toyota', 'https://img.supercarros.com/AdsPhotos/1024x768/0/11749825.jpg?wmo=d107654c193750e83d149feb0bf0428adc3155d10a6e939125385f8bcd4b92376e14dafb117d96011d1a8c539443313c402e189cfc521b69ebd280f2981bc384',34000, 1),
('Camioneta', 'Azul', 'Ranger', 'Ford', 'https://www.elcarrocolombiano.com/wp-content/uploads/2019/09/20190923-FORD-RANGER-2020-COLOMBIA-PRECIOS-VERSIONES-CARACTERISTICAS-01.jpg',25000, 1),
('SUV', 'Rojo', 'CX-5', 'Mazda', 'https://www.motor.com.co/__export/1649991314913/sites/motor/img/2022/04/14/20220414_085514641_5996019d8c4b5_r_1503009443280_384-0-1464-540.jpeg_554688468.jpeg',28000, 1),
('Sedan', 'Gris', 'Altima', 'Nissan', 'https://noticias.coches.com/wp-content/uploads/2014/07/nissan_altima-hybrid-2010_r19.jpg',23000, 1),
('Convertible', 'Plateado', 'SLC', 'Mercedes-Benz', 'https://automercol.com.co/wp-content/uploads/2020/11/Colores_SLC.jpg',60000, 1),
('Coupé', 'Blanco', 'LC', 'Lexus', 'https://hips.hearstapps.com/hmg-prod/images/1-final-1655192842.jpg?crop=0.7504145936981759xw:1xh;center,top&resize=1200:*',90000, 1),
('SUV', 'Negro', 'Grand Cherokee', 'Jeep', 'https://img.remediosdigitales.com/03730d/jeep-grand-cherokee-s-2018--2-/1366_2000.jpg',40000, 1),
('Sedan', 'Azul', 'Optima', 'Kia', 'https://www.rutamotor.com/wp-content/uploads/2019/12/2021-Kia-K5-Korean-spec-2.jpg',22000, 1),
('Hatchback', 'Rojo', 'Mazda3', 'Mazda', 'https://upload.wikimedia.org/wikipedia/commons/8/88/2019_Mazda3_SE-L_2.0_Front.jpg',24000, 1),
('Camioneta', 'Gris', 'Colorado', 'Chevrolet', 'https://elcarrocolombiano.com/wp-content/uploads/2017/12/20171212-CHEVROLET-COLORADO-2018-COLOMBIA-02.jpg',26000, 1),
('Convertible', 'Amarillo', 'Beetle', 'Volkswagen', 'https://estaticos-cdn.prensaiberica.es/clip/5a1986f1-2f34-4013-8227-3b86740b49e2_source-aspect-ratio_default_0.jpg',30000, 1),
('Sedan', 'Blanco', 'Impreza', 'Subaru', 'https://www.rutamotor.com/wp-content/uploads/2016/01/Subaru-WRX-STi-2016-Test-Drive-12.jpg',21000, 1);

CREATE TABLE IF NOT EXISTS usuarios (
    id SERIAL PRIMARY KEY,
    nombre_usuario VARCHAR(50) NOT NULL UNIQUE,
    contrasena VARCHAR(255) NOT NULL,
    correo_electronico VARCHAR(100) NOT NULL UNIQUE
);

INSERT INTO usuarios (nombre_usuario, contrasena, correo_electronico) 
VALUES
('admin', 'admin', 'admin@outlook.com');
('admin', 'admin', 'admin@outlook.es');