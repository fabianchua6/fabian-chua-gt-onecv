-- public.students definition

-- Drop table

-- DROP TABLE public.students;

CREATE TABLE public.students (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" text NULL,
	email text NULL,
	suspended bool NULL DEFAULT false,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT students_email_key UNIQUE (email),
	CONSTRAINT students_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_students_deleted_at ON public.students USING btree (deleted_at);


-- public.teachers definition

-- Drop table

-- DROP TABLE public.teachers;

CREATE TABLE public.teachers (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	"name" text NULL,
	email text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT teachers_email_key UNIQUE (email),
	CONSTRAINT teachers_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_teachers_deleted_at ON public.teachers USING btree (deleted_at);


-- public.registrations definition

-- Drop table

-- DROP TABLE public.registrations;

CREATE TABLE public.registrations (
	teacher_email text NOT NULL,
	student_email text NOT NULL,
	CONSTRAINT registrations_pkey PRIMARY KEY (teacher_email, student_email),
	CONSTRAINT registrations_student_email_fk FOREIGN KEY (student_email) REFERENCES public.students(email) ON UPDATE CASCADE,
	CONSTRAINT registrations_teacher_email_fk FOREIGN KEY (teacher_email) REFERENCES public.teachers(email) ON UPDATE CASCADE
);


INSERT INTO public.teachers (id,"name",email,created_at,updated_at,deleted_at) VALUES
	 ('58c891b5-0396-4a3e-9d1b-2fb3af099bbd',NULL,'teachereyung@gmail.com',NULL,NULL,NULL),
	 ('16fab90f-5caf-44e5-82f0-095c2ec9b7b0',NULL,'teachersensei@gmail.com',NULL,NULL,NULL),
	 ('212a0083-f5c6-461b-a9d4-401a29916a02',NULL,'teacherpaul@gmail.com',NULL,NULL,NULL),
	 ('5264a0ed-28aa-4d34-abed-16f57713ed77',NULL,'teacheranne@gmail.com',NULL,NULL,NULL),
	 ('11f0ff78-226a-4d7f-bc68-a1a51a4e522e',NULL,'teachercarl@gmail.com',NULL,NULL,NULL),
	 ('0d169257-c1ba-42ea-8794-a530b52375df',NULL,'teacherhelen@gmail.com',NULL,NULL,NULL),
	 ('12293dd1-8608-4123-83fb-bf0948c74e37',NULL,'teacherjoe@gmail.com',NULL,NULL,NULL),
	 ('92b6fc8d-a099-4f5f-9518-67d14cb9ffa1',NULL,'teacherken@gmail.com',NULL,NULL,NULL),
	 ('f0878baa-8e71-41db-8109-d7ddbf30d6f5',NULL,'teacherbrian@gmail.com',NULL,NULL,NULL),
	 ('8f4b1dd4-a28a-4140-ac97-aa3e3389797b',NULL,'teachernancy@gmail.com',NULL,NULL,NULL);

INSERT INTO public.students (id,"name",email,suspended,created_at,updated_at,deleted_at) VALUES
	 ('eb36205e-7968-4575-92f9-e998b00c15e1',NULL,'studentjon@gmail.com',false,NULL,NULL,NULL),
	 ('3994aa8a-c3df-4a14-98e3-81cdbbfc4617',NULL,'studenthon@gmail.com',false,NULL,NULL,NULL),
	 ('1011bfb0-fe18-4c47-8088-be663d91ef3e',NULL,'student_only_under_teacher_ken@gmail.com',false,NULL,NULL,NULL),
	 ('efd11f33-4d32-4428-bd23-2611871ca44d',NULL,'studentbob@gmail.com',false,NULL,NULL,NULL),
	 ('4f4a0343-91be-4545-bbd5-0661eda05d8c',NULL,'studentagnes@gmail.com',false,NULL,NULL,NULL),
	 ('f56e5740-4552-4d0b-8aba-846dec5fef8a',NULL,'commonstudent1@gmail.com',false,NULL,NULL,NULL),
	 ('d8962b05-1c6f-4591-b196-1665e624fd20',NULL,'commonstudent2@gmail.com',false,NULL,NULL,NULL),
	 ('c36bcc12-8d5a-4a35-895a-f033bf4f71e8',NULL,'studentfabian@gmail.com',false,NULL,NULL,NULL),
	 ('727fc8cb-6885-4121-b252-b3e35104a973',NULL,'studentmary@gmail.com',false,NULL,NULL,NULL),
	 ('ebc08138-55cf-40b3-8383-e6f7cf331393',NULL,'studentmiche@gmail.com',false,NULL,NULL,NULL);

INSERT INTO public.registrations (teacher_email,student_email) VALUES
	 ('teacherjoe@gmail.com','commonstudent2@gmail.com'),
	 ('teacherjoe@gmail.com','commonstudent1@gmail.com'),
	 ('teacherken@gmail.com','commonstudent1@gmail.com'),
	 ('teacherken@gmail.com','commonstudent2@gmail.com'),
	 ('teacherken@gmail.com','student_only_under_teacher_ken@gmail.com');
