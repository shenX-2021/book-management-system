import axios from './lib/axios';

export const listApi = (params) => axios.get('/book/list', { params });
export const createApi = (data) => axios.post('/book/create', data);
export const editApi = (id, data) => axios.patch(`/book/${id}`, data);
export const deleteApi = (id) => axios.delete(`/book/${id}`);
