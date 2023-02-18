<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { loggedIn, selectedSchools, token } from '../../stores';
	import type { generateResponse } from '../../types';
	import { _submitGenerateRequest } from './+page';

  let response: generateResponse;
  let ready: boolean = false;

  onMount( async () => {
    if (!$loggedIn) {
			goto('/login');
		} else if (!$token) {
				goto('/');
		} else {
      const schools = $selectedSchools;
      const tokenString = $token.value;
      response = await _submitGenerateRequest(schools, tokenString);
      ready = true;
      console.log(response)
    }
		
  })

</script>

<div class="w-full h-full">
  {#if ready}
    <iframe title="report" src={response.url}/>
  {/if}
</div>
