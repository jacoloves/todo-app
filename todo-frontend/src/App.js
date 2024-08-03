import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTimes } from '@fortawesome/free-solid-svg-icons';

function App() {
  const [tasks, setTasks] = useState([]);
  const [newTask, setNewTask] = useState('');

  useEffect(() => {
    axios.get('http://localhost:8081/tasks')
      .then(response => setTasks(response.data))
      .catch(error => console.error('Error fetching tasks:', error))
  }, []);

  const addTask = () => {
    if (newTask.trim() === '') return;
    axios.post('http://localhost:8081/tasks', { text: newTask })
      .then(response => {
        setTasks([...tasks, response.data]);
        setNewTask('');
      })
      .catch(error => console.error('Error adding task:', error));
  };

  const toggleTask = (id) => {
    axios.put(`http://localhost:8081/tasks/${id}/toggle`)
      .then(response => {
        setTasks(tasks.map(task => task.id === id ? response.data : task));
      })
      .catch(error => console.error('Error toggling task:', error));
  };

  const deleteTask = (id) => {
    axios.delete(`http://localhost:8081/tasks/${id}`)
      .then(() => {
        setTasks(tasks.filter(task => task.id !== id));
      })
      .catch(error => console.error('Error deleteing task:', error));
  };

  const handleKeyPress = (e) => {
    if (e.key === 'Enter') {
      addTask();
    }
  };

  return (
    <div className="App container mx-auto p-4">
      <h1 className="text-3xl font-bold mb-4">TODO App</h1>
      <div className="flex mb-4">
        <input
          type="text"
          value={newTask}
          onChange={(e) => setNewTask(e.target.value)}
          onKeyPress={handleKeyPress}
          className="border-2 border-black rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline mr-2"
          placeholder="Add a new task"
        />
        <button onClick={addTask} className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline">Add</button>
      </div>
      <ul className="list-disc pl-5 border-t border-gray-300 pt-4">
        {tasks.map(task => (
          <li key={task.id} className="flex items-center mb-2">
            <input
              type="checkbox"
              checked={task.completed}
              onChange={() => toggleTask(task.id)}
              className="mr-2"
            />
            <span className={task.completed ? 'line-through' : ''}>{task.text}</span>
            <button onClick={() => deleteTask(task.id)} className="bg-red-500 hover:bg-red-700 text-white py-1 px-2 rounded ml-2">
              <FontAwesomeIcon icon={faTimes} />
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;
