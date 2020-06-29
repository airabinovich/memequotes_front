import React from 'react';
import Header from "./components/Header";
import Characters from "./components/Characters";
import './App.css';

export default () => {
    let cols = [
        {
            "name": "id",
            "visualName": "ID"
        }, {
            "name": "name",
            "visualName": "Character"
        }, {
            "name": "fun_facts",
            "visualName": "Curiosidades"
        }
    ];
    let rows = [
        {
            "id": 1,
            "name": "Comandante",
            "fun_facts": "usa el cutucuchillo"
        },
        {
            "id": 2,
            "name": "Menem",
            "fun_facts": "Tiene un cohete que te manda a la estratosfera"
        },
        {
            "id": 3,
            "name": "El Chino Cirujano",
            "fun_facts": "Junta la latita comueso"
        }
    ];
    return (
        <div className="App">
            <Header/>
            <Characters/>
        </div>
    )
}

/*function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}
*/

