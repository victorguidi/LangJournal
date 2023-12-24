-- PostgreSQL database
CREATE DATABASE LangJournal;

CREATE TYPE QuestionCategory AS ENUM ('Base', 'Advanced', 'Expert', 'Master', 'Easy', 'Medium', 'Hard', 'Very Hard');

CREATE TABLE IF NOT EXISTS DefaultQuestions (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL  ,
    question TEXT NOT NULL,
    category QuestionCategory NOT NULL,
    difficulty INTEGER NOT NULL
);


CREATE TABLE IF NOT EXISTS Answers (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL,
    question_id UUID NOT NULL,
    answer TEXT NOT NULL,
    is_correct BOOLEAN NOT NULL

    CONSTRAINT fk_question
      FOREIGN KEY(question_id)
      REFERENCES defaultquestions(id)
      ON DELETE CASCADE
);

-- Test data
INSERT INTO defaultquestions (question, category, difficulty) values ('What is the capital of the United States?', 'Base', 1);

INSERT INTO answers (question_id, answer, is_correct) values ('95d7c1c8-2334-4516-870c-c714326583e3', 'Washington, D.C.', true);
INSERT INTO answers (question_id, answer, is_correct) values ('95d7c1c8-2334-4516-870c-c714326583e3', 'Washington, D.D.', false);
INSERT INTO answers (question_id, answer, is_correct) values ('95d7c1c8-2334-4516-870c-c714326583e3', 'Washington, DC.C.', false);
INSERT INTO answers (question_id, answer, is_correct) values ('95d7c1c8-2334-4516-870c-c714326583e3', 'Washington, C.C.', false);
