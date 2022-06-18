import React, {useEffect, useState} from 'react';
import axios from 'axios'
import './App.css';
import Todo from './components/todo'
import TodoForm from './components/todoForm'

function App() {

  //Fetch data
  const [todos, setTodos] = useState([])

  useEffect(() => {
    fetchData();
  },[]);

  const fetchData = async () => {
    // debugger;
    const result = await axios.get(
      '/tasks',
    );
    console.log("init fetch")
    console.log(result.data)
    setTodos(result.data);
  };


  const addTodo = async text =>{
    //Create a task object based on text
    var newTask = {
      text: text,
      isComplete : false
    }

    //Post it to server
    try {
      // debugger
      const resp = await axios.post('/tasks', newTask);
      console.log("add")
      // console.log(resp)
      setTodos(resp.data)  
    } catch (err) {
      // Handle Error Here
      console.error(err);
    }
  }

  const completeTodo = async index =>{
    setTodos()


  }

  const deleteTodo = async index =>{
    const deletedTodo = todos[index];
    console.log(deletedTodo)
    try {
      // debugger
      const resp = await axios.delete('/tasks', {data: {id: deletedTodo.id}});
      console.log("delete")
      // console.log(resp)
      setTodos(resp.data)  
    } catch (err) {
      // Handle Error Here
      console.error(err);
    }
  }
  
  

  return (
    <div className="app">

          <div className="todo-list">
            {todos.map((todo, index) => (
              <Todo
                key={index}
                index={index}
                todo={todo}
                completeTodo={completeTodo}
                removeTodo={deleteTodo}
              />
            ))}

          <TodoForm addTodo={addTodo}/>
          </div>
        </div>
  );
}

export default App;
