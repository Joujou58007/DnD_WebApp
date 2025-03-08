<script>
	import { onMount } from 'svelte';
  
	let message = 'Loading...';
  
	onMount(async () => {
	  const response = await fetch('http://localhost:8080/api/hello');
	  const data = await response.text();
	  message = data;
	});
  </script>
  
  <main>
	<h1>Hello from Svelte!</h1>
	<p>{message}</p>
	<p>LOULOU</p>
	<p>TOUPINOU</p>
	<p>LIAMOU</p>
  </main>
  
  <style>
	main {
	  text-align: center;
	  padding: 1em;
	  max-width: 240px;
	  margin: 0 auto;
	}
	p {
	  color: red;
	}
  </style>