<script lang="ts">
	import { goto } from '$app/navigation';
  import { loggedIn } from '../stores';
  import "../app.css";
	import { onMount } from 'svelte';

  let loginButtonText: string;
  let showButton: boolean;

  const buttonClickHandler = (e: MouseEvent & { currentTarget: EventTarget & HTMLButtonElement;}) => {
    console.log($loggedIn)
		e.preventDefault;
    if ($loggedIn) {
      loggedIn.set(false);
      loginButtonText = 'login'
      goto('/');
	  } else {
      goto('/login');
      showButton = false;
    }
  }
  onMount(() => {
    showButton = !(window.location.pathname === '/login');
  })
</script>


<div id="header" class="flex flex-row pb-8 pt-2 mb-8 justify-end">
	<nav>
    {#if showButton}
      <button class="border-2 mr-4 ml-4 pl-2 pr-2 rounded-sm"
      on:click={(e) => buttonClickHandler(e)}>
        {$loggedIn ? 'logout' : 'login'}
      </button>
    {/if}
  </nav>
</div>
<slot>

</slot>