<template>
  <div class="container">
    <n-space class="pb20">
      <n-input
        v-model:value="formData.name"
        clearable
        :show-button="false"
        placeholder="Search By Name"
      />

      <n-button type="info" @click="search"> Search </n-button>
      <n-button type="success" @click="add"> Add </n-button>
    </n-space>
    <n-data-table
      :columns="columns"
      :data="data"
      :pagination="pagination"
      :bordered="false"
    />
  </div>
  <add-or-edit-book
    mode="add"
    v-model:show="isShowAdd"
    @update-books="search"
  />
  <add-or-edit-book
    mode="edit"
    :book="editBook"
    v-model:show="isShowEdit"
    @update-books="search"
  />
</template>

<script setup>
import AddOrEditBook from './components/AddOrEditBook.vue';
import { NButton, NSpace, NPopconfirm } from 'naive-ui';
import { h, ref, reactive } from 'vue';
import { deleteApi, listApi } from './http';

// table prop
const columns = [
  {
    title: 'ID',
    key: 'id',
  },
  {
    title: 'Name',
    key: 'name',
  },
  {
    title: 'Author',
    key: 'author',
  },
  {
    title: 'Published',
    key: 'published',
  },
  {
    title: 'ISBN',
    key: 'ISBN',
  },
  {
    title: 'Action',
    key: 'actions',
    render(row) {
      return h(NSpace, () => [
        h(
          NButton,
          {
            strong: true,
            size: 'small',
            type: 'info',
            onClick: () => {
              editBook.value = row;
              isShowEdit.value = true;
            },
          },
          { default: () => 'Edit' },
        ),
        h(
          NPopconfirm,
          {
            'negative-text': 'cancel',
            'positive-text': 'confirm',
            onPositiveClick: () => {
              del(row.id);
            },
            onnNegativeClick: () => {},
          },
          {
            default: () => `Are you sure you want to delete it?`,
            trigger: () =>
              h(
                NButton,
                {
                  strong: true,
                  size: 'small',
                  type: 'error',
                  // onClick: () => del(row.id),
                },
                { default: () => 'Delete' },
              ),
          },
        ),
        ,
      ]);
    },
  },
];
const pagination = ref(false);
const data = ref([]);
const formData = reactive({
  name: '',
});

// add book prop
const isShowAdd = ref(false);

// edit book prop
const isShowEdit = ref(false);
const editBook = ref();

// add book data
async function add() {
  isShowAdd.value = true;
}

// serach all book
async function search() {
  data.value = await listApi(formData);
}

// delete book by id, and refresh book data
async function del(id) {
  await deleteApi(id);
  await search();
}

async function onCreated() {
  await search();
}
onCreated();
</script>

<style scoped lang="scss">
.container {
  padding: 30px 60px;
}
.pb20 {
  padding-bottom: 20px;
}
</style>
