import { injectGlobal } from 'styled-components';
import GothamFont from '../fonts/gotham-book.ttf';
import GothamMediumFont from '../fonts/GothamMedium_1.ttf';

const globalCss = injectGlobal`
  @font-face {
    font-family: Gotham;
    src: url('${GothamFont}') format('opentype');
  }

  @font-face {
    font-family: Gotham-Medium;
    src: url('${GothamMediumFont}') format('opentype');
  }

  html, body {
    font-size: 62.5% !important; //1rem, 1px
    font-family: 'Gotham', sans-serif !important;
    width: 100%;
    height: 100%;
    margin: 0;
    padding: 0;
    position: relative;
    /* overflow-x: hidden; */
    -webkit-overflow-scrolling: touch;
  }

  #app {
    height: 100%;
  }

  ::-webkit-scrollbar {
    width: 0px;  /* remove scrollbar space */
    background: transparent;  /* optional: just make scrollbar invisible */
  }
  /* optional: show position indicator in red */
  ::-webkit-scrollbar-thumb {
    background: #FF0000;
  }

  .ant-form-explain {
    display: none !important;
  }
  
  .modal-button {
    font-size: 13px;
    font-size: 1.3rem;
  }


  .full-height {
    height: 100%;
  }

  .switch-wrapper {
    position: relative;
    width: 100%;
    min-height: 100%;
  }

  .switch-wrapper > div {
    position: absolute;
    background: #fff;
    min-width: 100%;
    min-height: 100%;
    margin: 0 auto;
    box-shadow: 2px 0 30px rgba(0,0,0,.2);
    z-index: 1;
  }

  .ant-form-explain {
    font-size: 11px;
  }

  .ant-form-item {
    margin-bottom: 0px;
  }
`;

export default globalCss;
