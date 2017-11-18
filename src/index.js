import 'babel-polyfill';
import 'defiant';
import 'whatwg-fetch';
import { render } from 'react-dom';
import 'normalize.css';
import 'react-toast-mobile/lib/react-toast-mobile.css';
import routes from './route';

render(routes, document.querySelector('#app'));
