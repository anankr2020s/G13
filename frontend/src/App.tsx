import React from 'react';
import Navbar from './components/Navbar';
import './App.css';
import Input from './components/Input';
import { Button } from '@material-ui/core';
import Home from './components/home';



function App() {
  return (
    <div className = "app" >
      <Navbar/>
      <Input/>
    </div>
  );
}

export default App;
