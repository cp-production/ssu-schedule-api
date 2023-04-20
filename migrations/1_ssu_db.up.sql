CREATE TABLE departments (
    url varchar not null,
    name varchar not null
);

CREATE TABLE groups (
    dep varchar not null,
    ed-form varchar not null,
    groupNum varchar not null,
    CONSTRAINT constraint_groups PRIMARY KEY(dep, edd-form)
)
