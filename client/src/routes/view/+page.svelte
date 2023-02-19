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
    }
		
  })

</script>

<div class="w-full h-screen">
  {#if ready}
    <object title="report" data={response.url} type='application/pdf' class="w-full h-full"/>
  {/if}
</div>
