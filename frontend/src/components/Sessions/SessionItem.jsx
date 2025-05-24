import React from 'react'

export default function SessionItem({ session, onEdit, onDelete }) {
  return (
    <div className="bg-white p-4 rounded shadow transition transform hover:scale-[1.01]">
      <div className="flex justify-between items-center">
        <div>
          <h3 className="text-lg font-semibold">
            Date: {session.date} — {session.duration_minutes} min
          </h3>
          {session.observations && (
            <p className="text-sm text-gray-500">
              Obs: {session.observations}
            </p>
          )}
        </div>
        <div className="space-x-2">
          <button
            onClick={onEdit}
            className="px-3 py-1 bg-yellow-500 text-white rounded transition transform active:scale-95"
          >
            Edit
          </button>
          <button
            onClick={onDelete}
            className="px-3 py-1 bg-red-600 text-white rounded transition transform active:scale-95"
          >
            Delete
          </button>
        </div>
      </div>
      {session.exercises.length > 0 && (
        <ul className="mt-3 list-disc list-inside text-sm">
          {session.exercises.map(se => (
            <li key={se.id}>
              {/* name resolved client-side */}
              Exercise ID {se.exercise_id}: {se.sets}×{se.reps} @ {se.weight}kg
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}
