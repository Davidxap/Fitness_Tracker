-- 1. Users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 2. WorkoutSessions
CREATE TABLE IF NOT EXISTS workout_sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    date DATE NOT NULL,
    duration_minutes INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 3. Exercises (catálogo)
CREATE TABLE IF NOT EXISTS exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    muscle_group VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. SessionExercises (unión N–N con datos)
CREATE TABLE IF NOT EXISTS session_exercises (
    id SERIAL PRIMARY KEY,
    session_id INTEGER REFERENCES workout_sessions(id) ON DELETE CASCADE,
    exercise_id INTEGER REFERENCES exercises(id) ON DELETE CASCADE,
    sets INTEGER NOT NULL,
    reps INTEGER NOT NULL,
    weight FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- Insterto un listado de ejercicios tipicos
INSERT INTO exercises (name, description, muscle_group)
SELECT 'Squat', 'Basic squat exercise', 'Legs'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Squat');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Bench Press', 'Barbell bench press', 'Chest'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Bench Press');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Deadlift', 'Barbell deadlift', 'Back'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Deadlift');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Pull-Up', 'Bodyweight pull-up', 'Back'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Pull-Up');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Push-Up', 'Bodyweight push-up', 'Chest'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Push-Up');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Overhead Press', 'Standing barbell press', 'Shoulders'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Overhead Press');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Barbell Row', 'Bent-over barbell row', 'Back'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Barbell Row');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Lunges', 'Walking lunges', 'Legs'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Lunges');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Bicep Curl', 'Dumbbell bicep curl', 'Arms'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Bicep Curl');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Tricep Dip', 'Bodyweight tricep dip', 'Arms'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Tricep Dip');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Plank', 'Core stabilization exercise', 'Core'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Plank');

INSERT INTO exercises (name, description, muscle_group)
SELECT 'Jumping Jack', 'Cardio warm-up', 'Full Body'
WHERE NOT EXISTS (SELECT 1 FROM exercises WHERE name='Jumping Jack');
