<div class={styles.main}>
  <h1> Login Form</h1>
  <input type="text" bind:value={username} required={true}> 
  <input type="password" bind:value={password} required={true}>
  <button on:click={handleLogin}>
    submit
  </button>
</div>

<script lang="ts">
  import { setContext } from 'svelte';
	import { goto } from '$app/navigation';
  import type { loginResponse } from '../../types';
  import styles from './login.module.css';
	import { loggedIn, token } from '../../stores';
  let username: string, password: string;

  const handleLogin =async () => {
    const destructuredToken = await login(username, password);

    if (destructuredToken) {
      token.set(destructuredToken);
      loggedIn.set(true);
      goto('/submit');
    } else {
      window.alert('Unable to log in');
    }
  }

  export const login = async (user: string, pass: string): Promise<loginResponse | null> => {
    const apiResponse = await fetch('http://localhost:3000/api/login', {
      method: 'POST',
      body: JSON.stringify({
        username: user,
        password: pass
      })
    })
    if (apiResponse.ok) {
      const response: loginResponse = await apiResponse.json();
      return response;
    }
    return null;
  }
</script>
