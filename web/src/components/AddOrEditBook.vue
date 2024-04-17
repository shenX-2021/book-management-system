<template>
  <n-modal
    :show="props.show"
    :mask-closable="false"
    preset="dialog"
    :title="title"
    positive-text="Submit"
    negative-text="Cancel"
    @positive-click="onPositiveClick"
    @negative-click="onNegativeClick"
  >
    <template #default>
      <n-form ref="formRef" :model="formData" :rules="rules">
        <div class="n-form-item-label n-form-item-label--right-mark id">
          ID: {{ formData.id }}
        </div>
        <n-form-item path="name" label="Name">
          <n-input
            v-model:value="formData.name"
            placeholder="Please input Name"
            clearable
            first
            @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="author" label="Author">
          <n-input
            v-model:value="formData.author"
            placeholder="Please input Author"
            clearable
            first
            @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="ISBN" label="ISBN">
          <n-input-number
            v-model:value="formData.ISBN"
            :show-button="false"
            placeholder="Please input ISBN"
            style="width: 100%"
            clearable
            first
            @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="published" label="Published">
          <n-input
            v-model:value="formData.published"
            placeholder="Please input Published Date"
            clearable
            first
            @keydown.enter.prevent
          />
        </n-form-item>
      </n-form>
    </template>

    <template #close><div></div></template>
  </n-modal>
</template>

<script setup>
import { defineProps, defineEmits, watch, reactive, toRaw } from 'vue';
import { createApi, editApi } from '../http';

const MODE = {
  EDIT: 'edit',
  ADD: 'add',
};

const props = defineProps({
  show: Boolean,
  book: Object,
  mode: String,
});
const emit = defineEmits(['update:show', 'update-books']);
const formData = reactive({});
const title = ref('');
const formRef = ref();

const rules = {
  name: [
    {
      required: true,
      message: 'Please input Name',
    },
  ],
  author: [
    {
      required: true,
      message: 'Please input Author',
    },
  ],
  ISBN: [
    {
      required: true,
      message: 'Please input ISBN',
    },
    {
      validator: (rule, value) => {
        if (value === undefined || value === null) return true;
        const str = value.toString();
        return str.length === 10 || str.length === 13;
      },
      message: 'Lenght of ISBN must be 10 or 13',
    },
  ],
  published: [
    {
      required: true,
      message: 'Please input Published Date',
    },
    {
      validator: (rule, value) => {
        if (value === undefined || value === null) return true;
        return /^(\d{4})([/-]\d{1,2})?([/-]\d{1,2})?$/.test(value);
      },
      message: 'format of Published Date is wrong',
    },
  ],
};

function onNegativeClick() {
  emit('update:show', false);
  clearFormData();
}
async function onPositiveClick() {
  formRef.value.validate(async (errors) => {
    if (!errors) {
      // handle with different mode
      if (props.mode === MODE.ADD) {
        // await bookEntity.add(toRaw(formData));
        await createApi(toRaw(formData));
      } else {
        await editApi(formData.id, toRaw(formData));
      }

      emit('update:show', false);
      emit('update-books');
      clearFormData();
    } else {
      console.log(errors);
    }
  });
}

function clearFormData() {
  delete formData.id;
  delete formData.name;
  delete formData.author;
  delete formData.ISBN;
  delete formData.published;
}

watch(
  () => props.show,
  (newVal) => {
    if (newVal) {
      title.value = props.mode === MODE.ADD ? 'Add Book' : 'Edit Book';

      if (props.mode === MODE.EDIT && props.book) {
        Object.assign(formData, props.book);
      }
    }
  },
);
</script>

<style lang="scss" scoped>
.id {
  margin-bottom: 8px;
  font-size: 16px;
}
</style>
