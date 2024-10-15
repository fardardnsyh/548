PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE Users (
  uuid TEXT PRIMARY KEY,
  username text NOT NULL unique CHECK(LENGTH(username) <= 40),
  email text NOT NULL unique,
  password_hash text NOT NULL
);
INSERT INTO Users VALUES('34022c09-4f77-448c-90b7-55e4934c2b71','EvilgeniuS1982','stevenpearson1982@gmail.com','$2a$14$fSl2CqNUNsDQhzK6PWOiu.f7MNnZf9izXbuKGmHDkkkbs60zzqLN.');
CREATE TABLE Session (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid text NOT NULL,
  auth_uuid TEXT,
  FOREIGN KEY (auth_uuid) REFERENCES user (uuid)
);
INSERT INTO Session VALUES(2,'f0a065f1-a337-47c9-8aa1-d143a273f5a6','34022c09-4f77-448c-90b7-55e4934c2b71');
CREATE TABLE Cat_posts (
  id integer primary key autoincrement,
  post_id integer not null,
  category TEXT not null,
  foreign key (post_id) references posts (id)
);
CREATE TABLE Posts (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL,
  body TEXT NOT NULL,
  user_id TEXT NOT NULL,
  amount_likes INTEGER NOT NULL,
  amount_dislikes INTEGER NOT NULL,
  categorys TEXT NOT NULL,
  posted_on TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES user (uuid)
);
INSERT INTO Posts VALUES(1,'gghhyydt',',ljsbdljjbsb','34022c09-4f77-448c-90b7-55e4934c2b71',0,0,'','06-01-2022 18:10:20');
CREATE TABLE Comments (
  id integer PRIMARY KEY,
  comment text,
  auth_id TEXT NOT NULL,
  post_id integer NOT NULL,
  amount_likes INTEGER NOT NULL DEFAULT 0,
  amount_dislikes INTEGER NOT NULL DEFAULT 0,
  commented_on TEXT NOT NULL,
  FOREIGN KEY (auth_id) REFERENCES user (uuid) FOREIGN KEY (post_id) REFERENCES posts (id)
);
INSERT INTO Comments VALUES(1,'hello','34022c09-4f77-448c-90b7-55e4934c2b71',1,0,0,'06-01-2022 18:11:08');
INSERT INTO Comments VALUES(2,'hello2','34022c09-4f77-448c-90b7-55e4934c2b71',1,0,0,'06-01-2022 18:12:15');
CREATE TABLE likes (id integer primary key autoincrement, user_uuid text not null references Users (uuid), post_id integer references Posts (id), com_id integer references Comments (id));
CREATE TABLE dislikes (id integer primary key autoincrement, 
user_uuid text not null references Users (uuid), 
post_id integer references Posts (id),
com_id integer references Comments (id)
);
DELETE FROM sqlite_sequence;
INSERT INTO sqlite_sequence VALUES('Session',2);
INSERT INTO sqlite_sequence VALUES('Posts',1);
COMMIT;
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
COMMIT;
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE likes (id integer primary key autoincrement, user_uuid text not null references Users (uuid), post_id integer not null references Posts (id), com_id integer not null references Comments (id));
COMMIT;
Cat_posts  Comments   Posts      Session    Users      dislikes   likes    
Cat_posts  Comments   Posts      Session    Users      dislikes   likes    
