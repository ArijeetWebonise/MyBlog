import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

ReactDOM.render(<App />, document.getElementById('app-root'));

if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/service-worker.js');
}
