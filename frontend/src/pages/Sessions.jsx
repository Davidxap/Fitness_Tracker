import React, { useEffect, useState } from 'react'
import api from '../api/api'
import useAuth from '../hooks/useAuth'
import SessionList from '../components/Sessions/SessionList'
import SessionForm from '../components/Sessions/SessionForm'

export default function Sessions() {
  const { user } = useAuth()
  const [sessions, setSessions] = useState([])
  const [editing, setEditing] = useState(null)
  const [showForm, setShowForm] = useState(false)
  const [message, setMessage] = useState('')

  const loadSessions = async () => {
    try {
      const [sRes, seRes] = await Promise.all([
        api.get('/sessions'),
        api.get('/session-exercises')
      ])
      // Filtramos por user.id
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

  useEffect(() => {
    if (user) loadSessions()
  }, [user])

  const handleCreate = () => {
    setEditing(null)
    setShowForm(true)
  }

  const handleSave = async data => {
    try {
      let sessionId
      if (editing) {
        await api.put(`/sessions/${editing.id}`, {
          date: data.date,
          duration_minutes: data.duration_minutes,
          observations: data.observations,
          user_id: user.id
        })
        sessionId = editing.id
        // Borramos previos
        await Promise.all(
          editing.exercises.map(e =>
            api.delete(`/session-exercises/${e.id}`)
          )
        )
      } else {
        const r = await api.post('/sessions', {
          date: data.date,
          duration_minutes: data.duration_minutes,
          observations: data.observations,
          user_id: user.id
        })
        sessionId = r.data.id
      }
      // Creamos nuevos session-exercises
      await Promise.all(
        data.exercises.map(ent =>
          api.post('/session-exercises', {
            session_id: sessionId,
            exercise_id: ent.exercise_id,
            sets: ent.sets,
            reps: ent.reps,
            weight: ent.weight
          })
        )
      )
      setMessage('Session saved successfully!')
      setTimeout(() => setMessage(''), 3000)
      setShowForm(false)
      setEditing(null)
      loadSessions()
    } catch (err) {
      console.error(err)
    }
  }

  const handleEdit = session => {
    setEditing(session)
    setShowForm(true)
  }

  const handleDelete = async id => {
    if (!window.confirm('Delete this session?')) return
    try {
      const toDel = sessions.find(s => s.id === id).exercises
      await Promise.all(toDel.map(e => api.delete(`/session-exercises/${e.id}`)))
      await api.delete(`/sessions/${id}`)
      loadSessions()
    } catch (err) {
      console.error(err)
    }
  }

  if (!user) return <p className="p-6">Loading...</p>

  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-3xl font-bold">Your Sessions</h1>
        <button
          onClick={handleCreate}
          className="bg-green-600 text-white px-4 py-2 rounded transition transform active:scale-95"
        >
          + New Session
        </button>
      </div>

      {message && (
        <div className="mb-4 p-3 bg-green-100 text-green-800 rounded">
          {message}
        </div>
      )}

      {showForm && (
        <SessionForm
          key={editing ? editing.id : 'new'}
          initialData={editing}
          onSave={handleSave}
          onCancel={() => {
            setShowForm(false)
            setEditing(null)
          }}
        />
      )}

      <SessionList
        sessions={sessions}
        onEdit={handleEdit}
        onDelete={handleDelete}
      />
    </div>
  )
}
