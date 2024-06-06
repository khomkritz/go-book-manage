<template>
    <div v-if="books">
      <h2>Books List</h2>
      <ul>
        <li v-for="book in books" :key="book.id">
          {{ book.title }} - {{ book.author }}
        </li>
      </ul>
    </div>
    <div v-else>Loading books...</div>
  </template>
  
  <script>
  import axios from 'axios';
  
  export default {
    data() {
      return {
        books: null,
      };
    },
    created() {
      this.fetchBooks();
    },
    methods: {
      async fetchBooks() {
        try {
          const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTc2NjgyOTcsImlzcyI6IjEifQ.PEZnIqgNARr2VCDNFWQ6o_Cj6RyditBc7RypAIOoqOk"
          const response = await axios.get('http://localhost:8000/api/books', {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
          this.books = response.data;
        } catch (error) {
          console.error('Error fetching books:', error);
        }
      },
    },
  };
  </script>
  
  <style scoped>
  /* Your styles for the books list */
  </style>
  