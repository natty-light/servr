<div class={styles.main}>
  <h1> Login Form</h1>
  <input type="text" bind:value={username} required={true}> 
  <input type="password" bind:value={password} required={true}>
  <button on:click={handleLogin}>
    submit
  </button>
</div>

<script lang="ts">
	import { goto } from '$app/navigation';
  import { login } from './+page.server';
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

</script>
