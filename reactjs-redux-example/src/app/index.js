import React, { Component } from 'react';
import Counter from './components/Counter';
import { Provider } from 'react-redux';
import store from './store';
class App extends Component {
  render() {
    return (
      <Provider store={store}>
        <div className="App">
          <Counter />
        </div>
      </Provider>
    );
  }
}

export default App;
