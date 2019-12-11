CREATE TABLE public.t_customer
(
  id serial NOT NULL PRIMARY KEY,
  name character varying(255) NOT NULL,
  document character varying(20) NOT NULL,

  created_at timestamp(0) without time zone NOT NULL DEFAULT now(),
  updated_at timestamp(0) without time zone
);

CREATE TABLE test_table (
	id int(20) NOT NULL,
	name varchar(20) COLLATE utf8_bin NOT NULL,
	created_at datetime DEFAULT NULL,
	updated_at datetime DEFAULT NULL,
	sort_no varchar(45) COLLATE utf8_bin DEFAULT NULL,
	code varchar(45) COLLATE utf8_bin DEFAULT NULL,
	PRIMARY KEY (id)
  ) 