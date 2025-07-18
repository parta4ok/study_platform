BEGIN;

CREATE SCHEMA IF NOT EXISTS kvs;

CREATE TABLE IF NOT EXISTS kvs.question_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS kvs.topics (
    id SERIAL PRIMARY KEY,
    topic_id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS kvs.questions (
    id BIGSERIAL PRIMARY KEY,
    question_id SERIAL NOT NULL UNIQUE,
    question_type_id SERIAL NOT NULL REFERENCES kvs.question_types(id),
    topic_id SERIAL NOT NULL REFERENCES kvs.topics(topic_id),
    subject VARCHAR(255) NOT NULL,
    variants TEXT[] NOT NULL,
    correct_answers TEXT[] NOT NULL,
    usage_count INT NOT NULL DEFAULT 0
);

END;