import React from 'react';

export default function FormInput({ label, type, value, onChange, name, error }) {
  return (
    <div className="mb-4">
      <label className="block mb-1">{label}</label>
      <input
        className="w-full p-2 border rounded"
        type={type}
        name={name}
        value={value}
        onChange={onChange}
      />
      {error && <p className="text-red-600 text-sm mt-1">{error}</p>}
    </div>
  );
}
