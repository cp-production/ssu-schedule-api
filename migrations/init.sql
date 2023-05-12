CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    full_name varchar not null,
    short_name varchar not null,
    url varchar not null
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    education_form varchar not null,
    group_num varchar not null,
    department_id int not null
);

CREATE TABLE subgroups (
    id SERIAL PRIMARY KEY,
    subgroup_name varchar not null,
    group_id SERIAL not null
);

CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    full_name varchar not null
);

CREATE TABLE students_schedule (
    id SERIAL PRIMARY KEY,
    group_id int not null,
    day_num varchar not null, 
    week_type varchar not null, 
    lesson_type varchar not null,
    lesson_name varchar not null,
    teacher varchar not null,
    lesson_place varchar not null,
    subgroup_name varchar not null    
);

CREATE TABLE teachers_schedule (
    id SERIAL PRIMARY KEY,
    day_num varchar not null,
    week_type varchar not null,
    lesson_type varchar not null,
    lesson_name varchar not null,
    group_num varchar not null,
    lesson_place varchar not null,
    subgroup_name varchar not null
)
ALTER TABLE groups
ADD CONSTRAINT groups_dep_id_fkey
FOREIGN KEY (department_id)
REFERENCES departments(id);

ALTER TABLE subgroups
ADD CONSTRAINT subgroups_group_id_fkey
FOREIGN KEY (group_id)
REFERENCES groups(id);
