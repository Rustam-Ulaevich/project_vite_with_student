import { useState, useEffect } from 'react';

// Определяем структуру пользователя
interface User {
  id: number;
  name: string;
  email: string;
}

// Определяем структуру ответа API (если нужно)
interface ApiResponse {
  message?: string;
  error?: string;
}

const API_URL = 'http://localhost:3001/api';

export function App() {
  // ✅ Явно указываем тип массива пользователей
  const [users, setUsers] = useState<User[]>([]);
  const [newUserName, setNewUserName] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>('');

  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}/users`);
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      const data: User[] = await response.json(); // ✅ Явно указываем тип данных
      setUsers(data);
    } catch (error) {
      console.error('Ошибка загрузки пользователей:', error);
      setError('Не удалось загрузить пользователей');
    } finally {
      setLoading(false);
    }
  };

  const createUser = async () => {
    if (!newUserName.trim()) {
      setError('Введите имя пользователя');
      return;
    }

    setLoading(true);
    try {
      const response = await fetch(`${API_URL}/users`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ 
          name: newUserName.trim(), 
          email: `${newUserName.trim().toLowerCase()}@example.com` 
        })
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      const newUser: User = await response.json(); // ✅ Явно указываем тип
      setUsers([...users, newUser]); // ✅ Теперь TypeScript доволен
      setNewUserName('');
      setError('');
    } catch (error) {
      console.error('Ошибка создания пользователя:', error);
      setError('Не удалось создать пользователя');
    } finally {
      setLoading(false);
    }
  };

  const deleteUser = async (id: number) => {
    setLoading(true);
    try {
      const response = await fetch(`${API_URL}/users/${id}`, { 
        method: 'DELETE' 
      });
      
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      
      setUsers(users.filter(user => user.id !== id)); // ✅ TypeScript понимает, что user - это User
      setError('');
    } catch (error) {
      console.error('Ошибка удаления пользователя:', error);
      setError('Не удалось удалить пользователя');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'Arial' }}>
      <h1>Мой Mock API</h1>
      
      {error && (
        <div style={{ color: 'red', marginBottom: '10px' }}>
          {error}
        </div>
      )}

      <div style={{ marginBottom: '20px' }}>
        <input 
          value={newUserName} 
          onChange={(e) => setNewUserName(e.target.value)}
          placeholder="Имя пользователя"
          disabled={loading}
          style={{ marginRight: '10px', padding: '5px' }}
        />
        <button 
          onClick={createUser} 
          disabled={loading}
          style={{ padding: '5px 10px' }}
        >
          {loading ? 'Загрузка...' : 'Добавить пользователя'}
        </button>
      </div>

      <h2>Список пользователей:</h2>
      {loading && users.length === 0 ? (
        <p>Загрузка...</p>
      ) : (
        <ul>
          {users.map(user => (
            <li key={user.id} style={{ marginBottom: '10px' }}>
              {user.name} ({user.email})
              <button 
                onClick={() => deleteUser(user.id)} 
                disabled={loading}
                style={{ marginLeft: '10px', padding: '2px 8px' }}
              >
                Удалить
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
