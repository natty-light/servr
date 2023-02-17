<div class={styles.main}>
  <h1>
    Select your schools
  </h1>
</div>

<script lang="ts">
  import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { loggedIn, token } from '../../stores';
  import styles from './submit.module.css';

  onMount( () => {
    if (!$loggedIn) {
      goto('/login')
    } else {
      if (!$token) {
        goto('/')
      }
    }
  })

</script>