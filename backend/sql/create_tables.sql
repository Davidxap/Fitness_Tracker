
-- 1) Users con age y weight
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(100) NOT NULL,
  age INTEGER,
  weight FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2) Sessions con name y observations
CREATE TABLE workout_sessions (
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(100) NOT NULL,
  date DATE NOT NULL,
  duration_minutes INTEGER NOT NULL,
  observations TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3) Exercises
CREATE TABLE exercises (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  muscle_group VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4) Session_Exercises
CREATE TABLE session_exercises (
  id SERIAL PRIMARY KEY,
  session_id INTEGER REFERENCES workout_sessions(id) ON DELETE CASCADE,
  exercise_id INTEGER REFERENCES exercises(id) ON DELETE CASCADE,
  sets INTEGER NOT NULL,
  reps INTEGER NOT NULL,
  weight FLOAT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Seed 30 exercises (name, description, muscle_group)
INSERT INTO exercises (name, description, muscle_group) VALUES
-- 1–10
('Squat', 'Basic squat', 'Legs'),
('Bench Press', 'Barbell bench press', 'Chest'),
('Deadlift', 'Barbell deadlift', 'Back'),
('Pull-Up', 'Bodyweight pull-up', 'Back'),
('Push-Up', 'Bodyweight push-up', 'Chest'),
('Overhead Press', 'Standing barbell press', 'Shoulders'),
('Barbell Row', 'Bent-over row', 'Back'),
('Lunges', 'Walking lunges', 'Legs'),
('Bicep Curl', 'Dumbbell curl', 'Arms'),
('Tricep Dip', 'Bodyweight dip', 'Arms'),
-- 11–20
('Plank', 'Core hold', 'Core'),
('Jumping Jack', 'Cardio warm-up', 'Full Body'),
('Lat Pulldown', 'Cable lat pulldown', 'Back'),
('Leg Press', 'Machine leg press', 'Legs'),
('Chest Fly', 'Dumbbell fly', 'Chest'),
('Leg Curl', 'Machine curl', 'Legs'),
('Calf Raise', 'Standing calf raise', 'Legs'),
('Shoulder Shrug', 'Barbell shrug', 'Shoulders'),
('Hammer Curl', 'Dumbbell hammer curl', 'Arms'),
('Cable Row', 'Seated cable row', 'Back'),
-- 21–30
('Hip Thrust', 'Barbell hip thrust', 'Glutes'),
('Front Squat', 'Barbell front squat', 'Legs'),
('Incline Bench', 'Incline press', 'Chest'),
('Decline Bench', 'Decline press', 'Chest'),
('Dumbbell Press', 'Dumbbell chest press', 'Chest'),
('Face Pull', 'Cable face pull', 'Back'),
('Russian Twist', 'Core twist', 'Core'),
('Mountain Climber', 'Cardio/core', 'Full Body'),
('Burpee', 'Full body cardio', 'Full Body'),
('Side Plank', 'Core side hold', 'Core');
