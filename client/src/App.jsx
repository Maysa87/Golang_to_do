import React from "react";
import "./App.css";
import {Container} from 'semantic-ui-react';
import Todolist from './To-do';

function App(){
  return (
    <div>
      <Container>
        <Todolist></Todolist>
      </Container>
    </div>
  )
}

export default App;
