// frontend/src/components/Sessions/SessionItem.jsx
import React from 'react'

/**
 * Muestra una sesión con:
 * - Name
 * - Date
 * - Duration
 * - Observations
 * - Lista de ejercicios (Exercise, Sets, Reps, Weight kg)
 */
export default function SessionItem({
  session,
  formatDate,
  onEdit,
  onDelete
}) {
  return (
    <div className="bg-white p-4 rounded shadow mb-4">
      {/* Sesión - encabezado */}
      <div className="flex justify-between items-start mb-3">
        <div>
          <p><span className="font-semibold">Name:</span> {session.name}</p>
          <p><span className="font-semibold">Date:</span> {formatDate(session.date)}</p>
          <p><span className="font-semibold">Duration:</span> {session.duration_minutes} min</p>
        </div>
        <div className="space-x-2">
          <button
            onClick={onEdit}
            className="px-3 py-1 bg-yellow-500 text-white rounded transition active:scale-95"
          >
            Edit
          </button>
          <button
            onClick={onDelete}
            className="px-3 py-1 bg-red-600 text-white rounded transition active:scale-95"
          >
            Delete
          </button>
        </div>
      </div>

      {/* Observations */}
      {session.observations && (
        <p className="mb-3">
          <span className="font-semibold">Observations:</span> {session.observations}
        </p>
      )}

      {/* Exercises list */}
      {session.exercises.length > 0 && (
        <div>
          <p className="font-semibold mb-1">Exercises:</p>
          <ul className="list-disc list-inside text-sm space-y-1">
            {session.exercises.map((se, idx) => (
              <li key={idx}>
                <span className="font-semibold">Exercise:</span> {se.exercise_name} —{' '}
                <span className="font-semibold">Sets:</span> {se.sets},{' '}
                <span className="font-semibold">Reps:</span> {se.reps},{' '}
                <span className="font-semibold">Weight:</span> {se.weight} kg
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  )
}
