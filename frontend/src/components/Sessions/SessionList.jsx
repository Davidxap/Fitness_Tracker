import React from 'react'
import SessionItem from './SessionItem'

export default function SessionList({ sessions, onEdit, onDelete }) {
  if (sessions.length === 0) {
    return <p className="text-gray-600">You have no sessions yet.</p>
  }
  return (
    <div className="space-y-4">
      {sessions.map(s => (
        <SessionItem
          key={s.id}
          session={s}
          onEdit={() => onEdit(s)}
          onDelete={() => onDelete(s.id)}
        />
      ))}
    </div>
  )
}
