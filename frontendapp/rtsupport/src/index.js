import React from 'react';
import ReactDOM from 'react-dom';
import './app.css';
//import App from './App';
import App from './components/App.jsx';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<App />, document.getElementById('root'));
registerServiceWorker();
