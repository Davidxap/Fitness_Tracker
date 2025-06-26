// frontend/src/pages/Exercises.jsx
import React, { useEffect, useState } from 'react'
import api from '../api/api'

export default function Exercises() {
  const [list, setList] = useState([])

  useEffect(() => {
    api.get('/exercises').then(res => setList(res.data))
  }, [])

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-4">Exercises</h1>
      <ul className="space-y-2">
        {list.map(ex => (
          <li
            key={ex.id}
            className="bg-white p-4 rounded shadow flex justify-between items-center"
          >
            <div>
              <p><span className="font-semibold">Name:</span> {ex.name}</p>
              <p><span className="font-semibold">Muscle Group:</span> {ex.muscle_group}</p>
              {ex.description && (
                <p><span className="font-semibold">Description:</span> {ex.description}</p>
              )}
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}