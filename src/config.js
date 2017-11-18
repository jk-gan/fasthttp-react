let API_ROOT_URL = '';
if (/revenuemonster.my/gi.test(window.location.hostname)) {
  API_ROOT_URL = 'https://dev-api.revenuemonster.my';
} else {
  API_ROOT_URL = 'https://dev-api.revenuemonster.my';
  // API_ROOT_URL = 'https://14ca70ed.ngrok.io';
}
const VERSION = 'api/v1';

export default {
  API_ROOT: API_ROOT_URL,
  API_ROOT_URL: `${API_ROOT_URL}/${VERSION}`,
  QRCODE_UPLOAD_URL: `${API_ROOT_URL}/${VERSION}/loyalty/member/upload/qrcode/redeem`,
  AVATAR_UPLOAD_URL: `${API_ROOT_URL}/${VERSION}/user/avatar`,
};
