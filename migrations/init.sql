CREATE TABLE departments (
    id SERIAL PRIMARY KEY,
    url varchar not null,
    fullName varchar not null
);

CREATE TABLE groups (
    id SERIAL PRIMARY KEY,
    edForm varchar not null,
    groupNum varchar not null,
    dep_id int not null
);

CREATE TABLE subgroups (
    id SERIAL PRIMARY KEY,
    subgroupName varchar not null,
    group_id SERIAL not null
);

CREATE TABLE teachers (
    id SERIAL PRIMARY KEY,
    fullName varchar not null
);

CREATE TABLE studentsSchedule (
    id SERIAL PRIMARY KEY,
    group_id int not null,
    dayNum varchar not null, 
    weekType varchar not null, 
    lessonType varchar not null,
    lessonName varchar not null,
    teacher varchar not null,
    lessonPlace varchar not null,
    subgroupName varchar not null    
);

CREATE TABLE teachersSchedule (
    id SERIAL PRIMARY KEY,
    dayNum varchar not null,
    weekType varchar not null,
    lessonType varchar not null,
    lessonName varchar not null,
    groupNum varchar not null,
    lessonPlace varchar not null,
    subgroupName varchar not null
)
ALTER TABLE groups
ADD CONSTRAINT groups_dep_id_fkey
FOREIGN KEY (dep_id)
REFERENCES departments(id);

ALTER TABLE subgroups
ADD CONSTRAINT subgroups_group_id_fkey
FOREIGN KEY (group_id)
REFERENCES groups(id);
