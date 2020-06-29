import React from 'react';
import Header from "./components/Header";
import Characters from "./components/Characters";
import './App.css';

export default () => {
    return (
        <div className="App">
            <Header/>
            <Characters/>
        </div>
    )
}

