<script lang="ts">
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import { loggedIn, selectedSchools, token } from '../../stores';
	import type { generateRequest } from '../../types';

  let url: string
  let ready: boolean = false;

  onMount( async () => {
    if (!$loggedIn) {
			goto('/login');
		} else if (!$token) {
				goto('/');
		} else {
      const schools = $selectedSchools;
      const tokenString = $token.value;
      const generateRequestBody: generateRequest = {schools, tokenString};
      const response = await fetch('/server/view', {
        method: 'POST',
        body: JSON.stringify(generateRequestBody),
      })
      url = URL.createObjectURL(await response.blob())
      ready = true;
    }
		
  })

</script>

<div class="w-full h-screen">
  {#if ready}
    <object title="report" data={url} type='application/pdf' class="w-full h-full"/>
  {/if}
</div>
