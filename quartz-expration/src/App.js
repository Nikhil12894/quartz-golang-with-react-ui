import './App.css';
import React, { Component } from 'react'
import Cron from './croncomponent/cron';
export class App extends Component{
  render() {
    return (
        <div>
            <Cron/>
        </div>
    );
}
}

export default App;
