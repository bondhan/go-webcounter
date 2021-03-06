CREATE SEQUENCE m_visitor_seq;
CREATE SEQUENCE m_visitor_counter;

CREATE TABLE m_visitor (
  id int check (id > 0) NOT NULL DEFAULT NEXTVAL ('m_visitor_seq'),
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  counter int check (counter >= 0) NOT NULL DEFAULT NEXTVAL ('m_visitor_counter'),  
  PRIMARY KEY (id),
  CONSTRAINT m_visitor_counter_unique UNIQUE  (counter)
);

CREATE INDEX m_visitor_deleted_at ON m_visitor (deleted_at);

insert into m_visitor (id, created_at, updated_at, counter ) values (1, now(), now(), 0);