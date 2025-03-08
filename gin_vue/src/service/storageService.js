// 本地缓存服务
const PREFIX = 'gin_vue';
// user 模块
const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token`;
const USER_INFO = `${USER_PREFIX}info`;

// 储存
const set = (key, date) => {
  localStorage.setItem(key, date);
};

// 读取
const get = (key) => localStorage.getItem(key);

export default {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};
