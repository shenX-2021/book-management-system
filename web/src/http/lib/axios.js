import AxiosStatic from 'axios';
import { createDiscreteApi } from 'naive-ui';

const { message } = createDiscreteApi(['message']);

const axios = AxiosStatic.create({
  baseURL: '/api',
  timeout: 10000,
});

axios.interceptors.request.use((config) => {
  return config;
});

axios.interceptors.response.use(
  (res) => {
    return res.data;
  },
  (err) => {
    if (err?.message.startsWith('timeout of ')) {
      message.error('请求超时，请稍后再试');
    } else if (err?.message?.startsWith('Network Error')) {
      message.error('网络异常');
    } else {
      message.error(err?.response?.data ?? '未知异常');
    }

    return Promise.reject(err.response.data);
  },
);

export default axios;
