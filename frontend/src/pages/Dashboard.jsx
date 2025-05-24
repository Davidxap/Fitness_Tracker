import React, { useEffect, useState } from 'react'
import api from '../api/api'
import useAuth from '../hooks/useAuth'

export default function Dashboard() {
  const { user } = useAuth()
  const [sessions, setSessions] = useState([])

  useEffect(() => {
    const load = async () => {
      try {
        // Ya tenemos user
        // Cargamos sesiones
        const [sRes, seRes] = await Promise.all([
          api.get('/sessions'),
          api.get('/session-exercises')
        ])
        const own = sRes.data.filter(s => s.user_id === user.id)
        const combined = own.map(s => ({
          ...s,
          exercises: seRes.data.filter(se => se.session_id === s.id)
        }))
        setSessions(combined)
      } catch (err) {
        console.error(err)
      }
    }
    if (user) load()
  }, [user])

  if (!user) return <p className="p-6">Loading dashboard...</p>

  return (
    <div className="p-6">
      <h1 className="text-3xl font-bold mb-2">
        Welcome, {user.name}!
      </h1>
      <p className="mb-6">Email: {user.email}</p>

      <h2 className="text-2xl font-semibold mb-2">Your Sessions</h2>
      {sessions.length === 0 ? (
        <p>You have no saved sessions yet.</p>
      ) : (
        <div className="space-y-4">
          {sessions.map(s => (
            <div key={s.id} className="bg-white p-4 rounded shadow">
              <div className="font-medium">
                {s.date} — {s.duration_minutes} min
              </div>
              {s.exercises.length > 0 && (
                <ul className="mt-2 list-disc list-inside text-sm">
                  {s.exercises.map(se => (
                    <li key={se.id}>
                      ID {se.exercise_id}: {se.sets}×{se.reps} @ {se.weight}kg
                    </li>
                  ))}
                </ul>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
