// frontend/src/utils/jwt_utils.js

/**
 * Decodifica el payload de un JWT sin validar firma
 * para extraer el claim user_id
 */
export function getUserIdFromToken() {
  const token = localStorage.getItem('token');
  if (!token) return null;
  try {
    const payload = token.split('.')[1];
    const decoded = JSON.parse(atob(payload));
    return decoded.user_id;        // coincide con claim "user_id"
  } catch {
    return null;
  }
}
