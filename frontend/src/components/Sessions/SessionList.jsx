// frontend/src/components/Sessions/SessionList.jsx
import React from 'react'
import SessionItem from './SessionItem'

/**
 * Lista todas las sesiónes con formato
 */
export default function SessionList({
  sessions,
  onEdit,
  onDelete,
  formatDate
}) {
  if (sessions.length === 0) {
    return <p className="text-gray-600">You have no sessions yet.</p>
  }
  return (
    <div>
      {sessions.map(s => (
        <SessionItem
          key={s.id}
          session={s}
          formatDate={formatDate}
          onEdit={() => onEdit(s)}
          onDelete={() => onDelete(s.id)}
        />
      ))}
    </div>
  )
}
