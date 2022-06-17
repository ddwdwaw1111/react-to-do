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
      'http://localhost:8080/tasks',
    );

    console.log(result)

    setTodos(result.data);
  };


  const addTodo = async text =>{
    // type Task struct {
    //   Id   int    `json:"id"`
    //   Text string `json:"text"`
    //   // Tags []string  `json:"tags"`
    //   Due        time.Time `json:"due"`
    //   IsComplete bool      `json:"isComplete"`
    // }

    //Create a task object based on text
    var newTask = {
      text: text,
      isComplete : false
    }
    //Post it to server
    try {
      // debugger
      const resp = await axios.post('http://localhost:8080/tasks', newTask);
      console.log(resp);
  } catch (err) {
      // Handle Error Here
      console.error(err);
  }
    //Call Fetch function again

    // const newTodos = [...todos, {text}];
    // setTodos(newTodos);
  }

  const completeTodo = index =>{
    const newTodos = [...todos];
    newTodos[index].isComplete = true;
    setTodos(newTodos)

  }

  const removeTodo = index =>{
    const newTodos = [...todos];
    newTodos.splice(index,1)
    setTodos(newTodos)
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
                removeTodo={removeTodo}
              />
            ))}

          <TodoForm addTodo={addTodo}/>
          </div>
        </div>
  );
}

export default App;
