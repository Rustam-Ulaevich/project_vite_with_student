// Подключаем установленные пакеты
const express = require('express');
const cors = require('cors');

// Создаём экземпляр приложения express
const app = express();

// Настройка middleware
app.use(cors());              // Разрешаем запросы с любых доменов
app.use(express.json());     // Автоматически парсим JSON в теле запроса

// Определяем порт, на котором будет работать сервер
// process.env.PORT нужен для облачных хостингов, 3001 - локальный порт по умолчанию
const PORT = process.env.PORT || 3001;

// ВРЕМЕННОЕ ХРАНИЛИЩЕ ДАННЫХ (в памяти сервера)
// При остановке сервера все данные исчезнут
const database = {
    users: [
        { id: 1, name: 'Анна', email: 'anna@example.com' },
        { id: 2, name: 'Петр', email: 'petr@example.com' }
    ],
    posts: [
        { id: 1, title: 'Первый пост', content: 'Содержание', userId: 1 }
    ]
};

// ВСПОМОГАТЕЛЬНАЯ ФУНКЦИЯ для генерации ID
// Используем Date.now() для уникальности
const generateId = () => Date.now();


// МАРШРУТЫ (ENDPOINTS)

// Базовый маршрут для проверки работы сервера
app.get('/', (req, res) => {
    res.json({ 
        message: 'Мой Mock API работает!',
        endpoints: {
            users: '/api/users',
            posts: '/api/posts',
            dynamic: '/api/:resource'
        }
    });
});

// ПОЛУЧИТЬ ВСЕ ЗАПИСИ ИЗ РЕСУРСА
// Пример: GET /api/users
app.get('/api/:resource', (req, res) => {
    const { resource } = req.params;
    
    // Проверяем, существует ли такой ресурс в нашей базе
    if (database[resource]) {
        res.json(database[resource]);
    } else {
        // Если ресурса нет, возвращаем ошибку 404
        res.status(404).json({ 
            error: `Ресурс '${resource}' не найден. Доступные ресурсы: ${Object.keys(database).join(', ')}`
        });
    }
});

// ПОЛУЧИТЬ ОДНУ ЗАПИСЬ ПО ID
// Пример: GET /api/users/1
app.get('/api/:resource/:id', (req, res) => {
    const { resource, id } = req.params;
    
    if (!database[resource]) {
        return res.status(404).json({ error: `Ресурс '${resource}' не найден` });
    }
    
    const item = database[resource].find(item => item.id === parseInt(id));
    
    if (item) {
        res.json(item);
    } else {
        res.status(404).json({ error: `Запись с id ${id} не найдена в ресурсе '${resource}'` });
    }
});

// ➕ СОЗДАТЬ НОВУЮ ЗАПИСЬ
// Пример: POST /api/users с телом { "name": "Новый", "email": "new@example.com" }
app.post('/api/:resource', (req, res) => {
    const { resource } = req.params;
    
    if (!database[resource]) {
        return res.status(404).json({ error: `Ресурс '${resource}' не найден` });
    }
    
    // Создаём новую запись: добавляем ID и объединяем с данными из запроса
    const newItem = {
        id: generateId(),
        ...req.body
    };
    
    database[resource].push(newItem);
    res.status(201).json(newItem); // 201 = Created
});

// ОБНОВИТЬ ЗАПИСЬ
// Пример: PUT /api/users/1 с телом { "name": "Обновлённое имя" }
app.put('/api/:resource/:id', (req, res) => {
    const { resource, id } = req.params;
    
    if (!database[resource]) {
        return res.status(404).json({ error: `Ресурс '${resource}' не найден` });
    }
    
    const index = database[resource].findIndex(item => item.id === parseInt(id));
    
    if (index === -1) {
        return res.status(404).json({ error: `Запись с id ${id} не найдена` });
    }
    
    // Обновляем запись: сохраняем старый ID, обновляем остальные поля
    const updatedItem = {
        ...database[resource][index],
        ...req.body,
        id: parseInt(id) // гарантируем, что ID не изменится
    };
    
    database[resource][index] = updatedItem;
    res.json(updatedItem);
});

// УДАЛИТЬ ЗАПИСЬ
// Пример: DELETE /api/users/1
app.delete('/api/:resource/:id', (req, res) => {
    const { resource, id } = req.params;
    
    if (!database[resource]) {
        return res.status(404).json({ error: `Ресурс '${resource}' не найден` });
    }
    
    const initialLength = database[resource].length;
    database[resource] = database[resource].filter(item => item.id !== parseInt(id));
    
    if (database[resource].length === initialLength) {
        return res.status(404).json({ error: `Запись с id ${id} не найдена` });
    }
    
    res.status(204).send(); // 204 = No Content (успешно удалено)
});




// 🆕 СОЗДАТЬ НОВЫЙ РЕСУРС (динамически)
// Пример: POST /api/resources с телом { "name": "products", "schema": ["name", "price"] }
app.post('/api/resources', (req, res) => {
    const { name, initialData = [] } = req.body;
    
    // Проверяем, не существует ли уже такой ресурс
    if (database[name]) {
        return res.status(400).json({ error: `Ресурс '${name}' уже существует` });
    }
    
    // Создаём новый ресурс с начальными данными
    database[name] = initialData;
    
    res.status(201).json({
        message: `Ресурс '${name}' создан`,
        availableResources: Object.keys(database)
    });
});

// 📋 ПОЛУЧИТЬ СПИСОК ВСЕХ ДОСТУПНЫХ РЕСУРСОВ
app.get('/api/resources', (req, res) => {
    res.json({
        resources: Object.keys(database),
        endpoints: Object.keys(database).map(r => `/api/${r}`)
    });
});




// ЗАПУСК СЕРВЕРА
app.listen(PORT, () => {
    console.log(`🚀 Сервер запущен на http://localhost:${PORT}`);
    console.log(`📝 Доступные ресурсы: ${Object.keys(database).join(', ')}`);
    console.log(`📋 Пример запроса: GET http://localhost:${PORT}/api/users`);
});