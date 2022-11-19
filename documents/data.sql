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
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp
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
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp,
  lastlogin timestamp
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
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp,
  lastlogin timestamp
);

create table roles (
  roleid varchar(40) primary key,
  rolename varchar(255) not null,
  status char(1) not null,
  remark varchar(255),
  createdby varchar(40),
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp
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
  time timestamp,
  status varchar(255),
  remark varchar(255)
);

insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('dashboard','Dashboard','A','/dashboard','dashboard','assignments',1,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('admin','Admin','A','/admin','admin','contacts',2,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('setup','Setup','A','/setup','setup','settings',3,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('report','Report','A','/report','report','pie_chart',4,'');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('user','User Management','A','/admin/users','user','person',1,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('role','Role Management','A','/admin/roles','role','credit_card',2,'admin');
insert into modules (moduleid,modulename,status,path,resourcekey,icon,sequence,parent) values ('audit_log','Audit Log','A','/admin/audit-logs','audit_log','zoom_in',4,'admin');

insert into roles (roleid, rolename, status, remark) values ('admin','Admin','A','Admin');
insert into roles (roleid, rolename, status, remark) values ('call_center','Call Center','A','Call Center');
insert into roles (roleid, rolename, status, remark) values ('it_support','IT Support','A','IT Support');
insert into roles (roleid, rolename, status, remark) values ('operator','Operator Group','A','Operator Group');

insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00001','gareth.bale.corp','gareth.bale@gmail.com','Gareth Bale Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00002','cristiano.ronaldo.corp','cristiano.ronaldo@gmail.com','Cristiano Ronaldo Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00003','james.rodriguez.corp','james.rodriguez@gmail.com','James Rodríguez Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00004','zinedine.zidane.corp','zinedine.zidane@gmail.com','Zinedine Zidane Corp','https://upload.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','0987654321');
insert into entities (entityid,entityname,email,displayname,imageurl,status,phone) values ('00005','kaka.corp','kaka@gmail.com','Kaká Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','0987654321');

insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00001','gareth.bale','gareth.bale@gmail.com','Gareth Bale','https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00002','cristiano.ronaldo','cristiano.ronaldo@gmail.com','Cristiano Ronaldo','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00003','james.rodriguez','james.rodriguez@gmail.com','James Rodríguez','https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00004','zinedine.zidane','zinedine.zidane@gmail.com','Zinedine Zidane','https://upload.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00005','kaka','kaka@gmail.com','Kaká','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00006','luis.figo','luis.figo@gmail.com','Luís Figo','https://upload.wikimedia.org/wikipedia/commons/thumb/6/63/UEFA_TT_7209.jpg/440px-UEFA_TT_7209.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00007','ronaldo','ronaldo@gmail.com','Ronaldo','https://upload.wikimedia.org/wikipedia/commons/c/c8/Real_Valladolid-Valencia_CF%2C_2019-05-18_%2890%29_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00008','thibaut.courtois','thibaut.courtois@gmail.com','Thibaut Courtois','https://upload.wikimedia.org/wikipedia/commons/thumb/c/c4/Courtois_2018_%28cropped%29.jpg/440px-Courtois_2018_%28cropped%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00009','luka.modric','luka.modric@gmail.com','Luka Modrić','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e9/ISL-HRV_%287%29.jpg/440px-ISL-HRV_%287%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00010','xabi.alonso','xabi.alonso@gmail.com','Xabi Alonso','https://upload.wikimedia.org/wikipedia/commons/thumb/4/4a/Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg/440px-Xabi_Alonso_Training_2017-03_FC_Bayern_Muenchen-3_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00011','karim.benzema','karim.benzema@gmail.com','Karim Benzema','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e4/Karim_Benzema_2018.jpg/440px-Karim_Benzema_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00012','marc-andre.ter.stegen','marc-andre.ter.stegen@gmail.com','Marc-André ter Stegen','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e1/Marc-Andr%C3%A9_ter_Stegen.jpg/500px-Marc-Andr%C3%A9_ter_Stegen.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00013','sergino.dest','sergino.dest@gmail.com','Sergiño Dest','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6e/Sergino_Dest.jpg/440px-Sergino_Dest.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00014','gerard.pique','gerard.pique@gmail.com','Gerard Piqué','https://upload.wikimedia.org/wikipedia/commons/4/4e/Gerard_Piqu%C3%A9_2018.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00015','ronald.araujo','ronald.araujo@gmail.com@gmail.com','Ronald Araújo','https://pbs.twimg.com/media/EtnqxaEU0AAc6A6.jpg','A','M','0987654321','Mr','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00016','sergio.busquets','sergio.busquets@gmail.com@gmail.com','Sergio Busquets','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fd/Sergio_Busquets_2018.jpg/440px-Sergio_Busquets_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00017','antoine.griezmann','antoine.griezmann@gmail.com@gmail.com','Antoine Griezmann','https://upload.wikimedia.org/wikipedia/commons/thumb/f/fc/Antoine_Griezmann_in_2018_%28cropped%29.jpg/440px-Antoine_Griezmann_in_2018_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00018','miralem.pjanic','miralem.pjanic@gmail.com@gmail.com','Miralem Pjanić','https://upload.wikimedia.org/wikipedia/commons/thumb/d/d4/20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg/440px-20150331_2025_AUT_BIH_2130_Miralem_Pjani%C4%87.jpg','A','M','0987654321','Mrs','M');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00019','martin.braithwaite','martin.braithwaite@gmail.com@gmail.com','Martin Braithwaite','https://img.a.transfermarkt.technology/portrait/header/95732-1583334177.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00020','ousmane.dembele','ousmane.dembele@gmail.com@gmail.com','Ousmane Dembélé','https://upload.wikimedia.org/wikipedia/commons/7/77/Ousmane_Demb%C3%A9l%C3%A9_2018.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00021','riqui.puig','riqui.puig@gmail.com@gmail.com','Riqui Puig','https://upload.wikimedia.org/wikipedia/commons/thumb/a/ae/Bar%C3%A7a_Napoli_12_%28cropped%29.jpg/440px-Bar%C3%A7a_Napoli_12_%28cropped%29.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00022','philip.coutinho','philip.coutinho@gmail.com@gmail.com','Philip Coutinho','https://upload.wikimedia.org/wikipedia/commons/thumb/9/96/Norberto_Murara_Neto_2019.jpg/440px-Norberto_Murara_Neto_2019.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00023','victor.lindelof','victor.lindelof@gmail.com@gmail.com','Victor Lindelöf','https://upload.wikimedia.org/wikipedia/commons/thumb/c/cc/CSKA-MU_2017_%286%29.jpg/440px-CSKA-MU_2017_%286%29.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00024','eric.bailly','eric.bailly@gmail.com@gmail.com','Eric Bailly','https://upload.wikimedia.org/wikipedia/commons/c/cf/Eric_Bailly_-_ManUtd.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00025','phil.jones','phil.jones@gmail.com@gmail.com','Phil Jones','https://upload.wikimedia.org/wikipedia/commons/f/fa/Phil_Jones_2018-06-28_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00026','harry.maguire','harry.maguire@gmail.com@gmail.com','Harry Maguire','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00027','paul.pogba','paul.pogba@gmail.com@gmail.com','Paul Pogba','https://upload.wikimedia.org/wikipedia/commons/b/be/Harry_Maguire_2018-07-11_1.jpg','I','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00028','edinson.cavani','edinson.cavani@gmail.com@gmail.com','Edinson Cavani','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Edinson_Cavani_2018.jpg/440px-Edinson_Cavani_2018.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00029','juan.mata','juan.mata@gmail.com@gmail.com','Juan Mata','https://upload.wikimedia.org/wikipedia/commons/7/70/Ukr-Spain2015_%286%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00030','anthony.martial','anthony.martial@gmail.com@gmail.com','Anthony Martial','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Anthony_Martial_27_September_2017_cropped.jpg/440px-Anthony_Martial_27_September_2017_cropped.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00031','marcus.rashford','marcus.rashford@gmail.com@gmail.com','Marcus Rashford','https://upload.wikimedia.org/wikipedia/commons/5/5e/Press_Tren_CSKA_-_MU_%283%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00032','mason.greenwood','mason.greenwood@gmail.com@gmail.com','Mason Greenwood','https://upload.wikimedia.org/wikipedia/commons/thumb/e/e0/Mason_Greenwood.jpeg/440px-Mason_Greenwood.jpeg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00033','lee.grant','lee.grant@gmail.com@gmail.com','Lee Grant','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8e/LeeGrant09.jpg/400px-LeeGrant09.jpg','A','M','0987654321','Ms','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00034','jesse.lingard','jesse.lingard@gmail.com@gmail.com','Jesse Lingard','https://upload.wikimedia.org/wikipedia/commons/thumb/1/11/Jesse_Lingard_2018-06-13_1.jpg/440px-Jesse_Lingard_2018-06-13_1.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00035','keylor.navas','keylor.navas@gmail.com@gmail.com','Keylor Navas','https://upload.wikimedia.org/wikipedia/commons/thumb/c/c3/Liver-RM_%288%29_%28cropped%29.jpg/440px-Liver-RM_%288%29_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00036','achraf.hakimi','achraf.hakimi@gmail.com@gmail.com','Achraf Hakimi','https://upload.wikimedia.org/wikipedia/commons/9/91/Iran-Morocco_by_soccer.ru_14_%28Achraf_Hakimi%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00037','presnel.kimpembe','presnel.kimpembe@gmail.com@gmail.com','Presnel Kimpembe','https://upload.wikimedia.org/wikipedia/commons/thumb/0/0e/Presnel_Kimpembe.jpg/400px-Presnel_Kimpembe.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00038','sergio.ramos','sergio.ramos@gmail.com@gmail.com','Sergio Ramos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg/440px-FC_RB_Salzburg_versus_Real_Madrid_%28Testspiel%2C_7._August_2019%29_09.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00039','marquinhos','marquinhos@gmail.com@gmail.com','Marquinhos','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg/440px-Brasil_conquista_primeiro_ouro_ol%C3%ADmpico_nos_penaltis_1039278-20082016-_mg_4916_%28cropped%29.jpg','A','M','0987654321','Mr','E');
insert into users (userid,username,email,displayname,imageurl,status,gender,phone,title,position) values ('00040','marco.verratti','marco.verratti@gmail.com@gmail.com','Marco Verratti','https://upload.wikimedia.org/wikipedia/commons/d/d0/Kiev-PSG_%289%29.jpg','A','M','0987654321','Mr','E');

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

insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'dashboard', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'admin', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'user', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'role', 7);
insert into rolemodules(roleid, moduleid, permissions) values ('it_support', 'audit_log', 7);


create table companies (
  companyid varchar(40) primary key,
  companyname varchar(255) not null,
  email varchar(255) not null,
  displayname varchar(255) not null,
  status char(1) not null,
  phone varchar(20),
  imageurl varchar(500),
  createdby varchar(40),
  createdat timestamp,
  updatedby varchar(40),
  updatedat timestamp,
  lastlogin timestamp
);
create table companyusers (
  companyid varchar(40) not null,
  userid varchar(40) not null,
  primary key (companyid, userid)
);

insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00001','gareth.bale.corp','gareth.bale@gmail.com','Gareth Bale Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/4/41/Liver-RM_%282%29_%28cropped%29.jpg/440px-Liver-RM_%282%29_%28cropped%29.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00002','cristiano.ronaldo.corp','cristiano.ronaldo@gmail.com','Cristiano Ronaldo Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/8/8c/Cristiano_Ronaldo_2018.jpg/400px-Cristiano_Ronaldo_2018.jpg','I','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00003','james.rodriguez.corp','james.rodriguez@gmail.com','James Rodríguez Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/7/79/James_Rodriguez_2018.jpg/440px-James_Rodriguez_2018.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00004','zinedine.zidane.corp','zinedine.zidane@gmail.com','Zinedine Zidane Corp','https://upload.wikimedia.org/wikipedia/commons/f/f3/Zinedine_Zidane_by_Tasnim_03.jpg','A','0987654321');
insert into companies (companyid,companyname,email,displayname,imageurl,status,phone) values ('00005','kaka.corp','kaka@gmail.com','Kaká Corp','https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Kak%C3%A1_visited_Stadium_St._Petersburg.jpg/500px-Kak%C3%A1_visited_Stadium_St._Petersburg.jpg','A','0987654321');

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

-- ----------------------------
-- Table structure for questions
-- ----------------------------
CREATE TABLE questions (
   id varchar(50) NULL,
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
INSERT INTO "questions" VALUES ('f4Z9KQ_aM', '', 'What will happen after the brunch?', 'f', '{3}', NULL, 'TEST_02', '{"The people will be served drinks.","An artist will show his drawings.","The guests will introduce themselves.","A poet will read a poem."}');
INSERT INTO "questions" VALUES ('Io2cK1Vtp', '', 'What do guests get to take home?', 'f', '{2}', NULL, 'TEST_02', '{Lunch,"A book of mythology","A collection of poems","A painting"}');
INSERT INTO "questions" VALUES ('QTjMlQ_ap', 'test2', 'You want to identify which tasks in the planned learning activity incorporate the use of the computer lab.
Which task should you identify?
This item is part of a case study. To view the case study information, click on the Case Study button below.', 'f', '{3}', NULL, 'TEST_01', '{"The students must hang the posters in the entrance hall of the school.","The students must assemble the images on a wall poster.","The students must select and edit  digital images that illustrate the theme or the event.","The students must split into small groups and each group must identify a theme or a historical event."}');
INSERT INTO "questions" VALUES ('gJL-XQVaM', '', 'As an extension activity to the planned learning activity, you ask some of the students to contact external experts to gather additional information about the historical events.
You want to ensure that the students use an appropriate means of communication when they contact the experts.
What should you advise the students to use?
This item is part of a case study. To view the case study information, click on the Case Study button below.', 'f', '{2}', NULL, 'TEST_01', '{"their personal mobile phone","an anonymous online forum"," a generic email account of the school"," their social network account"}');
INSERT INTO "questions" VALUES ('VFfZX1_ap', '', 'You teach a class of 20 students. Your classroom contains six computers.
You run an activity that is carried out mostly on paper.
Before the end of the activity, the students will require access to the computers.
Which lesson structure should provide the best learning experience for the students?', 'f', '{3}', NULL, 'TEST_01', '{"Create several different activities that require the computers at different stages during the lesson."," Execute the computer-based portion of the activity as a class demonstration.","Modify the tasks of the lesson so that only some of the students require access to the computers.","Have all of the students perform the same activity and instruct the students to take turns using the computers."}');
INSERT INTO "questions" VALUES ('tP8GXQVaM', '', 'What are the man and woman mainly discussing?', 'f', '{3}', NULL, 'TEST_02', '{"A vacation","A budget","A company policy","A conference"}');
INSERT INTO "questions" VALUES ('NQL_XQ_tp', '', 'How is the woman traveling?', 'f', '{1}', NULL, 'TEST_02', '{"By plane","By bus","By taxi","By car"}');
INSERT INTO "questions" VALUES ('9i-QvQ_tp', '', 'Why aren''t the man and woman going together?', 'f', '{2}', NULL, 'TEST_02', '{"The woman needs to arrive earlier.","The man has to work overtime.","The woman dislikes air travel.","The man has to go to the bank first."}');
INSERT INTO "questions" VALUES ('VzkRv1Vtp', '', 'What does the man have to do today?', 'f', '{2}', NULL, 'TEST_02', '{"Visit his lawyer","Get a massage","Go to the doctor.","Make an appointment."}');
INSERT INTO "questions" VALUES ('T_xCX1_ap', '', 'What can be inferred from the conversation?', 'f', '{2}', NULL, 'TEST_02', '{"The woman is the man''s receptionist.","The lawyer works in the same building.","The woman has no deadlines today.","The man and woman have a meeting this afternoon."}');
INSERT INTO "questions" VALUES ('VWuoK1_tM', '', 'What does the woman offer to do for the man?', 'f', '{0}', NULL, 'TEST_02', '{"Answer his phone","Call his lawyer","Pick up the newspaper","Take notes at the meeting"}');
INSERT INTO "questions" VALUES ('aZOHd1Vtp', '', 'What should the passengers do before exiting the ship?', 'f', '{2}', NULL, 'TEST_02', '{"Welcome the visitors","Check the time","Collect their personal items","Take a picture"}');
INSERT INTO "questions" VALUES ('S6J0KQVaM', '', 'What does the speaker imply?', 'f', '{1}', NULL, 'TEST_02', '{"The water was rough.","The weather was poor yesterday.","The tour went faster than usual.","There is only one way to exit."}');
INSERT INTO "questions" VALUES ('cJWfK1_ap', '', 'What will happen in five minutes?', 'f', '{0}', NULL, 'TEST_02', '{"The ship will arrive at the dock.","The passengers will go shopping.","The passengers will take photos of the ship.","The market will open."}');
INSERT INTO "questions" VALUES ('f_XaK1VtM', '', 'How often does this event take place?', 'f', '{2}', NULL, 'TEST_02', '{Monthly,"Four times a year","Once a year","Every four years"}');
INSERT INTO "questions" VALUES ('UkkSd1_tM', '', 'Where would you hear this type of talk?', 'f', '{2}', NULL, 'TEST_02', '{"In an airport","At a bus station","On a telephone","On television"}');
INSERT INTO "questions" VALUES ('cDiEdQ_tp', '', 'Who should visit the website?', 'f', '{1}', NULL, 'TEST_02', '{"Staff members of Speedy Reservations","Callers looking for special schedules","Callers looking for staff members","People looking for fight times"}');
INSERT INTO "questions" VALUES ('hs2iK1VtM', '', 'What is the quickest way to get connected to a representative?', 'f', '{1}', NULL, 'TEST_02', '{"Hang up and call again","Call back later"," Wait patiently","Call a different number"}');
INSERT INTO "questions" VALUES ('IzqAd1Vap', '', 'Do you think the boss will take us out for lunch today?', 'f', '{1}', '{part2}', 'TEST_02', '{"I had a hamburger and fries.","He''s probably too busy today.","I didn''t take the bus."}');
INSERT INTO "questions" VALUES ('owb_KQ_tM', '', 'Did you fax the letter to the client?', 'f', '{1}', '{part2}', 'TEST_02', '{"I''ll type it after lunch.","No, I sent it by email.","I collected some important facts."}');
INSERT INTO "questions" VALUES ('Lk7LKQ_tp', '', 'Did you have to wait very long?', 'f', '{0}', '{part2}', 'TEST_02', '{"No more than an hour.","I hate long line-ups.","I''ve gained twenty pounds."}');
INSERT INTO "questions" VALUES ('7t9IrRVap', '', 'What is the last step to be followed to determine whether a person is a RFI and thus has reporting obligations', 'f', '{2}', NULL, 'FATCA_Mock_Test', '{"Is the Entity a Financial Institution?","Is it an Entity?","Is the Financial Institution in India?","Is the Financial Institution a Non-Reporting Financial Institution?"}');
INSERT INTO "questions" VALUES ('dlbV9RVtp', '', 'An Active NFE should have', 'f', '{2}', NULL, 'FATCA_Mock_Test', '{"if the entity during the last year: More than (or equal to) 25% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 50% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 50% of the entity’s gross income consisted of Active Income","if the entity during the last year: More than (or equal to) 75% of the entity’s gross income consisted of Active Income"}');
INSERT INTO "questions" VALUES ('_1RQ9RVap', '', 'An listed entity under FATCA, is technically classified as', 'f', '{2}', NULL, 'FATCA_Mock_Test', '{"Active NFE","Account Holder","Active NFFE",nominee}');
INSERT INTO "questions" VALUES ('OZdm9R_aM', '', 'Under FATCA, each country entered into a separate bilateral intergovernmental agreement with the United States and what objective was achieved', 'f', '{2}', NULL, 'FATCA_Mock_Test', '{"The local financial institutions complies with FATCA reporting standards  and will not breach local data protection laws ","Financial institutions in complying with FATCA, will not breach local data protection laws only","None of these","The local financial institutions complies with FATCA reporting standards only"}');
INSERT INTO "questions" VALUES ('csVonL_tp', '', 'Which of the following is an exempt beneficial owner', 'f', '{2}', NULL, 'FATCA_Mock_Test', '{"All of these","Any International Organisation or any wholly owned agency or instrumentality thereof","Any Foreign Central Bank of Issue","Any Foreign Government, any political subdivision of a Foreign Government, or any wholly owned agency or instrumentality of any one or more of the foregoing"}');

-- ----------------------------
-- Table structure for tests
-- ----------------------------
CREATE TABLE "tests" (
  "testid" varchar ,
  "title" varchar ,
  "effectivedate" timestamp(6),
  "questions" _varchar ,
  "tags" _varchar ,
  "categoryid" varchar 
);

-- ----------------------------
-- Records of tests
-- ----------------------------
INSERT INTO "tests" VALUES ('2E1cy1_tp', 'Q1', '2023-06-04 00:00:00', '{aZOHd1Vtp}', NULL, 'FA_01');
INSERT INTO "tests" VALUES ('f6MRyQ_tM', 'Q1', '2023-06-04 00:00:00', '{aZOHd1Vtp}', NULL, 'FA_01');
INSERT INTO "tests" VALUES ('vEMky1_tp', 'Q2', '2023-06-04 00:00:00', '{S6J0KQVaM}', NULL, 'FA_01');
-- ----------------------------
-- Table structure for tickets
-- ----------------------------
DROP TABLE IF EXISTS "tickets";
CREATE TABLE "tickets" (
   "id" varchar(40) ,
   "title" varchar(100) ,
   "body" text ,
   "categoryid" varchar(40) ,
   "requestor" varchar(20) ,
   "requested_at" timestamp(6),
   "approver" varchar(20) ,
   "approved_at" timestamp(6),
   "assignee" varchar(40) ,
   "completed_at" timestamp(6),
   "status" varchar(1) 
);
-- ----------------------------
-- Records of tickets
-- ----------------------------
INSERT INTO "tickets" VALUES ('s319KmVtM', 'Sample Title', 'Sample body text.', 'sampleCategoryId', 'sampleRequestor', '2023-06-14 10:30:00', 'sampleApprover', '2023-06-14 11:00:00', 'sampleAssignee', '2023-06-14 12:00:00', 'S');
INSERT INTO "tickets" VALUES ('iTJWdm_aM', 'Sample Ticket', 'This is a sample ticket description.', 'category456', 'john', '2023-06-14 10:30:00', 'approver123', '2023-06-14 11:00:00', 'assignee456', '2023-06-14 12:00:00', 'S');

/*
alter table userroles add foreign key (userid) references users (userid);
alter table userroles add foreign key (roleid) references roles (roleid);

alter table modules add foreign key (parent) references modules (moduleid);

alter table rolemodules add foreign key (roleid) references roles (roleid);
alter table rolemodules add foreign key (moduleid) references modules (moduleid);
*/
