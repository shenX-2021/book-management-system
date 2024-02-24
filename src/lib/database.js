const BOOKS = [
  {
    name: 'The Moon and Sixpence',
    author: 'Maugham, W. Somerset',
    ISBN: '9781420951929',
    published: '2016-01-05',
  },
  {
    name: 'The Old Man and the Sea',
    author: 'Hemingway, Ernest',
    ISBN: '9780684801223',
    published: '1995-05-05',
  },
  {
    name: 'The Nightingale and the Rose',
    author: 'Wilde, Oscar',
    ISBN: '9798461534783',
    published: '2021',
  },
  {
    name: 'Naissance Du Purgatoire',
    author: 'Le Jacques, Goff',
    ISBN: '9782070326440',
    published: '1991-09-01',
  },
];

class BookEntity {
  /**
   * init database, and add init book data
   * @returns undefined
   */
  async init() {
    return new Promise((resolve, reject) => {
      let indexedDB =
        window.indexedDB ||
        window.mozIndexedDB ||
        window.webkitIndexedDB ||
        window.msIndexedDB;
      // open database, or create if not exists
      const request = indexedDB.open(BookEntity.DB_NAME, BookEntity.DB_VERSION);
      request.onsuccess = (event) => {
        let db = event.target.result;
        this.db = db;
        resolve();
        console.log('open database success');
      };
      // open database failed
      request.onerror = (event) => {
        console.log('open database failed:', event.target.errorCode);
        reject(
          new Error(
            `can not open the database! errorCode is ${event.target.errorCode}`,
          ),
        );
      };
      // database upgrade callback
      request.onupgradeneeded = (event) => {
        console.log('onupgradeneeded');
        this.db = event.target.result;
        let objectStore;

        objectStore = this.db.createObjectStore(BookEntity.STORE_NAME, {
          keyPath: 'id',
          autoIncrement: true,
        });

        objectStore.createIndex('ISBN', 'ISBN', { unique: true });
        objectStore.createIndex('name', 'name', { unique: false });

        objectStore.transaction.oncomplete = () => {
          const bookObjectStore = this.db
            .transaction(BookEntity.STORE_NAME, 'readwrite')
            .objectStore(BookEntity.STORE_NAME);
          BOOKS.forEach((book) => {
            bookObjectStore.add(book);
          });
        };
      };
    });
  }

  /**
   * search all book
   * @returns Book[]
   */
  findAll() {
    return new Promise((resolve) => {
      const books = [];

      const objectStore = this.db
        .transaction(BookEntity.STORE_NAME)
        .objectStore(BookEntity.STORE_NAME);

      objectStore.openCursor().onsuccess = (event) => {
        const cursor = event.target.result;
        if (cursor) {
          books.push(cursor.value);
          cursor.continue();
        } else {
          resolve(books);
        }
      };
    });
  }

  /**
   * search book by id
   * @param {number} id book id
   * @returns Book
   */
  findById(id) {
    return new Promise((resolve) => {
      const transaction = this.db.transaction([BookEntity.STORE_NAME]);
      const objectStore = transaction.objectStore(BookEntity.STORE_NAME);
      const request = objectStore.get(id);
      request.onerror = () => {
        // todo: error handle
      };
      request.onsuccess = () => {
        console.log(`id [${id}] result:`, request.result);
        resolve(request.result);
      };
    });
  }

  /**
   * add book
   * @param {Book} book book data
   * @returns undefined
   */
  add(book) {
    return new Promise((resolve) => {
      const transaction = this.db.transaction(
        [BookEntity.STORE_NAME],
        'readwrite',
      );

      transaction.oncomplete = () => {
        resolve();
      };

      transaction.onerror = () => {
        // todo: error handle
      };

      console.log(book);
      const objectStore = transaction.objectStore(BookEntity.STORE_NAME);
      const request = objectStore.add(book);
      request.onerror = () => {
        // todo: error handle
      };
    });
  }

  /**
   * delete book by id
   * @param {number} id book id
   * @returns undefined
   */
  delById(id) {
    return new Promise((resolve) => {
      const request = this.db
        .transaction([BookEntity.STORE_NAME], 'readwrite')
        .objectStore(BookEntity.STORE_NAME)
        .delete(id);

      request.onsuccess = () => {
        resolve();
      };
    });
  }

  /**
   * edit book by id
   * @param {number} id book id
   * @param {book} id book data
   * @returns undefined
   */
  editById(id, book) {
    return new Promise((resolve) => {
      const objectStore = this.db
        .transaction([BookEntity.STORE_NAME], 'readwrite')
        .objectStore(BookEntity.STORE_NAME);
      const request = objectStore.get(id);
      request.onerror = () => {
        // todo: error handle
      };
      request.onsuccess = () => {
        // make sure not change other book data
        const data = Object.assign({}, book, { id });
        const requestUpdate = objectStore.put(data);
        requestUpdate.onerror = () => {
          // todo: error handle
        };
        requestUpdate.onsuccess = () => {
          resolve();
        };
      };
    });
  }
}

BookEntity.DB_NAME = 'Book';
BookEntity.DB_VERSION = 3;
BookEntity.STORE_NAME = 'ook';

export const bookEntity = new BookEntity();
