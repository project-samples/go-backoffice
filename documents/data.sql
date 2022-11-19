create table modules (
  moduleid varchar(40) primary key,
  modulename varchar(255) not null,
  status char(1) not null,
  path varchar(255),
  resourcekey varchar(255),
  icon varchar(255),
  sequence int not null,
  parent varchar(40),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz
);

create table entities (
  entityid varchar(40) primary key,
  entityname varchar(255) not null,
  email varchar(255) not null,
  displayname varchar(255) not null,
  status char(1) not null,
  phone varchar(20),
  imageurl varchar(500),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz
);

create table users (
  userid varchar(40) primary key,
  username varchar(255) not null,
  email varchar(255) not null,
  displayname varchar(255) not null,
  status char(1) not null,
  gender char(1),
  phone varchar(20),
  title varchar(10),
  position varchar(40),
  imageurl varchar(500),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz,
  lastlogin timestamptz
);

create table roles (
  roleid varchar(40) primary key,
  rolename varchar(255) not null,
  status char(1) not null,
  remark varchar(255),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamptz
);

create table entityusers (
  entityid varchar(40) not null,
  userid varchar(40) not null,
  primary key (entityid, userid)
);

create table userroles (
  userid varchar(40) not null,
  roleid varchar(40) not null,
  primary key (userid, roleid)
);

create table rolemodules (
  roleid varchar(40) not null,
  moduleid varchar(40) not null,
  permissions int not null,
  primary key (roleid, moduleid)
);

create table auditlog (
  id varchar(255) primary key,
  resource varchar(255),
  userid varchar(255),
  ip varchar(255),
  action varchar(255),
  time timestamptz,
  status varchar(255),
  remark varchar(255)
);

INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('6xydt3Qap', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-02 21:00:06.811', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gRAIVh1tM', 'term', '00005', '188.239.138.226', 'patch', '2023-07-03 12:09:51.659', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('d8sQRO1ap', 'entity', '00005', '188.239.138.226', 'patch', '2023-07-03 13:04:20.950', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gMu1Rh1aM', 'entity', '00005', '188.239.138.226', 'patch', '2023-07-03 13:04:24.491', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('jrFkzsQaM', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-03 16:00:42.627', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('4lVacs1aM', 'authentication', '00001', '::1', 'authenticate', '2023-07-03 16:22:13.157', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('a8Y-cbQtM', 'product', '00001', '95.194.49.166', 'patch', '2023-07-03 16:22:23.430', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Wvc4Us1aM', 'term', '00001', '95.194.49.166', 'patch', '2023-07-03 20:43:31.757', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('tcztIsQap', 'term', '00001', '::1', 'create', '2023-07-03 20:44:02.086', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('dO7zIb1ap', 'entity', '00001', '::1', 'patch', '2023-07-03 20:44:47.349', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('K-KcIbQtp', 'company', '00001', '::1', 'patch', '2023-07-03 20:45:55.702', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('G5JcIsQap', 'company', '00001', '::1', 'patch', '2023-07-03 20:45:59.129', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('HaLnIb1tM', 'company', '00001', '::1', 'patch', '2023-07-03 20:46:02.818', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('h_kcUbQap', 'company', '00001', '219.62.20.91', 'patch', '2023-07-03 20:46:05.519', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('jpTZIbQtM', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:07.779', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('UH_ZUsQtp', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:32.408', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('wP1SUsQtp', 'company', '00001', '70.182.126.53', 'patch', '2023-07-03 20:46:34.747', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('BxYPUb1aM', 'role', '00001', '::1', 'patch', '2023-07-03 20:46:42.944', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('rjegUs1tM', 'role', '00001', '::1', 'patch', '2023-07-03 20:47:02.120', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('lbmgUbQtM', 'role', '00001', '::1', 'patch', '2023-07-03 20:47:09.713', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('5o7-JsQap', 'role', '00001', '::1', 'patch', '2023-07-03 21:02:15.442', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('6eTFGbQap', 'role', '00001', '::1', 'patch', '2023-07-03 21:05:48.155', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('14S3JsQaM', 'role', '00001', '::1', 'patch', '2023-07-03 21:05:55.771', 'fail', 'pq: duplicate key text violates unique constraint "rolemodules_pkey"');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('DOYhJb1tp', 'article', '00001', '::1', 'patch', '2023-07-03 21:06:22.692', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('gKzOGs1tp', 'article', '00001', '::1', 'patch', '2023-07-03 21:06:25.995', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('SD3OJsQaM', 'authentication', '00005', '188.239.138.226', 'authenticate', '2023-07-03 21:06:32.586', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('wD-7GbQaM', 'term', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:36.507', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('n3x7Js1tp', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:41.929', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Jm2NJbQap', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:47.577', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('mHJNJbQtM', 'product', '00005', '188.239.138.226', 'patch', '2023-07-03 21:08:54.878', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('u2RuJs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:09:32.212', 'success', '');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('2GrXJb1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:09:43.729', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('tx0dJsQtM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:10.950', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('Ua9dJbQaM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:15.896', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('QD3KJb1tp', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:21.980', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('CU5dGs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:26.719', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('UnAKJs1tM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:31.352', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('SiyKGs1ap', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:10:38.634', 'fail', 'Data Validation Failed');
INSERT INTO auditlog (id, resource, userid, ip, "action", "time", status, remark) values('yYReJsQaM', 'user', '00005', '188.239.138.226', 'patch', '2023-07-03 21:11:10.110', 'success', '');


insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('dashboard','Dashboard','A','/dashboard','dashboard','assignments',1,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('admin','Admin','A','/admin','admin','contacts',2,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('setup','Setup','A','/setup','setup','settings',3,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('report','Report','A','/report','report','pie_chart',4,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('user','User Management','A','/admin/users','user','person',1,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('role','Role Management','A','/admin/roles','role','credit_card',2,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('audit_log','Audit Log','A','/admin/audit-logs','audit_log','zoom_in',4,'admin');
insert into modules (moduleid, modulename, status,path, resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('article', 'Article', 'A', '/article', 'article', 'zoom_in', 4, '', NULL, NULL, NULL, NULL);
insert into modules (moduleid, modulename, status,path, resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('entity', 'Entities', 'A', '/entities', 'entity', 'zoom_in', 4, '', NULL, NULL, NULL, NULL);
insert into modules (moduleid, modulename, status,path, resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('company', 'Company', 'A', '/companies', 'company', 'zoom_in', 4, '', NULL, NULL, NULL, NULL);
insert into modules (moduleid, modulename, status,path, resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('test', 'Test', 'A', '/test', 'test', 'zoom_in', 4, '', NULL, NULL, NULL, NULL);
insert into modules (moduleid, modulename, status,path, resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('ticket', 'Ticket Support', 'A', '/ticket', 'ticket', 'support', 4, '', NULL, NULL, NULL, NULL);
INSERT INTO modules (moduleid, modulename, status, "path", resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('term', 'Term', 'A', '/terms', 'term', 'support', 4, '', NULL, NULL, NULL, NULL);
INSERT INTO modules (moduleid, modulename, status, "path", resourcekey, icon, "sequence", parent, createdby, createdat, updatedby, updatedat)values('product', 'Product', 'A', '/products', 'product', 'support', 4, '', NULL, NULL, NULL, NULL);

insert into roles (roleid, rolename, status, remark) values ('admin','Admin','A','Admin');
insert into roles (roleid, rolename, status, remark) values ('call_center','Call Center','A','Call Center');
insert into roles (roleid, rolename, status, remark) values ('it_support','IT Support','A','IT Support');
insert into roles (roleid, rolename, status, remark) values ('operator','Operator Group','A','Operator Group');

insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00001','gareth.bale.corp','gareth.bale@gmail.com','Gareth Bale Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00002','cristiano.ronaldo.corp','cristiano.ronaldo@gmail.com','Cristiano Ronaldo Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00003','james.rodriguez.corp','james.rodriguez@gmail.com','James Rodríguez Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00004','zinedine.zidane.corp','zinedine.zidane@gmail.com','Zinedine Zidane Corp','https://uploads.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00005','kaka.corp','kaka@gmail.com','Kaká Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','0987654321');

insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00001','gareth.bale','gareth.bale@gmail.com','Gareth Bale','https://uploads.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00002','cristiano.ronaldo','cristiano.ronaldo@gmail.com','Cristiano Ronaldo','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00003','james.rodriguez','james.rodriguez@gmail.com','James Rodríguez','https://uploads.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00004','zinedine.zidane','zinedine.zidane@gmail.com','Zinedine Zidane','https://uploads.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00005','kaka','kaka@gmail.com','Kaká','https://uploads.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00006','luis.figo','luis.figo@gmail.com','Luís Figo','https://uploads.wikimedia.org/wikipedia/commons/thumb/6/63/UEFA_TT_7209.jpg/440px-UEFA_TT_7209.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00007','ronaldo','ronaldo@gmail.com','Ronaldo','https://uploads.wikimedia.org/wikipedia/commons/c/c8/Real_Valladolid-Valencia_CF%2C_2019-05-18_%2890%29_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00008','thibaut.courtois','thibaut.courtois@gmail.com','Thibaut Courtois','https://uploads.wikimedia.org/wikipedia/commons/thumb/c/c4/Courtois_2018_%28cropped%29.jpg/440px-Courtois_2018_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00009','luka.modric','luka.modric@gmail.com','Luka Modrić','https://uploads.wikimedia.org/wikipedia/commons/thumb/e/e9/ISL-HRV_%287%29.jpg/440px-ISL-HRV_%287%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00010','xabi.alonso','xabi.alonso@gmail.com','Xabi Alonso','https://uploads.wikimedia.org/wikipedia/commons/thumb/4/4a/Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg/440px-Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00011','karim.benzema','karim.benzema@gmail.com','Karim Benzema','https://uploads.wikimedia.org/wikipedia/commons/thumb/e/e4/Karim_Benzema_2018.jpg/440px-Karim_Benzema_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00012','marc-andre.ter.stegen','marc-andre.ter.stegen@gmail.com','Marc-André ter Stegen','https://uploads.wikimedia.org/wikipedia/commons/thumb/e/e1/Marc-Andr%C3%A9_ter_Stegen.jpg/500px-Marc-Andr%C3%A9_ter_Stegen.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00013','sergino.dest','sergino.dest@gmail.com','Sergiño Dest','https://uploads.wikimedia.org/wikipedia/commons/thumb/6/6e/Sergino_Dest.jpg/440px-Sergino_Dest.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00014','gerard.pique','gerard.pique@gmail.com','Gerard Piqué','https://uploads.wikimedia.org/wikipedia/commons/4/4e/Gerard_Piqu%C3%A9_2018.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00015','ronald.araujo','ronald.araujo@gmail.com@gmail.com','Ronald Araújo','https://pbs.twimg.com/media/EtnqxaEU0AAc6A6.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00016','sergio.busquets','sergio.busquets@gmail.com@gmail.com','Sergio Busquets','https://uploads.wikimedia.org/wikipedia/commons/thumb/f/fd/Sergio_Busquets_2018.jpg/440px-Sergio_Busquets_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00017','antoine.griezmann','antoine.griezmann@gmail.com@gmail.com','Antoine Griezmann','https://uploads.wikimedia.org/wikipedia/commons/thumb/f/fc/Antoine_Griezmann_in_2018_%28cropped%29.jpg/440px-Antoine_Griezmann_in_2018_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00018','miralem.pjanic','miralem.pjanic@gmail.com@gmail.com','Miralem Pjanić','https://uploads.wikimedia.org/wikipedia/commons/thumb/d/d4/20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg/440px-20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg','A','M','0987654321','Mrs','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00019','martin.braithwaite','martin.braithwaite@gmail.com@gmail.com','Martin Braithwaite','https://img.a.transfermarkt.technology/portrait/header/95732-1583334177.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00020','ousmane.dembele','ousmane.dembele@gmail.com@gmail.com','Ousmane Dembélé','https://uploads.wikimedia.org/wikipedia/commons/7/77/Ousmane_Demb%C3%A9l%C3%A9_2018.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00021','riqui.puig','riqui.puig@gmail.com@gmail.com','Riqui Puig','https://uploads.wikimedia.org/wikipedia/commons/thumb/a/ae/Bar%C3%A7a_Napoli_12_%28cropped%29.jpg/440px-Bar%C3%A7a_Napoli_12_%28cropped%29.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00022','philip.coutinho','philip.coutinho@gmail.com@gmail.com','Philip Coutinho','https://uploads.wikimedia.org/wikipedia/commons/thumb/9/96/Norberto_Murara_Neto_2019.jpg/440px-Norberto_Murara_Neto_2019.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00023','victor.lindelof','victor.lindelof@gmail.com@gmail.com','Victor Lindelöf','https://uploads.wikimedia.org/wikipedia/commons/thumb/c/cc/CSKA-MU_2017_%286%29.jpg/440px-CSKA-MU_2017_%286%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00024','eric.bailly','eric.bailly@gmail.com@gmail.com','Eric Bailly','https://uploads.wikimedia.org/wikipedia/commons/c/cf/Eric_Bailly_-_ManUtd.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00025','phil.jones','phil.jones@gmail.com@gmail.com','Phil Jones','https://uploads.wikimedia.org/wikipedia/commons/f/fa/Phil_Jones_2018-06-28_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00026','harry.maguire','harry.maguire@gmail.com@gmail.com','Harry Maguire','https://uploads.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00027','paul.pogba','paul.pogba@gmail.com@gmail.com','Paul Pogba','https://uploads.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00028','edinson.cavani','edinson.cavani@gmail.com@gmail.com','Edinson Cavani','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/88/Edinson_Cavani_2018.jpg/440px-Edinson_Cavani_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00029','juan.mata','juan.mata@gmail.com@gmail.com','Juan Mata','https://uploads.wikimedia.org/wikipedia/commons/7/70/Ukr-Spain2015_%286%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00030','anthony.martial','anthony.martial@gmail.com@gmail.com','Anthony Martial','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/88/Anthony_Martial_27_September_2017_cropped.jpg/440px-Anthony_Martial_27_September_2017_cropped.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00031','marcus.rashford','marcus.rashford@gmail.com@gmail.com','Marcus Rashford','https://uploads.wikimedia.org/wikipedia/commons/5/5e/Press_Tren_CSKA_-_MU_%283%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00032','mason.greenwood','mason.greenwood@gmail.com@gmail.com','Mason Greenwood','https://uploads.wikimedia.org/wikipedia/commons/thumb/e/e0/Mason_Greenwood.jpeg/440px-Mason_Greenwood.jpeg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00033','lee.grant','lee.grant@gmail.com@gmail.com','Lee Grant','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/8e/LeeGrant09.jpg/400px-LeeGrant09.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00034','jesse.lingard','jesse.lingard@gmail.com@gmail.com','Jesse Lingard','https://uploads.wikimedia.org/wikipedia/commons/thumb/1/11/Jesse_Lingard_2018-06-13_1.jpg/440px-Jesse_Lingard_2018-06-13_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00035','keylor.navas','keylor.navas@gmail.com@gmail.com','Keylor Navas','https://uploads.wikimedia.org/wikipedia/commons/thumb/c/c3/Liver-RM_%288%29_%28cropped%29.jpg/440px-Liver-RM_%288%29_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00036','achraf.hakimi','achraf.hakimi@gmail.com@gmail.com','Achraf Hakimi','https://uploads.wikimedia.org/wikipedia/commons/9/91/Iran-Morocco_by_soccer.ru_14_%28Achraf_Hakimi%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00037','presnel.kimpembe','presnel.kimpembe@gmail.com@gmail.com','Presnel Kimpembe','https://uploads.wikimedia.org/wikipedia/commons/thumb/0/0e/Presnel_Kimpembe.jpg/400px-Presnel_Kimpembe.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00038','sergio.ramos','sergio.ramos@gmail.com@gmail.com','Sergio Ramos','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/88/FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg/440px-FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00039','marquinhos','marquinhos@gmail.com@gmail.com','Marquinhos','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/8c/Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg/440px-Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00040','marco.verratti','marco.verratti@gmail.com@gmail.com','Marco Verratti','https://uploads.wikimedia.org/wikipedia/commons/d/d0/Kiev-PSG_%289%29.jpg','A','M','0987654321','Mr','E');

insert into entityusers(entityid, userid) values ('00001','00001');
insert into entityusers(entityid, userid) values ('00001','00002');
insert into entityusers(entityid, userid) values ('00001','00003');
insert into entityusers(entityid, userid) values ('00001','00004');
insert into entityusers(entityid, userid) values ('00002','00005');
insert into entityusers(entityid, userid) values ('00003','00006');
insert into entityusers(entityid, userid) values ('00004','00007');
insert into entityusers(entityid, userid) values ('00005','00008');
insert into entityusers(entityid, userid) values ('00005','00009');
insert into entityusers(entityid, userid) values ('00005','00010');

insert into userroles(userid, roleid) values ('00001','admin');
insert into userroles(userid, roleid) values ('00003','admin');
insert into userroles(userid, roleid) values ('00004','admin');
insert into userroles(userid, roleid) values ('00005','it_support');
insert into userroles(userid, roleid) values ('00007','admin');
insert into userroles(userid, roleid) values ('00008','call_center');
insert into userroles(userid, roleid) values ('00009','it_support');
insert into userroles(userid, roleid) values ('00010','call_center');
insert into userroles(userid, roleid) values ('00011','it_support');
insert into userroles(userid, roleid) values ('00012','call_center');
insert into userroles(userid, roleid) values ('00012','it_support');

insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'setup', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'report', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'audit_log', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'ticket', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'term', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('admin', 'product', 7);

insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'audit_log', 7);


create table companies (
  companyid varchar(40) primary key,
  companyname varchar(255) not null,
  email varchar(255) not null,
  status char(1) not null,
  phone varchar(20),
  imageurl varchar(500),
  createdby varchar(40),
  createdat timestamptz,
  updatedby varchar(40),
  updatedat timestamp,
  lastlogin timestamptz
);
create table companyusers (
  companyid varchar(40) not null,
  userid varchar(40) not null,
  primary key (companyid, userid)
);

insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00001','gareth.bale.corp','gareth.bale@gmail.com','Gareth Bale Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00002','cristiano.ronaldo.corp','cristiano.ronaldo@gmail.com','Cristiano Ronaldo Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00003','james.rodriguez.corp','james.rodriguez@gmail.com','James Rodríguez Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00004','zinedine.zidane.corp','zinedine.zidane@gmail.com','Zinedine Zidane Corp','https://uploads.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00005','kaka.corp','kaka@gmail.com','Kaká Corp','https://uploads.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','0987654321');

insert into companyusers(companyid, userid) values ('00001','00001');
insert into companyusers(companyid, userid) values ('00001','00002');
insert into companyusers(companyid, userid) values ('00001','00003');
insert into companyusers(companyid, userid) values ('00001','00004');
insert into companyusers(companyid, userid) values ('00002','00005');
insert into companyusers(companyid, userid) values ('00003','00006');
insert into companyusers(companyid, userid) values ('00004','00007');
insert into companyusers(companyid, userid) values ('00005','00008');
insert into companyusers(companyid, userid) values ('00005','00009');
insert into companyusers(companyid, userid) values ('00005','00010');

CREATE TABLE questions (
   id varchar(50) PRIMARY KEY,
   title varchar(255) NULL,
   body varchar(2000) NULL,
   mixed bool NULL,
   answers _int8 NULL,
   tags _varchar NULL,
   categoryid varchar NULL,
   "options" _varchar NULL
);
-- ----------------------------
-- Records of questions
-- ----------------------------
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('n2N0UM6tM', 'test2', 'You want to identify which tasks in the planned learning activity incorporate the use of the computer lab.
Which task should you identify?
This item is part of a case study. To view the case study information, click on the Case Study button below.', false, NULL, 'TEST_01', NULL, NULL, NULL, NULL, NULL, '{"{\"body\": \"The students must select and edit digital images that illustrate the theme or the event.\", \"correct\": true}"}');
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('mbdiwMwap', 'aaaa', 'aaaaa', false, NULL, 'TEST_01', NULL, NULL, NULL, NULL, NULL, '{"{\"body\": \"Yes\", \"correct\": true}","{\"body\": \"No\", \"correct\": false}"}');
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('od19Lp6ap', 'rwerwer', 'sdfsdfsd', false, '{dfasdfsdf,sdfsdfsdf,sdfsdfdsf}', 'TEST_02', NULL, NULL, NULL, NULL, NULL, '{"{\"body\": \"sdfsdfsdf\", \"correct\": true}","{\"body\": \"sdfsdfsdfsdf\", \"correct\": false}"}');
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('f4Z9KQ_aM', 'adsdfsdfsdf', 'What will happen after the brunch?', false, NULL, 'TEST_02', '{"The people will be served drinks.","An artist will show his drawings.","The guests will introduce themselves.","A poet will read a poem."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', '{"{\"body\": \"asdfssdf\", \"correct\": false}","{\"body\": \"sdfdsfsdfsdf\", \"correct\": false}"}');
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('Io2cK1Vtp', '', 'What do guests get to take home?', false, NULL, 'TEST_02', '{Lunch,"A book of mythology","A collection of poems","A painting"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('QTjMlQ_ap', 'test2', 'You want to identify which tasks in the planned learning activity incorporate the use of the computer lab.
Which task should you identify?
This item is part of a case study. To view the case study information, click on the Case Study button below.', false, NULL, 'TEST_01', '{"The students must hang the posters in the entrance hall of the school.","The students must assemble the images on a wall poster.","The students must select and edit  digital images that illustrate the theme or the event.","The students must split into small groups and each group must identify a theme or a historical event."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('gJL-XQVaM', '', 'As an extension activity to the planned learning activity, you ask some of the students to contact external experts to gather additional information about the historical events.
You want to ensure that the students use an appropriate means of communication when they contact the experts.
What should you advise the students to use?
This item is part of a case study. To view the case study information, click on the Case Study button below.', false, NULL, 'TEST_01', '{"their personal mobile phone","an anonymous online forum"," a generic email account of the school"," their social network account"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('VFfZX1_ap', '', 'You teach a class of 20 students. Your classroom contains six computers.
You run an activity that is carried out mostly on paper.
Before the end of the activity, the students will require access to the computers.
Which lesson structure should provide the best learning experience for the students?', false, NULL, 'TEST_01', '{"Create several different activities that require the computers at different stages during the lesson."," Execute the computer-based portion of the activity as a class demonstration.","Modify the tasks of the lesson so that only some of the students require access to the computers.","Have all of the students perform the same activity and instruct the students to take turns using the computers."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('tP8GXQVaM', '', 'What are the man and woman mainly discussing?', false, NULL, 'TEST_02', '{"A vacation","A budget","A company policy","A conference"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('NQL_XQ_tp', '', 'How is the woman traveling?', false, NULL, 'TEST_02', '{"By plane","By bus","By taxi","By car"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('9i-QvQ_tp', '', 'Why aren''t the man and woman going together?', false, NULL, 'TEST_02', '{"The woman needs to arrive earlier.","The man has to work overtime.","The woman dislikes air travel.","The man has to go to the bank first."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('VzkRv1Vtp', '', 'What does the man have to do today?', false, NULL, 'TEST_02', '{"Visit his lawyer","Get a massage","Go to the doctor.","Make an appointment."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('T_xCX1_ap', '', 'What can be inferred from the conversation?', false, NULL, 'TEST_02', '{"The woman is the man''s receptionist.","The lawyer works in the same building.","The woman has no deadlines today.","The man and woman have a meeting this afternoon."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('VWuoK1_tM', '', 'What does the woman offer to do for the man?', false, NULL, 'TEST_02', '{"Answer his phone","Call his lawyer","Pick up the newspaper","Take notes at the meeting"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('aZOHd1Vtp', '', 'What should the passengers do before exiting the ship?', false, NULL, 'TEST_02', '{"Welcome the visitors","Check the time","Collect their personal items","Take a picture"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('S6J0KQVaM', '', 'What does the speaker imply?', false, NULL, 'TEST_02', '{"The water was rough.","The weather was poor yesterday.","The tour went faster than usual.","There is only one way to exit."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('cJWfK1_ap', '', 'What will happen in five minutes?', false, NULL, 'TEST_02', '{"The ship will arrive at the dock.","The passengers will go shopping.","The passengers will take photos of the ship.","The market will open."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('f_XaK1VtM', '', 'How often does this event take place?', false, NULL, 'TEST_02', '{Monthly,"Four times a year","Once a year","Every four years"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('UkkSd1_tM', '', 'Where would you hear this type of talk?', false, NULL, 'TEST_02', '{"In an airport","At a bus station","On a telephone","On television"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('cDiEdQ_tp', '', 'Who should visit the website?', false, NULL, 'TEST_02', '{"Staff members of Speedy Reservations","Callers looking for special schedules","Callers looking for staff members","People looking for fight times"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('hs2iK1VtM', '', 'What is the quickest way to get connected to a representative?', false, NULL, 'TEST_02', '{"Hang up and call again","Call back later"," Wait patiently","Call a different number"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('Lk7LKQ_tp', '', 'Did you have to wait very long?', false, '{part2}', 'TEST_02', '{"No more than an hour.","I hate long line-ups.","I''ve gained twenty pounds."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('IzqAd1Vap', '', 'Do you think the boss will take us out for lunch today?', false, '{part2}', 'TEST_02', '{"I had a hamburger and fries.","He''s probably too busy today.","I didn''t take the bus."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('_1RQ9RVap', '', 'An listed entity under FATCA, is technically classified as', false, NULL, 'FATCA_Mock_Test', '{"Active NFE","Account Holder","Active NFFE",nominee}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('owb_KQ_tM', '', 'Did you fax the letter to the client?', false, '{part2}', 'TEST_02', '{"I''ll type it after lunch.","No, I sent it by email.","I collected some important facts."}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('7t9IrRVap', '', 'What is the last step to be followed to determine whether a person is a RFI and thus has reporting obligations', false, NULL, 'FATCA_Mock_Test', '{"Is the Entity a Financial Institution?","Is it an Entity?","Is the Financial Institution in India?","Is the Financial Institution a Non-Reporting Financial Institution?"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('dlbV9RVtp', '', 'An Active NFE should have', false, NULL, 'FATCA_Mock_Test', '{"if the entity during the last year: More than (or equal to) 25% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 50% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 50% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 75% of the entity’s gross income consisted of Active Income"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('OZdm9R_aM', '', 'Under FATCA, each country entered into a separate bilateral intergovernmental agreement with the United States and what objective was achieved', false, NULL, 'FATCA_Mock_Test', '{"The local financial institutions complies with FATCA reporting standards  and will not breach local data protection laws ","Financial institutions in complying with FATCA, will not breach local data protection laws only","None of these","The local financial institutions complies with FATCA reporting standards only"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);
insert into questions
(id, title, body, mixed, tags, categoryid, "options", createdby, createdat, updatedby, updatedat, answers)
values('csVonL_tp', '', 'Which of the following is an exempt beneficial owner', false, NULL, 'FATCA_Mock_Test', '{"All of these","Any International Organisation or any wholly owned agency or instrumentality thereof","Any Foreign Central Bank of Issue","Any Foreign Government, any political subdivision of a Foreign Government, or any wholly owned agency or instrumentality of any one or more of the foregoing"}', NULL, '2023-06-15 13:38:23.570', NULL, '2023-06-15 13:37:53.430', NULL);

CREATE TABLE "tests" (
  "testid" varchar PRIMARY KEY ,
  "title" varchar ,
  "effectivedate" timestamp(6),
  "questions" _varchar ,
  "tags" _varchar ,
  "categoryid" varchar,
  createdBy VARCHAR(20),
  createdAt timestamptz ,
  updatedBy VARCHAR(20),
  updatedAt timestamptz
);

INSERT INTO tests(testid, title, effectivedate, questions, tags, categoryid, createdby, createdat, updatedby, updatedat)
VALUES('32uPgd6aM', 'Test 2', '2023-06-30 07:39:00.000', '{}', NULL, 'Test_02', NULL, NULL, NULL, NULL);
INSERT INTO tests(testid, title, effectivedate, questions, tags, categoryid, createdby, createdat, updatedby, updatedat)
VALUES('cQeB3YwaM', 'Test 3', '2023-06-30 17:33:00.000', '{"QTjMlQ_ap","NQL_XQ_tp"}', '{"asd"}', 'Test_02', NULL, NULL, NULL, NULL);
INSERT INTO tests(testid, title, effectivedate, questions, tags, categoryid, createdby, createdat, updatedby, updatedat)
VALUES('GzP3EYwap', 'Test 4', '2023-06-22 01:12:00.000', '{}', '{"sdfsdfsdfsdf"}', 'Test_01', NULL, NULL, NULL, NULL);
INSERT INTO tests(testid, title, effectivedate, questions, tags, categoryid, createdby, createdat, updatedby, updatedat)
VALUES('uYDyiY6tp', 'Test 1', '2023-07-26 21:00:00.000', '{}', NULL, 'Test_01', NULL, NULL, NULL, NULL);

DROP TABLE IF EXISTS "tickets";
CREATE TABLE "tickets" (
   "id" varchar(40) PRIMARY KEY ,
   "title" varchar(100) ,
   "body" text ,
   "categoryid" varchar(40) ,
   "requestor" varchar(20) ,
   "requestedat" timestamp(6),
   "approver" varchar(20) ,
   "approvedat" timestamp(6),
   "assignee" varchar(40) ,
   "completedat" timestamp(6),
   "status" varchar(1) 
);

insert into "tickets" values ('s319KmVtM', 'Sample Title', 'Sample body text.', 'sampleCategoryId', 'sampleRequestor', '2023-06-14 10:30:00', 'sampleApprover', '2023-06-14 11:00:00', 'sampleAssignee', '2023-06-14 12:00:00', 'S');
insert into "tickets" values ('iTJWdm_aM', 'Sample Ticket', 'This is a sample ticket description.', 'category456', 'john', '2023-06-14 10:30:00', 'approver123', '2023-06-14 11:00:00', 'assignee456', '2023-06-14 12:00:00', 'S');

CREATE TABLE articles (
   id VARCHAR(255) PRIMARY KEY,
   title VARCHAR(255),
   authorid VARCHAR(255),
   name VARCHAR(255),
   description TEXT,
   publishedat timestamptz,
   type VARCHAR(255) NOT NULL,
   body TEXT NOT NULL,
   tags _varchar,
   status VARCHAR(10),
   createdBy VARCHAR(20),
   createdAt timestamptz ,
   updatedBy VARCHAR(20),
   updatedAt timestamptz
);


CREATE TABLE terms (
  id VARCHAR(50) PRIMARY KEY,
  title VARCHAR(255),
  description VARCHAR(255),
  effectivedate timestamptz  NOT NULL,
  publishedat timestamptz ,
  type VARCHAR(10) NOT NULL,
  body TEXT NOT NULL,
  tags _varchar,
  status VARCHAR(10),
  createdby VARCHAR(50),
  createdat timestamptz ,
  updatedby VARCHAR(50),
  updatedat timestamptz
);

INSERT INTO terms
(id, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('Ym8AWpwaM', '4. Where do I display my Terms and Conditions agreement?', 'This is a sample term', '2023-06-07 00:00:00.000', '2023-06-15 10:00:00.000', 'definition', '<p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">Display your Terms and Conditions agreement in the following places, where applicable:</p><ol style="margin-bottom: 1.5rem; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;"><li>Via a static link to your&nbsp;<span style="font-weight: 700;">website footer</span></li><li><span style="font-weight: 700;">In a menu</span>&nbsp;within your mobile app (typically in an About, Legal, Info menu)</li><li>On your "<span style="font-weight: 700;">Create Account</span>" or similar type of page</li><li>On any&nbsp;<span style="font-weight: 700;">checkout</span>&nbsp;or order finalization pages</li></ol>', '{"sample","term"}', 'active', '00001', '2023-06-15 11:36:58.436', '00005', '2023-06-28 17:10:23.233');
INSERT INTO terms
(id, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('vwRJOBQaM', '5. How do I make my Terms and Conditions agreement enforceable?', ' ', '2023-06-27 06:09:27.605', '2023-06-25 06:09:27.000', 'sdfsdf', '<a href="https://www.termsfeed.com/blog/sample-terms-and-conditions-template/#5__How_do_I_make_my_Terms_and_Conditions_agreement_enforceable" class="accordion-button" style="text-decoration-line: underline; color: rgb(26, 84, 169); display: block; position: relative; padding-right: 1rem; cursor: pointer; font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; font-weight: 700; letter-spacing: normal; background-color: rgb(255, 255, 255);"><span style="color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-weight: 400;">To make your Terms &amp; Conditions agreement enforceable, place an&nbsp;</span><span style="color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif;">un-ticked checkbox</span><span style="color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-weight: 400;">&nbsp;next to a link to your agreement and a statement that says something along the lines of, "</span><em style="color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-weight: 400;">By checking this box, you agree to be bound by our Terms and Conditions agreement.</em><span style="color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-weight: 400;">"</span><br></a>', NULL, 'A', '00005', '2023-06-27 13:09:39.701', '00005', '2023-06-28 17:10:50.371');
INSERT INTO terms
(id, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('KXNAgM6ap', '2. What are the benefits of having a Terms and Conditions agreement?', 'What are the benefits of having a Terms and Conditions agreement?', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'definition', '<p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">The main benefit of having a Terms and Conditions agreement is that&nbsp;<span style="font-weight: 700;">you maintain the highest level of control over your website/business</span>.</p><p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">Your Terms and Conditions agreement is where you''re able to list your rules when it comes to the use of your website. You''re also able to maintain the right to terminate abusive accounts, disclaim warranties and limit your liability.</p>', '{"sample","term"}', 'active', 'kaka', '2023-06-15 11:02:02.037', '00005', '2023-06-28 17:09:27.513');
INSERT INTO terms
(id, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('5FN7cpwap', '1. Is a Terms and Conditions agreement required?', 'Is a Terms and Conditions agreement required?', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'definition', '<p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">A Terms and Conditions agreement is&nbsp;<span style="font-weight: 700;">not</span>&nbsp;legally required. However, having one comes with a number of important benefits for both you and your users/customers.</p><p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">The benefits include increasing your control over your business/platform, while helping your users understand your rules, requirements and restrictions.</p>', '{"sample","term"}', 'active', 'Jane Smith', '2023-06-14 15:30:00.000', '00005', '2023-06-28 17:09:34.228');
INSERT INTO terms
(id, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('iBXKPM6tM', '3. What clauses should be in my Terms and Conditions agreement?', 'This is a sample term', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'definition', '<p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">Some clauses are specific to certain types of businesses and won''t be found in all Terms and Conditions agreements. For example, you won''t need a clause about subscription payment terms if you don''t offer paid subscriptions.</p><p style="margin-right: 0px; margin-bottom: 1.5rem; margin-left: 0px; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;">In general, almost every Terms and Conditions agreement should include the following clauses:</p><ol style="margin-bottom: 1.5rem; color: rgb(4, 12, 26); font-family: &quot;Helvetica Neue&quot;, Helvetica, Arial, sans-serif; font-size: 19.2px; letter-spacing: normal;"><li>Introduction</li><li>Right to make changes to the agreement</li><li>User guidelines (rules, restrictions, requirements)</li><li>Copyright and intellectual property</li><li>Governing law</li><li>Warranty disclaimer</li><li>Limitation of liability</li><li>Termination of accounts/service</li><li>Contact information</li></ol>', '{"sample","term"}', 'A', '', '2023-06-15 11:01:30.637', '00005', '2023-06-28 17:10:02.635');


CREATE TABLE products (
  productid VARCHAR(50) PRIMARY KEY,
  productname VARCHAR(50) NOT NULL,
  title VARCHAR(255),
  description VARCHAR(255),
  effectivedate timestamptz NOT NULL,
  publishedat timestamptz,
  type VARCHAR(10) NOT NULL,
  body TEXT NOT NULL,
  tags _varchar,
  status VARCHAR(10),
  createdby VARCHAR(50),
  createdat timestamptz ,
  updatedby VARCHAR(50),
  updatedat timestamptz
);

INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('304lgp6tM', 'Laptop Asus VivoBook X515KA-EJ135Wt', 'Laptop Asus VivoBook X515KA-EJ135W', 'Laptop Asus VivoBook X515KA-EJ135W', '2023-03-16 00:00:00.000', '2023-06-15 10:00:00.000', 'laptop', '<h1 style="box-sizing: inherit; margin-right: 10px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 18px; font-weight: 700; -webkit-line-clamp: 3; -webkit-box-orient: vertical; color: rgb(10, 38, 60); display: -webkit-box; line-height: 2; overflow: hidden; font-family: sans-serif; letter-spacing: normal;">Laptop Asus VivoBook X515KA-EJ135W</h1>', '{"sample","product"}', 'active', '00005', '2023-06-15 10:58:30.510', '00005', '2023-06-28 16:20:19.142');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('OKUEgM6tM', 'MacBook Air M1 256GB 2020', 'MacBook Air M1 256GB 2020 ', 'MacBook Air M1 256GB 2020 I Chính hãng Apple Việt Nam', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'sample', '<h2 align="center" style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 21px; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;">Macbook Air M1 2020 - Sang trọng, tinh tế, hiệu năng khủng</span></h2><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;">Macbook Air M1</span>&nbsp;là dòng sản phẩm có thiết kế mỏng nhẹ, sang trọng và tinh tế cùng với đó là giá thành phải chăng nên MacBook Air đã thu hút được đông đảo người dùng yêu thích và lựa chọn. Đây cũng là một trong những phiên bản Macbook Air mới nhất mà khách hàng không thể bỏ qua khi đến với CellphoneS. Dưới đây là chi tiết về thiết kế, cấu hình của máy.</p>', '{"sample","product"}', 'A', '00005', '2023-06-15 10:56:06.734', '00005', '2023-06-28 16:17:09.613');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('xo0pgp6tM', 'Sample ProductMac mini M2 2023 ', 'Mac mini M2 2023 (8 CPU - 10 GPU - 16GB - 256GB)', 'Mac mini M2 2023 (8 CPU - 10 GPU - 16GB - 256GB)', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'laptop', '<h2 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 21px; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;">Mac Mini M2 16GB 2023 – Đa nhiệm tốt, hoạt động hiệu quả</span></h2><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;">Mac Mini M2 16GB 2023 mang trong mình nhiều đặc điểm vượt trội hơn hẳn thế hệ tiền nhiệm. Thông qua nguồn “sức mạnh” của con chip M2, sản phẩm&nbsp;<a href="https://cellphones.com.vn/laptop/mac/mini.html" title="Mac mini chính hãng" target="_blank" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer;"><span style="box-sizing: inherit; color: currentcolor; font-weight: 700;">Mac mini</span></a>&nbsp;này chắc chắn bạn sẽ hài lòng với hiệu quả mà phiên bản mang đến trong quá trình hoạt động.</p>', '{"sample","product"}', 'A', '00005', '2023-06-15 10:49:44.076', '00005', '2023-06-28 16:18:03.115');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('o3HXxpwaM', ' MacBook Pro 13 M2 2022 24GB 512GB', 'Apple MacBook Pro 13 M2 2022 24GB 512GB', 'Apple MacBook Pro 13 M2 2022 24GB 512GB', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'laptop', '<h1 style="box-sizing: inherit; margin-right: 10px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 18px; font-weight: 700; -webkit-line-clamp: 3; -webkit-box-orient: vertical; color: rgb(10, 38, 60); display: -webkit-box; line-height: 2; overflow: hidden; font-family: sans-serif; letter-spacing: normal;">Apple MacBook Pro 13 M2 2022 24GB 512GB</h1>', '{"sample","product"}', 'A', '', '2023-06-15 11:18:07.422', '00005', '2023-06-28 16:19:19.678');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('H6bCgpwtp', 'Apple MacBook Pro 13 M2 2022 8GB 512GB', 'Apple MacBook Pro 13 M2 2022 8GB 512GB', 'Apple MacBook Pro 13 M2 2022 8GB 512GB', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'sample', '', '{"sample","product"}', 'I', 'kaka', '2023-06-15 11:06:02.752', '00005', '2023-06-28 17:02:03.903');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('ozxzEM6ap', 'Samsung Galaxy S22 Ultra (8GB - 128GB)', 'Samsung Galaxy S22 Ultra (8GB - 128GB)', 'Samsung Galaxy S22 Ultra (8GB - 128GB)', '2023-06-15 00:00:00.000', '2023-05-16 10:00:00.000', 'samsung', '<span style="color: rgb(74, 74, 74); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">Đúng như các thông tin được đồn đoán trước đó, mẫu flagship mới của gả khổng lồ Hàn Quốc được ra mắt với tên gọi là Samsung Galaxy S22 Ultra với nhiều cải tiến đáng giá.</span>', '{"sample","product"}', 'A', 'kaka', '2023-06-15 11:10:38.688', '00005', '2023-06-28 16:23:20.974');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('C7dWEM6ap', 'iPhone 14 Pro Max 128GB | Chính hãng VN/A', 'Bitcoin', 'This is a sample product', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'Crypto', 'sdfsfsdfsdfqsdqweqwdwedqwer', '{"sample","product"}', 'I', '', '2023-06-15 11:14:05.951', '00005', '2023-06-28 16:24:04.893');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('wn1XxMwaM', 'Laptop ASUS X515MA-BR481W N', 'Laptop ASUS X515MA-BR481W N', 'Laptop ASUS X515MA-BR481W N', '2023-06-15 00:00:00.000', '2023-06-15 13:00:00.000', 'sample', '<h2 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 21px; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;">Laptop Asus X515MA-BR481W - Gọn Nhẹ Với Thiết Kế Năng Động</h2><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;">Laptop Asus X515MA-BR481W là chiếc&nbsp;<a href="https://cellphones.com.vn/laptop/asus/vivobook.html" title="Laptop Asus Vivobook chính hãng" target="_blank" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer;"><span style="box-sizing: inherit; color: currentcolor; font-weight: 700;">laptop Asus Vivobook</span></a>&nbsp;được thiết kế năng động với hiệu năng mạnh mẽ và phù hợp với túi tiền của nhiều người dùng. Asus X515MA-BR481W sẽ là một lựa chọn không thể tuyệt vời hơn dành cho các bạn học sinh, sinh viên.</p>', '{"sample","product"}', 'active', '00005', '2023-06-15 11:18:33.295', '00005', '2023-06-28 16:19:55.619');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('PQFZPM6ap', 'Laptop ASUS Gaming ROG Zephyrus G14 GA401QC-K2199W', 'Laptop ASUS Gaming ROG Zephyrus G14 GA401QC-K2199W', 'Laptop ASUS Gaming ROG Zephyrus G14 GA401QC-K2199W', '2023-05-04 00:00:00.000', '2023-06-13 10:00:00.000', 'sample', '<h2 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 21px; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;">Laptop ASUS Gaming ROG Zephyrus G14 GA401QC-K2199W – Cấu hình mạnh mẽ</span></h2><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;"><a href="https://cellphones.com.vn/laptop/asus/gaming.html" title="Laptop Asus gaming chính hãng" target="_blank" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer;">Laptop ASUS Gaming</a>&nbsp;ROG Zephyrus G14 GA401QC-K2199W&nbsp;</span>là chiếc laptop mà bạn nên sở hữu nếu muốn một cấu hình máy tính chơi game với độ trải nghiệm hình ảnh đồ họa mượt mà, âm thanh sống động. Ngoài ra còn có thiết kế nhỏ gọn và những tính năng vượt trội.</p><h3 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 8px 0px; font-size: 16px; font-weight: 700; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54);">Thiết kế nổi bật, độ phân giải WQHD sắc nét</span></h3><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;">Laptop ASUS Gaming ROG Zephyrus G14 GA401QC-K2199W là chiếc laptop của thương hiệu ASUS sản xuất với phân khúc thị trường chính là phục vụ nhu cầu Gaming, vì vậy mà thiết kế đậm chất Gaming. Laptop có màu xám.</p>', '{"sample","product"}', 'active', '00005', '2023-06-15 10:54:51.146', '00005', '2023-06-28 16:21:48.265');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('8VhiPpwtp', 'iPhone 14 Pro Max 128GB ', 'iPhone 14 Pro Max 128GB ', 'iPhone 14 Pro Max 128GB ', '2023-06-15 00:00:00.000', '2023-06-15 10:00:00.000', 'iphone', '<blockquote><a href="https://cellphones.com.vn/iphone-14-pro-max.html" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer; font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify; background-color: rgb(255, 255, 255);"><span style="box-sizing: inherit; color: currentcolor; font-weight: 700;">iPhone 14 Pro Max</span></a><span style="color: rgb(100, 100, 100); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">&nbsp;sở hữu thiết kế&nbsp;</span><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700; font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">màn hình Dynamic Island</span><span style="color: rgb(100, 100, 100); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">&nbsp;ấn tượng cùng màn hình&nbsp;</span><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700; font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">OLED 6,7 inch</span><span style="color: rgb(100, 100, 100); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">&nbsp;hỗ trợ always-on display và hiệu năng vượt trội với&nbsp;</span><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700; font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">chip A16 Bionic</span><span style="color: rgb(100, 100, 100); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">.</span></blockquote>', '{"sample","product"}', 'A', '00005', '2023-06-15 10:56:30.418', '00005', '2023-06-28 16:22:43.552');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('027ZPM6tM', 'OPPO A95', 'OPPO A95', 'This is a sample product', '2023-04-20 00:00:00.000', '2023-06-28 10:00:00.000', 'OPPO', '<a href="https://cellphones.com.vn/oppo-a95.html" title="OPPO A95" target="_blank" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer; font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify; background-color: rgb(255, 255, 255);"><span style="box-sizing: inherit; color: currentcolor; font-weight: 700;">OPPO A95&nbsp;</span></a><span style="color: rgb(74, 74, 74); font-family: sans-serif; font-size: 15px; letter-spacing: normal; text-align: justify;">là dòng chiến binh nghìn máu phân khúc bình dân đến từ OPPO. Hiệu năng mạnh mẽ được đảm bảo bởi bộ vi xử lý Snapdragon 662 cùng bộ nhớ RAM 8GB kết hợp công nghệ mở rộng 5GB ấn tượng giúp OPPO A95 dễ dàng cân đẹp mọi thao tác đa nhiệm</span>', '{"sample","product"}', 'A', '00005', '2023-06-15 10:54:55.907', '00005', '2023-06-28 16:23:54.403');
INSERT INTO products
(productid, productname, title, description, effectivedate, publishedat, "type", body, tags, status, createdby, createdat, updatedby, updatedat)
VALUES('e6gbZMwtM', 'Laptop gaming MSI GF63 Thin 11SC 664VN I', 'Laptop gaming MSI GF63 Thin 11SC 664VN I', 'Laptop gaming MSI GF63 Thin 11SC 664VN I', '2023-06-16 00:00:00.000', '2023-06-15 10:00:00.000', 'laptop', '<h2 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 0px; font-size: 21px; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;">Laptop MSI Gaming GF63 Thin 11SC-664VN - dòng laptop Gaming với thiết kế mỏng nhẹ</span></h2><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54); font-weight: 700;"><a href="https://cellphones.com.vn/laptop/msi.html" title="Laptop MSI chính hãng" target="_self" style="box-sizing: inherit; color: rgb(215, 0, 24); cursor: pointer;">Laptop MSI</a>&nbsp;Gaming GF63 Thin 11SC-664VN</span>&nbsp;là sản phẩm thuộc phân khúc giá tầm trung, phù hợp với những game thủ. Đây là dòng laptop có vẻ ngoài nhỏ gọn nhưng vẫn mang đậm phong cách gaming.</p><h3 style="box-sizing: inherit; margin-right: 0px; margin-bottom: 0px; margin-left: 0px; padding: 8px 0px; font-size: 16px; font-weight: 700; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;"><span style="box-sizing: inherit; color: rgb(54, 54, 54);">Vẻ ngoài thanh mảnh nhưng vẫn đậm chất gaming</span></h3><p style="box-sizing: inherit; margin-right: 0px; margin-bottom: 10px; margin-left: 0px; padding: 0px; font-size: 15px; line-height: 1.5; color: rgb(74, 74, 74); font-family: sans-serif; letter-spacing: normal; text-align: justify;">Nhìn chung, chiếc laptop mang một vẻ ngoài khá thon gọn, tổng bề dày chỉ 21,7 mm và có khối lượng khoảng 1,86kg. Tuy nhiên vì là dòng máy gaming nên ở phần giữa của nắp máy được in chìm logo rồng đỏ quen thuộc nhằm tạo điểm nhấn. Đặc biệt hơn, phần thông gió ở phía dưới được in ẩn hình chữ X khá độc đáo.</p>', '{"sample","product"}', 'A', '00005', '2023-06-14 15:30:00.000', '00005', '2023-06-28 17:01:52.083');

/*
alter table userroles add foreign key (userid) references users (userid);
alter table userroles add foreign key (roleid) references roles (roleid);

alter table modules add foreign key (parent) references modules (moduleid);

alter table rolemodules add foreign key (roleid) references roles (roleid);
alter table rolemodules add foreign key (moduleid) references modules (moduleid);
*/

CREATE TABLE ticketcommentthread (
  commentid varchar(40) NOT NULL,
  id varchar(40) NULL,
  author varchar(255) NULL,
  "comment" text NULL,
  "time" timestamptz(6) NULL,
  updatedat timestamptz(6) NULL,
  histories _jsonb NULL,
  CONSTRAINT ticketcommentthread_pk PRIMARY KEY (commentid)
);

CREATE TABLE ticketcommentthreadinfo (
   commentid varchar(40) NOT NULL,
   replycount int4 NULL DEFAULT 0,
   usefulcount int4 NULL DEFAULT 0,
   CONSTRAINT ticketcommentthreadinfo_pk PRIMARY KEY (commentid)
);

CREATE TABLE ticketreplycomment (
  commentid varchar(40) NOT NULL,
  commentthreadid varchar(40) NOT NULL,
  id varchar(40) NULL,
  author varchar(255) NULL,
  "comment" text NULL,
  "time" timestamptz(6) NULL,
  updatedat timestamptz(6) NULL,
  histories _jsonb NULL,
  CONSTRAINT ticketreplycomment_pk PRIMARY KEY (commentid)
);

CREATE TABLE ticketreplycommentinfo (
   commentid varchar(40) NOT NULL,
   replycount int4 NULL DEFAULT 0,
   usefulcount int4 NULL DEFAULT 0,
   CONSTRAINT ticketreplycommentinfo_pk PRIMARY KEY (commentid)
);

CREATE TABLE tickets (
   id varchar(40) NULL,
   title varchar(100) NULL,
   body text NULL,
   categoryid varchar(40) NULL,
   requestor varchar(20) NULL,
   requested_at timestamp(6) NULL,
   approver varchar(20) NULL,
   approved_at timestamp(6) NULL,
   assignee varchar(40) NULL,
   completed_at timestamp(6) NULL,
   status varchar(1) NULL,
   attachments _jsonb NULL
);

INSERT INTO tickets
(id, title, body, categoryid, requestor, requested_at, approver, approved_at, assignee, completed_at, status, attachments)
VALUES('s319KmVtM', '[Bug] IssueLabelBot no longer works ', 'Describe the bug
The document references https://github.com/marketplace/issue-label-bot but the owner shut it down:
https://github.com/machine-learning-apps/Issue-Label-Bot

The Issue label bot is no longer live. You can still view the code, but we have taken down the bot because of related infrastructure costs', '3', 'sampleRequestor', '2023-06-14 10:30:00.000', 'sampleApprover', '2023-06-14 11:00:00.000', 'sampleAssignee', '2023-06-14 12:00:00.000', 'H', NULL);
INSERT INTO tickets
(id, title, body, categoryid, requestor, requested_at, approver, approved_at, assignee, completed_at, status, attachments)
VALUES('gjuIgMwtp', 'Markdown syntax checker / parser before ', 'What is the current behavior?

useCustomCompareEffect does not give a chance to prevent running the effect on mount.
The code says that depsEqual are not even called, as first part of !ref.current || !depsEqual(deps, ref.current) expression is falsy.

What is the expected behavior?
depsEqual should be called to allow deciding on not running the effect on mount, as for every later state change.

A little about versions:

OS: Linux
Browser (vendor and version): Firefox
React: 18.2
react-use: 17.4.0
Did this worked in the previous package version?: I don''t think so.
', '2', 'john', '2023-06-14 10:30:00.000', 'approver123', '2023-06-14 11:00:00.000', 'assignee456', '2023-06-14 12:00:00.000', 'I', NULL);
INSERT INTO tickets
(id, title, body, categoryid, requestor, requested_at, approver, approved_at, assignee, completed_at, status, attachments)
VALUES('iTJWdm_aM', 'Guidelines too opinionated about squashi', 'This is a sample ticket description.', '4', 'john', '2023-06-14 10:30:00.000', 'approver123', '2023-06-14 11:00:00.000', 'assignee456', '2023-06-14 12:00:00.000', 'O', '{{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/vEOmOqwaM_nice_beach.jpg", "type": "image", "fileName": "vEOmOqwaM_nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/Bq0Ybuwtp_nice_beach.jpg", "type": "image", "fileName": "Bq0Ybuwtp_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/k9ptbqwaM_nice_beach.jpg", "type": "image", "fileName": "k9ptbqwaM_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/28Sgsu6tp_nice_beach.jpg", "type": "image", "fileName": "28Sgsu6tp_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/UI7Sdq6aM_nice_beach.jpg", "type": "image", "fileName": "UI7Sdq6aM_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/D9dBLvwtM_nice_beach.jpg", "type": "image", "fileName": "D9dBLvwtM_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/h80RrdwtM_nice_beach.jpg", "type": "image", "fileName": "h80RrdwtM_nice_beach.jpg", "originalFileName": "nice_beach.jpg"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/QiwNYBQaM_file_example_XLS_10.xls", "type": "application", "fileName": "QiwNYBQaM_file_example_XLS_10.xls", "originalFileName": "file_example_XLS_10.xls"},{"url": "https:/storage.googleapis.com/go-firestore-rest-api.appspot.com/sub/FRX8toQap_file_example_XLS_10.xls", "size": 8704, "type": "application/vnd.ms-excel", "fileName": "FRX8toQap_file_example_XLS_10.xls", "originalFileName": "file_example_XLS_10.xls"}}');

ALTER TABLE ticketcommentthread ADD CONSTRAINT ticketcommentthread_pk PRIMARY KEY (commentid);
ALTER TABLE ticketcommentthreadinfo ADD CONSTRAINT ticketcommentthreadinfo_pk PRIMARY KEY (commentid);
ALTER TABLE ticketreplycomment ADD CONSTRAINT ticketreplycomment_pk PRIMARY KEY (commentid);
ALTER TABLE ticketreplycommentinfo ADD CONSTRAINT ticketreplycommentinfo_pk PRIMARY KEY (commentid);
